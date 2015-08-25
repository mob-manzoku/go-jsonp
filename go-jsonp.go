package gojsonp

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

var exceptCallback = regexp.MustCompile(`callback\((.*)\)`)
var jsonize = regexp.MustCompile("([0-9a-zA-Z]*):([0-9a-zA-Z]*)")

func GetStringFromJsonpString(str string) string {
	excepted := exceptCallback.FindAllStringSubmatch(str, -1)[0][1]

	return jsonize.ReplaceAllString(excepted, "\"${1}\":${2}")
}

func GetStringFromURL(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	str := string(body)

	return GetStringFromJsonpString(str), nil
}
