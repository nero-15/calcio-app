package apifootball

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
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

type Status struct {
	Get        string        `json:"get"`
	Parameters []interface{} `json:"parameters"`
	Errors     []interface{} `json:"errors"`
	Results    int           `json:"results"`
	Paging     struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response struct {
		Account struct {
			Firstname string `json:"firstname"`
			Lastname  string `json:"lastname"`
			Email     string `json:"email"`
		} `json:"account"`
		Subscription struct {
			Plan   string    `json:"plan"`
			End    time.Time `json:"end"`
			Active bool      `json:"active"`
		} `json:"subscription"`
		Requests struct {
			Current  int `json:"current"`
			LimitDay int `json:"limit_day"`
		} `json:"requests"`
	} `json:"response"`
}

func (api *APIClient) GetStatus() (Status, error) {
	resp, err := api.doRequest("status", map[string]string{})
	var status Status
	if err != nil {
		return status, err
	}
	json.Unmarshal(resp, &status)
	return status, nil
}

func (api *APIClient) GetLeagues() []byte {
	resp, _ := api.doRequest("leagues", map[string]string{
		"code":   "IT",
		"season": "2021",
	})
	return resp
}

func (api *APIClient) GetLeagueByLeagueId(leagueId string) []byte {
	resp, _ := api.doRequest("leagues", map[string]string{
		"code":   "IT",
		"season": "2021",
		"id":     leagueId,
	})
	return resp
}

func (api *APIClient) GetStandingsByLeagueId(leagueId string) []byte {
	resp, _ := api.doRequest("standings", map[string]string{
		"season": "2021",
		"league": leagueId,
	})
	return resp
}

func (api *APIClient) GetTopscorersByLeagueId(leagueId string) []byte {
	resp, _ := api.doRequest("players/topscorers", map[string]string{
		"season": "2021",
		"league": leagueId,
	})
	return resp
}

func (api *APIClient) GetTopassistsByLeagueId(leagueId string) []byte {
	resp, _ := api.doRequest("players/topassists", map[string]string{
		"season": "2021",
		"league": leagueId,
	})
	return resp
}
