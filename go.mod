module github.com/buzzology/go-microservices-tutorial/shippy-service-vessel/v0.2

go 1.16

replace google.golang.org/grpc v1.38.0 => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.5.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/pkg/errors v0.9.1
	go.mongodb.org/mongo-driver v1.5.2
	google.golang.org/protobuf v1.26.0
)
