package tests

import (
	"context"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

// Prueba end to end para crear una materia
func TestEndToEndCreateSubject(t *testing.T) {
	// Crea el contexto de navegador
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Establece un timeout
	ctx, timeoutCancel := context.WithTimeout(ctx, 15*time.Second)
	defer timeoutCancel()

	// Mensaje de exito
	var successMessage string

	// Ejecuta las acciones de la prueba
	err := chromedp.Run(
		ctx,
		// Navega al formulario
		chromedp.Navigate("http://localhost:8080/"),
		// Espera que el campo ID este visible
		chromedp.WaitVisible(`#subjectId`, chromedp.ByID),
		// Llena los campos del formulario
		chromedp.SetValue(`#subjectId`,          "probando-id",          chromedp.ByID),
		chromedp.SetValue(`#subjectTitle`,       "probando título",      chromedp.ByID),
		chromedp.SetValue(`#subjectDescription`, "probando descripción", chromedp.ByID),
		// Hace clic en boton de enviar
		chromedp.Click(`#submitBtn`, chromedp.NodeVisible),
		// Espera el mensaje de exito
		chromedp.WaitVisible(`#successMessage`, chromedp.ByID),
		// Extrae el texto del mensaje
		chromedp.Text(`#successMessage`, &successMessage, chromedp.ByID),
	)

	// Valida si hubo fallas durante la prueba
	if err != nil {
		t.Fatal("Fallo la prueba end to end:", err)
	}

	// Valida que el mensaje de exito sea el esperado
	if successMessage != "Creado correctamente" {
		t.Errorf("Mensaje inesperado: '%s'", successMessage)
	}
}