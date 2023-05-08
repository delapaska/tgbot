package animals

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

type RandomImage struct {
	Message string
	Status  string
}

var statusCodesAll = []int{
	100, 101, 102, 103, 200, 201, 202, 203, 204, 206,
	207, 300, 301, 302, 303, 304, 305, 306, 307, 308,
	400, 401, 402, 403, 404, 405, 406, 407, 408, 409,
	410, 411, 412, 413, 414, 415, 416, 417, 418, 420,
	421, 422, 423, 424, 425, 426, 429, 431, 444, 450,
	451, 500, 501, 502, 503, 504, 505, 506, 507, 508,
	510, 511}

func GetRandomCatStatusCode() int {
	randomIndex := rand.Intn(len(statusCodesAll))
	pick := statusCodesAll[randomIndex]
	return pick
}

func GetRandomDogImageUrl() string {
	response, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data RandomImage
	json.Unmarshal(responseData, &data)

	return data.Message
}
