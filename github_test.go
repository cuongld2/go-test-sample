package go_example

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"os"
	"testing"
)

func getReposList(token string, username string) *resty.Response {

	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", "Bearer "+token).
		Get("https://api.github.com/users/" + username + "/repos")

	if err != nil {

		log.Fatal("Error happens")

	}

	return resp
}
func TestIntMinTableDriven(t *testing.T) {

	token := os.Getenv("TOKEN")
	username := os.Getenv("USERNAME")
	response := getReposList(token, username)
	assert.Equal(t, response.StatusCode(), http.StatusOK, "response status code should be 200")
	t.Log(string(response.Body()))

}
