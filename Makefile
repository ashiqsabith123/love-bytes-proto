genauth:
	protoc --go_out=. --go-grpc_out=. --proto_path=./auth/pb auth.proto

genmatch:
	protoc --go_out=. --go-grpc_out=. --proto_path=./match/pb match.proto
