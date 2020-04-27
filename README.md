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
```


## Complile It using protoc

```bash
protoc --go_out=. *.proto
```