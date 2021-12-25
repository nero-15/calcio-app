package main

import (
	"bytes"
	"log"
	"net/http"
	"net/url"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// // TemplateRenderer is a custom html/template renderer for Echo framework
// type TemplateRenderer struct {
// 	templates *template.Template
// }

// // Render renders a template document
// func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

// 	// Add global methods if data is a map
// 	if viewContext, isMap := data.(map[string]interface{}); isMap {
// 		viewContext["reverse"] = c.Echo().Reverse
// 	}

// 	return t.templates.ExecuteTemplate(w, name, data)
// }

const baseURL = "https://api.football-data.org/v2/"

func main() {
	e := echo.New()

	// renderer := &TemplateRenderer{
	// 	templates: template.Must(template.New("").Delims("[[", "]]").ParseGlob("views/*.html")), // vue.jsとdelimsがかぶるので変更
	// }
	// e.Renderer = renderer

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `"time":"${time_rfc3339}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status},"error":"${error}"` + "\n",
	}))
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")

		url, _ := url.Parse(baseURL)
		//url.Path = path.Join(url.Path, "search")
		queryParams := url.Query()
		queryParams.Set("hogehoge", "hugahuga")

		url.RawQuery = queryParams.Encode()
		resp, err := http.Get(url.String())
		if err != nil {
			log.Fatal(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		defer resp.Body.Close()

		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		jsonBytes := ([]byte)(buf.String())

		return c.JSON(http.StatusOK, jsonBytes)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
