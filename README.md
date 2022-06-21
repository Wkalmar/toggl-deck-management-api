## Building the application
Running the code is as simple as
```
go run main.go
```

### Using Docker
Alternatively if you want to leverage pre-installed environment you can take advantage of supplied Dockerfile. In such a case you should perform
```
docker build -t <image-name> .
```
and then
```
docker run -it --rm -p 8080:8080 <image-name>
```

## Geging to know the application
The software leverages OpenAPI so you can access the documentation at http://localhost:8080/swagger/index.html