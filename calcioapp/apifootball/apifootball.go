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
	Team   `json:"team"`
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

type Fixtures struct {
	CommonResponse
	Response []struct {
		Fixture struct {
			ID        int       `json:"id"`
			Referee   string    `json:"referee"`
			Timezone  string    `json:"timezone"`
			Date      time.Time `json:"date"`
			Timestamp int       `json:"timestamp"`
			Periods   struct {
				First  int `json:"first"`
				Second int `json:"second"`
			} `json:"periods"`
			Venue struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				City string `json:"city"`
			} `json:"venue"`
			Status struct {
				Long    string `json:"long"`
				Short   string `json:"short"`
				Elapsed int    `json:"elapsed"`
			} `json:"status"`
		} `json:"fixture"`
		League `json:"league"`
		Teams  struct {
			Home struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Logo   string `json:"logo"`
				Winner bool   `json:"winner"`
			} `json:"home"`
			Away struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Logo   string `json:"logo"`
				Winner bool   `json:"winner"`
			} `json:"away"`
		} `json:"teams"`
		Goals struct {
			Home int `json:"home"`
			Away int `json:"away"`
		} `json:"goals"`
		Score struct {
			Halftime struct {
				Home int `json:"home"`
				Away int `json:"away"`
			} `json:"halftime"`
			Fulltime struct {
				Home int `json:"home"`
				Away int `json:"away"`
			} `json:"fulltime"`
			Extratime struct {
				Home interface{} `json:"home"`
				Away interface{} `json:"away"`
			} `json:"extratime"`
			Penalty struct {
				Home interface{} `json:"home"`
				Away interface{} `json:"away"`
			} `json:"penalty"`
		} `json:"score"`
	} `json:"response"`
}

type Injuries struct {
	CommonResponse
	Response []struct {
		Player  `json:"player"`
		Team    `json:"team"`
		Fixture struct {
			ID        int       `json:"id"`
			Timezone  string    `json:"timezone"`
			Date      time.Time `json:"date"`
			Timestamp int       `json:"timestamp"`
		} `json:"fixture"`
		League `json:"league"`
	} `json:"response"`
}

type League struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Logo    string `json:"logo"`
	Flag    string `json:"flag"`
	Season  int    `json:"season"`
	Round   string `json:"round"`
}

type Leagues struct {
	CommonResponse
	Response []struct {
		League  `json:"league"`
		Country `json:"country"`
		Seasons []Season `json:"seasons"`
	} `json:"response"`
}

type Minute struct {
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
}

type Players struct {
	CommonResponse
	Response []struct {
		Player     `json:"player"`
		Statistics []Statistic `json:"statistics"`
	} `json:"response"`
}

type Statistics struct {
	CommonResponse
	Response struct {
		League   `json:"league"`
		Team     `json:"team"`
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
				Minute `json:"minute"`
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
				Minute `json:"minute"`
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
			Yellow Minute `json:"yellow"`
			Red    Minute `json:"red"`
		} `json:"cards"`
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

type Team struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	Country  string `json:"country"`
	Founded  int    `json:"founded"`
	National bool   `json:"national"`
	Logo     string `json:"logo"`
}

type Teams struct {
	CommonResponse
	Response []struct {
		Team  `json:"team"`
		Venue `json:"venue"`
	} `json:"response"`
}

type Trophies struct {
	CommonResponse
	Response []struct {
		Trophy
	} `json:"response"`
}

type Trophy struct {
	League  string `json:"league"`
	Country string `json:"country"`
	Season  string `json:"season"`
	Place   string `json:"place"`
}

