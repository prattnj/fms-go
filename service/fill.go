package service

import (
	"encoding/json"
	"fmt"
	"github.com/prattnj/fms-go/dal"
	"github.com/prattnj/fms-go/model"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var fillPeople []*model.Person
var fillEvents []*model.Event
var fillUser model.User
var totalGen int

var maleNames *[]string
var femaleNames *[]string
var lastNames *[]string
var locations *[]model.Location

func Fill(username string, generations string) model.GenericResponse {

	// Validate generations
	if username == "" {
		return model.GenericResponse{Success: false, Message: "Error: blank username"}
	}
	numGen, err := strconv.Atoi(generations)
	if err != nil {
		return model.GenericResponse{Success: false, Message: "Error: generation number must be an integer"}
	}
	if numGen < 0 {
		return model.GenericResponse{Success: false, Message: "Error: negative number of generations"}
	} else if numGen > 12 {
		return model.GenericResponse{Success: false, Message: "Error: too many generations"}
	}

	// Validate username
	db := dal.Db()
	if db == nil {
		return serverError
	}
	tx, err := db.Begin()
	if err != nil {
		return serverError
	}
	fillUser, err = dal.U_find(tx, username)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return serverError
		}
		return serverError
	}
	if fillUser.Username == "" {
		return model.GenericResponse{Success: false, Message: "Error: invalid username"}
	}

	// Clear persons and events belonging to user
	err = dal.P_clearForUser(tx, username)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return serverError
		}
		return serverError
	}
	err = dal.E_clearForUser(tx, username)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return serverError
		}
		return serverError
	}

	// Fill the database
	err = generateData(numGen, fillUser.Gender)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return serverError
		}
		return serverError
	}
	for person := range fillPeople {
		err := dal.P_insert(tx, fillPeople[person])
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return serverError
			}
			return serverError
		}
	}
	for event := range fillEvents {
		err := dal.E_insert(tx, fillEvents[event])
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return serverError
			}
			return serverError
		}
	}

	err = tx.Commit()
	if err != nil {
		return serverError
	}
	err = dal.DbClose(db)
	if err != nil {
		return serverError
	}

	return model.GenericResponse{Success: true, Message: "Successfully added " + strconv.Itoa(len(fillPeople)) +
		" persons and " + strconv.Itoa(len(fillEvents)) + " events to the database"}
}

// Wrapper method for generatePerson, to clear the arrays
func generateData(numGen int, gender string) error {

	fillPeople = nil
	fillEvents = nil
	totalGen = numGen

	var err error
	maleNames, femaleNames, lastNames, locations, err = instantiateData()
	if err != nil {
		return err
	}

	rootPerson := generatePerson(numGen, gender)
	rootPerson.PersonID = fillUser.PersonID
	rootPerson.FirstName = fillUser.FirstName
	rootPerson.LastName = fillUser.LastName
	generateUserBirth()
	fixLastNames(rootPerson)

	return nil
}

func generatePerson(numGen int, gender string) *model.Person {

	var father model.Person
	var mother model.Person

	if numGen > 0 {

		// Recursively generate parents
		father = *generatePerson(numGen-1, "m")
		mother = *generatePerson(numGen-1, "f")

		updateSpouseID(mother.PersonID, father.PersonID)
		updateSpouseID(father.PersonID, mother.PersonID)

		generateMarriage(&father, &mother)
	}

	personID := generateID(defaultIDLength)

	// Generate first name
	firstName := ""
	if gender == "m" {
		firstName = getMaleName()
	} else {
		firstName = getFemaleName()
	}

	// Generate last name
	lastName := ""
	if numGen > 0 {
		lastName = father.LastName
	} else {
		lastName = getLastName()
	}

	// Create and add person
	person := model.Person{
		PersonID:           personID,
		AssociatedUsername: fillUser.Username,
		FirstName:          firstName,
		LastName:           lastName,
		Gender:             gender,
		FatherID:           "",
		MotherID:           "",
		SpouseID:           "",
	}
	if numGen > 0 {
		person.FatherID = father.PersonID
		person.MotherID = mother.PersonID
	}
	fillPeople = append(fillPeople, &person)

	// Generate events for everyone but user
	if numGen != totalGen {
		birthYear := generateBirth(personID, numGen)
		generateDeath(personID, birthYear)
	}

	return &person
}

func updateSpouseID(personID string, spouseID string) {

	for person := range fillPeople {
		if fillPeople[person].PersonID == personID {
			fillPeople[person].SpouseID = spouseID
			break
		}
	}
}

func generateBirth(personID string, numGen int) int {

	location := getLocation()

	// All birth years are multiples of 29, going backwards from ~2000
	birthYear := 1996 - ((totalGen - numGen) * 29)
	birthYear += rand.Intn(7)

	// Create and add event
	birth := model.Event{
		EventID:            generateID(defaultIDLength),
		AssociatedUsername: fillUser.Username,
		PersonID:           personID,
		Latitude:           location.Latitude,
		Longitude:          location.Longitude,
		Country:            location.Country,
		City:               location.City,
		EventType:          "birth",
		Year:               birthYear,
	}
	fillEvents = append(fillEvents, &birth)
	return birthYear
}

