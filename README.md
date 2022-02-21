# Go API Seed

This is a Go lang startup project for API's using gin framework

There's an endpoint /health defined, commonly used to check the status of the service

The service is configured to run at PORT 8002

## Local
```
go ./cmd/main.go
```

## Docker
```
docker build -t ${your_tag} .
docker run -p 8002:8002 --name ${your_docker_name} -d ${image_tag}
```