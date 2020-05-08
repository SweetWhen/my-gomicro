cd Services/protos
protoc --micro_out=../ --go_out=../ test.proto
protoc --micro_out=../ --go_out=../ models.proto
protoc --micro_out=../ --go_out=../ prodService.proto
protoc --micro_out=../ --go_out=../ UserService.proto
protoc-go-inject-tag -input=../models.pb.go
protoc-go-inject-tag -input=../prodService.pb.go
protoc-go-inject-tag -input=../UserService.pb.go
cd .. && cd ..+