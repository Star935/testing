package tests

import (
	"encoding/json"
	"math/rand"
	"testing"
	"time"

	. "locustselenium/models"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

// Generador aleatorio de strings recibiendo un numero de la cantidad de caracteres
func randomString(characterSize int) string {
	// Opciones de letras para randomizar
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	// Crea un slice del string
	characterSlice := make([]rune, characterSize)
	// itera por la cantidad de caracteres
	for i := range characterSlice {
		characterSlice[i] = letters[rand.Intn(len(letters))]
	}

	// Retorna el slice de caracteres casteado a un string
	return string(characterSlice)
}

// Generador de peticiones
func generateTarget(t *testing.T) vegeta.Target {
	// Crea nueva materia
	subject := Subject{
		ID          : randomString(8),
		Title       : randomString(2),
		Description : randomString(15),
	}

	// Codifica la materia
	body, _ := json.Marshal(subject)

	// Imprime en la consola la materia generado
	t.Logf("Creando la materia: %+v", subject)

	// Retorna la ruta donde realizara las peticiones
	return vegeta.Target{
		Method : "POST",
		URL    : "http://localhost:8080/subjects",
		Body   : body,
		Header : map[string][]string{
			"Content-Type": {"application/json"},
		},
	}
}

// Prueba de performance para creacion de materias
func TestPerformanceCreateSubjects(t *testing.T) {
	// Define la tasa de creacion de materias
	rate := vegeta.Rate{
		// Frecuencia de creacion de materias
		Freq: 10, 
		// Tiempo entre creacion de materias
		Per : time.Second,
	}

	// Duracion de la prueba
	duration := 10 * time.Second

	// Crea nuevo atacante que generara las peticiones
	attacker := vegeta.NewAttacker()

	// Metricas de la prueba
	var metrics vegeta.Metrics

	// Define la forma en que se generara cada request
	targeter := func() vegeta.Targeter {
		return func(tgt *vegeta.Target) error {
			// Genera una peticion
			*tgt = generateTarget(t)
			return nil
		}
	}()

	// Realiza las peticiones
	for res := range attacker.Attack(targeter, rate, duration, "Comenzando la creacion de materias") {
		// Agrega los resultados de cada peticion a las metricas
		metrics.Add(res)
	}
	// Cierra las metricas despues de la ejecucion de las peticiones
	metrics.Close()

	// Imprime en la terminal del testing los resultados
	t.Logf("Resultados:")
	t.Logf("Peticiones totales: %d", metrics.Requests)
	t.Logf("Porcentaje de exito: %.2f%%", metrics.Success * 100)
	t.Logf("Latencia: %s", metrics.Latencies.Total)
	t.Logf("Peticiones por segundo: %.2f req/s", metrics.Throughput)
}