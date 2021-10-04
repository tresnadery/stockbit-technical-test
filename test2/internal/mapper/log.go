package mapper

import (
	"stockbit-technical-test/test2/internal/domain"
	"time"
)

func ToLogModel(request, response string) *domain.Log {
	currentTime := time.Now()
	return &domain.Log{
		Request:   request,
		Response:  response,
		CreatedAt: &currentTime,
	}
}
