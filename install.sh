#!/bin/bash

set -eo pipefail

# Check for required commands
if ! command -v curl >/dev/null 2>&1; then
    echo "Error: curl is required but not installed. Please install curl first."
    exit 1
fi

if ! command -v sha256sum >/dev/null 2>&1; then
    echo "Error: sha256sum is required but not installed. Please install it first."
    exit 1
fi

# Check for sudo/root
if [ "$(id -u)" -ne 0 ] && ! command -v sudo >/dev/null 2>&1; then
    echo "Error: This script requires sudo privileges. Please install sudo or run as root."
    exit 1
fi

# Improved OS detection for Windows
OS="$(uname | tr '[:upper:]' '[:lower:]')"
if [[ "$OS" != "linux" && "$OS" != "darwin" ]]; then
    # Check for Windows more thoroughly
    case "$(uname -s)" in
        MINGW*|MSYS*|CYGWIN*)
            OS="windows"
            ;;
        *)
            if [[ "$OS" == *"windows"* ]]; then
                OS="windows"
            else
                echo "Error: Unsupported operating system: $OS"
                exit 1
            fi
            ;;
    esac
fi

# Windows-specific adjustments
if [ "$OS" = "windows" ]; then
    # Use more reliable way to get user profile on Windows
    INSTALL_DIR="$(cygpath "$USERPROFILE" 2>/dev/null || echo "$USERPROFILE")/AppData/Local/Microsoft/WindowsApps"
    # Convert path separators for Windows
    INSTALL_DIR=$(echo "$INSTALL_DIR" | sed 's/\\/\//g')
    BINARY_NAME="gopilot-windows-${ARCH}.exe"  # Name of the release artifact
    FINAL_NAME="gopilot.exe"  # Final name after installation
else
    INSTALL_DIR="/usr/local/bin"
    BINARY_NAME="gopilot-${OS}-${ARCH}"
    FINAL_NAME="gopilot"  # Final name after installation
fi

ARCH="$(uname -m)"

# Map architecture names
case "${ARCH}" in
    x86_64) ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
    armv8*) ARCH="arm64" ;;
    *)
        echo "Error: Unsupported architecture: ${ARCH}"
        exit 1
        ;;
esac

# Create installation directory if needed
mkdir -p "${INSTALL_DIR}" || {
    echo "Error: Failed to create installation directory"
    exit 1
}

# Set GitHub repository
REPO="zacharyiles/gopilot"

# Get latest version with better error handling
echo "Fetching latest version..."
RESPONSE=$(curl -sL "https://api.github.com/repos/${REPO}/releases/latest")
if [[ $(echo "$RESPONSE" | grep -c "API rate limit exceeded") -ne 0 ]]; then
    echo "Error: GitHub API rate limit exceeded. Please try again later."
    exit 1
fi

VERSION=$(echo "$RESPONSE" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
if [ -z "${VERSION}" ]; then
    echo "Error: Could not determine latest version"
    echo "API Response: $RESPONSE"
    exit 1
fi

# Download URLs
BINARY_URL="https://github.com/${REPO}/releases/download/${VERSION}/${BINARY_NAME}"
CHECKSUM_URL="${BINARY_URL}.sha256"

# Create temporary directory
TMP_DIR=$(mktemp -d)
trap 'rm -rf ${TMP_DIR}' EXIT

# Download binary and checksum
echo "Downloading GoPilot ${VERSION}..."
if ! curl -sL -f "${BINARY_URL}" -o "${TMP_DIR}/${BINARY_NAME}"; then
    echo "Error: Failed to download binary from ${BINARY_URL}"
    echo "Please check if the release exists and you have internet connectivity"
    exit 1
fi

if ! curl -sL -f "${CHECKSUM_URL}" -o "${TMP_DIR}/${BINARY_NAME}.sha256"; then
    echo "Error: Failed to download checksum from ${CHECKSUM_URL}"
    echo "Please check if the release exists and you have internet connectivity"
    exit 1
fi

# Verify checksum
echo "Verifying checksum..."
cd "${TMP_DIR}"
if ! sha256sum -c "${BINARY_NAME}.sha256"; then
    echo "Error: Checksum verification failed"
    exit 1
fi

# Modify installation for Windows
if [ "$OS" = "windows" ]; then
    # Create directory without sudo on Windows
    mkdir -p "${INSTALL_DIR}" || {
        echo "Error: Failed to create installation directory"
        exit 1
    }
    
    # Install without sudo on Windows, renaming to gopilot.exe
    mv "${TMP_DIR}/${BINARY_NAME}" "${INSTALL_DIR}/${FINAL_NAME}" || {
        echo "Error: Failed to install binary. Make sure you have write permissions."
        exit 1
    }
    
    chmod +x "${INSTALL_DIR}/${FINAL_NAME}" || true  # chmod may not work on Windows
    
    # Verify PATH on Windows
    if ! echo "$PATH" | tr ':' '\n' | grep -q "${INSTALL_DIR}"; then
        echo "Warning: Installation directory is not in PATH"
        echo "Please add the following directory to your PATH:"
        echo "${INSTALL_DIR}"
    fi
else
    # Unix installation
    if ! sudo mv "${TMP_DIR}/${BINARY_NAME}" "${INSTALL_DIR}/${FINAL_NAME}"; then
        echo "Error: Failed to install binary"
        exit 1
    fi

    if ! sudo chmod +x "${INSTALL_DIR}/${FINAL_NAME}"; then
        echo "Error: Failed to set executable permissions"
        exit 1
    fi
fi

# Verify installation
if command -v gopilot >/dev/null 2>&1; then
    echo "GoPilot ${VERSION} has been installed successfully!"
    echo "You can now use it by running 'gopilot' from your terminal."
else
    echo "GoPilot has been installed to: ${INSTALL_DIR}"
    if [ "$OS" = "windows" ]; then
        echo "You may need to:"
        echo "1. Close and reopen your terminal"
        echo "2. Add the installation directory to your PATH if it's not already there"
    fi
fi