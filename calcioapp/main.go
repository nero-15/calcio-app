package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
	})

	e.GET("/inter", func(c echo.Context) error {
		// 取得したいデータのURL作成
		url, _ := url.Parse(config.Config.FootballDataBaseUrl)
		url.Path = path.Join(url.Path, "teams", "108")

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("X-Auth-Token", config.Config.FootballDataApiToken) // アカウント登録時に送られてきたAPIトークンをリクエストヘッダーに追加
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(byteArray))
		return c.JSON(http.StatusOK, string(byteArray))
	})

	e.GET("/api/status", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "status")
		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(byteArray))

		return c.String(http.StatusOK, "api-football")
	})

	e.GET("/api/countries", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "countries")

		queryParams := url.Query()
		queryParams.Set("code", "IT")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)

		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/leagues", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "leagues")

		queryParams := url.Query()
		queryParams.Set("code", "IT")
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

	e.GET("/api/SerieA/teams", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "teams")

		queryParams := url.Query()
		queryParams.Set("league", "135")
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

	e.GET("/api/SerieB/teams", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "teams")

		queryParams := url.Query()
		queryParams.Set("league", "136")
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

	e.GET("/api/teams/statistics/inter", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "teams", "statistics")

		queryParams := url.Query()
		queryParams.Set("league", "135")
		queryParams.Set("season", "2021")
		queryParams.Set("team", "505")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)

		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/teams/seasons", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "teams", "seasons")

		queryParams := url.Query()
		queryParams.Set("team", "505")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)

		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/venues", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "venues")

		queryParams := url.Query()
		queryParams.Set("id", "907")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)

		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/standings/SerieA", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "standings")

		queryParams := url.Query()
		queryParams.Set("league", "135")
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

	e.GET("/api/standings/SerieB", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "standings")

		queryParams := url.Query()
		queryParams.Set("league", "136")
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

	e.GET("/api/fixtures/rounds", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "fixtures", "rounds")

		queryParams := url.Query()
		queryParams.Set("league", "135")
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

	e.GET("/api/fixtures", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "fixtures")

		queryParams := url.Query()
		queryParams.Set("league", "135")
		queryParams.Set("season", "2021")
		queryParams.Set("team", "505")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)

		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/fixtures/headtohead", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "fixtures", "headtohead")

		queryParams := url.Query()
		queryParams.Set("league", "135")
		queryParams.Set("season", "2021")
		queryParams.Set("h2h", "505-489")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)

		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/fixtures/statistics", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "fixtures", "statistics")

		queryParams := url.Query()
		queryParams.Set("fixture", "731698")
		queryParams.Set("team", "505")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)

		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/fixtures/events", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "fixtures", "events")

		queryParams := url.Query()
		queryParams.Set("fixture", "731698")
		queryParams.Set("team", "505")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)

		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/fixtures/lineups", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "fixtures", "lineups")

		queryParams := url.Query()
		queryParams.Set("fixture", "731698")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/fixtures/players", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "fixtures", "players")

		queryParams := url.Query()
		queryParams.Set("fixture", "731698")
		queryParams.Set("team", "505")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/injuries", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "injuries")

		queryParams := url.Query()
		queryParams.Set("season", "2021")
		queryParams.Set("team", "505")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/predictions", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "predictions")

		queryParams := url.Query()
		queryParams.Set("fixture", "731698")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/coachs", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "coachs")

		queryParams := url.Query()
		queryParams.Set("team", "505")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/players/seasons", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "players", "seasons")

		queryParams := url.Query()
		queryParams.Set("player", "217")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/players", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "players")

		queryParams := url.Query()
		queryParams.Set("id", "217")
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

	e.GET("/api/players/squads", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "players", "squads")

		queryParams := url.Query()
		queryParams.Set("team", "505")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/players/topscorers", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "players", "topscorers")

		queryParams := url.Query()
		queryParams.Set("league", "135")
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

	e.GET("/api/players/topassists", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "players", "topassists")

		queryParams := url.Query()
		queryParams.Set("league", "135")
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

	e.GET("/api/players/topyellowcards", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "players", "topyellowcards")

		queryParams := url.Query()
		queryParams.Set("league", "135")
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

	e.GET("/api/players/topredcards", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "players", "topredcards")

		queryParams := url.Query()
		queryParams.Set("league", "135")
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

	e.GET("/api/transfers", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "transfers")

		queryParams := url.Query()
		queryParams.Set("player", "30558")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/trophies", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "trophies")

		queryParams := url.Query()
		queryParams.Set("player", "30558")
		url.RawQuery = queryParams.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		req.Header.Add("x-apisports-key", config.Config.ApiFootballApiToken)
		client := new(http.Client)
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(byteArray))
	})

	e.GET("/api/sidelined", func(c echo.Context) error {
		url, _ := url.Parse(config.Config.ApiFootballBaseUrl)
		url.Path = path.Join(url.Path, "sidelined")

		queryParams := url.Query()
		queryParams.Set("player", "201")
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
