package main

import (
	"fmt"
	"log"
	"github.com/golang/protobuf/proto"
)

func main(){
	fmt.Println("Hello World")

	vijay := &Person{
		Name: "Vijay",
		Age: 26,
	}

	data, err := proto.Marshal(vijay)

	if(err != nil){
		log.Fatal("Marshal error : ", err)
	}

	fmt.Println(data)
}