package external

import (
	"encoding/json"
	"net/http"
	"stockbit-technical-test/test2/internal/domain"
	"stockbit-technical-test/test2/internal/dto"
	"stockbit-technical-test/test2/internal/mapper"
	"strconv"
	"strings"
)

const InternalServerErrorCode = 500

type omdbIntegration struct{}

func NewOmdbIntegration() domain.OmdbIntegration {
	return &omdbIntegration{}
}

func (omdbIntegration) FetchMovie(req *dto.OmdbRequest) (interface{}, int, error) {
	var err error
	var client = &http.Client{}
	data := make(map[string]interface{})
	url := mapper.GenerateUrlOmdb(req)

	request, err := http.NewRequest("GET", strings.ReplaceAll(url, " ", "%20"), nil)
	if err != nil {
		return nil, InternalServerErrorCode, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, InternalServerErrorCode, err
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, InternalServerErrorCode, err
	}
	splitStatus := strings.Split(response.Status, " ")
	convStatusToInt, _ := strconv.Atoi(splitStatus[0])
	return &data, convStatusToInt, nil
}
