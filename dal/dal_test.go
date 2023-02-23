package dal

import (
	"fmt"
	"github.com/prattnj/fms-go/model"
	"testing"
)

func TestAuthtoken(t *testing.T) {

	// Start transaction / set up
	db := Db()
	tx, err := db.Begin()
	handleTestError(t, err, 14)
	token := model.AuthToken{AuthToken: "test", Username: "njpratt"}
	err = T_clear(tx)
	handleTestError(t, err, 17)

	// Test all methods
	err = T_insert(tx, token)
	handleTestError(t, err, 21)
	err = T_insert(tx, token)
	if err == nil {
		t.Error("Duplicate token inserted, but error not thrown")
	}
	token2, err := T_find(tx, "test")
	handleTestError(t, err, 27)
	if token2.AuthToken != token.AuthToken {
		t.Error("Token not found")
	}
	err = T_clear(tx)
	handleTestError(t, err, 32)
	token3, err := T_find(tx, "test")
	if token3.AuthToken != "" {
		t.Error("Token not cleared")
	}
	err = T_insert(tx, token)
	handleTestError(t, err, 38)
	username2, err := T_getUsername(tx, "test")
	if username2 != token.Username {
		t.Error("Username not found")
	}

	// Clean up
	err = tx.Rollback()
	handleTestError(t, err, 46)
	err = DbClose(db)
	handleTestError(t, err, 48)
}

func TestUser(t *testing.T) {

	// Start transaction / set up
	db := Db()
	tx, err := db.Begin()
	handleTestError(t, err, 56)
	userToAdd := model.User{
		Username:  "njpratt",
		Password:  "fake",
		Email:     "pratnt@gmail.com",
		FirstName: "Noah",
		LastName:  "Pratt",
		Gender:    "m",
		PersonID:  "fakeID",
	}
	err = U_clear(tx)
	handleTestError(t, err, 67)

	// Test all methods
	err = U_insert(tx, userToAdd)
	handleTestError(t, err, 71)
	err = U_insert(tx, userToAdd)
	if err == nil {
		t.Error("Duplicate user inserted, but error not thrown")
	}
	user2, err := U_find(tx, "njpratt")
	handleTestError(t, err, 77)
	if user2.Username != userToAdd.Username {
		t.Error("User not found")
	}
	gender, err := U_getGender(tx, "njpratt")
	handleTestError(t, err, 82)
	if gender != "m" {
		t.Error("Invalid gender found")
	}
	valid, err := U_validate(tx, "njpratt", "fake")
	handleTestError(t, err, 87)
	if !valid {
		t.Error("User not validated when they should have been")
	}
	valid, err = U_validate(tx, "njpratt", "wrong")
	handleTestError(t, err, 92)
	if valid {
		t.Error("User validated when they shouldn't have been")
	}
	err = U_clear(tx)
	handleTestError(t, err, 97)
	user3, err := U_find(tx, "njpratt")
	if user3.Username != "" {
		t.Error("User not cleared")
	}

	// Clean up
	err = tx.Rollback()
	handleTestError(t, err, 105)
	err = DbClose(db)
	handleTestError(t, err, 107)

}

func TestPerson(t *testing.T) {

	// Start transaction / set up
	db := Db()
	tx, err := db.Begin()
	handleTestError(t, err, 104)
	personToAdd := model.Person{
		PersonID:           "fakeID",
		AssociatedUsername: "njpratt",
		FirstName:          "Noah",
		LastName:           "Pratt",
		Gender:             "m",
		FatherID:           "",
		MotherID:           "",
		SpouseID:           "",
	}
	err = P_clear(tx)
	handleTestError(t, err, 128)

	// Test all methods
	err = P_insert(tx, personToAdd)
	handleTestError(t, err, 132)
	err = P_insert(tx, personToAdd)
	if err == nil {
		t.Error("Duplicate person inserted, but error not thrown")
	}
	person2, err := P_find(tx, "fakeID")
	handleTestError(t, err, 138)
	if person2.PersonID != personToAdd.PersonID {
		t.Error("Person not found")
	}
	persons, err := P_getForUsername(tx, "njpratt")
	handleTestError(t, err, 143)
	if len(persons) != 1 {
		t.Error("Invalid number of persons found")
	}
	if persons[0].PersonID != personToAdd.PersonID {
		t.Error("Invalid person found")
	}
	err = P_clear(tx)
	handleTestError(t, err, 151)
	person3, err := P_find(tx, "fakeID")
	if person3.PersonID != "" {
		t.Error("Person not cleared")
	}
	err = P_insert(tx, personToAdd)
	handleTestError(t, err, 157)
	err = P_clearForUser(tx, "njpratt")
	handleTestError(t, err, 159)
	person4, err := P_find(tx, "fakeID")
	if person4.PersonID != "" {
		t.Error("Person not cleared")
	}

	// Clean up
	err = tx.Rollback()
	handleTestError(t, err, 167)
	err = DbClose(db)
	handleTestError(t, err, 169)

}

func TestEvent(t *testing.T) {

	// Start transaction / set up
	db := Db()
	tx, err := db.Begin()
	handleTestError(t, err, 178)
	eventToAdd := model.Event{
		EventID:            "fakeID",
		AssociatedUsername: "njpratt",
		PersonID:           "fakeID",
		Latitude:           0.0,
		Longitude:          0.0,
		Country:            "USA",
		City:               "Provo",
		EventType:          "Birth",
		Year:               1997,
	}
	err = E_clear(tx)
	handleTestError(t, err, 191)

	// Test all methods
	err = E_insert(tx, eventToAdd)
	handleTestError(t, err, 195)
	err = E_insert(tx, eventToAdd)
	if err == nil {
		t.Error("Duplicate event inserted, but error not thrown")
	}
	event2, err := E_find(tx, "fakeID")
	handleTestError(t, err, 201)
	if event2.EventID != eventToAdd.EventID {
		t.Error("Event not found")
	}
	events, err := E_getForUsername(tx, "njpratt")
	handleTestError(t, err, 206)
	if len(events) != 1 {
		t.Error("Invalid number of events found")
	}
	if events[0].EventID != eventToAdd.EventID {
		t.Error("Invalid event found")
	}
	err = E_clear(tx)
	handleTestError(t, err, 214)
	event3, err := E_find(tx, "fakeID")
	if event3.EventID != "" {
		t.Error("Event not cleared")
	}
	err = E_insert(tx, eventToAdd)
	handleTestError(t, err, 220)
	err = E_clearForUser(tx, "njpratt")
	handleTestError(t, err, 222)
	event4, err := E_find(tx, "fakeID")
	if event4.EventID != "" {
		t.Error("Event not cleared")
	}
	err = E_insert(tx, eventToAdd)
	handleTestError(t, err, 228)
	event5, err := E_findBirth(tx, "fakeID")
	handleTestError(t, err, 230)
	if event5.EventID != eventToAdd.EventID {
		t.Error("Birth event not found")
	}

	// Clean up
	err = tx.Rollback()
	handleTestError(t, err, 237)
	err = DbClose(db)
	handleTestError(t, err, 239)

}

func handleTestError(t *testing.T, err error, code int) {
	if err != nil {
		t.Error(err)
		fmt.Println(err)
		fmt.Printf("RETURNING (%d)...\n", code)
	}
}
