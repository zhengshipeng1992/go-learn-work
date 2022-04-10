package main

// Injectors from wire.go:

//func InitProtoServe(collection *mongo.Collection) *service.Hello {
//	hello := data.NewMongoRepo(collection)
//	helloRequest := biz.NewHelloRequest(hello)
//	serviceHello := service.NewHello(helloRequest)
//	return serviceHello
//}
//
//func InitClient(ctx context.Context, uri string) (*mongo.Client, error) {
//	config := mongodb.NewConfig(ctx, uri)
//	client, err := mongodb.CreateClient(config)
//	if err != nil {
//		return nil, err
//	}
//	return client, nil
//}
