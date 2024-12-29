# Contributing to GoPilot

Thank you for your interest in contributing to GoPilot! This document provides guidelines and instructions for contributing to the project.

## Development Setup

1. Fork and clone the repository:
```bash
git clone https://github.com/yourusername/gopilot.git
cd gopilot
```

2. Install Go 1.22.3 or later

3. Build the project:
```bash
go build -o gopilot ./cmd/main.go
```

## Making Changes

1. Create a new branch for your changes:
```bash
git checkout -b feature/your-feature-name
```

2. Make your changes and commit them:
```bash
git add .
git commit -m "Description of your changes"
```

3. Push to your fork:
```bash
git push origin feature/your-feature-name
```

4. Create a Pull Request from your fork to the main repository

## Release Process

### Creating a New Release

1. Ensure all tests pass and the main branch is stable

2. Create and push a new tag:
```bash
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0
```

### Automated Release Process

When a tag is pushed, the GitHub Actions workflow (.github/workflows/release.yml) automatically:

1. Builds binaries for all supported platforms:
   - Windows (amd64, arm64)
   - macOS (amd64, arm64)
   - Linux (amd64, arm64)

2. Generates SHA256 checksums for each binary

3. Creates a GitHub release with:
   - All platform binaries
   - Checksum files
   - Auto-generated release notes

### Post-Release Verification

After the release workflow completes:

1. Check the [Actions tab](https://github.com/zacharyiles/gopilot/actions) to verify the build succeeded

2. Verify the [release page](https://github.com/zacharyiles/gopilot/releases) contains:
   - All platform binaries
   - SHA256 checksum files
   - Generated release notes

3. Test the installation script:
```bash
curl -sSL https://raw.githubusercontent.com/zacharyiles/gopilot/main/install.sh | bash
```

4. Verify the installed binary works correctly:
```bash
gopilot --version
```

## Code Style

- Follow standard Go formatting conventions
- Use `gofmt` to format your code
- Add comments for exported functions and packages
- Write tests for new functionality

## Testing

Run tests before submitting a PR:

```bash
go test ./...
```

## Documentation

- Update README.md for user-facing changes
- Update code comments for API changes
- Include examples for new features

## Need Help?

- Open an issue for bugs or feature requests
- Ask questions in the discussions section
- Review existing issues and PRs before creating new ones

## License

By contributing to GoPilot, you agree that your contributions will be licensed under the MIT License. 