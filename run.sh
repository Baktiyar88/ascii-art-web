#!/bin/bash

# Set image name
IMAGE_NAME="ascii-art-web-dockerize"

echo "🔨 Building Docker image..."
docker build -t $IMAGE_NAME .

echo "🚀 Running Docker container..."
docker run -p 7777:7777 --rm $IMAGE_NAME
