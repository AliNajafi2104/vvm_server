module github.com/AliNajafi2104/vvm_server/gateway

go 1.24.2

replace github.com/AliNajafi2104/vvm_server/common => ../common

require (
	github.com/AliNajafi2104/vvm_server/common v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.1
	google.golang.org/grpc v1.72.0
)

require (
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)
