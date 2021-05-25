package main

import (
	"context"
	"fmt"
	"log"
	"os"
	pb "github.com/buzzology/go-microservices-tutorial/shippy-service-vessel/proto/vessel"
	micro "github.com/micro/go-micro/v2"
)

const defaultHost = "datastore:27017"

func main() {

	// The name must match the proto package name
	service := micro.NewService(micro.Name("shippy.service.vessel"))
	service.Init()

	uri := os.GetEnv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	// Create a mongo db connection
	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}

	defer client.Disconnect(context.Background()) // Defer means it won't run until this function returns

	// Prepare repository and handlers
	vesselCollection := client.Database("shippy").Collection("vessels")
	repository := &MongoRepository{vesselCollection}
	h := &handler{repository}

	// Register handlers
	if err := pb.RegisterVesselServiceHandler(service.Server(), h); err != nil {
		log.Panic(err)
	}

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

	// Run the server
	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
