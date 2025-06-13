# 📘 API de Gestión de Materias

Esta es una API sencillas desarrollada en Golang utilizando el framework **Echo**, que permite realizar operaciones CRUD sobre materias . Incluye pruebas de **performance** con [Vegeta](https://github.com/tsenart/vegeta) y pruebas **End-to-End** con [chromedp](https://github.com/chromedp/chromedp).

---

## Ejecución del proyecto
El proyecto corre en localhost

### Requisitos

- Go 1.21+
- Navegador instalado (solo para pruebas E2E)

### Puerto
:8080

### Recursos web disponibles
- GET "/subjects" Recupera todas las materias
- POST "/subjects" Crea una nueva materia
- GET "/subject/:id" Recupera una materia por su id
- PUT "/subject/:id" Actualiza una materia por su id
- DELETE "/subject/:id" Elimina una materia por su id

### Comando para iniciar el servidor:
```bash
go run main.go
```
### Comando para ejecutar todas las pruebas:
```bash
go test ./tests -v -count=1
```

### Comando para ejecutar las pruebas de performance:
```bash
go test ./tests -run TestPerformanceCreateSubjects -v -count=1
```

### Comando para ejecutar las pruebas end to end:
```bash
go test ./tests -run TestEndToEndCreateSubject -v -count=1
```