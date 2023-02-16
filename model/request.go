package model

type LoadRequest struct {
	Users   []User   `json:"users"`
	Persons []Person `json:"persons"`
	Events  []Event  `json:"events"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Gender    string `json:"gender"`
}
