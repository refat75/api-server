# api-server-docker


## Build the Docker Image
```shell
  docker build -t api-server-docker -f Dockerfile .
```

## Run the Container
```shell
    docker run -d -p 8080:8080 api-server-docker
```

## Check the Running Container
```shell
    docker ps
```

## Stop the Container
```shell
    docker stop <container_id>
```

## Remove the Container
```shell
    docker rm <container_id>
```

## Login Credentials
{
"username": "admin",
"password": "123456"
}