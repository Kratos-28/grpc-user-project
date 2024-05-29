package models

type User struct {
	ID      int32
	Fname   string
	City    string
	Phone   int64
	Height  float32
	Married bool
}

var Users = []User{
	{ID: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
	{ID: 2, Fname: "John", City: "NY", Phone: 2345678901, Height: 6.1, Married: false},
}
