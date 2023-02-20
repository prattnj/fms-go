package dal

import (
	"github.com/prattnj/fms-go/model"
	"testing"
)

func TestAuthtoken(t *testing.T) {
	db := Db()
	tx, err := db.Begin()
	if err != nil {
		t.Error(err)
	}
	token := model.AuthToken{AuthToken: "test", Username: "njpratt"}
	err = T_insert(tx, token)
	if err != nil {
		t.Error(err)
	}
	token2, err := T_find(tx, "test")
	if err != nil {
		t.Error(err)
	}
	if token2.AuthToken != token.AuthToken {
		t.Error("Token not found")
	}
	err = T_clear(tx)
	if err != nil {
		t.Error(err)
	}
	err = tx.Rollback()
	err = DbClose(db)
	if err != nil {
		t.Error(err)
	}
}
