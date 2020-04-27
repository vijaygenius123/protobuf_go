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
