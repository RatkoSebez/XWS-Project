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
	// users

	user1 := model.User{"Dusan", "Markovic", "dusan@gmail.com", "dulecar", "dulecar123", false}
	user2 := model.User{"Bojan", "Prodanovic", "bojan@gmail.com", "bp21", "123", false}
	user3 := model.User{"Ratko", "Sebez", "ratko@gmail.com", "ratko", "123", true}

	collection := client.Database("agent").Collection("users")
	_, err := collection.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	_, err = collection.InsertOne(context.TODO(), user1)
	if err != nil {
		panic(err)
	}
	_, err = collection.InsertOne(context.TODO(), user2)
	if err != nil {
		panic(err)
	}
	_, err = collection.InsertOne(context.TODO(), user3)
	if err != nil {
		panic(err)
	}

	// companies

	company1 := model.Company{"kompanija1", "mail1@gmail.com", "adresa1", "4681563153", "opis1", "dusan@gmail.com", false}
	company2 := model.Company{"kompanija2", "mail2@gmail.com", "adresa2", "3286454686", "opis2", "dusan@gmail.com", false}
	company3 := model.Company{"kompanija3", "mail3@gmail.com", "adresa3", "9894842315", "opis3", "dusan@gmail.com", false}
	company4 := model.Company{"kompanija4", "mail4@gmail.com", "adresa4", "5457422315", "opis4", "bojan@gmail.com", false}

	collection = client.Database("agent").Collection("company")
	_, err = collection.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	_, err = collection.InsertOne(context.TODO(), company1)
	if err != nil {
		panic(err)
	}
	_, err = collection.InsertOne(context.TODO(), company2)
	if err != nil {
		panic(err)
	}
	_, err = collection.InsertOne(context.TODO(), company3)
	if err != nil {
		panic(err)
	}
	_, err = collection.InsertOne(context.TODO(), company4)
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

func initRepo(clientDB *mongo.Client) *repository.AgentRepository {
	return &repository.AgentRepository{Client: clientDB}
}

func initService(repo *repository.AgentRepository) *service.AgentService {
	return &service.AgentService{Repository: repo}
}

func initHandler(authenticationService *service.AgentService) *handlers.AgentHandler {
	return &handlers.AgentHandler{Service: authenticationService}
}

func handlerFunc(handler *handlers.AgentHandler) {
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
	// auth
	router.HandleFunc("/api/test", handler.Preflight).Methods("OPTIONS")
	router.HandleFunc("/api/test", handler.Test).Methods("GET")
	router.HandleFunc("/api/login", handler.Preflight).Methods("OPTIONS")
	router.HandleFunc("/api/login", handler.Login).Methods("POST")
	router.HandleFunc("/api/register", handler.Preflight).Methods("OPTIONS")
	router.HandleFunc("/api/register", handler.Register).Methods("POST")
	// company
	router.HandleFunc("/api/company", handler.Preflight).Methods("OPTIONS")
	router.HandleFunc("/api/company", handler.CreateCompany).Methods("POST")
	router.HandleFunc("/api/company", handler.Preflight).Methods("OPTIONS")
	router.HandleFunc("/api/company", handler.GetAllCompanies).Methods("GET")
	router.HandleFunc("/api/company", handler.Preflight).Methods("OPTIONS")
	router.HandleFunc("/api/company", handler.EditCompany).Methods("PUT")
	//router.HandleFunc("/api/company/{email}", handler.Preflight).Methods("OPTIONS")
	//router.HandleFunc("/api/company/{email}", handler.GetAllUsersCompanies).Methods("GET")
	router.HandleFunc("/api/company/approve", handler.Preflight).Methods("OPTIONS")
	router.HandleFunc("/api/company/approve", handler.ApproveCompany).Methods("POST")

	log.Fatal(http.ListenAndServe(":8090", router))
}

func main() {
	utilities.TracerInit("agent app")
	dbClient := initDB()
	repo := initRepo(dbClient)
	service := initService(repo)
	handler := initHandler(service)
	handlerFunc(handler)
}
