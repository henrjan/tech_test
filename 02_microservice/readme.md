# Microservice for search movie

```
update : so I decided to add gRPC and gRPC-gateway as v2, by default when running docker-compose v1 will be used.
To use v2, you need to change "v1" on dockerfile to "v2" and bind ports :8081 on docker host to port :8081 on docker container.
gRPC server use port :8080 and gRPC-gateway use port 8081.
```

Before running the appliaction, you need to change configuration for Mysql at "docker-compose.yaml" file,
for Mysql user, password, address, and port.

You also need to configure ports application will listen to, defaults to :8080.

Create empty Mysql database using "db.sql"

To run docker compose, use this command : 

    "docker-compose up --build -d"

Default port used on docker-compose file is :8080 and container internal port is :8080.

To stop the docker container use following command :

    "docker-compose down"

To access search movie API, use `/v1/movie` endpoints with url query parameters `search_word` and `page`.