package api

import (
	"go_final_project/pkg/db"
	"net/http"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		writeJson(w, http.StatusBadRequest, map[string]string{
			"error": "Не указан идентификатор",
		})
		return
	}

	err := db.DeleteTask(id)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}
	writeJson(w, http.StatusOK, map[string]string{})
}