type Transfers struct {
	CommonResponse
	Response []struct {
		Player struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"player"`
		Update    time.Time  `json:"update"`
		Transfers []Transfer `json:"transfers"`
	} `json:"response"`
}

type Transfer struct {
	Date  string `json:"date"`
	Type  string `json:"type"`
	Teams struct {
		In struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Logo string `json:"logo"`
		} `json:"in"`
		Out struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Logo string `json:"logo"`
		} `json:"out"`
	} `json:"teams"`
}

type Venue struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Capacity int    `json:"capacity"`
	Surface  string `json:"surface"`
	Image    string `json:"image"`
}

type Venues struct {
	CommonResponse
	Response []struct {
		Venue
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

func (api *APIClient) GetPlayersByLeagueIdAndTeamId(leagueId string, teamId string) (Players, error) {
	resp, err := api.doRequest("players", map[string]string{
		"season": "2021",
		"league": leagueId,
		"team":   teamId,
	})
	var players Players
	if err != nil {
		return players, err
	}
	json.Unmarshal(resp, &players)
	return players, nil
}

func (api *APIClient) GetFixturesByLeagueIdAndTeamId(leagueId string, teamId string) (Fixtures, error) {
	resp, err := api.doRequest("fixtures", map[string]string{
		"season": "2021",
		"league": leagueId,
		"team":   teamId,
	})
	var fixtures Fixtures
	if err != nil {
		return fixtures, err
	}
	json.Unmarshal(resp, &fixtures)
	return fixtures, nil
}

func (api *APIClient) GetFixtureByFixtureId(fixtureId string) (Fixtures, error) {
	resp, err := api.doRequest("fixtures", map[string]string{
		"id": fixtureId,
	})
	var fixtures Fixtures
	if err != nil {
		return fixtures, err
	}
	json.Unmarshal(resp, &fixtures)
	return fixtures, nil
}

func (api *APIClient) GetInjuriesByLeagueIdAndTeamIdAndFixtureId(leagueId string, teamId string, fixtureId string) (Injuries, error) {
	resp, err := api.doRequest("injuries", map[string]string{
		"season":  "2021",
		"league":  leagueId,
		"team":    teamId,
		"fixture": fixtureId,
	})
	var injuries Injuries
	if err != nil {
		return injuries, err
	}
	json.Unmarshal(resp, &injuries)
	return injuries, nil
}

func (api *APIClient) GetStatisticsByTeamIdAndFixtureId(teamId string, fixtureId string) (FixturesStatistics, error) {
	resp, err := api.doRequest("fixtures/statistics", map[string]string{
		"team":    teamId,
		"fixture": fixtureId,
	})
	var fixturesStatistics FixturesStatistics
	if err != nil {
		return fixturesStatistics, err
	}
	json.Unmarshal(resp, &fixturesStatistics)
	return fixturesStatistics, nil
}

type FixturesStatistics struct {
	Get        string `json:"get"`
	Parameters struct {
		Fixture string `json:"fixture"`
		Team    string `json:"team"`
	} `json:"parameters"`
	Errors  []interface{} `json:"errors"`
	Results int           `json:"results"`
	Paging  struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []struct {
		Team struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Logo string `json:"logo"`
		} `json:"team"`
		Statistics []struct {
			Type  string `json:"type"`
			Value int    `json:"value"`
		} `json:"statistics"`
	} `json:"response"`
}

func (api *APIClient) GetEventsByTeamIdAndFixtureId(teamId string, fixtureId string) (Events, error) {
	resp, err := api.doRequest("fixtures/events", map[string]string{
		"team":    teamId,
		"fixture": fixtureId,
	})
	var events Events
	if err != nil {
		return events, err
	}
	json.Unmarshal(resp, &events)
	return events, nil
}

type Events struct {
	Get        string `json:"get"`
	Parameters struct {
		Fixture string `json:"fixture"`
		Team    string `json:"team"`
	} `json:"parameters"`
	Errors  []interface{} `json:"errors"`
	Results int           `json:"results"`
	Paging  struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []struct {
		Time struct {
			Elapsed int         `json:"elapsed"`
			Extra   interface{} `json:"extra"`
		} `json:"time"`
		Team struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Logo string `json:"logo"`
		} `json:"team"`
		Player struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"player"`
		Assist struct {
			ID   interface{} `json:"id"`
			Name interface{} `json:"name"`
		} `json:"assist"`
		Type     string      `json:"type"`
		Detail   string      `json:"detail"`
		Comments interface{} `json:"comments"`
	} `json:"response"`
}

func (api *APIClient) GetLineupsByTeamIdAndFixtureId(teamId string, fixtureId string) (Lineups, error) {
	resp, err := api.doRequest("fixtures/lineups", map[string]string{
		"team":    teamId,
		"fixture": fixtureId,
	})
	var lineups Lineups
	if err != nil {
		return lineups, err
	}
	json.Unmarshal(resp, &lineups)
	return lineups, nil
}

