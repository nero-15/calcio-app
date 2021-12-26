package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

func main() {
	url, _ := url.Parse("https://api.football-data.org/v2/") //baseUrl
	url.Path = path.Join(url.Path, "matches")
	queryParams := url.Query()
	//queryParams.Set("hogehoge", "hugahuga")

	url.RawQuery = queryParams.Encode()
	req, _ := http.NewRequest("GET", url.String(), nil)
	req.Header.Add("X-Auth-Token", "Your API token")
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
}
