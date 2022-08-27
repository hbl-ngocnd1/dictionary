package infrastructure

import (
	"errors"
	"html/template"
	"io"
	"net/http"

	"github.com/hbl-ngocnd1/dictionary/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRegistry struct {
	templates map[string]*template.Template
}

// Render renders a template document
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func SetupServer() *echo.Echo {
	e := echo.New()
	templates := make(map[string]*template.Template)
	templates["home.html"] = template.Must(template.ParseFiles("public/views/home.html", "public/views/base.html"))
	templates["upload.html"] = template.Must(template.ParseFiles("public/views/upload.html", "public/views/base.html"))
	templates["dictionary.html"] = template.Must(template.ParseFiles("public/views/dictionary.html", "public/views/base.html"))

	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "./static")
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "home.html", map[string]interface{}{"router": "home"})
	})
	visitorHandler := controller.NewVisitorHandler(GetDB())
	e.POST("/api/visitors", visitorHandler.CreateVisitor)
	e.GET("/api/visitors", visitorHandler.ListVisitor)
	fileHandler := controller.NewFileHandler()
	e.GET("/upload", fileHandler.UploadFiles)
	e.POST("/api/upload", fileHandler.ApiUpload)
	dictHandler := controller.NewDictHandler()
	e.GET("/dictionary", dictHandler.Dict)
	e.GET("/api/dictionary", dictHandler.ApiDict)
	e.PUT("api/dictionary/:index", dictHandler.ApiGetDetail)
	return e
}
