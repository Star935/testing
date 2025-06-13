package handlers

import (
	. "locustselenium/models"
	"sync"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Mapa de materias
var subjects = map[string]Subject{}

// Controlador de GoRoutines
var lock = sync.Mutex{}

// Metodo para recuperar todas las materias
func GetSubjects(c echo.Context) error {
	// Controla los goroutines para que se ejecute uno a la vez y evitar corrupciones en los datos
	lock.Lock()

	// Libera los goroutines cuando finalice la ejecucion del metodo
	defer lock.Unlock()

	// Instancia de arreglo de materias
	var subjectList []Subject

	// Itera sobre el mapa de materias
	for _, subject := range subjects {
		// Añade la materia en el arreglo de material
		subjectList = append(subjectList, subject)
	}

	// Retorna el arreglo de materias
	return c.JSON(http.StatusOK, subjectList)
}

// Metodo para recuperar materia por su id
func GetSubjectById(c echo.Context) error {
	// Controla los goroutines para que se ejecute uno a la vez y evitar corrupciones en los datos
	lock.Lock()

	// Libera los goroutines cuando finalice la ejecucion del metodo
	defer lock.Unlock()

	// Recupera el id del parametro de la url
	id := c.Param("id")

	// Busca la materia con base al id recuperado de la peticion
	subject, ok := subjects[id]
	// Valida si la encontro
	if !ok {
		return c.NoContent(http.StatusNotFound)
	}

	// Retorna la materia recuperada
	return c.JSON(http.StatusCreated, subject)
}

// Metodo para crear nuevas materias
func CreateSubject(c echo.Context) error {
	// Controla los goroutines para que se ejecute uno a la vez y evitar corrupciones en los datos
	lock.Lock()

	// Libera los goroutines cuando finalice la ejecucion del metodo
	defer lock.Unlock()

	// Instancia de materia
	var subject Subject

	// Agrega los datos de la materia en la instancia, en su espacio de memoria 
	if err := c.Bind(&subject); err != nil {
		return err
	}

	// Agrega al mapa de materias la nueva materia
	subjects[subject.ID] = subject

	// Retorna la materia creada
	return c.JSON(http.StatusCreated, subject)
}

// Metodo para actualizar una materia existente
func UpdateSubject(c echo.Context) error {
	// Controla los goroutines para que se ejecute uno a la vez y evitar corrupciones en los datos
	lock.Lock()

	// Libera los goroutines cuando finalice la ejecucion del metodo
	defer lock.Unlock()

	// Recupera el id del parametro de la url
	id := c.Param("id")

	// Instancia de actualizada materia
	var updatedSubject Subject

	// Agrega los datos de la materia en la instancia de la materia actualizada, en su espacio de memoria 
	if err := c.Bind(&updatedSubject); err != nil {
		return err
	}

	// Actualiza la materia almacenada en el mapa de materias la materia actualizada, buscandola mediante el id
	subjects[id] = updatedSubject

	// Retorna la materia actualizada
	return c.JSON(http.StatusOK, updatedSubject)
}

// Metodo para eliminar una materia
func DeleteSubject(c echo.Context) error {
	// Controla los goroutines para que se ejecute uno a la vez y evitar corrupciones en los datos
	lock.Lock()

	// Libera los goroutines cuando finalice la ejecucion del metodo
	defer lock.Unlock()

	// Recupera el id del parametro de la url
	id := c.Param("id")

	// Elimina del mapa de materias la materia especificada con el id
	delete(subjects, id)

	// Retorna respuesta http
	return c.NoContent(http.StatusOK)
}