package main

import (
	"context"
	"github.com/google/wire"
	"go-learn-work/work03/internel/biz"
	"go-learn-work/work03/internel/data"
	"go-learn-work/work03/internel/pkg/database/mongodb"
	"go-learn-work/work03/internel/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitProtoServe(collection *mongo.Collection) *service.Hello {
	wire.Build(data.NewMongoRepo, biz.NewHelloRequest, service.NewHello)
	return &service.Hello{}
}

func InitClient(ctx context.Context, uri string) (*mongo.Client, error) {
	wire.Build(mongodb.NewConfig, mongodb.CreateClient)
	return nil, nil
}
