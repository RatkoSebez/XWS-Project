package main

import (
	"XWS-Project/registration/handlers"
	"XWS-Project/registration/repository"
	"XWS-Project/registration/service"
	"XWS-Project/utilities"
	"net/http"

	"context"
	"fmt"
	"log"

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

func initRepo(clientDB *mongo.Client) *repository.RegistrationRepository {
	return &repository.RegistrationRepository{Client: clientDB}
}

func initService(repo *repository.RegistrationRepository) *service.RegistrationService {
	return &service.RegistrationService{RegistrationRepository: repo}
}

func initHandler(service *service.RegistrationService) *handlers.RegistrationHandler {
	return &handlers.RegistrationHandler{RegistrationService: service}
}

func handlerFunc(handler *handlers.RegistrationHandler) {
	fmt.Println("Registration server started...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/register", handler.Register).Methods("POST")
	router.HandleFunc("/register", handler.Preflight).Methods("OPTIONS")

	log.Fatal(http.ListenAndServe(":8081", router))
}
func main() {
	utilities.TracerInit("registration")
	dbClient := initDatabase()
	regRepo := initRepo(dbClient)
	regService := initService(regRepo)
	regHandler := initHandler(regService)
	handlerFunc(regHandler)
}
