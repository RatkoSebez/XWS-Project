package main

import (
	"XWS-Project/agent_backend/handlers"
	"XWS-Project/agent_backend/model"
	"XWS-Project/agent_backend/repository"
	"XWS-Project/agent_backend/service"
	"XWS-Project/utilities"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	//"github.com/gorilla/mux"
	//"net/http"

	//"time"
	//"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func addDataToDB(client *mongo.Client) {
	user1 := model.User{"Dusan", "Markovic", "dusan@gmail.com", "1212", "dulecar", "dulecar123"}
	user2 := model.User{"Bojan", "Prodanovic", "bojan@gmail.com", "121232", "bp21", "123"}

	collection := client.Database("agent").Collection("users")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"email": "dusan@gmail.com"})
	if err != nil {
		panic(err)
	}
	_, err = collection.DeleteOne(context.TODO(), bson.M{"email": "bojan@gmail.com"})
	if err != nil {
		panic(err)
	}
	insertRes, err := collection.InsertOne(context.TODO(), user1)
	if err != nil {
		panic(err)
	}
	insertRes1, err := collection.InsertOne(context.TODO(), user2)
	fmt.Println(insertRes)
	fmt.Println(insertRes1)
	if err != nil {
		panic(err)
	}
}

func initDB() *mongo.Client {
	//mongoUri := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI("mongodb://" + "localhost:27017" + "/?connect=direct")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	addDataToDB(client)
	fmt.Println("Connected to MongoDB!")

	return client
}

func initRepo(clientDB *mongo.Client) *repository.AuthenticationRepository {
	return &repository.AuthenticationRepository{Client: clientDB}
}

func initService(repo *repository.AuthenticationRepository) *service.AuthenticationService {
	return &service.AuthenticationService{AuthRepo: repo}
}

func initHandler(service *service.AuthenticationService) *handlers.AuthenticationHandler {
	return &handlers.AuthenticationHandler{AuthenticationService: service}
}

func handlerFunc(handler *handlers.AuthenticationHandler) {
	//listener, err := net.Listen("tcp", fmt.Sprintf(":%s", "8081"))
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//grpcServer := grpc.NewServer()
	//pb.RegisterAuthenticationServer(grpcServer, handler)
	//if err := grpcServer.Serve(listener); err != nil {
	//	log.Fatalf("failed to serve: %s", err)
	//}
	fmt.Println("Agent application started...")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/test", handler.Preflight).Methods("OPTIONS")
	router.HandleFunc("/api/test", handler.Test).Methods("GET")
	router.HandleFunc("/api/login", handler.Preflight).Methods("OPTIONS")
	router.HandleFunc("/api/login", handler.Login).Methods("POST")
	router.HandleFunc("/api/register", handler.Register).Methods("POST")

	log.Fatal(http.ListenAndServe(":8090", router))
}

func main() {
	utilities.TracerInit("authentication")
	dbClient := initDB()
	autRepo := initRepo(dbClient)
	autService := initService(autRepo)
	autHandler := initHandler(autService)
	handlerFunc(autHandler)
}
