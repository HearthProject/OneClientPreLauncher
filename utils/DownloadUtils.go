package utils

import (
	"os"
	"net/http"
	"io"
	"io/ioutil"
	"errors"
)

func DownloadFile(filepath string, url string) (err error) {
	out, err := os.Create(filepath)
	if err != nil  {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil  {
		return err
	}

	return nil
}

func GetString(url string) (string, error) {
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		return  "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 { // OK
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			return  "", err2
		}
		bodyString := string(bodyBytes)
		return bodyString, nil
	}

	return  "", errors.New("Failed to download file")
}
