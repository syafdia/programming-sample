# OMDB API Clean Architecture
Clean architecture implementation on Go programming language with multiple presentation layer (HTTP and gRPC).

### How to run
- HTTP Server, on terminal run `go run cmd/http/main.go`, HTTP server will be served on `0.0.0.0:8001`
- gRPC Server, on terminal run `go run cmd/grpc/main.go`, gRPC server will be served on `0.0.0.0:8002`
- HTTP client example, on terminal run `go run examples/http/main.go
- gRPC client example, on terminal run `go run examples/grpc/main.go

### License
MIT License. Copyright 2021 syafdia@gmail.com