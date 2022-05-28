package main

import (
	"XWS-Project/job_offer/handlers"
	"XWS-Project/job_offer/repository"
	"XWS-Project/job_offer/service"
	pb "XWS-Project/proto/job_offer_service"
	"XWS-Project/utilities"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
)

func initDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://" + "localhost:27017" + "/?connect=direct")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	return client
}

func initRepo(job_offerDB *mongo.Client) *repository.JobRepository {
	return &repository.JobRepository{Client: job_offerDB}
}

func initService(repo *repository.JobRepository) *service.JobService {
	return &service.JobService{JobRepo: repo}
}

func initHandler(service *service.JobService) *handlers.JobHandler {
	return &handlers.JobHandler{JobService: service}
}

func handlerFunc(handler *handlers.JobHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", "8086"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterJobOfferServiceServer(grpcServer, handler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func main() {
	utilities.TracerInit("job_offer")
	dbClient := initDB()
	jobRepo := initRepo(dbClient)
	jobService := initService(jobRepo)
	jobHandler := initHandler(jobService)
	handlerFunc(jobHandler)
}
