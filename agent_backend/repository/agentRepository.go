package repository

import (
	"XWS-Project/agent_backend/dto"
	"XWS-Project/agent_backend/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type AgentRepository struct {
	Client *mongo.Client
}

func (repository *AgentRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	collection := repository.Client.Database("agent").Collection("users")

	filter := bson.D{{"email", email}}
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *AgentRepository) RegisterUser(ctx context.Context, user *model.User) error {
	collection := repository.Client.Database("agent").Collection("users")
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User registrated: ", insertResult.InsertedID)

	return nil
}

func (repository *AgentRepository) CreateCompany(ctx context.Context, company *model.Company) error {
	company.IsApproved = false
	collection := repository.Client.Database("agent").Collection("company")
	insertResult, err := collection.InsertOne(context.TODO(), company)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Company registered: ", insertResult.InsertedID)

	return nil
}

func (repository *AgentRepository) GetCompanyByName(ctx context.Context, name string) (*model.Company, error) {
	company := &model.Company{}
	collection := repository.Client.Database("agent").Collection("company")

	filter := bson.D{{"name", name}}
	err := collection.FindOne(ctx, filter).Decode(&company)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (repository *AgentRepository) ApproveCompany(ctx context.Context, dto *dto.ApproveCompanyDTO) {
	company := &model.Company{}
	collection := repository.Client.Database("agent").Collection("company")
	filter := bson.D{{"name", dto.Name}}
	err := collection.FindOne(ctx, filter).Decode(&company)
	company.IsApproved = true
	update := bson.M{"$set": company}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
}

func (repository *AgentRepository) GetAll(ctx context.Context) []*model.Company {
	var companies []*model.Company
	collection := repository.Client.Database("agent").Collection("company")

	// Finding multiple documents returns a cursor
	cur, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the cursor
	for cur.Next(ctx) {
		var elem model.Company
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		companies = append(companies, &elem)
	}

	return companies
}
