package hibp

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func reqHIBP(reqApi string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, reqApi, nil)
	if err != nil {
		return "", fmt.Errorf("request init failed\n%s",
			err.Error())
	}

	c := http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed\n%s",
			err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("Unexpected response status code received: %d %s",
			res.StatusCode, http.StatusText(res.StatusCode))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("reading res.Body failed\n%s",
			err.Error())
	}
	return string(body), err
}
