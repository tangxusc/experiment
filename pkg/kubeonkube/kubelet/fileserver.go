package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)
//CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build fileserver.go
func main() {
	relativePath := os.Getenv("relativePath")
	if len(relativePath) <= 0 {
		relativePath = "/"
	}
	filepath := os.Getenv("filepath")
	if len(filepath) <= 0 {
		filepath = "/home/"
	}
	fileport := os.Getenv("fileport")
	if len(fileport) <= 0 {
		fileport = "8080"
	}
	fmt.Printf("relativePath:%s,filepath:%s,fileport:%v \n", relativePath, filepath, fileport)
	engine := gin.Default()
	engine.StaticFS(relativePath, gin.Dir(filepath, true))
	engine.Run(fmt.Sprintf(":%s", fileport))
}
