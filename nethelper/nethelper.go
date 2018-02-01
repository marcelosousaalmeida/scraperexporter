package nethelper

import (
	"net/http"
	"log"
	"io/ioutil"
	"net/url"
	"strconv"
	"time"
)

func GetHttpBody(url string) string {
	request, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		log.Panicf("Error making HTTP request: %s", err)
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicf("Error reading HTTP request body: %s", err)
		return ""
	}

	return string(body)
}

func BuildURL(urlRaw string, avoidCache ...string) string {
	u, err  := url.Parse(urlRaw)

	if err != nil {
		log.Fatalf("Error building URL: %s", err)
	}

	if len(avoidCache) > 0 && avoidCache[0] == "true" {
		queryString := strconv.FormatInt(time.Now().Unix(), 10)
		q := u.Query()
		q.Set("z" + queryString, queryString)
		u.RawQuery = q.Encode()
	}

	return u.String()
}