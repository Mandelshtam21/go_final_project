package api

import (
	"net/http"

	"go_final_project/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search")
	tasks, err := db.Tasks(50, search)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}
	writeJson(w, http.StatusOK, TasksResp{
		Tasks: tasks,
	})
}