func generateUserBirth() {

	// Custom birth event for yours truly
	var location model.Location
	if fillUser.Username == "njpratt" {
		location = model.Location{
			Latitude:  33.4484,
			Longitude: -112.074,
			Country:   "United States",
			City:      "Phoenix",
		}
	} else {
		location = getLocation()
	}

	var birthYear int
	if fillUser.Username == "njpratt" {
		birthYear = 2000
	} else {
		birthYear = 1996 + rand.Intn(7)
	}
	userBirth := model.Event{
		EventID:            generateID(defaultIDLength),
		AssociatedUsername: fillUser.Username,
		PersonID:           fillUser.PersonID,
		Latitude:           location.Latitude,
		Longitude:          location.Longitude,
		Country:            location.Country,
		City:               location.City,
		EventType:          "birth",
		Year:               birthYear,
	}
	fillEvents = append(fillEvents, &userBirth)

}

func generateDeath(personID string, birthYear int) {

	location := getLocation()

	// People live to be 75-95 years old
	deathYear := birthYear + 75 + rand.Intn(21)

	// Create and add event
	death := model.Event{
		EventID:            generateID(defaultIDLength),
		AssociatedUsername: fillUser.Username,
		PersonID:           personID,
		Latitude:           location.Latitude,
		Longitude:          location.Longitude,
		Country:            location.Country,
		City:               location.City,
		EventType:          "death",
		Year:               deathYear,
	}
	fillEvents = append(fillEvents, &death)

}

func generateMarriage(father *model.Person, mother *model.Person) {

	location := getLocation()

	// Calculate marriage year (between when the younger person 18 and 24 years old inclusive)
	motherBirthYear := findBirthYear(mother.PersonID)
	fatherBirthYear := findBirthYear(father.PersonID)
	var marriageYear int
	if motherBirthYear > fatherBirthYear {
		marriageYear = motherBirthYear + 18 + rand.Intn(7)
	} else {
		marriageYear = fatherBirthYear + 18 + rand.Intn(7)
	}

	// Create and add events
	motherMarriage := model.Event{
		EventID:            generateID(defaultIDLength),
		AssociatedUsername: fillUser.Username,
		PersonID:           mother.PersonID,
		Latitude:           location.Latitude,
		Longitude:          location.Longitude,
		Country:            location.Country,
		City:               location.City,
		EventType:          "marriage",
		Year:               marriageYear,
	}
	fatherMarriage := model.Event{
		EventID:            generateID(defaultIDLength),
		AssociatedUsername: fillUser.Username,
		PersonID:           father.PersonID,
		Latitude:           location.Latitude,
		Longitude:          location.Longitude,
		Country:            location.Country,
		City:               location.City,
		EventType:          "marriage",
		Year:               marriageYear,
	}
	fillEvents = append(fillEvents, &motherMarriage)
	fillEvents = append(fillEvents, &fatherMarriage)
}

func fixLastNames(rootPerson *model.Person) {
	// Sets last names of people directly up the paternal line to be the user's
	p := rootPerson
	for {
		p = findFather(p)
		if p == nil {
			break
		}
		p.LastName = fillUser.LastName
	}
}

func findBirthYear(personID string) int {
	for _, event := range fillEvents {
		if event.PersonID == personID && event.EventType == "birth" {
			return event.Year
		}
	}
	return -1
}

func findFather(child *model.Person) *model.Person {
	if child.FatherID == "" {
		return nil
	}
	for _, person := range fillPeople {
		if person.PersonID == child.FatherID {
			return person
		}
	}
	return nil
}

func getMaleName() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return (*maleNames)[r.Intn(len(*maleNames))]
}

func getFemaleName() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return (*femaleNames)[r.Intn(len(*femaleNames))]
}

func getLastName() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return (*lastNames)[r.Intn(len(*lastNames))]
}

func getLocation() model.Location {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return (*locations)[r.Intn(len(*locations))]
}

func instantiateData() (*[]string, *[]string, *[]string, *[]model.Location, error) {

	maleFile := "json/mnames.json"
	femaleFile := "json/fnames.json"
	lastFile := "json/snames.json"
	locationFile := "json/locations.json"
	maleJson, err1 := os.ReadFile(maleFile)
	femaleJson, err2 := os.ReadFile(femaleFile)
	lastJson, err3 := os.ReadFile(lastFile)
	locationJson, err4 := os.ReadFile(locationFile)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return nil, nil, nil, nil, fmt.Errorf("error reading json files")
	}

	var maleNames []string
	var femaleNames []string
	var lastNames []string
	var locations []model.Location

	err1 = json.Unmarshal(maleJson, &maleNames)
	err2 = json.Unmarshal(femaleJson, &femaleNames)
	err3 = json.Unmarshal(lastJson, &lastNames)
	err4 = json.Unmarshal(locationJson, &locations)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return nil, nil, nil, nil, fmt.Errorf("error unmarshalling json")
	}

	return &maleNames, &femaleNames, &lastNames, &locations, nil
}
