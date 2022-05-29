package model

import "XWS-Project/agent_backend/dto"

type Company struct {
	Name              string                  `json:"name"`
	Email             string                  `json:"email"`
	Address           string                  `json:"address"`
	PhoneNumber       string                  `json:"phoneNumber"`
	Description       string                  `json:"description"`
	OwnerEmail        string                  `json:"ownerEmail" bson:"ownerEmail"`
	IsApproved        bool                    `json:"isApproved"`
	Comments          []string                `json:"comments"`
	InterviewReviews  []string                `json:"interviewReviews"`
	SalaryForPosition []dto.SalaryForPosition `json:"salaryForPosition"`
	Jobs              []Job                   `json:"jobs"`
}
