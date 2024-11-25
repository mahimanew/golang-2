
```
mkdir go-docker-app
cd go-docker-app
go mod init go-docker-app
go run main.go

http://localhost:9090

docker build -t go-docker-app .
docker run -p 9090:9090 go-docker-app

docker login
docker tag go-docker-app mahi2029/golang-repo:go-docker-app

docker push mahi2029/golang-repo:go-docker-app
```
