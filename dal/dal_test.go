package dal

import (
	"fmt"
	"github.com/prattnj/fms-go/model"
	"testing"
)

func TestAuthtoken(t *testing.T) {
	db := Db()
	tx, err := db.Begin()
	handleTestError(t, err)
	token := model.AuthToken{AuthToken: "test", Username: "njpratt"}
	err = T_clear(tx)
	handleTestError(t, err)
	err = T_insert(tx, token)
	handleTestError(t, err)
	token2, err := T_find(tx, "test")
	handleTestError(t, err)
	if token2.AuthToken != token.AuthToken {
		t.Error("Token not found")
	}
	err = T_clear(tx)
	fmt.Print("hello6\n")
	handleTestError(t, err)
	fmt.Print("hello7\n")
	err = tx.Rollback()
	err = DbClose(db)
	handleTestError(t, err)
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
		fmt.Println(err)
		fmt.Print("RETURNING...\n")
	}
}
