package models

type UserLocation struct{
	Country string `json:"country"`
	City string `json:"city"`
}

type UserModel struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// Location UserLocation	`json:"location"`
}