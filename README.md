# 📘 API de Gestión de Materias

Este proyecto es una API simple escrita en Go utilizando el framework **Echo**, que permite realizar operaciones CRUD sobre materias (`Subjects`). Además, incluye pruebas de **performance** con [Vegeta](https://github.com/tsenart/vegeta) y pruebas **End-to-End** con [chromedp](https://github.com/chromedp/chromedp).

---

## Ejecución del proyecto

### Requisitos

- Go 1.21+
- Navegador instalado (solo para pruebas E2E)

### Comando para iniciar el servidor:

```bash
go run main.go
```

```bash
go test ./tests -v -count=1
```

```bash
go test ./tests -run TestPerformanceCreateSubjects -v -count=1
```

```bash
go test ./tests -run TestEndToEndCreateSubject -v -count=1
```