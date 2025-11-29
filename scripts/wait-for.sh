until [ "$(docker inspect -f '{{.State.Status}}' $1)" == "running" ]; do
  echo "Waiting for container to be running..."
  sleep 0.1
done
echo "Container is running."
