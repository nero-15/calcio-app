package apifootball

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/nero-15/calcio-app/config"
)

type APIClient struct {
	token      string
	baseUrl    string
	httpClient *http.Client
}

func GetStatus() string {
	url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
	url.Path = path.Join(url.Path, "status")

	req, _ := http.NewRequest("GET", url.String(), nil)
	req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return string(byteArray)
}
