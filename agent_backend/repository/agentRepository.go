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
	var comments []string
	company.Comments = comments
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

func (repository *AgentRepository) GetAllCompanies(ctx context.Context) []*model.Company {
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

func (repository *AgentRepository) EditCompany(ctx context.Context, newCompany *model.Company) {
	company := &model.Company{}
	collection := repository.Client.Database("agent").Collection("company")
	filter := bson.D{{"name", newCompany.Name}}
	err := collection.FindOne(ctx, filter).Decode(&company)
	company.Address = newCompany.Address
	company.PhoneNumber = newCompany.PhoneNumber
	company.Email = newCompany.Email
	company.Description = newCompany.Description
	update := bson.M{"$set": company}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
}

func (repository *AgentRepository) CreateCompanyComment(ctx context.Context, dto *dto.CreateCommentDTO) {
	company := &model.Company{}
	collection := repository.Client.Database("agent").Collection("company")
	filter := bson.D{{"name", dto.Name}}
	err := collection.FindOne(ctx, filter).Decode(&company)
	company.Comments = append(company.Comments, dto.Comment)
	update := bson.M{"$set": company}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
}

func (repository *AgentRepository) CreateCompanyInterviewReview(ctx context.Context, dto *dto.CreateInterviewReviewDTO) {
	company := &model.Company{}
	collection := repository.Client.Database("agent").Collection("company")
	filter := bson.D{{"name", dto.Name}}
	err := collection.FindOne(ctx, filter).Decode(&company)
	company.InterviewReviews = append(company.InterviewReviews, dto.InterviewReview)
	update := bson.M{"$set": company}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
}

func (repository *AgentRepository) CreateCompanySalary(ctx context.Context, data *dto.CreateSalaryForPositionDTO) {
	company := &model.Company{}
	collection := repository.Client.Database("agent").Collection("company")
	filter := bson.D{{"name", data.Name}}
	err := collection.FindOne(ctx, filter).Decode(&company)
	company.SalaryForPosition = append(company.SalaryForPosition, dto.SalaryForPosition{data.JobPosition, data.Salary})
	update := bson.M{"$set": company}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
}

//func (repository *AgentRepository) GetAllUsersCompanies(ctx context.Context, userEmail string) []*model.Company {
//	var companies []*model.Company
//	collection := repository.Client.Database("agent").Collection("company")
//
//	// Finding multiple documents returns a cursor
//	cur, err := collection.Find(ctx, bson.D{{"ownerEmail", userEmail}})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Iterate through the cursor
//	for cur.Next(ctx) {
//		var elem model.Company
//		err := cur.Decode(&elem)
//		if err != nil {
//			log.Fatal(err)
//		}
//		companies = append(companies, &elem)
//	}
//
//	return companies
//}
