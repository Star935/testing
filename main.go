package main

import (
	. "locustselenium/handlers"
	"github.com/labstack/echo/v4"
)

// Entrada del programa
func main() {
	// Genera instancia de echo
	e := echo.New()

	// Utiliza el html en la raiz
	e.GET("/", func(c echo.Context) error {
		return c.File("index.html")
	})

	// Rutas para los recursos web
	e.GET("/subjects", GetSubjects)
	e.POST("/subjects", CreateSubject)
	e.GET("/subject/:id", GetSubjectById)
	e.PUT("/subject/:id", UpdateSubject)
	e.DELETE("/subject/:id", DeleteSubject)

	e.Logger.Fatal(e.Start(":8080"))
}