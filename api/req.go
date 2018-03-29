package api

import steamErr "github.com/edmundluk3/steam-api/error"
import (
	"net/http"
	"strings"
	"io/ioutil"
	"fmt"
	"log"
)


func ApiGet(apiUrl string, endPoint string, path string, version string, payload map[string]interface{}) ([]byte, error) {
	var payloadStr string

	for k, v := range payload {
		payloadStr += fmt.Sprintf("%s=%v&", k,v)
	}

	urlList := []string{apiUrl, endPoint, path, version, "?" + payloadStr}

	url := strings.Join(urlList, "")

	log.Print(url)

	resp, err := http.Get(url)

	if err != nil {
		// handle error
		return nil, err
	}

	code := resp.StatusCode

	if code < 200 || code > 299 {
		return nil, handleError(code)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body, nil
}

func ApiPost() {

}


func handleError (code int) error{
	for k, v := range steamErr.ErrorMap {
		if k == code {
			return &steamErr.SteamAPIError{
				Code: code,
				Msg: v,
			}
		}
	}

	return &steamErr.SteamAPIError{
		Code: code,
		Msg: steamErr.CommonErr,
	}
}
