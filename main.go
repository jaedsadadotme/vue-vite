package main

import (
	"io"
	"net/http"
	"text/template"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TemplateRenderer renders a template document
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("dist/*.html")),
	}
	e.Renderer = renderer
	e.Use(middleware.Static("dist"))
	e.GET("*", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{"title": c.Request().RequestURI})
	})
	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}
