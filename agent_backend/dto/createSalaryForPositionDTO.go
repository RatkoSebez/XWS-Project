package dto

type CreateSalaryForPositionDTO struct {
	Name        string `json:"name"`
	Salary      int    `json:"salary"`
	JobPosition string `json:"jobPosition"`
}
