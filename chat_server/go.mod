module github.com/Yashikab/chat_sample/chat_server

go 1.20

replace github.com/Yashikab/chat_sample/chat_protobuf => ../chat_protobuf

require github.com/Yashikab/chat_sample/chat_protobuf v0.0.0-00010101000000-000000000000

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230711160842-782d3b101e98 // indirect
	google.golang.org/grpc v1.58.2 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)
