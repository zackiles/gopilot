# GoPilot: A Unified No-Frills LLM CLI in Go

GoPilot is a simple, Golang-based CLI tool for interacting with a variety of AI models. It automatically maintains chat history for continuous and context-aware conversations, streamlining interactions without any extra setup. Users can easily choose from text, streaming, or structured input options. GoPilot supports a wide range of AI providers, allowing for seamless integration and quick results across different models, all with minimal configuration. Its simplicity and flexibility make it a perfect choice for developers seeking efficient AI integration directly from the terminal.

## Installing

### Quick Install (Linux/macOS)

To install the CLI, run the following command to download and execute the installation script:

```bash
curl -sSL https://github.com/zacharyiles/gopilot/raw/main/install.sh | bash
```

This script will:
- Detect your OS and architecture
- Download the appropriate binary from the latest release
- Verify the SHA256 checksum
- Install it to `/usr/local/bin`

## Usage

Interact with AI models via the CLI using a prompt. Chat history is maintained by default and stored on disk. Structured prompts or text-based files are automatically detected and handled accordingly.

To run a basic prompt:

```bash
gopilot "Write a short poem about AI."
```

For structured prompts:

```bash
gopilot '{"message": "Write a short poem about AI"}'
```

Using file input:

```bash
gopilot /path/to/prompt.md
```

### Options:

Flags must be passed **after** the prompt. For example, you cannot use the command like `gopilot -n "What's the weather today?"`. Instead, it should be:

```bash
gopilot "What's the weather today?" -n
```

Here are the available flags:

- `--stream` or `-s`: Stream the response compatible with Unix pipelines.
- `--new` or `-n`: Erases previous chat history and starts a new one for this message.
- `--one-shot` or `-o`: Disables chat history for this message. History is not passed as context to the LLM and is not saved.
- `--with-context` or `-w`: Pass a string or one or more paths to text-based files (comma-separated). The contents will be extracted and used as additional context for the model. This is useful for tasks like analyzing or making changes to code files.
- `--config` or `-c`: Specify a configuration file path. This overrides other configuration methods.
- `--version` or `-v`: Show version information

### Output:

The tool will return the response from the model.

Example:

```bash
$ gopilot "What's the weather today?"
```

Response:

```
The weather today is sunny with a high of 25 degrees.
```

For structured response:

```bash
$ gopilot '{"message": "What's the weather today?"}'
```

Response:

```json
{
  "message": "What's the weather today?",
  "response": {
    "forecast": "sunny",
    "temperature": "25 degrees"
  }
}
```

## Configuration

GoPilot uses configuration values to determine how it interacts with AI providers and their models. These configuration values can be set in several ways, allowing flexibility for different workflows and use cases.

### Configuration Options

The following configuration options are supported:

- **`PROVIDER`**: Specifies the AI provider to use. Supported values include:
  - `openai`
  - `cohere-ai`
  - `@huggingface/inference`
  - `langchain`
  - `openrouter`
  - `anthropic`
  
- **`API_KEY`**: The API key for the selected provider. You must supply this for authentication when using the providers.
  
- **`MODEL`**: The model to use with the selected provider. Refer to the provider's documentation for available models. For example, OpenAI's `text-davinci-003`.

### Setting Configuration Values

Configuration values can be set in one of the following ways:

#### 1. **Environment Variables**
   - Define configuration values directly in your shell or in a `.env` file located in the same directory where GoPilot is being executed. Values from the `.env` file are loaded automatically.

#### 2. **Configuration Files**
   - Create a JSON or YAML file with the following names: `gopilot.config.json` or `gopilot.config.yml`. Place the configuration file in the directory from which you run the GoPilot CLI, as this is where GoPilot will look for the configuration. Example of a JSON configuration file:

   ```json
   {
     "PROVIDER": "openai",
     "API_KEY": "your-api-key",
     "MODEL": "text-davinci-003"
   }
   ```

#### 3. **Command-Line Arguments**
   - You can specify a configuration file path directly using the `--config` or `-c` flag when running the CLI. This method takes precedence over the environment variables and configuration files.

   Example usage:

   ```bash
   gopilot --config /path/to/gopilot.config.json
   ```

Each configuration method is optional, and values provided through later methods will override any earlier settings.

## Supported Models

- [openai](https://platform.openai.com/docs)
- [langchain](https://docs.langchain.com/)
- [openai-api](https://github.com/KartikTalwar/openai-api)
- [@huggingface/inference](https://huggingface.co/docs/api-inference/)
- [anthropic](https://console.anthropic.com/docs)
- [cohere-ai](https://docs.cohere.ai/)
- [openrouter](https://github.com/openrouter-ai/openrouter)

## Development

### Release Process

GoPilot uses GitHub Actions to automatically build and release binaries for multiple platforms when a new tag is pushed. To create a new release:

1. Tag the commit: `git tag -a v1.0.0 -m "Release v1.0.0"`
2. Push the tag: `git push origin v1.0.0`

The GitHub Action will automatically:
- Build binaries for Windows, macOS, and Linux (both amd64 and arm64)
- Generate SHA256 checksums for each binary
- Create a GitHub release with all artifacts
- Generate release notes

### Open Questions

1. How should complex JSON inputs (e.g., examples for classification) be handled in a CLI-friendly way?
2. Should the CLI support multiple configuration files for different providers or workflows?
3. How will the CLI gracefully handle rate-limiting or API errors across different providers?

## Building from Source

To build GoPilot from source:

```bash
git clone https://github.com/zacharyiles/gopilot
cd gopilot
go build -o gopilot ./cmd/main.go
```

For a release build with optimizations:
```bash
go build -trimpath -ldflags="-s -w" -o gopilot ./cmd/main.go
```

## Manual Installation

1. Visit the [releases page](https://github.com/zacharyiles/gopilot/releases)
2. Download the appropriate binary for your system:
   - `gopilot-darwin-amd64` for Intel macOS
   - `gopilot-darwin-arm64` for Apple Silicon macOS
   - `gopilot-linux-amd64` for 64-bit Linux
   - `gopilot-linux-arm64` for ARM64 Linux
   - `gopilot-windows-amd64.exe` for 64-bit Windows
   - `gopilot-windows-arm64.exe` for ARM64 Windows
3. Verify the checksum (recommended):
   ```bash
   sha256sum -c gopilot-<os>-<arch>.sha256
   ```
4. Make the binary executable (Linux/macOS):
   ```bash
   chmod +x gopilot-<os>-<arch>
   ```
5. Move to a location in your PATH (Linux/macOS):
   ```bash
   sudo mv gopilot-<os>-<arch> /usr/local/bin/gopilot
   ```
## Contribution

Contributions are welcome. Please open an issue or pull request for any feature requests, bug fixes, or improvements.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.