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

//TODO 404だった時の処理追加
func (api *APIClient) doRequest(urlPath string, query map[string]string) (body []byte, err error) {
	url, _ := url.Parse(api.baseUrl)
	url.Path = path.Join(url.Path, urlPath)

	queryParams := url.Query()
	for key, value := range query {
		queryParams.Set(key, value)
	}
	url.RawQuery = queryParams.Encode()

	req, _ := http.NewRequest("GET", url.String(), nil)
	req.Header.Add("x-apisports-key", api.token)
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (api *APIClient) GetStatus() []byte {
	resp, _ := api.doRequest("status", map[string]string{})
	return resp
}

func (api *APIClient) GetLeagues() string {
	body, _ := api.doRequest("leagues", map[string]string{
		"code":   "IT",
		"season": "2021",
	})
	return body
}

func (api *APIClient) GetLeagueByLeagueId(leagueId string) string {
	body, _ := api.doRequest("leagues", map[string]string{
		"code":   "IT",
		"season": "2021",
		"id":     leagueId,
	})
	return body
}

func (api *APIClient) GetStandingsByLeagueId(leagueId string) string {
	body, _ := api.doRequest("standings", map[string]string{
		"season": "2021",
		"league": leagueId,
	})
	return body
}

func (api *APIClient) GetTopscorersByLeagueId(leagueId string) string {
	body, _ := api.doRequest("players/topscorers", map[string]string{
		"season": "2021",
		"league": leagueId,
	})
	return body
}

func (api *APIClient) GetTopassistsByLeagueId(leagueId string) string {
	body, _ := api.doRequest("players/topassists", map[string]string{
		"season": "2021",
		"league": leagueId,
	})
	return body
}