type Lineups struct {
	Get        string `json:"get"`
	Parameters struct {
		Fixture string `json:"fixture"`
		Team    string `json:"team"`
	} `json:"parameters"`
	Errors  []interface{} `json:"errors"`
	Results int           `json:"results"`
	Paging  struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []struct {
		Team struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Logo   string `json:"logo"`
			Colors struct {
				Player struct {
					Primary string `json:"primary"`
					Number  string `json:"number"`
					Border  string `json:"border"`
				} `json:"player"`
				Goalkeeper struct {
					Primary string `json:"primary"`
					Number  string `json:"number"`
					Border  string `json:"border"`
				} `json:"goalkeeper"`
			} `json:"colors"`
		} `json:"team"`
		Coach struct {
			ID    int    `json:"id"`
			Name  string `json:"name"`
			Photo string `json:"photo"`
		} `json:"coach"`
		Formation string `json:"formation"`
		Startxi   []struct {
			Player struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Number int    `json:"number"`
				Pos    string `json:"pos"`
				Grid   string `json:"grid"`
			} `json:"player"`
		} `json:"startXI"`
		Substitutes []struct {
			Player struct {
				ID     int         `json:"id"`
				Name   string      `json:"name"`
				Number int         `json:"number"`
				Pos    string      `json:"pos"`
				Grid   interface{} `json:"grid"`
			} `json:"player"`
		} `json:"substitutes"`
	} `json:"response"`
}

func (api *APIClient) GetPlayersByTeamIdAndFixtureId(teamId string, fixtureId string) (FixturesPlayers, error) {
	resp, err := api.doRequest("fixtures/players", map[string]string{
		"team":    teamId,
		"fixture": fixtureId,
	})
	var fixturesPlayers FixturesPlayers
	if err != nil {
		return fixturesPlayers, err
	}
	json.Unmarshal(resp, &fixturesPlayers)
	return fixturesPlayers, nil
}

type FixturesPlayers struct {
	Get        string `json:"get"`
	Parameters struct {
		Fixture string `json:"fixture"`
		Team    string `json:"team"`
	} `json:"parameters"`
	Errors  []interface{} `json:"errors"`
	Results int           `json:"results"`
	Paging  struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []struct {
		Team struct {
			ID     int       `json:"id"`
			Name   string    `json:"name"`
			Logo   string    `json:"logo"`
			Update time.Time `json:"update"`
		} `json:"team"`
		Players []struct {
			Player struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Photo string `json:"photo"`
			} `json:"player"`
			Statistics []struct {
				Games struct {
					Minutes    int    `json:"minutes"`
					Number     int    `json:"number"`
					Position   string `json:"position"`
					Rating     string `json:"rating"`
					Captain    bool   `json:"captain"`
					Substitute bool   `json:"substitute"`
				} `json:"games"`
				Offsides interface{} `json:"offsides"`
				Shots    struct {
					Total interface{} `json:"total"`
					On    interface{} `json:"on"`
				} `json:"shots"`
				Goals struct {
					Total    interface{} `json:"total"`
					Conceded int         `json:"conceded"`
					Assists  interface{} `json:"assists"`
					Saves    interface{} `json:"saves"`
				} `json:"goals"`
				Passes struct {
					Total    int         `json:"total"`
					Key      interface{} `json:"key"`
					Accuracy string      `json:"accuracy"`
				} `json:"passes"`
				Tackles struct {
					Total         interface{} `json:"total"`
					Blocks        interface{} `json:"blocks"`
					Interceptions interface{} `json:"interceptions"`
				} `json:"tackles"`
				Duels struct {
					Total interface{} `json:"total"`
					Won   interface{} `json:"won"`
				} `json:"duels"`
				Dribbles struct {
					Attempts interface{} `json:"attempts"`
					Success  interface{} `json:"success"`
					Past     interface{} `json:"past"`
				} `json:"dribbles"`
				Fouls struct {
					Drawn     interface{} `json:"drawn"`
					Committed interface{} `json:"committed"`
				} `json:"fouls"`
				Cards struct {
					Yellow int `json:"yellow"`
					Red    int `json:"red"`
				} `json:"cards"`
				Penalty struct {
					Won      interface{} `json:"won"`
					Commited interface{} `json:"commited"`
					Scored   int         `json:"scored"`
					Missed   int         `json:"missed"`
					Saved    int         `json:"saved"`
				} `json:"penalty"`
			} `json:"statistics"`
		} `json:"players"`
	} `json:"response"`
}

func (api *APIClient) GetCoachsByTeamId(teamId string) (Coachs, error) {
	resp, err := api.doRequest("coachs", map[string]string{
		"team": teamId,
	})
	var coachs Coachs
	if err != nil {
		return coachs, err
	}
	json.Unmarshal(resp, &coachs)
	return coachs, nil
}

