#!/bin/bash
set -eo pipefail

# Default values
VERSION=""
TAG_MESSAGE=""
RELEASE=false

# Help message
show_help() {
    echo "Usage: ./build.sh [options]"
    echo "Options:"
    echo "  -v, --version VERSION    Version to build/tag (e.g., 1.0.0)"
    echo "  -m, --message MESSAGE    Tag message (required for release)"
    echo "  -r, --release           Create and push a release tag"
    echo "  -h, --help              Show this help message"
}

# Parse arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        -v|--version)
            VERSION="$2"
            shift 2
            ;;
        -m|--message)
            TAG_MESSAGE="$2"
            shift 2
            ;;
        -r|--release)
            RELEASE=true
            shift
            ;;
        -h|--help)
            show_help
            exit 0
            ;;
        *)
            echo "Unknown option: $1"
            show_help
            exit 1
            ;;
    esac
done

# If no version specified, use git describe or default to dev
if [ -z "$VERSION" ]; then
    if git describe --tags >/dev/null 2>&1; then
        VERSION=$(git describe --tags)
    else
        VERSION="dev"
    fi
fi

# Build the binary with version information
echo "Building GoPilot version ${VERSION}..."
go build -trimpath -ldflags="-s -w -X main.Version=${VERSION}" -o gopilot ./cmd/main.go

# If this is a release build
if [ "$RELEASE" = true ]; then
    if [[ ! $VERSION =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
        echo "Error: Release version must be in semver format (e.g., 1.0.0)"
        exit 1
    fi
    
    if [ -z "$TAG_MESSAGE" ]; then
        echo "Error: Tag message is required for release. Use -m or --message"
        exit 1
    fi

    # Create and push tag
    echo "Creating release tag v${VERSION}..."
    git tag -a "v${VERSION}" -m "${TAG_MESSAGE}"
    git push origin "v${VERSION}"
    
    echo "Release tag pushed. GitHub Actions will now:"
    echo "1. Build binaries for all platforms"
    echo "2. Generate checksums"
    echo "3. Create a GitHub release"
    echo "4. Upload all artifacts"
fi

echo "Build complete!" 