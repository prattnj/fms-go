package dal

import (
	"fmt"
	"github.com/prattnj/fms-go/model"
	"testing"
)

func TestAuthtoken(t *testing.T) {
	db := Db()
	fmt.Print("hello1\n")
	tx, err := db.Begin()
	fmt.Print("hello2\n")
	handleTestError(t, err)
	token := model.AuthToken{AuthToken: "test", Username: "njpratt"}
	err = T_clear(tx)
	handleTestError(t, err)
	err = T_insert(tx, token)
	fmt.Print("hello4\n")
	handleTestError(t, err)
	token2, err := T_find(tx, "test")
	handleTestError(t, err)
	if token2.AuthToken != token.AuthToken {
		t.Error("Token not found")
	}
	err = T_clear(tx)
	handleTestError(t, err)
	fmt.Print("hello7\n")
	err = tx.Rollback()
	err = DbClose(db)
	fmt.Print("hello9\n")
	handleTestError(t, err)

	fmt.Print("hello20\n")
}

func TestEvent(t *testing.T) {

}

func TestPerson(t *testing.T) {

}

func TestUser(t *testing.T) {

}

func handleTestError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
		fmt.Print("RETURNING...\n")
	}
}
