# Checking docker images
echo "List of images..."
docker images

# New line
echo ""

# Checking docker container
echo "List of containers..."
docker ps -a

# Checking docker file system
# It doesn't work and must be fixed
# docker exec -it "ascii-art-web" /bin/bash