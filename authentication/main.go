package main

import (
	"XWS-Project/authentication/handlers"
	"XWS-Project/authentication/model"
	"XWS-Project/authentication/repository"
	"XWS-Project/authentication/service"
	"XWS-Project/utilities"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	//"time"
	//"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func addDataToDB(client *mongo.Client) {
	user1 := model.User{1111, "admin@dlk.com", "admin", "admin", true}
	user2 := model.User{1112, "user@dlk.com", "user", "user", true}

	collection := client.Database("dislinkt").Collection("users")
	insertRes, err := collection.InsertOne(context.TODO(), user1)
	insertRes1, err := collection.InsertOne(context.TODO(), user2)
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
	/*err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}*/
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
	fmt.Println("Auth server started...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/login", handler.Preflight).Methods("OPTIONS")
	router.HandleFunc("/login", handler.Login).Methods("POST")
	router.HandleFunc("/test", handler.Test).Methods("GET", "OPTIONS")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	utilities.TracerInit("authentication")
	dbClient := initDB()
	autRepo := initRepo(dbClient)
	autService := initService(autRepo)
	autHandler := initHandler(autService)
	handlerFunc(autHandler)
}
