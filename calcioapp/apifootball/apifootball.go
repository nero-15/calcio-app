package apifootball

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

type APIClient struct {
	token      string
	baseUrl    string
	httpClient *http.Client
}

func New(token string, baseUrl string) *APIClient {
	apiClient := &APIClient{token, baseUrl, &http.Client{}}
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
