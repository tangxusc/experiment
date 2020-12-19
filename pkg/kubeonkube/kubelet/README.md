## go build
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build fileserver.go
## node install
curl -o install.sh http://fileserver:8080/install.sh && sh http://fileserver:8080