# Go Signup


Signup orchestrator

## Dependencies

- protoc

```bash
go get -u -v google.golang.org/protobuf/...

go get -u -v google.golang.org/grpc/...

go install github.com\golang\protobuf\protoc-gen-go

go get -u github.com/gofiber/fiber/v2

go get -u github.com/swaggo/swag/cmd/swag
# 1.16 or newer
go install github.com/swaggo/swag/cmd/swag@latest

go get -u github.com/gofiber/swagger

go get github.com/go-playground/validator/v10
```