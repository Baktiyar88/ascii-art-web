# Set image, container and port
IMAGE_NAME="ascii-art-web-dockerize"
CONTAINER_NAME="ascii-art-web"

# Stop and remove old container if it exists
if docker ps -a --format '{{.Names}}' | grep -Eq "^$CONTAINER_NAME\$"; then
  echo "🛑 Stopping old container..."
  docker container stop $CONTAINER_NAME
  echo "🗑️ Removing old container..."
  docker container rm $CONTAINER_NAME
fi

echo "🗑️ Removing image..."
docker image rm $IMAGE_NAME