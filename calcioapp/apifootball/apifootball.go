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

type CommonResponse struct {
	Get        string        `json:"get"`
	Parameters []interface{} `json:"parameters"`
	Errors     []interface{} `json:"errors"`
	Results    int           `json:"results"`
	Paging     struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
}

type Status struct {
	CommonResponse
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

type Leagues struct {
	CommonResponse
	Response []struct {
		League struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
			Logo string `json:"logo"`
		} `json:"league"`
		Country struct {
			Name string `json:"name"`
			Code string `json:"code"`
			Flag string `json:"flag"`
		} `json:"country"`
		Seasons []struct {
			Year     int    `json:"year"`
			Start    string `json:"start"`
			End      string `json:"end"`
			Current  bool   `json:"current"`
			Coverage struct {
				Fixtures struct {
					Events             bool `json:"events"`
					Lineups            bool `json:"lineups"`
					StatisticsFixtures bool `json:"statistics_fixtures"`
					StatisticsPlayers  bool `json:"statistics_players"`
				} `json:"fixtures"`
				Standings   bool `json:"standings"`
				Players     bool `json:"players"`
				TopScorers  bool `json:"top_scorers"`
				TopAssists  bool `json:"top_assists"`
				TopCards    bool `json:"top_cards"`
				Injuries    bool `json:"injuries"`
				Predictions bool `json:"predictions"`
				Odds        bool `json:"odds"`
			} `json:"coverage"`
		} `json:"seasons"`
	} `json:"response"`
}

func (api *APIClient) GetLeagues() (Leagues, error) {
	resp, err := api.doRequest("leagues", map[string]string{
		"code":   "IT",
		"season": "2021",
	})
	var leagues Leagues
	if err != nil {
		return leagues, err
	}
	json.Unmarshal(resp, &leagues)
	return leagues, nil
}

func (api *APIClient) GetLeagueByLeagueId(leagueId string) (Leagues, error) {
	resp, err := api.doRequest("leagues", map[string]string{
		"code":   "IT",
		"season": "2021",
		"id":     leagueId,
	})
	var leagues Leagues
	if err != nil {
		return leagues, err
	}
	json.Unmarshal(resp, &leagues)
	return leagues, nil
}

type Standings struct {
	CommonResponse
	Response []struct {
		League struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Country   string `json:"country"`
			Logo      string `json:"logo"`
			Flag      string `json:"flag"`
			Season    int    `json:"season"`
			Standings [][]struct {
				Rank int `json:"rank"`
				Team struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
					Logo string `json:"logo"`
				} `json:"team"`
				Points      int    `json:"points"`
				Goalsdiff   int    `json:"goalsDiff"`
				Group       string `json:"group"`
				Form        string `json:"form"`
				Status      string `json:"status"`
				Description string `json:"description"`
				All         struct {
					Played int `json:"played"`
					Win    int `json:"win"`
					Draw   int `json:"draw"`
					Lose   int `json:"lose"`
					Goals  struct {
						For     int `json:"for"`
						Against int `json:"against"`
					} `json:"goals"`
				} `json:"all"`
				Home struct {
					Played int `json:"played"`
					Win    int `json:"win"`
					Draw   int `json:"draw"`
					Lose   int `json:"lose"`
					Goals  struct {
						For     int `json:"for"`
						Against int `json:"against"`
					} `json:"goals"`
				} `json:"home"`
				Away struct {
					Played int `json:"played"`
					Win    int `json:"win"`
					Draw   int `json:"draw"`
					Lose   int `json:"lose"`
					Goals  struct {
						For     int `json:"for"`
						Against int `json:"against"`
					} `json:"goals"`
				} `json:"away"`
				Update time.Time `json:"update"`
			} `json:"standings"`
		} `json:"league"`
	} `json:"response"`
}

func (api *APIClient) GetStandingsByLeagueId(leagueId string) (Standings, error) {
	resp, err := api.doRequest("standings", map[string]string{
		"season": "2021",
		"league": leagueId,
	})
	var standings Standings
	if err != nil {
		return standings, err
	}
	json.Unmarshal(resp, &standings)
	return standings, nil
}

