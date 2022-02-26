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

type League struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Logo    string `json:"logo"`
	Flag    string `json:"flag"`
	Season  int    `json:"season"`
}

type Country struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Flag string `json:"flag"`
}

type Season struct {
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
}

type Leagues struct {
	CommonResponse
	Response []struct {
		League  `json:"league"`
		Country `json:"country"`
		Seasons []Season `json:"seasons"`
	} `json:"response"`
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

type Player struct {
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
}

type Statistic struct {
	Team struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Logo string `json:"logo"`
	} `json:"team"`
	League `json:"league"`
	Games  struct {
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
}

type Topscorers struct {
	CommonResponse
	Response []struct {
		Player     `json:"player"`
		Statistics []Statistic `json:"statistics"`
	} `json:"response"`
}

type Topassists struct {
	CommonResponse
	Response []struct {
		Player     `json:"player"`
		Statistics []Statistic `json:"statistics"`
	} `json:"response"`
}

type Topyellowcards struct {
	CommonResponse
	Response []struct {
		Player
		Statistics []Statistic `json:"statistics"`
	} `json:"response"`
}

type Topredcards struct {
	CommonResponse
	Response []struct {
		Player
		Statistics []Statistic `json:"statistics"`
	} `json:"response"`
}

type Teams struct {
	CommonResponse
	Response []struct {
		Team struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Code     string `json:"code"`
			Country  string `json:"country"`
			Founded  int    `json:"founded"`
			National bool   `json:"national"`
			Logo     string `json:"logo"`
		} `json:"team"`
		Venue struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Address  string `json:"address"`
			City     string `json:"city"`
			Capacity int    `json:"capacity"`
			Surface  string `json:"surface"`
			Image    string `json:"image"`
		} `json:"venue"`
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

func (api *APIClient) GetTopyellowcardsByLeagueId(leagueId string) (Topyellowcards, error) {
	resp, err := api.doRequest("players/topyellowcards", map[string]string{
		"season": "2021",
		"league": leagueId,
	})
	var topyellowcards Topyellowcards
	if err != nil {
		return topyellowcards, err
	}
	json.Unmarshal(resp, &topyellowcards)
	return topyellowcards, nil
}

func (api *APIClient) GetTopredcardsByLeagueId(leagueId string) (Topredcards, error) {
	resp, err := api.doRequest("players/topredcards", map[string]string{
		"season": "2021",
		"league": leagueId,
	})
	var topyellowcards Topredcards
	if err != nil {
		return topyellowcards, err
	}
	json.Unmarshal(resp, &topyellowcards)
	return topyellowcards, nil
}

func (api *APIClient) GetTeamsByLeagueId(leagueId string) (Teams, error) {
	resp, err := api.doRequest("teams", map[string]string{
		"season": "2021",
		"league": leagueId,
	})
	var teams Teams
	if err != nil {
		return teams, err
	}
	json.Unmarshal(resp, &teams)
	return teams, nil
}

func (api *APIClient) GetTeamsByLeagueIdAndTeamId(leagueId string, teamId string) (Teams, error) {
	resp, err := api.doRequest("teams", map[string]string{
		"season": "2021",
		"league": leagueId,
		"id":     teamId,
	})
	var teams Teams
	if err != nil {
		return teams, err
	}
	json.Unmarshal(resp, &teams)
	return teams, nil
}

func (api *APIClient) GetStatisticsByLeagueIdAndTeamId(leagueId string, teamId string) (Statistics, error) {
	resp, err := api.doRequest("teams/statistics", map[string]string{
		"season": "2021",
		"league": leagueId,
		"team":   teamId,
	})
	var statistics Statistics
	if err != nil {
		return statistics, err
	}
	json.Unmarshal(resp, &statistics)
	return statistics, nil
}

type Statistics struct {
	CommonResponse
	Response struct {
		League `json:"league"`
		Team   struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Logo string `json:"logo"`
		} `json:"team"`
		Form     string `json:"form"`
		Fixtures struct {
			Played struct {
				Home  int `json:"home"`
				Away  int `json:"away"`
				Total int `json:"total"`
			} `json:"played"`
			Wins struct {
				Home  int `json:"home"`
				Away  int `json:"away"`
				Total int `json:"total"`
			} `json:"wins"`
			Draws struct {
				Home  int `json:"home"`
				Away  int `json:"away"`
				Total int `json:"total"`
			} `json:"draws"`
			Loses struct {
				Home  int `json:"home"`
				Away  int `json:"away"`
				Total int `json:"total"`
			} `json:"loses"`
		} `json:"fixtures"`
		Goals struct {
			For struct {
				Total struct {
					Home  int `json:"home"`
					Away  int `json:"away"`
					Total int `json:"total"`
				} `json:"total"`
				Average struct {
					Home  string `json:"home"`
					Away  string `json:"away"`
					Total string `json:"total"`
				} `json:"average"`
				Minute struct {
					Zero15 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"0-15"`
					One630 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"16-30"`
					Three145 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"31-45"`
					Four660 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"46-60"`
					Six175 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"61-75"`
					Seven690 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"76-90"`
					Nine1105 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"91-105"`
					One06120 struct {
						Total      interface{} `json:"total"`
						Percentage interface{} `json:"percentage"`
					} `json:"106-120"`
				} `json:"minute"`
			} `json:"for"`
			Against struct {
				Total struct {
					Home  int `json:"home"`
					Away  int `json:"away"`
					Total int `json:"total"`
				} `json:"total"`
				Average struct {
					Home  string `json:"home"`
					Away  string `json:"away"`
					Total string `json:"total"`
				} `json:"average"`
				Minute struct {
					Zero15 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"0-15"`
					One630 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"16-30"`
					Three145 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"31-45"`
					Four660 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"46-60"`
					Six175 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"61-75"`
					Seven690 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"76-90"`
					Nine1105 struct {
						Total      int    `json:"total"`
						Percentage string `json:"percentage"`
					} `json:"91-105"`
					One06120 struct {
						Total      interface{} `json:"total"`
						Percentage interface{} `json:"percentage"`
					} `json:"106-120"`
				} `json:"minute"`
			} `json:"against"`
		} `json:"goals"`
		Biggest struct {
			Streak struct {
				Wins  int `json:"wins"`
				Draws int `json:"draws"`
				Loses int `json:"loses"`
			} `json:"streak"`
			Wins struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"wins"`
			Loses struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"loses"`
			Goals struct {
				For struct {
					Home int `json:"home"`
					Away int `json:"away"`
				} `json:"for"`
				Against struct {
					Home int `json:"home"`
					Away int `json:"away"`
				} `json:"against"`
			} `json:"goals"`
		} `json:"biggest"`
		CleanSheet struct {
			Home  int `json:"home"`
			Away  int `json:"away"`
			Total int `json:"total"`
		} `json:"clean_sheet"`
		FailedToScore struct {
			Home  int `json:"home"`
			Away  int `json:"away"`
			Total int `json:"total"`
		} `json:"failed_to_score"`
		Penalty struct {
			Scored struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			} `json:"scored"`
			Missed struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			} `json:"missed"`
			Total int `json:"total"`
		} `json:"penalty"`
		Lineups []struct {
			Formation string `json:"formation"`
			Played    int    `json:"played"`
		} `json:"lineups"`
		Cards struct {
			Yellow struct {
				Zero15 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"0-15"`
				One630 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"16-30"`
				Three145 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"31-45"`
				Four660 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"46-60"`
				Six175 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"61-75"`
				Seven690 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"76-90"`
				Nine1105 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"91-105"`
				One06120 struct {
					Total      interface{} `json:"total"`
					Percentage interface{} `json:"percentage"`
				} `json:"106-120"`
			} `json:"yellow"`
			Red struct {
				Zero15 struct {
					Total      interface{} `json:"total"`
					Percentage interface{} `json:"percentage"`
				} `json:"0-15"`
				One630 struct {
					Total      interface{} `json:"total"`
					Percentage interface{} `json:"percentage"`
				} `json:"16-30"`
				Three145 struct {
					Total      interface{} `json:"total"`
					Percentage interface{} `json:"percentage"`
				} `json:"31-45"`
				Four660 struct {
					Total      interface{} `json:"total"`
					Percentage interface{} `json:"percentage"`
				} `json:"46-60"`
				Six175 struct {
					Total      interface{} `json:"total"`
					Percentage interface{} `json:"percentage"`
				} `json:"61-75"`
				Seven690 struct {
					Total      interface{} `json:"total"`
					Percentage interface{} `json:"percentage"`
				} `json:"76-90"`
				Nine1105 struct {
					Total      interface{} `json:"total"`
					Percentage interface{} `json:"percentage"`
				} `json:"91-105"`
				One06120 struct {
					Total      interface{} `json:"total"`
					Percentage interface{} `json:"percentage"`
				} `json:"106-120"`
			} `json:"red"`
		} `json:"cards"`
	} `json:"response"`
}

func (api *APIClient) GetPlayersByLeagueIdAndTeamId(leagueId string, teamId string) ([]byte, error) {
	resp, err := api.doRequest("players", map[string]string{
		"season": "2021",
		"league": leagueId,
		"team":   teamId,
	})
	return resp, err
}

func (api *APIClient) GetFixturesByLeagueIdAndTeamId(leagueId string, teamId string) ([]byte, error) {
	resp, err := api.doRequest("fixtures", map[string]string{
		"season": "2021",
		"league": leagueId,
		"team":   teamId,
	})
	return resp, err
}

func (api *APIClient) GetFixtureByFixtureId(fixtureId string) ([]byte, error) {
	resp, err := api.doRequest("fixtures", map[string]string{
		"id": fixtureId,
	})
	return resp, err
}

func (api *APIClient) GetInjuriesByLeagueIdAndTeamIdAndFixtureId(leagueId string, teamId string, fixtureId string) ([]byte, error) {
	resp, err := api.doRequest("injuries", map[string]string{
		"season":  "2021",
		"league":  leagueId,
		"team":    teamId,
		"fixture": fixtureId,
	})
	return resp, err
}

func (api *APIClient) GetStatisticsByTeamIdAndFixtureId(teamId string, fixtureId string) ([]byte, error) {
	resp, err := api.doRequest("fixtures/statistics", map[string]string{
		"team":    teamId,
		"fixture": fixtureId,
	})
	return resp, err
}

func (api *APIClient) GetEventsByTeamIdAndFixtureId(teamId string, fixtureId string) ([]byte, error) {
	resp, err := api.doRequest("fixtures/events", map[string]string{
		"team":    teamId,
		"fixture": fixtureId,
	})
	return resp, err
}

func (api *APIClient) GetLineupsByTeamIdAndFixtureId(teamId string, fixtureId string) ([]byte, error) {
	resp, err := api.doRequest("fixtures/lineups", map[string]string{
		"team":    teamId,
		"fixture": fixtureId,
	})
	return resp, err
}

func (api *APIClient) GetPlayersByTeamIdAndFixtureId(teamId string, fixtureId string) ([]byte, error) {
	resp, err := api.doRequest("fixtures/players", map[string]string{
		"team":    teamId,
		"fixture": fixtureId,
	})
	return resp, err
}

func (api *APIClient) GetCoachsByTeamId(teamId string) ([]byte, error) {
	resp, err := api.doRequest("coachs", map[string]string{
		"team": teamId,
	})
	return resp, err
}

func (api *APIClient) GetSquadsByTeamId(teamId string) ([]byte, error) {
	resp, err := api.doRequest("players/squads", map[string]string{
		"team": teamId,
	})
	return resp, err
}

func (api *APIClient) GetHeadtoheadByLeagueIdAndH2hId(leagueId string, h2hId string) ([]byte, error) {
	resp, err := api.doRequest("fixtures/headtohead", map[string]string{
		"league": leagueId,
		"h2h":    h2hId,
		"season": "2021",
	})
	return resp, err
}

func (api *APIClient) GetVenues() ([]byte, error) {
	resp, err := api.doRequest("venues", map[string]string{
		"country": "Italy",
	})
	return resp, err
}

func (api *APIClient) GetVenueByVenueId(venueId string) ([]byte, error) {
	resp, err := api.doRequest("venues", map[string]string{
		"country": "Italy",
		"id":      venueId,
	})
	return resp, err
}

func (api *APIClient) GetPredictionsByFixtureId(fixtureId string) ([]byte, error) {
	resp, err := api.doRequest("predictions", map[string]string{
		"fixture": fixtureId,
	})
	return resp, err
}

func (api *APIClient) GetPlayersByPlayerId(playerId string) ([]byte, error) {
	resp, err := api.doRequest("players", map[string]string{
		"id":     playerId,
		"season": "2021",
	})
	return resp, err
}

func (api *APIClient) GetTransfersByPlayerId(playerId string) ([]byte, error) {
	resp, err := api.doRequest("transfers", map[string]string{
		"player": playerId,
	})
	return resp, err
}

func (api *APIClient) GetTrophiesByPlayerId(playerId string) ([]byte, error) {
	resp, err := api.doRequest("trophies", map[string]string{
		"player": playerId,
	})
	return resp, err
}

func (api *APIClient) GetSidelinedByPlayerId(playerId string) ([]byte, error) {
	resp, err := api.doRequest("sidelined", map[string]string{
		"player": playerId,
	})
	return resp, err
}
