package model

type AuthToken struct {
	AuthToken string `json:"authToken"`
	Username  string `json:"username"`
}

type Event struct {
	EventID            string  `json:"eventID"`
	AssociatedUsername string  `json:"associatedUsername"`
	PersonID           string  `json:"personID"`
	Latitude           float32 `json:"latitude"`
	Longitude          float32 `json:"longitude"`
	Country            string  `json:"country"`
	City               string  `json:"city"`
	EventType          string  `json:"eventType"`
	Year               int     `json:"year"`
}

type Person struct {
	PersonID           string `json:"personID"`
	AssociatedUsername string `json:"associatedUsername"`
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	Gender             string `json:"gender"`
	FatherID           string `json:"fatherID,omitempty"`
	MotherID           string `json:"motherID,omitempty"`
	SpouseID           string `json:"spouseID,omitempty"`
}

type User struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Gender    string `json:"gender"`
	PersonID  string `json:"personID"`
}

type Log struct {
	Timestamp int64  `json:"timestamp"`
	Endpoint  string `json:"endpoint"`
	IPv4      string `json:"ipv4"`
	Success   bool   `json:"success"`
}

type Location struct {
	Country   string  `json:"country"`
	City      string  `json:"city"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
