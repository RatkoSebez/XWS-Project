package main

import (
	"XWS-Project/authentication/handlers"
	"XWS-Project/authentication/model"
	"XWS-Project/authentication/repository"
	"XWS-Project/authentication/service"
	pb "XWS-Project/proto/login_service"
	"XWS-Project/utilities"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	//"github.com/gorilla/mux"
	//"net/http"

	//"time"
	//"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func addDataToDB(client *mongo.Client) {
	user1 := model.User{"Dusan", "Markovic", "dusan@gmail.com", "1212", "male", 23, 10, 1999, "dulecar", "dulecar123", "dobar", []string{"iskustvo1", "iskustvo2"}, []string{"scepa", "scepa2"}, []string{"interesovanje1", "interesovanje12"}, []string{"vjestina1", "vjestina12"}, true}
	user2 := model.User{"Bojan", "Prodanovic", "bojan@gmail.com", "121232", "male", 21, 01, 1999, "bp21", "123", "biografija", []string{"iskustvo1", "iskustvo2"}, []string{"scepa", "scepa2"}, []string{"interesovanje11", "interesovanje112"}, []string{"vjestina123", "vjestina1s2"}, true}

	collection := client.Database("dislinkt").Collection("users")
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
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", "8081"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthenticationServer(grpcServer, handler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	/*fmt.Println("Auth server started...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/login", handler.Preflight).Methods("OPTIONS")
	router.HandleFunc("/login", handler.Login).Methods("POST")
	router.HandleFunc("/test", handler.Test).Methods("GET", "OPTIONS")
	log.Fatal(http.ListenAndServe(":8081", router))*/
}

func main() {
	utilities.TracerInit("authentication")
	dbClient := initDB()
	autRepo := initRepo(dbClient)
	autService := initService(autRepo)
	autHandler := initHandler(autService)
	handlerFunc(autHandler)
}
