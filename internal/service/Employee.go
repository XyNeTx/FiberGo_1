package service

import (
	"fiberGo_1/internal/model"
)

func GetEmployeeByID(id string) (model.Employee, error) {
	emp := model.Employee{}

	switch id {
	case "1":
		emp = model.Employee{
			EmployeeID: "1",
			Name:       "John Doe",
			Age:        30,
			Email:      "vO4l3@example.com",
			Salary:     50000,
		}
	case "2":
		emp = model.Employee{
			EmployeeID: "2",
			Name:       "Jane Doe",
			Age:        25,
			Email:      "oV4e9@example.com",
			Salary:     60000,
		}
	default:
		return model.Employee{}, nil
	}
	return emp, nil
}
