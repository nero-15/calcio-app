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

func New() *APIClient {
	apiClient := &APIClient{config.Config.ApiFootballApiToken, config.Config.ApiFootballBaseUrl, &http.Client{}}
	return apiClient
}

func (api *APIClient) GetStatus() string {
	url, _ := url.Parse(api.baseUrl)
	url.Path = path.Join(url.Path, "status")

	req, _ := http.NewRequest("GET", url.String(), nil)
	req.Header.Add("x-apisports-key", api.token)
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return string(byteArray)
}
