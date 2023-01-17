package models

type User struct {
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Name string `json:"name" bson:"name"`
	Age int `json:"age" bson:"age"`
	Address `json:"address" bson:"address"`
}

type Address struct {
	State string `json:"state" bson:"state"`
	City  string `json:"city" bson:"city"`
	Pincode int `json:"pincode" bson:"pincode"`
}

