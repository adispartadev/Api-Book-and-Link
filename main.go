package main

import (
	"api.go/bootstrap"
	"fmt"
)

func init() {
	fmt.Println("Starting APP")
	bootstrap.BootApplication()
}

func main() {
	fmt.Println("2023-09-08")
}
