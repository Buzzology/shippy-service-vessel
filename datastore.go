package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"fmt"
)

func CreateClient(ctx context.Context, uri string, retry int32) (*mongo.Client, error) {

	appliedOptions := options.Client().ApplyURI(uri)
	conn, err := mongo.Connect(ctx, appliedOptions)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx, nil); err != nil {

		fmt.Println(err)

		if retry >= 3 {
			return nil, err
		}

		retry += 1
		time.Sleep(time.Second * 2)
		return CreateClient(ctx, uri, retry)
	}

	return conn, err
}