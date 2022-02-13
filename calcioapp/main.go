package main

import (
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/nero-15/calcio-app/apifootball"
	"github.com/nero-15/calcio-app/config"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	renderer := &TemplateRenderer{
		templates: template.Must(template.New("").Delims("[[", "]]").ParseGlob("views/*.html")), // vue.jsとdelimsがかぶるので変更
	}
	e.Renderer = renderer

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `"time":"${time_rfc3339}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status},"error":"${error}"` + "\n",
	}))
	e.Use(middleware.Recover())

	apifootball := apifootball.New(config.Config.ApiFootballApiToken, config.Config.ApiFootballBaseUrl)

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
	})

	e.GET("api/footballData/teams/:teamId", func(c echo.Context) error {
		//inter = 108
		teamId := c.Param("teamId")

		// 取得したいデータのURL作成
		url, _ := url.Parse(config.Config.FootballDataBaseUrl)
		url.Path = path.Join(url.Path, "teams", teamId)

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("X-Auth-Token", config.Config.FootballDataApiToken) // アカウント登録時に送られてきたAPIトークンをリクエストヘッダーに追加
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/apiFootball/status", func(c echo.Context) error {
		status, err := apifootball.GetStatus()
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		if status.Results == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		statusByteArray, _ := json.Marshal(status)
		return c.String(http.StatusOK, string(statusByteArray)) //JSON形式で返す場合はc.JSON()メソッドを使う
	})

	e.GET("/api/apiFootball/leagues", func(c echo.Context) error {
		leagues, err := apifootball.GetLeagues()
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		if leagues.Results == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		leaguesByteArray, _ := json.Marshal(leagues)
		return c.String(http.StatusOK, string(leaguesByteArray))
	})

	e.GET("/api/apiFootball/league/:leagueId", func(c echo.Context) error {
		league, err := apifootball.GetLeagueByLeagueId(c.Param("leagueId"))
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		if league.Results == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		leagueByteArray, _ := json.Marshal(league)
		return c.String(http.StatusOK, string(leagueByteArray))
	})

	e.GET("/api/apiFootball/league/:leagueId/standings", func(c echo.Context) error {
		leagueId := c.Param("leagueId")
		standings, err := apifootball.GetStandingsByLeagueId(leagueId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		if standings.Results == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		standingsByteArray, _ := json.Marshal(standings)
		return c.String(http.StatusOK, string(standingsByteArray))
	})

	e.GET("/api/apiFootball/league/:leagueId/topscorers", func(c echo.Context) error {
		leagueId := c.Param("leagueId")
		topscorers, err := apifootball.GetTopscorersByLeagueId(leagueId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		if topscorers.Results == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		topscorersByteArray, _ := json.Marshal(topscorers)
		return c.String(http.StatusOK, string(topscorersByteArray))
	})

	e.GET("/api/apiFootball/league/:leagueId/topassists", func(c echo.Context) error {
		leagueId := c.Param("leagueId")
		topassists, err := apifootball.GetTopassistsByLeagueId(leagueId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		if topassists.Results == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		topassistsByteArray, _ := json.Marshal(topassists)
		return c.String(http.StatusOK, string(topassistsByteArray))
	})

	e.GET("/api/apiFootball/league/:leagueId/topyellowcards", func(c echo.Context) error {
		leagueId := c.Param("leagueId")
		topyellowcards, err := apifootball.GetTopyellowcardsByLeagueId(leagueId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		if topyellowcards.Results == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		topyellowcardsByteArray, _ := json.Marshal(topyellowcards)
		return c.String(http.StatusOK, string(topyellowcardsByteArray))
	})

	e.GET("/api/apiFootball/league/:leagueId/topredcards", func(c echo.Context) error {
		leagueId := c.Param("leagueId")
		topredcards, err := apifootball.GetTopredcardsByLeagueId(leagueId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		if topredcards.Results == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		topyellowcardsByteArray, _ := json.Marshal(topredcards)
		return c.String(http.StatusOK, string(topyellowcardsByteArray))
	})

	e.GET("/api/apiFootball/league/:leagueId/teams", func(c echo.Context) error {
		leagueId := c.Param("leagueId") //SerieA: 135, SerieB: 136
		teams, _ := apifootball.GetTeamsByLeagueId(leagueId)
		teamsByteArray, _ := json.Marshal(teams)
		return c.String(http.StatusOK, string(teamsByteArray))
	})

	e.GET("/api/apiFootball/league/:leagueId/team/:teamId", func(c echo.Context) error {
		leagueId := c.Param("leagueId") //SerieA: 135, SerieB: 136
		teamId := c.Param("teamId")

		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "teams")

		queryParams := url.Query()
		queryParams.Set("league", leagueId)
		queryParams.Set("team", teamId)
		queryParams.Set("season", "2021")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/apiFootball/league/:leagueId/team/:teamId/statistics", func(c echo.Context) error {
		leagueId := c.Param("leagueId")
		teamId := c.Param("teamId") //inter: 505

		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "teams", "statistics")

		queryParams := url.Query()
		queryParams.Set("league", leagueId)
		queryParams.Set("team", teamId)
		queryParams.Set("season", "2021")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/apiFootball/league/:leagueId/team/:teamId/players", func(c echo.Context) error {
		leagueId := c.Param("leagueId")
		teamId := c.Param("teamId")

		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "players")

		queryParams := url.Query()
		queryParams.Set("league", leagueId)
		queryParams.Set("team", teamId)
		queryParams.Set("season", "2021")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/apiFootball/league/:leagueId/team/:teamId/fixtures", func(c echo.Context) error {
		leagueId := c.Param("leagueId")
		teamId := c.Param("teamId")

		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "fixtures")

		queryParams := url.Query()
		queryParams.Set("league", leagueId)
		queryParams.Set("team", teamId)
		queryParams.Set("season", "2021")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/apiFootball/league/:leagueId/team/:teamId/fixture/:fixtureId", func(c echo.Context) error {
		leagueId := c.Param("leagueId")
		teamId := c.Param("teamId")
		fixtureId := c.Param("fixtureId")

		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "fixtures")

		queryParams := url.Query()
		queryParams.Set("id", fixtureId)
		queryParams.Set("league", leagueId)
		queryParams.Set("team", teamId)
		queryParams.Set("season", "2021")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/apiFootball/league/:leagueId/team/:teamId/fixture/:fixtureId/injuries", func(c echo.Context) error {
		leagueId := c.Param("leagueId")
		teamId := c.Param("teamId")
		fixtureId := c.Param("fixtureId")

		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "injuries")

		queryParams := url.Query()
		queryParams.Set("fixture", fixtureId)
		queryParams.Set("league", leagueId)
		queryParams.Set("team", teamId)
		queryParams.Set("season", "2021")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/apiFootball/team/:teamId/fixture/:fixtureId/statistics", func(c echo.Context) error {
		teamId := c.Param("teamId")
		fixtureId := c.Param("fixtureId")

		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "fixtures", "statistics")

		queryParams := url.Query()
		queryParams.Set("fixture", fixtureId)
		queryParams.Set("team", teamId)
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/apiFootball/team/:teamId/fixture/:fixtureId/events", func(c echo.Context) error {
		teamId := c.Param("teamId")
		fixtureId := c.Param("fixtureId") //731698

		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "fixtures", "events")

		queryParams := url.Query()
		queryParams.Set("fixture", fixtureId)
		queryParams.Set("team", teamId)
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/apiFootball/team/:teamId/fixture/:fixtureId/lineups", func(c echo.Context) error {
		teamId := c.Param("teamId")
		fixtureId := c.Param("fixtureId") //731698

		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "fixtures", "lineups")

		queryParams := url.Query()
		queryParams.Set("fixture", fixtureId)
		queryParams.Set("team", teamId)
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/apiFootball/team/:teamId/fixture/:fixtureId/players", func(c echo.Context) error {
		teamId := c.Param("teamId")
		fixtureId := c.Param("fixtureId") //731698

		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "fixtures", "players")

		queryParams := url.Query()
		queryParams.Set("fixture", fixtureId)
		queryParams.Set("team", teamId)
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/apiFootball/team/:teamId/coachs", func(c echo.Context) error {
		teamId := c.Param("teamId")

		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "coachs")

		queryParams := url.Query()
		queryParams.Set("team", teamId)
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/apiFootball/team/:teamId/squads", func(c echo.Context) error {
		teamId := c.Param("teamId")

		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "players", "squads")

		queryParams := url.Query()
		queryParams.Set("team", teamId)
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/apiFootball/league/:leagueId/fixtures/headtohead/:h2h", func(c echo.Context) error {
		leagueId := c.Param("leagueId")
		h2h := c.Param("h2h")

		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "fixtures", "headtohead")

		queryParams := url.Query()
		queryParams.Set("league", leagueId)
		queryParams.Set("h2h", h2h)
		queryParams.Set("season", "2021")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/apiFootball/venues", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "venues")

		queryParams := url.Query()
		queryParams.Set("country", "Italy")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/apiFootball/venue/:venueId", func(c echo.Context) error {
		venueId := c.Param("venueId") //Stadio Giuseppe Meazza: 907

		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "venues")

		queryParams := url.Query()
		queryParams.Set("id", venueId)
		queryParams.Set("country", "Italy")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/apiFootball/predictions/:fixtureId", func(c echo.Context) error {
		fixtureId := c.Param("fixtureId")

		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "predictions")

		queryParams := url.Query()
		queryParams.Set("fixture", fixtureId)
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/apiFootball/player/:playerId", func(c echo.Context) error {
		playerId := c.Param("playerId") // M. Škriniar: 198

		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "players")

		queryParams := url.Query()
		queryParams.Set("id", playerId)
		queryParams.Set("season", "2021")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/apiFootball/player/:playerId/transfers", func(c echo.Context) error {
		playerId := c.Param("playerId")

		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "transfers")

		queryParams := url.Query()
		queryParams.Set("player", playerId)
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/apiFootball/player/:playerId/trophies", func(c echo.Context) error {
		playerId := c.Param("playerId")

		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "trophies")

		queryParams := url.Query()
		queryParams.Set("player", playerId)
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/apiFootball/player/:playerId/sidelined", func(c echo.Context) error {
		playerId := c.Param("playerId")

		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "sidelined")

		queryParams := url.Query()
		queryParams.Set("player", playerId)
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.Logger.Fatal(e.Start(":8080"))
}
