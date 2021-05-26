module github.com/buzzology/shippy-service-vessel

go 1.16

replace google.golang.org/grpc v1.38.0 => google.golang.org/grpc v1.26.0

// replace github.com/buzzology/shippy-service-vessel => ../shippy-service-vessel

require (
	github.com/golang/protobuf v1.5.2
	github.com/micro/go-micro/v2 v2.9.1
	go.mongodb.org/mongo-driver v1.5.2
	google.golang.org/protobuf v1.26.0
)
