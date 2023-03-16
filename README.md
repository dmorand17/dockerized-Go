Initialize module

```bash
# Use direct proxy
go env -w GOPROXY=direct
go mod init web-app
go mod tidy
```

Run application

```bash
go run main.go
```

Test app

```bash
curl http://localhost:9000/version
```

Build application

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./server
```

Run `./server` executable

```bash
./server
```

Send a different request

```bash
curl http://localhost:9000/version
```

## Build

Build and tag the image

```bash
docker build . -t web-app:0.0.1
```

## Run

Run container and forward the traffic to `9001` on the localhost

```bash
docker run --name web-app-001 -d -p 9001:9000 web-app:0.0.1
```

Test its working

```bash
curl http://localhost:9001/version
```
