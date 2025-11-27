until [ "$(docker inspect -f '{{.State.Status}}' app)" == "running" ]; do
  echo "Waiting for container to be running..."
  sleep 0.1
done
echo "Container is running."