type Topscorers struct {
	CommonResponse
	Response []struct {
		Player struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Firstname string `json:"firstname"`
			Lastname  string `json:"lastname"`
			Age       int    `json:"age"`
			Birth     struct {
				Date    string `json:"date"`
				Place   string `json:"place"`
				Country string `json:"country"`
			} `json:"birth"`
			Nationality string `json:"nationality"`
			Height      string `json:"height"`
			Weight      string `json:"weight"`
			Injured     bool   `json:"injured"`
			Photo       string `json:"photo"`
		} `json:"player"`
		Statistics []struct {
			Team struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Logo string `json:"logo"`
			} `json:"team"`
			League struct {
				ID      int    `json:"id"`
				Name    string `json:"name"`
				Country string `json:"country"`
				Logo    string `json:"logo"`
				Flag    string `json:"flag"`
				Season  int    `json:"season"`
			} `json:"league"`
			Games struct {
				Appearences int         `json:"appearences"`
				Lineups     int         `json:"lineups"`
				Minutes     int         `json:"minutes"`
				Number      interface{} `json:"number"`
				Position    string      `json:"position"`
				Rating      string      `json:"rating"`
				Captain     bool        `json:"captain"`
			} `json:"games"`
			Substitutes struct {
				In    int `json:"in"`
				Out   int `json:"out"`
				Bench int `json:"bench"`
			} `json:"substitutes"`
			Shots struct {
				Total int `json:"total"`
				On    int `json:"on"`
			} `json:"shots"`
			Goals struct {
				Total    int         `json:"total"`
				Conceded int         `json:"conceded"`
				Assists  int         `json:"assists"`
				Saves    interface{} `json:"saves"`
			} `json:"goals"`
			Passes struct {
				Total    int `json:"total"`
				Key      int `json:"key"`
				Accuracy int `json:"accuracy"`
			} `json:"passes"`
			Tackles struct {
				Total         int         `json:"total"`
				Blocks        interface{} `json:"blocks"`
				Interceptions int         `json:"interceptions"`
			} `json:"tackles"`
			Duels struct {
				Total int `json:"total"`
				Won   int `json:"won"`
			} `json:"duels"`
			Dribbles struct {
				Attempts int         `json:"attempts"`
				Success  int         `json:"success"`
				Past     interface{} `json:"past"`
			} `json:"dribbles"`
			Fouls struct {
				Drawn     int `json:"drawn"`
				Committed int `json:"committed"`
			} `json:"fouls"`
			Cards struct {
				Yellow    int `json:"yellow"`
				Yellowred int `json:"yellowred"`
				Red       int `json:"red"`
			} `json:"cards"`
			Penalty struct {
				Won      interface{} `json:"won"`
				Commited interface{} `json:"commited"`
				Scored   int         `json:"scored"`
				Missed   int         `json:"missed"`
				Saved    interface{} `json:"saved"`
			} `json:"penalty"`
		} `json:"statistics"`
	} `json:"response"`
}

func (api *APIClient) GetTopscorersByLeagueId(leagueId string) (Topscorers, error) {
	resp, err := api.doRequest("players/topscorers", map[string]string{
		"season": "2021",
		"league": leagueId,
	})
	var topscorers Topscorers
	if err != nil {
		return topscorers, err
	}
	json.Unmarshal(resp, &topscorers)
	return topscorers, nil
}

type Topassists struct {
	CommonResponse
	Response []struct {
		Player struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Firstname string `json:"firstname"`
			Lastname  string `json:"lastname"`
			Age       int    `json:"age"`
			Birth     struct {
				Date    string `json:"date"`
				Place   string `json:"place"`
				Country string `json:"country"`
			} `json:"birth"`
			Nationality string `json:"nationality"`
			Height      string `json:"height"`
			Weight      string `json:"weight"`
			Injured     bool   `json:"injured"`
			Photo       string `json:"photo"`
		} `json:"player"`
		Statistics []struct {
			Team struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Logo string `json:"logo"`
			} `json:"team"`
			League struct {
				ID      int    `json:"id"`
				Name    string `json:"name"`
				Country string `json:"country"`
				Logo    string `json:"logo"`
				Flag    string `json:"flag"`
				Season  int    `json:"season"`
			} `json:"league"`
			Games struct {
				Appearences int         `json:"appearences"`
				Lineups     int         `json:"lineups"`
				Minutes     int         `json:"minutes"`
				Number      interface{} `json:"number"`
				Position    string      `json:"position"`
				Rating      string      `json:"rating"`
				Captain     bool        `json:"captain"`
			} `json:"games"`
			Substitutes struct {
				In    int `json:"in"`
				Out   int `json:"out"`
				Bench int `json:"bench"`
			} `json:"substitutes"`
			Shots struct {
				Total int `json:"total"`
				On    int `json:"on"`
			} `json:"shots"`
			Goals struct {
				Total    int         `json:"total"`
				Conceded int         `json:"conceded"`
				Assists  int         `json:"assists"`
				Saves    interface{} `json:"saves"`
			} `json:"goals"`
			Passes struct {
				Total    int `json:"total"`
				Key      int `json:"key"`
				Accuracy int `json:"accuracy"`
			} `json:"passes"`
			Tackles struct {
				Total         int `json:"total"`
				Blocks        int `json:"blocks"`
				Interceptions int `json:"interceptions"`
			} `json:"tackles"`
			Duels struct {
				Total int `json:"total"`
				Won   int `json:"won"`
			} `json:"duels"`
			Dribbles struct {
				Attempts int         `json:"attempts"`
				Success  int         `json:"success"`
				Past     interface{} `json:"past"`
			} `json:"dribbles"`
			Fouls struct {
				Drawn     int `json:"drawn"`
				Committed int `json:"committed"`
			} `json:"fouls"`
			Cards struct {
				Yellow    int `json:"yellow"`
				Yellowred int `json:"yellowred"`
				Red       int `json:"red"`
			} `json:"cards"`
			Penalty struct {
				Won      interface{} `json:"won"`
				Commited interface{} `json:"commited"`
				Scored   int         `json:"scored"`
				Missed   int         `json:"missed"`
				Saved    interface{} `json:"saved"`
			} `json:"penalty"`
		} `json:"statistics"`
	} `json:"response"`
}

func (api *APIClient) GetTopassistsByLeagueId(leagueId string) (Topassists, error) {
	resp, err := api.doRequest("players/topassists", map[string]string{
		"season": "2021",
		"league": leagueId,
	})
	var topassists Topassists
	if err != nil {
		return topassists, err
	}
	json.Unmarshal(resp, &topassists)
	return topassists, nil
}
