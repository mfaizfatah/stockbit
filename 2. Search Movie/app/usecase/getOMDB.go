package usecase

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"project/app/config"

	"github.com/gomodul/envy"
)

func GetOMDB(searchword, pagination string) ([]byte, int, error) {
	url := fmt.Sprintf("%s/?apikey=%s&s=%s&page=%s", config.OmdbVar.Url, config.OmdbVar.ApiKey, searchword, pagination)
	log.Println(envy.Get("OMDB_APIKEY"), config.OmdbVar.ApiKey)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, res.StatusCode, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, res.StatusCode, err
	}

	return body, res.StatusCode, nil
}
