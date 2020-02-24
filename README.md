# HelloGo
test repository for golang

## Prerequisites
```
brew install fswatch <- lib for hot-reload
```

## Quick start
```
cd /path/to/HelloGo
go build main.go
./main
```

or, just run:
```
go run main.go
```
Access http://localhost:8000/test/hoge and see if it works!

## With docker
Without WIFI, this may slow down to install the necessary modules.
```
cd docker
// production environment
docker-compose up -d prod
// dev environment
docker-compose up -d dev
```

## Authorization
```
curl -X POST -d 'username=foo' -d 'password=bar' http://localhost:8000/login
curl -X GET http://localhost:8000/me -H "Authorization: Bearer ey..."
```

## Reference
Go Official documents
https://astaxie.gitbooks.io/build-web-application-with-golang/ja/03.2.html

Go Echo
https://github.com/labstack/echo
