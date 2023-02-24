package model

type GenericResponse struct {
	Message string `json:"message,omitempty"`
	Success bool   `json:"success"`
}

type LoginResponse struct {
	AuthToken string `json:"authtoken,omitempty"`
	UserName  string `json:"username,omitempty"`
	PersonID  string `json:"personID,omitempty"`
	Success   bool   `json:"success"`
	Message   string `json:"message,omitempty"`
}

type EventResponse struct {
	Data    []Event `json:"data,omitempty"`
	Success bool    `json:"success"`
	Message string  `json:"message,omitempty"`
}

type EventIDResponse struct {
	AssociatedUsername string  `json:"associatedUsername,omitempty"`
	EventID            string  `json:"eventID,omitempty"`
	PersonID           string  `json:"personID,omitempty"`
	Latitude           float32 `json:"latitude,omitempty"`
	Longitude          float32 `json:"longitude,omitempty"`
	Country            string  `json:"country,omitempty"`
	City               string  `json:"city,omitempty"`
	EventType          string  `json:"eventType,omitempty"`
	Year               int     `json:"year,omitempty"`
	Success            bool    `json:"success"`
	Message            string  `json:"message,omitempty"`
}

type PersonResponse struct {
	Data    []Person `json:"data,omitempty"`
	Success bool     `json:"success"`
	Message string   `json:"message,omitempty"`
}

type PersonIDResponse struct {
	AssociatedUsername string `json:"associatedUsername,omitempty"`
	PersonID           string `json:"personID,omitempty"`
	FirstName          string `json:"firstName,omitempty"`
	LastName           string `json:"lastName,omitempty"`
	Gender             string `json:"gender,omitempty"`
	FatherID           string `json:"fatherID,omitempty"`
	MotherID           string `json:"motherID,omitempty"`
	SpouseID           string `json:"spouseID,omitempty"`
	Success            bool   `json:"success"`
	Message            string `json:"message,omitempty"`
}
