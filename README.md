# Protobuf Example

## Install Dependencies

```bash
go get github.com/golang/protobuf
go get github.com/golang/protobuf/proto
```

## Create A .proto file

```proto
syntax="proto3";

package main;

message Person {
      string name = 1;
      int32 age = 2;
}

option go_package = ".;main";
```


## Complile It using protoc

```bash
protoc --go_out=. *.proto
```
If protoc command fails, install it using homebrew
```bash
brew install protobuf
```

[Install Protobuf](http://google.github.io/proto-lens/installing-protoc.html)

Also ensure you install protoc-go-gen
```bash
go get github.com/golang/protobuf/protoc-gen-go
```

## Run The Program


```go
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
```


Run the program using command below

```bash
go run main.go person.pb.go
```
The output will be as below

```bash
Hello World
[10 5 86 105 106 97 121 16 26]
```

### Unmarshal

```go 
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

	newVijay := &Person{}

	err  = proto.Unmarshal(data, newVijay)

	if err != nil{
		log.Fatal("Unmarshal error : ", err)
	}

	fmt.Println(newVijay.GetName())
	fmt.Println(newVijay.GetAge())
}
```
Output of Unmarshal
```bash
Hello World
[10 5 86 105 106 97 121 16 26]
Vijay
26
```