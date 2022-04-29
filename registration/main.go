package main

import (
	"XWS-Project/registration/handlers"
	"XWS-Project/registration/model"
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

func testData(client *mongo.Client) {
	user1 := model.User{1111, "Dusan", "Markovic", "dusan@gmail.com", "1212", "male", 23, 10, 1999, "dulecar", "dulecar123", "dobar", []string{"iskustvo1", "iskustvo2"}, []string{"scepa", "scepa2"}, true}
	//user2 := model.User{}
	collection := client.Database("admin").Collection("user")
	insertRes, err := collection.InsertOne(context.TODO(), user1)
	//insertRes1, err := collection.InsertOne(context.TODO(), user2)
	fmt.Println(insertRes)
	//fmt.Println(insertRes1)
	if err != nil {
		log.Fatal(err)
	}
}

func initDatabase() *mongo.Client {
	//mongoUri := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI("mongodb://" + "localhost:27017" + "/?connect=direct")

	// Connect to MongoDB
	//ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	testData(client)
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
