#!/bin/bash

# Set image, container and port
IMAGE_NAME="ascii-art-web-dockerize"
CONTAINER_NAME="ascii-art-web"
PORT=7777:7777

echo "🔨 Building Docker image..."
docker image build -f Dockerfile -t $IMAGE_NAME .

# Stop and remove old container if it exists
if docker ps -a --format '{{.Names}}' | grep -Eq "^$CONTAINER_NAME\$"; then
  echo "🛑 Stopping old container..."
  docker container stop $CONTAINER_NAME
  echo "🗑️ Removing old container..."
  docker container rm $CONTAINER_NAME
fi

# Build or rebuild container
echo "🚀 Running Docker container..."
docker container run -p $PORT --detach --name $CONTAINER_NAME $IMAGE_NAME

echo "🌐 App is now running at http://localhost:7777"