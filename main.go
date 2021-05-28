package main

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/guardian/automat-v2/data"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

//go:embed assets templates
var files embed.FS

func main() {

	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       ".",
		Filesystem: http.FS(files),
	}))

	t := template.Must(template.ParseFS(files, "templates/*.go.html"))

	e.GET("/", func(c echo.Context) error {
		bodyHTML, err := execTemplate(t, "slots.go.html", data.Slots)
		if err != nil {
			return err
		}

		pageHTML, err := execTemplate(t, "page.go.html", template.HTML(bodyHTML))
		if err != nil {
			return err
		}

		return c.HTML(200, pageHTML)
	})

	port := getPort()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

func execTemplate(t *template.Template, name string, data interface{}) (string, error) {
	buf := bytes.Buffer{}
	err := t.ExecuteTemplate(&buf, name, data)
	return buf.String(), err
}

func getPort() int {
	port := 8080
	if p, ok := os.LookupEnv("PORT"); ok {
		parsed, err := strconv.Atoi(p)
		if err != nil {
			log.Fatalf("Fatal error: %v", err)
		}
		port = parsed
	}
	return port
}
