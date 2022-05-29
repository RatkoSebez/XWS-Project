package main

import (
	"XWS-Project/agent_backend/dto"
	"XWS-Project/agent_backend/handlers"
	"XWS-Project/agent_backend/model"
	"XWS-Project/agent_backend/repository"
	"XWS-Project/agent_backend/service"
	"XWS-Project/agent_backend/utilities"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

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

	comments := []string{"I like this company", "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Voluptates, animi quaerat. Consequatur iure, ipsam molestiae at corrupti, mollitia quidem fugiat excepturi quae ratione deserunt ipsum molestias unde nemo est error sapiente accusantium eveniet possimus facilis, nostrum alias laborum autem doloribus. Aperiam, id odit mollitia distinctio hic soluta ipsam nisi temporibus veniam repellendus repudiandae modi! Autem deleniti obcaecati iure molestias totam dignissimos blanditiis culpa maiores repudiandae qui nihil consequuntur eos soluta magni alias officia quibusdam reiciendis, natus voluptatibus. Cupiditate, eveniet impedit rem eos quaerat ex, aliquam quos totam odit atque labore? Porro qui, explicabo saepe vero minus perspiciatis ad ullam reiciendis."}
	salaryForPosition := []dto.SalaryForPosition{dto.SalaryForPosition{"backend developer", 3500}}
	interviewReviews := []string{"Lorem ipsum, dolor sit amet consectetur adipisicing elit. Voluptates, animi quaerat. Consequatur iure, ipsam molestiae at corrupti, mollitia quidem fugiat excepturi quae ratione deserunt ipsum molestias unde nemo est error sapiente accusantium eveniet possimus facilis, nostrum alias laborum autem doloribus. Aperiam, id odit mollitia distinctio hic soluta ipsam nisi temporibus veniam repellendus repudiandae modi! Autem deleniti obcaecati iure molestias totam dignissimos blanditiis culpa maiores repudiandae qui nihil consequuntur eos soluta magni alias officia quibusdam reiciendis, natus voluptatibus. Cupiditate, eveniet impedit rem eos quaerat ex, aliquam quos totam odit atque labore? Porro qui, explicabo saepe vero minus perspiciatis ad ullam reiciendis."}

	company1 := model.Company{
		"Apple",
		"apple@gmail.com",
		"1 Apple Park Way, Cupertino, California, U.S.",
		"4681563153",
		"Apple Inc. is an American multinational technology company that specializes in consumer electronics, software and online services headquartered in Cupertino, California, United States. Apple is the largest information technology company by revenue (totaling US$365.8 billion in 2021) and as of May 2022, it is the world's valuable company, the fourth-largest personal computer vendor by unit sales and second-largest mobile phone manufacturer. It is one of the Big Five American information technology companies, alongside Alphabet, Amazon, Meta, and Microsoft.",
		"dusan@gmail.com",
		false,
		comments,
		interviewReviews,
		salaryForPosition,
	}
	company2 := model.Company{
		"Microsoft",
		"microsoft@gmail.com",
		"One Microsoft Way Redmond, Washington, U.S.",
		"3286454686",
		"Microsoft Corporation is an American multinational technology corporation which produces computer software, consumer electronics, personal computers, and related services. Its best-known software products are the Microsoft Windows line of operating systems, the Microsoft Office suite, and the Internet Explorer and Edge web browsers. Its flagship hardware products are the Xbox video game consoles and the Microsoft Surface lineup of touchscreen personal computers. Microsoft ranked No. 21 in the 2020 Fortune 500 rankings of the largest United States corporations by total revenue; it was the world's largest software maker by revenue as of 2016. It is one of the Big Five American information technology companies, alongside Google, Amazon, Apple, and Meta.",
		"dusan@gmail.com",
		false,
		comments,
		interviewReviews,
		salaryForPosition,
	}
	company3 := model.Company{
		"Google",
		"google@gmail.com",
		"1600 Amphitheatre Parkway, Mountain View, California, U.S.",
		"9894842315",
		"Google LLC is an American multinational technology company that focuses on artificial intelligence,[10] search engine, online advertising, cloud computing, computer software, quantum computing, e-commerce, and consumer electronics. It has been referred to as the \"most powerful company in the world\" and one of the world's most valuable brands due to its market dominance, data collection, and technological advantages in the area of artificial intelligence. It is considered one of the Big Five American information technology companies, alongside Amazon, Apple, Meta, and Microsoft.",
		"dusan@gmail.com",
		false,
		comments,
		interviewReviews,
		salaryForPosition,
	}
	company4 := model.Company{
		"Amazon",
		"amazon@gmail.com",
		"Seattle, Washington and Arlington, Virginia, U.S.",
		"5457422315",
		"Amazon.com, Inc. is an American multinational technology company which focuses on e-commerce, cloud computing, digital streaming, and artificial intelligence. It has been referred to as \"one of the most influential economic and cultural forces in the world\", and is one of the world's most valuable brands. It is one of the Big Five American information technology companies, alongside Alphabet, Apple, Meta, and Microsoft.",
		"bojan@gmail.com",
		false,
		comments,
		interviewReviews,
		salaryForPosition,
	}

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
	router.HandleFunc("/api/company/comment", handler.Preflight).Methods("OPTIONS")
	router.HandleFunc("/api/company/comment", handler.CreateCompanyComment).Methods("PUT")
	router.HandleFunc("/api/company/interview", handler.Preflight).Methods("OPTIONS")
	router.HandleFunc("/api/company/interview", handler.CreateCompanyInterviewReview).Methods("PUT")
	router.HandleFunc("/api/company/salary", handler.Preflight).Methods("OPTIONS")
	router.HandleFunc("/api/company/salary", handler.CreateCompanySalary).Methods("PUT")

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
