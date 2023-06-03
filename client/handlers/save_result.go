package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/woeatory/Adamantite-TypeRacer/client/http_client"
	"io/ioutil"
	"net/http"
)

type Record struct {
	WPM      int `json:"WPM"`
	Accuracy int `json:"Accuracy"`
	Typos    int `json:"Typos"`
}

func SaveResult(wpm, accuracy, typos int) error {
	record := Record{
		WPM:      wpm,
		Accuracy: accuracy,
		Typos:    typos,
	}
	payload, err := json.Marshal(record)
	if err != nil {
		return err
	}
	var PATH = "http://" + http_client.ADDRESS + http_client.SaveResult

	req, err := http.NewRequest(http.MethodPost, PATH, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http_client.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New("error while creating new record")
	}
	// Read the response body
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}
