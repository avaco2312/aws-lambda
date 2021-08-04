package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	os.Setenv("TABLE_NAME", "movies")
	request := events.APIGatewayProxyRequest{
		Headers: map[string]string{"Count": "3"},
	}
	var response events.APIGatewayProxyResponse
	response, err := findAll(request)
	err1 := err
	var movies []Movie
	if err == nil {
		err1 = json.Unmarshal([]byte(response.Body), &movies)
	}
	assert.IsType(t, nil, err)
	assert.Equal(t, 200, response.StatusCode)
	assert.IsType(t, nil, err1)
	assert.Equal(t, 3, len(movies))
}
