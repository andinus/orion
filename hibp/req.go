package hibp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func reqHIBP(reqApi string) (body string, err error) {
	c := http.Client{
		// TODO: timeout should be configurable by the user
		Timeout: time.Second * 64,
	}

	req, err := http.NewRequest(http.MethodGet, reqApi, nil)
	if err != nil {
		err = fmt.Errorf("hibp/req.go: Failed to create request\n%s",
			err.Error())
		return
	}

	// User-Agent should be passed with every request to
	// make work easier for the server handler. Include contact
	// information along with the project name so they could reach
	// you if required.
	req.Header.Set("User-Agent",
		"Andinus / Orion - https://andinus.nand.sh/orion")

	res, err := c.Do(req)
	if err != nil {
		err = fmt.Errorf("hibp/req.go: Failed to get response\n%s",
			err.Error())
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Errorf("hibp/req.go: Unexpected response status code received: %d %s",
			res.StatusCode,
			http.StatusText(res.StatusCode))
		return
	}

	// This will read everything to memory and is okay to use here
	// because the response is expected to be small.
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Errorf("hibp/req.go: Failed to read res.Body\n%s",
			err.Error())
		return
	}

	body = string(out)
	return
}
