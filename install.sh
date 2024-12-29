#!/bin/bash

# Determine OS and architecture
OS="$(uname)"
ARCH="$(uname -m)"

# Set binary name based on OS
if [ "$OS" = "Darwin" ]; then
    BINARY_NAME="gopilot-darwin"
elif [ "$OS" = "Linux" ]; then
    BINARY_NAME="gopilot-linux"
else
    echo "Unsupported operating system: $OS"
    exit 1
fi

# Add architecture suffix
if [ "$ARCH" = "x86_64" ]; then
    BINARY_NAME="${BINARY_NAME}-amd64"
elif [ "$ARCH" = "arm64" ]; then
    BINARY_NAME="${BINARY_NAME}-arm64"
else
    echo "Unsupported architecture: $ARCH"
    exit 1
fi

# Create installation directory
INSTALL_DIR="/usr/local/bin"
sudo mkdir -p "$INSTALL_DIR"

# Download latest release
LATEST_RELEASE_URL="https://github.com/yourusername/gopilot/releases/latest/download/${BINARY_NAME}"
sudo curl -L "$LATEST_RELEASE_URL" -o "$INSTALL_DIR/gopilot"

# Make binary executable
sudo chmod +x "$INSTALL_DIR/gopilot"

echo "GoPilot has been installed successfully!"
echo "You can now use it by running 'gopilot' from anywhere in your terminal." 