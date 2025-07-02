# ASCII Art Web Generator

A simple web application written in Go that generates professional ASCII art from user input using customizable banner styles. The project is containerized with Docker following best practices, allowing seamless build and deployment.

## Features

- Generate ASCII art from any text input
- Support for three banner styles: **standard**, **shadow**, **thinkertoy**
- Robust input validation (ASCII characters only)
- Custom error pages for 404, 400, 500 statuses
- Static assets (CSS, icons) served via `/static`
- Containerized with a multi-stage Dockerfile for a minimal final image
- Helper scripts to automate Docker image build, container run, and cleanup

## Prerequisites

- Go 1.24+ installed locally if you want to build outside Docker
- Docker 20.10+ installed and running
- Bash (for helper scripts)

## Usage

You can build and run the project **either** by using the provided helper shell scripts **or** by issuing Docker commands manually.

### Option 1: Using helper scripts

```bash
# Build the Docker image (cleanup old containers/images, then build)
bash build.sh

# Run the container on port 7777
bash run.sh

# Stop and remove the container
bash stop.sh

# Cleanup image and prune unused resources
bash cleanup.sh
```

### Option 2: Manually via Docker CLI

1. **Build the Docker image**

   ```bash
   docker image build -f Dockerfile -t ascii-art-web .
   ```

2. **Run the container** (detached, port 7777)

   ```bash
   docker container run -d -p 7777:7777 --name ascii-art ascii-art-web
   ```

3. **Verify** the container is running:

   ```bash
   docker ps -a
   ```

4. **Inspect inside the container** (optional):

   ```bash
   docker exec -it ascii-art /bin/bash
   ls -l /app
   ```

## Dockerfile Highlights

- **Multi-stage build**: using `golang:1.24.2` for compile, `alpine:latest` for runtime
- **Labels (metadata)**:
  ```dockerfile
  LABEL maintainer="Baktiyar <bzhaksyb>, Bek <bnurekey>, Dana <disabaev>"
  LABEL version="1.0"
  LABEL description="ASCII Art Generator web application in Go"
  ```
- **Minimal final image**: only the compiled binary and required assets

## Helper Scripts

- `build.sh`: Build image, stop and remove old container, remove old image
- `run.sh`: Run the latest image in a new container
- `stop.sh`: Stop and remove the running container
- `cleanup.sh`: Remove the Docker image and prune unused resources

## Good Practices

- Only standard library packages are used in Go
- Input validation to ensure safe ASCII content
- Clear project structure separating handlers, helpers, and templates
- Comprehensive error handling and user-friendly error pages

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.

