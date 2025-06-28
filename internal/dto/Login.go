package dto

type LoginRequest struct {
	EmployeeID string `json:"employee_id"`
	Password   string `json:"password"`
}
