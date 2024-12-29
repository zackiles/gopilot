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

# Determine OS and architecture
OS="$(uname | tr '[:upper:]' '[:lower:]')"
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

# Set binary name
BINARY_NAME="gopilot-${OS}-${ARCH}"
if [ "${OS}" = "windows" ]; then
    BINARY_NAME="${BINARY_NAME}.exe"
fi

# Set installation directory based on OS
if [ "${OS}" = "darwin" ] || [ "${OS}" = "linux" ]; then
    INSTALL_DIR="/usr/local/bin"
else
    echo "Error: Unsupported operating system: ${OS}"
    exit 1
fi

# Create installation directory if needed
sudo mkdir -p "${INSTALL_DIR}"

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

# Install binary
echo "Installing GoPilot..."
if ! sudo mv "${BINARY_NAME}" "${INSTALL_DIR}/gopilot"; then
    echo "Error: Failed to install binary"
    exit 1
fi

if ! sudo chmod +x "${INSTALL_DIR}/gopilot"; then
    echo "Error: Failed to set executable permissions"
    exit 1
fi

echo "GoPilot ${VERSION} has been installed successfully!"
echo "You can now use it by running 'gopilot' from anywhere in your terminal." 