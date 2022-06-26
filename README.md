## Using in your own build pipelines

1. Ensure `docker-compose` is installed on your build system. For details on how to do this, see: https://docs.docker.com/compose/install/

2. Use our `Dockerfile` and `docker-compose.yml` files as defaults

3. Run `docker-compose up -d --build` to run this service and the dependencies, like Database and Redis
