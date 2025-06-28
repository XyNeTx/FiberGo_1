package service

import (
	"fiberGo_1/internal/model"
)

var LoginDetailSlice []model.Login

func initLogin() {
	LoginDetail := model.Login{
		EmployeeID:  "1",
		Password:    "password",
		Token:       "token",
		Role:        "role",
		Expire_Date: "expire_date",
	}
	LoginDetailSlice = make([]model.Login, 0)
	LoginDetailSlice[0] = LoginDetail
}

func GetLoginDetail() []model.Login {
	initLogin()
	return LoginDetailSlice
}

func Register(obj model.Login) {
	LoginDetailSlice = append(LoginDetailSlice, obj)
}
