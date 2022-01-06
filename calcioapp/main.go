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

	e.Logger.Fatal(e.Start(":8080"))
}
