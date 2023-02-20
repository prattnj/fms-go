package dal

import (
	"github.com/prattnj/fms-go/model"
	"testing"
)

func TestAuthtoken(t *testing.T) {
	// todo test authtoken
	db := Db()
	token := model.AuthToken{AuthToken: "test", Username: "njpratt"}
	err := T_insert(db, token)
	if err != nil {
		t.Error(err)
	}
	token2, err := T_find(db, "test")
	if err != nil {
		t.Error(err)
	}
	if token2.AuthToken != token.AuthToken {
		t.Error("Token not found")
	}
	//err = T_clear(db)
	if err != nil {
		t.Error(err)
	}
	err = DbClose(db)
	if err != nil {
		t.Error(err)
	}
}
