package models

type User struct {
	Name string `json:"name" bson:"user_name"`
	Age int `json:"age" bson:"user_age"`
	Address `json:"address" bson:"user_address"`
}

type Address struct {
	State string `json:"state" bson:"state"`
	City  string `json:"city" bson:"city"`
	Pincode int `json:"pincode" bson:"pincode"`
}