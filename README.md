# GoPilot: A Unified No-Frills LLM CLI in Go

## Overview

GoPilot is a simple, Golang-based CLI tool for interacting with a variety of AI models. It automatically maintains chat history for continuous and context-aware conversations, streamlining interactions without any extra setup. Users can easily choose from text, streaming, or structured input options. GoPilot supports a wide range of AI providers, allowing for seamless integration and quick results across different models, all with minimal configuration. Its simplicity and flexibility make it a perfect choice for developers seeking efficient AI integration directly from the terminal.

### Installing

To install the CLI, run the following command to download and execute the installation script:

curl \-sSL https://github.com//install.sh | bash

This script will download the latest version of the CLI and place it in your system's PATH.

---

## Commands

### Chatting

Interact with AI models via the CLI using a prompt. Chat history is maintained by default and stored on disk. Structured prompts or text-based files are automatically detected and handled accordingly.

#### Usage:

To run a basic prompt:

./gopilot "Write a short poem about AI."

For structured prompts:

./gopilot '{"message": "Write a short poem about AI"}'

Using file input:

./gopilot /path/to/prompt.md

#### Options:

Flags must be passed **after** the prompt. For example, you cannot use the command like `./gopilot -n "What's the weather today?"`. Instead, it should be:

./gopilot "What's the weather today?" \-n

Here are the available flags:

- `--stream` or `-s`: Stream the response compatible with Unix pipelines.  
- `--new` or `-n`: Erases previous chat history and starts a new one for this message.  
- `--one-shot` or `-o`: Disables chat history for this message. History is not passed as context to the LLM and is not saved.  
- `--with-context` or `-w`: Pass a string or one or more paths to text-based files (comma-separated). The contents will be extracted and used as additional context for the model. This is useful for tasks like analyzing or making changes to code files.  
- `--config` or `-c`: Specify a configuration file path. This overrides other configuration methods.

#### Output:

The tool will return the response from the model.

Example:

$ ./gopilot "What's the weather today?"

Response:

The weather today is sunny with a high of 25 degrees.

For structured response:

$ ./gopilot '{"message": "What's the weather today?"}'

Response:

{

  "message": "What's the weather today?",

  "response": {

    "forecast": "sunny",

    "temperature": "25 degrees"

  }

}

---

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

#### 1\. **Environment Variables**

- Define configuration values directly in your shell or in a `.env` file located in the same directory where GoPilot is being executed. Values from the `.env` file are loaded automatically.

#### 2\. **Configuration Files**

- Create a JSON or YAML file with the following names: `gopilot.config.json` or `gopilot.config.yml`. Place the configuration file in the directory from which you run the GoPilot CLI, as this is where GoPilot will look for the configuration. Example of a JSON configuration file:

{

  "PROVIDER": "openai",

  "API\_KEY": "your-api-key",

  "MODEL": "text-davinci-003"

}

#### 3\. **Command-Line Arguments**

- You can specify a configuration file path directly using the `--config` or `-c` flag when running the CLI. This method takes precedence over the environment variables and configuration files.

Example usage:

./gopilot \--config /path/to/gopilot.config.json

Each configuration method is optional, and values provided through later methods will override any earlier settings.

---

## Supported Models

- [openai](https://platform.openai.com/docs)  
- [langchain](https://docs.langchain.com/)  
- [openai-api](https://github.com/KartikTalwar/openai-api)  
- [@huggingface/inference](https://huggingface.co/docs/api-inference/)  
- [anthropic](https://console.anthropic.com/docs)  
- [cohere-ai](https://docs.cohere.ai/)  
- [openrouter](https://github.com/openrouter-ai/openrouter)

---

## TODO

1. How should complex JSON inputs (e.g., examples for classification) be handled in a CLI-friendly way?  
2. Should the CLI support multiple configuration files for different providers or workflows?  
3. How will the CLI gracefully handle rate-limiting or API errors across different providers?

---

## Contribution

Contributions are welcome. Please open an issue or pull request for any feature requests, bug fixes, or improvements.

---
