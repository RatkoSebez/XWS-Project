package main

import (
	"XWS-Project/follow/handlers"
	"XWS-Project/follow/model"
	"XWS-Project/follow/repository"
	"XWS-Project/follow/service"
	pb "XWS-Project/proto/follow_service"
	"XWS-Project/utilities"
	"context"
	"fmt"
	//"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
	//"net/http"
)

func addDataToDB(client *mongo.Client) {
	follows := []string{"asd@gmail.com", "dfas@gmail.com"}
	followers := []string{"test@gmail.com"}
	userFollows1 := model.Follow{"admin@dlk.com", follows, followers}
	userFollows2 := model.Follow{"user@dlk.com", follows, followers}

	collection := client.Database("dislinkt").Collection("userFollows")
	insertRes, err := collection.InsertOne(context.TODO(), userFollows1)
	insertRes1, err := collection.InsertOne(context.TODO(), userFollows2)
	fmt.Println(insertRes)
	fmt.Println(insertRes1)
	if err != nil {
		log.Fatal(err)
	}
}

func initDB() *mongo.Client {
	//mongoUri := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI("mongodb://" + "localhost:27017" + "/?connect=direct")

	// Connect to MongoDB
	//ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	addDataToDB(client)
	fmt.Println("Connected to MongoDB!")

	return client
}

func initRepo(clientDB *mongo.Client) *repository.FollowRepository {
	return &repository.FollowRepository{Client: clientDB}
}

func initService(repo *repository.FollowRepository) *service.FollowService {
	return &service.FollowService{FollowRepo: repo}
}

func initHandler(service *service.FollowService) *handlers.FollowHandler {
	return &handlers.FollowHandler{FollowService: service}
}

func handlerFunc(handler *handlers.FollowHandler) {
	/*fmt.Println("Follow server started...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/followRequest/{email}", handler.GetFollowRequest).Methods("GET")
	router.HandleFunc("/followRequest", handler.CreateFollowRequest).Methods("POST")
	router.HandleFunc("/follow/{email}", handler.GetFollow).Methods("GET")
	router.HandleFunc("/follow/", handler.CreateFollow).Methods("POST")
	log.Fatal(http.ListenAndServe(":8085", router))*/
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", "8085"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterFollowServiceServer(grpcServer, handler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func main() {
	utilities.TracerInit("follow")
	dbClient := initDB()
	autRepo := initRepo(dbClient)
	autService := initService(autRepo)
	autHandler := initHandler(autService)
	handlerFunc(autHandler)
}
