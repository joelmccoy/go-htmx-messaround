package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/joelmccoy/go-htmx-messaround/view/templates"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := templates.Layout()
		err := c.Render(r.Context(), w)

		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
	})

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", handler)
	e.GET("/login", handle_login)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func handler(c echo.Context) error {
	return render(c, templates.Layout())
}

func handle_login(c echo.Context) error {
	return render(c, templates.Login())
}

func render(c echo.Context, t templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return t.Render(c.Request().Context(), c.Response().Writer)
}