type Coachs struct {
	Get        string `json:"get"`
	Parameters struct {
		Team string `json:"team"`
	} `json:"parameters"`
	Errors  []interface{} `json:"errors"`
	Results int           `json:"results"`
	Paging  struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Age       int    `json:"age"`
		Birth     struct {
			Date    string      `json:"date"`
			Place   interface{} `json:"place"`
			Country string      `json:"country"`
		} `json:"birth"`
		Nationality string      `json:"nationality"`
		Height      interface{} `json:"height"`
		Weight      interface{} `json:"weight"`
		Photo       string      `json:"photo"`
		Team        struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Logo string `json:"logo"`
		} `json:"team"`
		Career []struct {
			Team struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Logo string `json:"logo"`
			} `json:"team"`
			Start string `json:"start"`
			End   string `json:"end"`
		} `json:"career"`
	} `json:"response"`
}

func (api *APIClient) GetSquadsByTeamId(teamId string) (Squads, error) {
	resp, err := api.doRequest("players/squads", map[string]string{
		"team": teamId,
	})
	var squads Squads
	if err != nil {
		return squads, err
	}
	json.Unmarshal(resp, &squads)
	return squads, nil
}

type Squads struct {
	Get        string `json:"get"`
	Parameters struct {
		Team string `json:"team"`
	} `json:"parameters"`
	Errors  []interface{} `json:"errors"`
	Results int           `json:"results"`
	Paging  struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []struct {
		Team struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Logo string `json:"logo"`
		} `json:"team"`
		Players []struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Age      int    `json:"age"`
			Number   int    `json:"number"`
			Position string `json:"position"`
			Photo    string `json:"photo"`
		} `json:"players"`
	} `json:"response"`
}

func (api *APIClient) GetHeadtoheadByLeagueIdAndH2hId(leagueId string, h2hId string) ([]byte, error) {
	resp, err := api.doRequest("fixtures/headtohead", map[string]string{
		"league": leagueId,
		"h2h":    h2hId,
		"season": "2021",
	})
	return resp, err
}

func (api *APIClient) GetVenues() (Venues, error) {
	resp, err := api.doRequest("venues", map[string]string{
		"country": "Italy",
	})
	var venues Venues
	if err != nil {
		return venues, err
	}
	json.Unmarshal(resp, &venues)
	return venues, nil
}

func (api *APIClient) GetVenueByVenueId(venueId string) (Venues, error) {
	resp, err := api.doRequest("venues", map[string]string{
		"country": "Italy",
		"id":      venueId,
	})
	var venues Venues
	if err != nil {
		return venues, err
	}
	json.Unmarshal(resp, &venues)
	return venues, nil
}

func (api *APIClient) GetPredictionsByFixtureId(fixtureId string) (Predictions, error) {
	resp, err := api.doRequest("predictions", map[string]string{
		"fixture": fixtureId,
	})
	var predictions Predictions
	if err != nil {
		return predictions, err
	}
	json.Unmarshal(resp, &predictions)
	return predictions, nil
}

