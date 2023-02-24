package model

type GenericResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type LoginResponse struct {
	AuthToken string `json:"authtoken"`
	UserName  string `json:"username"`
	PersonID  string `json:"personID"`
	Success   bool   `json:"success"`
	Message   string `json:"message,omitempty"`
}

type EventResponse struct {
	Data    []Event `json:"data"`
	Success bool    `json:"success"`
	Message string  `json:"message"`
}

type EventIDResponse struct {
	AssociatedUsername string  `json:"associatedUsername"`
	EventID            string  `json:"eventID"`
	PersonID           string  `json:"personID"`
	Latitude           float32 `json:"latitude"`
	Longitude          float32 `json:"longitude"`
	Country            string  `json:"country"`
	City               string  `json:"city"`
	EventType          string  `json:"eventType"`
	Year               int     `json:"year"`
	Success            bool    `json:"success"`
	Message            string  `json:"message"`
}

type PersonResponse struct {
	Data    []Person `json:"data"`
	Success bool     `json:"success"`
	Message string   `json:"message"`
}

type PersonIDResponse struct {
	AssociatedUsername string `json:"associatedUsername"`
	PersonID           string `json:"personID"`
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	Gender             string `json:"gender"`
	FatherID           string `json:"fatherID"`
	MotherID           string `json:"motherID"`
	SpouseID           string `json:"spouseID"`
	Success            bool   `json:"success"`
	Message            string `json:"message"`
}
