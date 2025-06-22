#!/bin/bash

#Stop the container by name or ID
docker ps
echo "⛔ Stopping container..."
docker stop $(docker ps -q --filter ancestor=ascii-art-web-dockerize)