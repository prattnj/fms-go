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

	// Test insert and find
	err = T_insert(tx, token)
	handleTestError(t, err, 21)
	token2, err := T_find(tx, "test")
	handleTestError(t, err, 23)
	if token2.AuthToken != token.AuthToken {
		t.Error("Token not found")
	}

	// Test clear and find
	err = T_clear(tx)
	handleTestError(t, err, 30)
	token3, err := T_find(tx, "test")
	if token3.AuthToken != "" {
		t.Error("Token not cleared")
	}
	err = T_insert(tx, token)
	handleTestError(t, err, 36)
	username2, err := T_getUsername(tx, "test")
	if username2 != token.Username {
		t.Error("Username not found")
	}

	// Clean up
	err = tx.Rollback()
	handleTestError(t, err, 44)
	err = DbClose(db)
	handleTestError(t, err, 46)
}

func TestUser(t *testing.T) {

}

func TestPerson(t *testing.T) {

}

func TestEvent(t *testing.T) {

}

func handleTestError(t *testing.T, err error, code int) {
	if err != nil {
		t.Error(err)
		fmt.Println(err)
		fmt.Printf("RETURNING (%d)...\n", code)
	}
}
