module sample.com/server

go 1.16

replace sample.com/pb => ../pb

require (
	github.com/golang/protobuf v1.5.1 // indirect
	google.golang.org/grpc v1.36.1
	sample.com/pb v0.0.0-00010101000000-000000000000
)
