### prerequisites
- Go >=1.16
- Git
- Protobuf compiler
  

### Installation
    
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    go mod tidy
    export PATH="$PATH:$(go env GOPATH)/bin"

### Genereate the proto file
```protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.```
### Run
```go run .```

