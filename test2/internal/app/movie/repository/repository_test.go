package repository

import (
	"context"
	"stockbit-technical-test/test2/internal/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestStore(t *testing.T) {
	now := time.Now()
	log := &model.Log{
		Request:   "TestRequest",
		Response:  "Test",
		CreatedAt: &now,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "INSERT logs SET request=\\? , response=\\? , created_at=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(log.Request, log.Response, log.CreatedAt).WillReturnResult(sqlmock.NewResult(12, 1))

	a := NewRepository(db)

	err = a.StoreLog(context.TODO(), log)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), log.ID)
}
