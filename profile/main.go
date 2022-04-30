package main

import (
	"XWS-Project/profile/handlers"
	"XWS-Project/profile/repository"
	"XWS-Project/profile/service"
	"XWS-Project/utilities"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initDatabase() *mongo.Client {
	//mongoUri := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI("mongodb://" + "localhost:27017" + "/?connect=direct")

	// Connect to MongoDB
	//ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	//testData(client) //treba mi samo prvi put
	fmt.Println("Connected to MongoDB!")

	return client
}

func initRepo(clientDB *mongo.Client) *repository.ProfileRepository {
	return &repository.ProfileRepository{Client: clientDB}
}

func initService(repo *repository.ProfileRepository) *service.ProfileService {
	return &service.ProfileService{ProfileRepository: repo}
}

func initHandler(service *service.ProfileService) *handlers.ProfileHandler {
	return &handlers.ProfileHandler{ProfileService: service}
}

func handlerFunc(handler *handlers.ProfileHandler) {
	fmt.Println("Profile server started...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{username}", handler.Edit).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", router))
}
func main() {
	utilities.TracerInit("profile")
	dbClient := initDatabase()
	regRepo := initRepo(dbClient)
	regService := initService(regRepo)
	regHandler := initHandler(regService)
	handlerFunc(regHandler)
}
