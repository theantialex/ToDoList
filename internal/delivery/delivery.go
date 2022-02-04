package delivery

import (
	"net/http"
	"todolist/m/internal/usecase"

	"github.com/julienschmidt/httprouter"
)

type TaskDelivery interface {
	Add(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	//Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	//List(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	//Mark(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type taskDelivery struct {
	usecase usecase.TaskUsecase
}

func NewTaskDelivery(usecase usecase.TaskUsecase) TaskDelivery {
	return &taskDelivery{
		usecase: usecase,
	}
}

func (*taskDelivery) Add(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
