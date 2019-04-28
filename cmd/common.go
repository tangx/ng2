package cmd

import (
	"io/ioutil"
	"net/http"
	"time"

	browser "github.com/EDDYCJY/fake-useragent"

	"github.com/sirupsen/logrus"
)

func httpGet2(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		logrus.Errorf("%v", err)
		return nil
	}
	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("%v", err)
	}

	return bodyByte
}

func httpGet(url string) []byte {
	client := &http.Client{Timeout: 2 * time.Second}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logrus.Fatalf("%s", err)
	}

	req.Header.Set("User-Agent", browser.Random())

	resp, err := client.Do(req)
	if err != nil {
		logrus.Fatalf("%s", err)
	}
	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("%v", err)
	}

	return bodyByte
}
