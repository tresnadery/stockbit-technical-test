package mapper

import (
	"stockbit-technical-test/test2/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToLogModel(t *testing.T) {
	request := "request"
	response := "response"
	now := time.Now()
	expected := &domain.Log{
		Request:   request,
		Response:  response,
		CreatedAt: &now,
	}
	actual := ToLogModel(request, response)
	actual.CreatedAt = &now
	assert.Equal(t, actual, expected, "they should be equal")
}
