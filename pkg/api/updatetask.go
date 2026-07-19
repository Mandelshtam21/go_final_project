package api

import (
	"bytes"
	"encoding/json"
	"go_final_project/pkg/db"
	"net/http"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		writeJson(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	err = json.Unmarshal(buf.Bytes(), &task)
	if err != nil {
		writeJson(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if task.ID == "" {
		writeJson(w, map[string]string{
			"error": "Не указан идентификатор",
		})
		return
	}

	if task.Title == "" {
		writeJson(w, map[string]string{
			"error": "Не указан заголовок задачи",
		})
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeJson(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	err = db.UpdateTask(&task)
	if err != nil {
		writeJson(w, map[string]string{
			"error": err.Error(),
		})
		return
	}
	writeJson(w, map[string]string{})
}
