package delivery

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"todolist/m/internal/models"
	"todolist/m/internal/usecase"

	"github.com/julienschmidt/httprouter"
)

type TaskDelivery interface {
	Add(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	List(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Mark(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type taskDelivery struct {
	usecase usecase.TaskUsecase
}

func NewTaskDelivery(usecase usecase.TaskUsecase) TaskDelivery {
	return &taskDelivery{
		usecase: usecase,
	}
}

func (d *taskDelivery) Add(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	newTask := new(models.Task)

	err := decoder.Decode(newTask)
	if err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{}"))
		return
	}

	addedTask, err := d.usecase.AddTask(*newTask)
	if err != nil {
		log.Printf("error while adding task")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{}"))
		return
	}

	bytes, err := json.Marshal(addedTask)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{}"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (d *taskDelivery) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Printf("error while getting param")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = d.usecase.DeleteTask(id)
	if err != nil {
		log.Printf("error while deleting task: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (d *taskDelivery) List(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tasks, err := d.usecase.List()
	if err != nil {
		log.Printf("error while getting tasks: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{}"))
		return
	}

	bytes, err := json.Marshal(tasks)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{}"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (d *taskDelivery) Mark(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Printf("error while getting param")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	request := new(models.MarkRequest)

	err = decoder.Decode(request)
	if err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{}"))
		return
	}

	task, err := d.usecase.Mark(id, *request)
	if err != nil {
		log.Printf("error while adding task")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{}"))
		return
	}

	bytes, err := json.Marshal(task)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{}"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
