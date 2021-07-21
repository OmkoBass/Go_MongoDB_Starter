package models

type User struct {
	Id 		 string  `json:"id,omitempty" bson:"_id,omitempty"`
	Username string  `json:"username" validate:"required"`
	Password string  `json:"password"`
	Salary   float64 `json:"salary" validate:"gt=0"`
	Age 	 float64 `json:"age"`
}