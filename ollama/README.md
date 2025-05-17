# Ollama Setup on Local PC

This Docker file can be used to set up Ollama inside a container, run a WebUI, and host a self-hosted LLM.

## Docker Preparation

To ensure proper GPU support, you **must** install the correct Nvidia drivers.

1. **Download the Latest Nvidia Docker Toolkit:** Download the latest Nvidia docker toolkit from the official Nvidia [website](https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/latest/install-guide.html)
2. **Install the Driver:** Follow the installation instructions provided by Nvidia.
3. **Verify the Installation:** After the driver is installed, run the following command to check the driver version:
```
docker run --rm nvidia/cuda:12.0-base-ubuntu22.04 nvidia-smi
```
(This command should display information about your GPU).

## Ollama Preparation

Download Ollama from the official GitHub: [https://github.com/ollama/ollama](https://github.com/ollama/ollama)
Checkout the latest version.
Copy `docker-compose.yml` in the project root. Make sure the `docker-compose.yml` file is located in the project root directory.

## Download the Model

To download the model, run:
```
docker compose up --build -d
```
And then, checkout inside the ollama container:
```
docker exec -it ollama bash
```
On the official Ollama website, find the LLM you need and run:
```
ollama run gemma3:1b
```
After downloading and unpacking, the window could be closed.

## What to do next. VSCode Integration

Install the `Continue` extension

After: Add Chat model -> Provider: Ollama -> Model: Autodetect.

To enable the VSCode extension to discover the Ollama instance, ensure Ollama is running on `0.0.0.0:11434`. This allows the VSCode extension to connect to the LLM.

The typical config could be like this:
```yml
  - name: gemma3 12b
    provider: ollama
    model: gemma3:12b
    roles:
      - chat
      - edit
      - apply
      - autocomplete
```

## Important Notes
*   **Resource Requirements:** LLMs require significant computational resources, maybe for beginning `gemma3:4b` could be enough.
*   **Hallucinations:** LLMs can generate incorrect or misleading information (hallucinations). Verify any information provided by the model.

## Afterwords 
It's still raw. The Nvidia driver is always disabling or crashing, so we have to restart container or computer each time. Additionally, it's better to check how much each LLM consumes to fit it on your PC. 
