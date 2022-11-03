package ip

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetCurrentIp() string {
	resp, err := http.Get("https://ifconfig.io/ip")
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		panic("Incorrect response")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s", data)
}
