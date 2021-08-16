package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"logkar-test/service"
	"logkar-test/service/model"
	"net/http"
	"os"
)

type soal3Usecase struct {
}

func NewSoal3Usecase() service.ISoal3Usecase {
	return soal3Usecase{}
}

func (s soal3Usecase) GetMovieFromOmdbAPI(page int, keyword string) (res model.SearchResponse, err error) {
	url := fmt.Sprintf("%s?apikey=%s&s=%s&page=%d", os.Getenv("omdbUrl"), os.Getenv("omdbKey"), keyword, page)

	resp, err := http.Get(url)
	if err != nil {
		err = errors.New("while get data from omdbapi")
		return res, err
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return res, err
	}

	err = json.Unmarshal(responseData, &res)
	if err != nil {
		log.Println(err)
		return res, err
	}

	return res, err
}
