package dto

type CreateJobDTO struct {
	CompanyName  string `json:"companyName"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Requirements string `json:"requirements"`
	Benefits     string `json:"benefits"`
}
