# GoPilot: A Unified No-Frills LLM CLI in Go

GoPilot is the AI tool for Devops and Engineering teams. Seamlessly integerate with your codebases, release pipelines and team workflows to automate your typical workflows using the power of AI. GoPilot is like a DevOps engineer always on hand, being able to generate code, write documentation, configure environments, and even deploy your code. Comprised of a comprehensive suite of agentic "actions" that can be composed, extended, and triggered from anywhere, you can think of GoPilot as providing a suite of workflows like Github Actions, except with AI!

GoPilot supports a wide range of AI providers, allowing for seamless integration and quick results across different models, all with minimal configuration. Its simplicity and flexibility make it a perfect choice for developers seeking efficient AI integration without all the bells and whistles.

## Why GoPilot?

The idea of GoPilot was formed at my time in large engineering teams like at HashiCorp, where myself and teamates would continually look to automate away toil be desiging and refining highly-determinsitic prompts to solve day-to-day tasks on engineering teams. Managing and using those prompts was tedius, even though the results of years tweaking them produced great results. We would've easily adopted a product to do this, but the market consits of mostly large frameworks and tools that are either: not specific to DevOps automations, not minimal enough to fit nicely in a pipeline or deployed to the edge, and not optimized in a way that produces highly deterministic and traceable results as it required for most DevOps tasks. Some light inspiration was drawn from [Aider](https://github.com/Aider-AI/aider) which can be thought of as the human and application-development focused version of GoPilot. GoPilot was carefully designed to work outside of a human-in-the-loop, and it's speciality is infrastructure, engineering standards, and solving toil.

### Quick Install

To install the CLI, run the following command in your terminal:

```bash
curl -sSL https://github.com/zacharyiles/gopilot/raw/main/install.sh | bash
```

This script will:
- Detect your OS and architecture
- Download the appropriate binary from the latest release
- Verify the SHA256 checksum
- Install it to your system's appropriate location

### Windows Installation Note

For Windows users, the install script requires one of the following:
- Git Bash (recommended)
- WSL (Windows Subsystem for Linux)
- Cygwin
- MSYS2

To install using Git Bash (recommended):
1. Open Git Bash
2. Run the installation command:
   ```bash
   curl -sSL https://github.com/zacharyiles/gopilot/raw/main/install.sh | bash
   ```
3. If the installation directory is not in your PATH, you may need to add it manually:
   - The default installation location is `%USERPROFILE%\AppData\Local\Microsoft\WindowsApps`
   - You can add this to your PATH through System Properties > Advanced > Environment Variables

Alternatively, Windows users can perform a manual installation by downloading the appropriate .exe file from the releases page.

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
- `--action` or `-a`: Specify an action plugin to process inputs and outputs (e.g., --action=edit-code)

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

## Actions

Actions are the core of GoPilot. They represent discrete, composable, agentic workflows. Actions leverage the full GoPilot framework to accomplish a set of very specific AI-driven tasks, such as updating a README after a new commit to main, or reviewing a new PR, or generating a new client SDK after an API chanegs. Actions are purpose built to solve typical grunt work and tasks on engineering and Devops teams, but supercharge them with the power of AI. Actions are meant to be used in pipelines, automations, or as part of a local development workflow. Actions are meant to be highly deterministic and utilize structured inputs and outputs and are optimized to not require human intervention or triggering.

NOTE: This project is still in early development and many planned actions are not yet implemented.

### edit-code
The `edit-code` action optimizes the interaction for code editing tasks. It:
- Formats the AI's instructions to focus on code modifications
- Returns only the relevant changed sections of code
- Adds helpful comments explaining the changes

Example usage:
```bash
gopilot "Update the error handling in main.go" --action=edit-code --with-context=main.go
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