protoc:
	protoc -I=./contracts/ \
			--go_out=plugins=grpc:. \
			./contracts/auth/*.proto \
			./contracts/common/*.proto \
			./contracts/user/*.proto

test:
	go test --cover github.com/ungame/go-signup/...