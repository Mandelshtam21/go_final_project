package api

import (
	"go_final_project/pkg/db"
	"net/http"
	"time"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		writeJson(w, map[string]string{
			"error": "Не указан идентификатор",
		})
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeJson(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	now := time.Now()
	if task.Repeat == "" {
		err := db.DeleteTask(id)
		if err != nil {
			writeJson(w, map[string]string{
				"error": err.Error(),
			})
			return
		}
		writeJson(w, map[string]string{})
		return
	}
	next, err := NextDate(now, task.Date, task.Repeat)
	if err != nil {
		writeJson(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	err = db.UpdateDate(next, id)
	if err != nil {
		writeJson(w, map[string]string{
			"error": err.Error(),
		})
		return
	}
	writeJson(w, map[string]string{})
}
