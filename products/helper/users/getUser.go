package helper_users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Users struct {
	Message string `json:"message"`
	Result  struct {
		UUID     string `json:"UUID"`
		Name     string `json:"Name"`
		Email    string `json:"Email"`
		Password string `json:"Password"`
		Image    string `json:"Image"`
		Role     string `json:"Role"`
	} `json:"result"`
	Status int `json:"status"`
}

func GetUserUUID(uuid string, jwt string) Users {

	url := "http://localhost:8080/users/" + uuid
	method := "GET"

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)

	req.Header.Add("Authorization", "Bearer "+jwt)

	res, _ := client.Do(req)

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	fmt.Println(string(body))

	var result = Users{}

	json.Unmarshal(body, &result)

	return result
}
