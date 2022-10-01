package sing

import (
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/Jeffail/gabs"
	"math/rand"
	"errors"
)

var (
	ErrApiKeyNotFound = errors.New("api key not found")
)

func Sing() (interface{}, error) {
	terms := [5]string{"dog", "cat", "bug", "horse", "bat"}

	if os.Getenv("SHAZAM_RAPIDAPI_KEY") == "" {
		return "", ErrApiKeyNotFound
	}

	ind := rand.Intn(4)

	term := terms[ind]
	if os.Getenv("TERM") != "" {
		term = os.Getenv("TERM")
	}

	url := fmt.Sprintf("https://shazam.p.rapidapi.com/search?term=%s&locale=en-US&offset=0&limit=1", term)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", os.Getenv("SHAZAM_RAPIDAPI_KEY"))
	req.Header.Add("X-RapidAPI-Host", "shazam.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		return "", err
	}

	data := jsonParsed.Path("tracks.hits.track.share.href").Data()

	return data, nil
}