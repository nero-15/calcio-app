package main

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/nero-15/calcio-app/apifootball"
	"github.com/nero-15/calcio-app/config"
	"github.com/nero-15/calcio-app/footballData"
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
		footballData := footballData.New(config.Config.FootballDataApiToken, config.Config.FootballDataBaseUrl)
		resp, _ := footballData.DoRequest("teams", c.Param("teamId")) //inter = 108
		return c.String(http.StatusOK, string(resp))
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
		teams, err := apifootball.GetTeamsByLeagueId(leagueId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		if teams.Results == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		teamsByteArray, _ := json.Marshal(teams)
		return c.String(http.StatusOK, string(teamsByteArray))
	})

	e.GET("/api/apiFootball/league/:leagueId/team/:teamId", func(c echo.Context) error {
		leagueId := c.Param("leagueId") //SerieA: 135, SerieB: 136
		teamId := c.Param("teamId")

		teams, err := apifootball.GetTeamsByLeagueIdAndTeamId(leagueId, teamId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		if teams.Results == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		teamsByteArray, _ := json.Marshal(teams)
		return c.String(http.StatusOK, string(teamsByteArray))
	})

	e.GET("/api/apiFootball/league/:leagueId/team/:teamId/statistics", func(c echo.Context) error {
		leagueId := c.Param("leagueId")
		teamId := c.Param("teamId") //inter: 505
		statistics, err := apifootball.GetStatisticsByLeagueIdAndTeamId(leagueId, teamId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		if statistics.Results == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		statisticsByteArray, _ := json.Marshal(statistics)
		return c.String(http.StatusOK, string(statisticsByteArray))
	})

	e.GET("/api/apiFootball/league/:leagueId/team/:teamId/players", func(c echo.Context) error {
		leagueId := c.Param("leagueId")
		teamId := c.Param("teamId")
		players, err := apifootball.GetPlayersByLeagueIdAndTeamId(leagueId, teamId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		if players.Results == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		playersByteArray, _ := json.Marshal(players)
		return c.String(http.StatusOK, string(playersByteArray))
	})

	e.GET("/api/apiFootball/league/:leagueId/team/:teamId/fixtures", func(c echo.Context) error {
		leagueId := c.Param("leagueId")
		teamId := c.Param("teamId")
		fixtures, err := apifootball.GetFixturesByLeagueIdAndTeamId(leagueId, teamId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		if fixtures.Results == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		fixturesByteArray, _ := json.Marshal(fixtures)
		return c.String(http.StatusOK, string(fixturesByteArray))
	})

	e.GET("/api/apiFootball/league/:leagueId/team/:teamId/fixture/:fixtureId", func(c echo.Context) error {
		fixtureId := c.Param("fixtureId")
		fixtures, err := apifootball.GetFixtureByFixtureId(fixtureId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		if fixtures.Results == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		fixturesByteArray, _ := json.Marshal(fixtures)
		return c.String(http.StatusOK, string(fixturesByteArray))
	})

	e.GET("/api/apiFootball/league/:leagueId/team/:teamId/fixture/:fixtureId/injuries", func(c echo.Context) error {
		leagueId := c.Param("leagueId")
		teamId := c.Param("teamId")
		fixtureId := c.Param("fixtureId")

		injuries, err := apifootball.GetInjuriesByLeagueIdAndTeamIdAndFixtureId(leagueId, teamId, fixtureId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		if injuries.Results == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		injuriesByteArray, _ := json.Marshal(injuries)
		return c.String(http.StatusOK, string(injuriesByteArray))
	})

	e.GET("/api/apiFootball/team/:teamId/fixture/:fixtureId/statistics", func(c echo.Context) error {
		teamId := c.Param("teamId")
		fixtureId := c.Param("fixtureId")

		fixturesStatistics, err := apifootball.GetStatisticsByTeamIdAndFixtureId(teamId, fixtureId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		if fixturesStatistics.Results == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		fixturesStatisticsByteArray, _ := json.Marshal(fixturesStatistics)
		return c.String(http.StatusOK, string(fixturesStatisticsByteArray))
	})

	e.GET("/api/apiFootball/team/:teamId/fixture/:fixtureId/events", func(c echo.Context) error {
		teamId := c.Param("teamId")
		fixtureId := c.Param("fixtureId") //731698

		events, err := apifootball.GetEventsByTeamIdAndFixtureId(teamId, fixtureId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		if events.Results == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		eventsByteArray, _ := json.Marshal(events)
		return c.String(http.StatusOK, string(eventsByteArray))
	})

	e.GET("/api/apiFootball/team/:teamId/fixture/:fixtureId/lineups", func(c echo.Context) error {
		teamId := c.Param("teamId")
		fixtureId := c.Param("fixtureId") //731698

		lineups, err := apifootball.GetLineupsByTeamIdAndFixtureId(teamId, fixtureId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		if lineups.Results == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		lineupsByteArray, _ := json.Marshal(lineups)
		return c.String(http.StatusOK, string(lineupsByteArray))
	})

	e.GET("/api/apiFootball/team/:teamId/fixture/:fixtureId/players", func(c echo.Context) error {
		teamId := c.Param("teamId")
		fixtureId := c.Param("fixtureId") //731698

		fixturesPlayers, err := apifootball.GetPlayersByTeamIdAndFixtureId(teamId, fixtureId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		if fixturesPlayers.Results == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		fixturesPlayersByteArray, _ := json.Marshal(fixturesPlayers)
		return c.String(http.StatusOK, string(fixturesPlayersByteArray))
	})

	e.GET("/api/apiFootball/team/:teamId/coachs", func(c echo.Context) error {
		teamId := c.Param("teamId")
		coachs, err := apifootball.GetCoachsByTeamId(teamId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		if coachs.Results == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		coachsByteArray, _ := json.Marshal(coachs)
		return c.String(http.StatusOK, string(coachsByteArray))
	})

	e.GET("/api/apiFootball/team/:teamId/squads", func(c echo.Context) error {
		teamId := c.Param("teamId")
		squads, _ := apifootball.GetSquadsByTeamId(teamId)
		squadsByteArray, _ := json.Marshal(squads)
		return c.String(http.StatusOK, string(squadsByteArray))
	})

	e.GET("/api/apiFootball/league/:leagueId/fixtures/headtohead/:h2h", func(c echo.Context) error {
		leagueId := c.Param("leagueId")
		h2hId := c.Param("h2h")
		resp, _ := apifootball.GetHeadtoheadByLeagueIdAndH2hId(leagueId, h2hId)
		return c.String(http.StatusOK, string(resp))
	})

	e.GET("/api/apiFootball/venues", func(c echo.Context) error {
		resp, _ := apifootball.GetVenues()
		return c.String(http.StatusOK, string(resp))
	})

	e.GET("/api/apiFootball/venue/:venueId", func(c echo.Context) error {
		venueId := c.Param("venueId") //Stadio Giuseppe Meazza: 907
		resp, _ := apifootball.GetVenueByVenueId(venueId)
		return c.String(http.StatusOK, string(resp))
	})

	e.GET("/api/apiFootball/predictions/:fixtureId", func(c echo.Context) error {
		fixtureId := c.Param("fixtureId")
		resp, _ := apifootball.GetPredictionsByFixtureId(fixtureId)
		return c.String(http.StatusOK, string(resp))
	})

	e.GET("/api/apiFootball/player/:playerId", func(c echo.Context) error {
		playerId := c.Param("playerId") // M. Škriniar: 198
		resp, _ := apifootball.GetPlayersByPlayerId(playerId)
		return c.String(http.StatusOK, string(resp))
	})

	e.GET("/api/apiFootball/player/:playerId/transfers", func(c echo.Context) error {
		playerId := c.Param("playerId")
		resp, _ := apifootball.GetTransfersByPlayerId(playerId)
		return c.String(http.StatusOK, string(resp))
	})

	e.GET("/api/apiFootball/player/:playerId/trophies", func(c echo.Context) error {
		playerId := c.Param("playerId")
		resp, _ := apifootball.GetTrophiesByPlayerId(playerId)
		return c.String(http.StatusOK, string(resp))
	})

	e.GET("/api/apiFootball/player/:playerId/sidelined", func(c echo.Context) error {
		playerId := c.Param("playerId")

		resp, _ := apifootball.GetSidelinedByPlayerId(playerId)
		return c.String(http.StatusOK, string(resp))
	})

	e.Logger.Fatal(e.Start(":8080"))
}
