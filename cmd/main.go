package main

import (
	"database/sql"
	"log"
	"net/http"
	"todolist/m/internal/delivery"
	"todolist/m/internal/repository"
	"todolist/m/internal/usecase"

	_ "github.com/lib/pq"

	"github.com/julienschmidt/httprouter"
)

func headerMiddleware(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json; charset=utf8")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		handler(w, r, ps)
	}
}

func main() {
	db, err := sql.Open("postgres", "postgres://todolist:12345@host.docker.internal:5432/todolist")
	if err != nil {
		log.Fatalln("failed to connect to database", err.Error())
	}

	taskHandler := delivery.NewTaskDelivery(usecase.NewTaskUsecase(repository.NewTaskRepository(db)))

	r := httprouter.New()

	r.POST("/task", headerMiddleware(taskHandler.Add))
	r.DELETE("/task/:id", headerMiddleware(taskHandler.Delete))
	r.GET("/tasks", headerMiddleware(taskHandler.List))
	r.PATCH("/task/:id", headerMiddleware(taskHandler.Mark))

	log.Println("start serving :8080")
	http.ListenAndServe(":8080", r)
}
