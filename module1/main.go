package main

import (
	"log"

	"github.com/wgsaxton/go-grpc-class/module1/proto"
)

func main() {
	person := proto.Person{
		Name: "Garrett Saxton",
	}

	log.Println(person.GetName())
}