type Predictions struct {
	Get        string `json:"get"`
	Parameters struct {
		Fixture string `json:"fixture"`
	} `json:"parameters"`
	Errors  []interface{} `json:"errors"`
	Results int           `json:"results"`
	Paging  struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []struct {
		Predictions struct {
			Winner struct {
				ID      int    `json:"id"`
				Name    string `json:"name"`
				Comment string `json:"comment"`
			} `json:"winner"`
			WinOrDraw bool        `json:"win_or_draw"`
			UnderOver interface{} `json:"under_over"`
			Goals     struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"goals"`
			Advice  string `json:"advice"`
			Percent struct {
				Home string `json:"home"`
				Draw string `json:"draw"`
				Away string `json:"away"`
			} `json:"percent"`
		} `json:"predictions"`
		League struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Country string `json:"country"`
			Logo    string `json:"logo"`
			Flag    string `json:"flag"`
			Season  int    `json:"season"`
		} `json:"league"`
		Teams struct {
			Home struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Logo  string `json:"logo"`
				Last5 struct {
					Form  string `json:"form"`
					Att   string `json:"att"`
					Def   string `json:"def"`
					Goals struct {
						For struct {
							Total   int    `json:"total"`
							Average string `json:"average"`
						} `json:"for"`
						Against struct {
							Total   int    `json:"total"`
							Average string `json:"average"`
						} `json:"against"`
					} `json:"goals"`
				} `json:"last_5"`
				League struct {
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
									Total      interface{} `json:"total"`
									Percentage interface{} `json:"percentage"`
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
									Total      interface{} `json:"total"`
									Percentage interface{} `json:"percentage"`
								} `json:"31-45"`
								Four660 struct {
									Total      int    `json:"total"`
									Percentage string `json:"percentage"`
								} `json:"46-60"`
								Six175 struct {
									Total      interface{} `json:"total"`
									Percentage interface{} `json:"percentage"`
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
							Home interface{} `json:"home"`
							Away interface{} `json:"away"`
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
								Total      int    `json:"total"`
								Percentage string `json:"percentage"`
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
				} `json:"league"`
			} `json:"home"`
			Away struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Logo  string `json:"logo"`
				Last5 struct {
					Form  string `json:"form"`
					Att   string `json:"att"`
					Def   string `json:"def"`
					Goals struct {
						For struct {
							Total   int    `json:"total"`
							Average string `json:"average"`
						} `json:"for"`
						Against struct {
							Total   int    `json:"total"`
							Average string `json:"average"`
						} `json:"against"`
					} `json:"goals"`
				} `json:"last_5"`
				League struct {
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
							Home interface{} `json:"home"`
							Away string      `json:"away"`
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
				} `json:"league"`
			} `json:"away"`
		} `json:"teams"`
		Comparison struct {
			Form struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"form"`
			Att struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"att"`
			Def struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"def"`
			PoissonDistribution struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"poisson_distribution"`
			H2H struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"h2h"`
			Goals struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"goals"`
			Total struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"total"`
		} `json:"comparison"`
		H2H []struct {
			Fixture struct {
				ID        int       `json:"id"`
				Referee   string    `json:"referee"`
				Timezone  string    `json:"timezone"`
				Date      time.Time `json:"date"`
				Timestamp int       `json:"timestamp"`
				Periods   struct {
					First  int `json:"first"`
					Second int `json:"second"`
				} `json:"periods"`
				Venue struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
					City string `json:"city"`
				} `json:"venue"`
				Status struct {
					Long    string `json:"long"`
					Short   string `json:"short"`
					Elapsed int    `json:"elapsed"`
				} `json:"status"`
			} `json:"fixture"`
			League struct {
				ID      int    `json:"id"`
				Name    string `json:"name"`
				Country string `json:"country"`
				Logo    string `json:"logo"`
				Flag    string `json:"flag"`
				Season  int    `json:"season"`
				Round   string `json:"round"`
			} `json:"league"`
			Teams struct {
				Home struct {
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Logo   string `json:"logo"`
					Winner bool   `json:"winner"`
				} `json:"home"`
				Away struct {
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Logo   string `json:"logo"`
					Winner bool   `json:"winner"`
				} `json:"away"`
			} `json:"teams"`
			Goals struct {
				Home int `json:"home"`
				Away int `json:"away"`
			} `json:"goals"`
			Score struct {
				Halftime struct {
					Home int `json:"home"`
					Away int `json:"away"`
				} `json:"halftime"`
				Fulltime struct {
					Home int `json:"home"`
					Away int `json:"away"`
				} `json:"fulltime"`
				Extratime struct {
					Home interface{} `json:"home"`
					Away interface{} `json:"away"`
				} `json:"extratime"`
				Penalty struct {
					Home interface{} `json:"home"`
					Away interface{} `json:"away"`
				} `json:"penalty"`
			} `json:"score"`
		} `json:"h2h"`
	} `json:"response"`
}

func (api *APIClient) GetPlayersByPlayerId(playerId string) (Players2, error) {
	resp, err := api.doRequest("players", map[string]string{
		"id":     playerId,
		"season": "2021",
	})
	var players Players2
	if err != nil {
		return players, err
	}
	json.Unmarshal(resp, &players)
	return players, nil
}

type Players2 struct {
	Get        string `json:"get"`
	Parameters struct {
		ID     string `json:"id"`
		Season string `json:"season"`
	} `json:"parameters"`
	Errors  []interface{} `json:"errors"`
	Results int           `json:"results"`
	Paging  struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
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
				Assists  interface{} `json:"assists"`
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

func (api *APIClient) GetTransfersByPlayerId(playerId string) (Transfers, error) {
	resp, err := api.doRequest("transfers", map[string]string{
		"player": playerId,
	})
	var transfers Transfers
	if err != nil {
		return transfers, err
	}
	json.Unmarshal(resp, &transfers)
	return transfers, nil
}

func (api *APIClient) GetTrophiesByPlayerId(playerId string) (Trophies, error) {
	resp, err := api.doRequest("trophies", map[string]string{
		"player": playerId,
	})
	var trophies Trophies
	if err != nil {
		return trophies, err
	}
	json.Unmarshal(resp, &trophies)
	return trophies, nil
}

func (api *APIClient) GetSidelinedByPlayerId(playerId string) ([]byte, error) {
	resp, err := api.doRequest("sidelined", map[string]string{
		"player": playerId,
	})
	return resp, err
}
