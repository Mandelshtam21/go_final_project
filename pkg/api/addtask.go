package api

import (
	"encoding/json"
	"go_final_project/pkg/db"
	"log"
	"net/http"
	"strconv"
	"time"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	//var buf bytes.Buffer

	/**	_, err := buf.ReadFrom(r.Body)
		if err != nil {
			writeJson(w, http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}

		err = json.Unmarshal(buf.Bytes(), &task)
		if err != nil {
			writeJson(w, http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}
	**/
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	if task.Title == "" {
		writeJson(w, http.StatusBadRequest, map[string]string{
			"error": "Не указан заголовок задачи",
		})
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}
	idStr := strconv.Itoa(int(id))
	writeJson(w, http.StatusCreated, map[string]string{
		"id": idStr,
	})
}

func checkDate(task *db.Task) error {
	var next string
	now := time.Now()
	now, err := time.Parse(DateFormat, now.Format(DateFormat))
	if err != nil {
		return err
	}

	if task.Date == "" {
		task.Date = now.Format(DateFormat)
	}

	t, err := time.Parse(DateFormat, task.Date)
	if err != nil {
		return err
	}

	if task.Repeat != "" {
		next, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
	}

	if afterNow(now, t) {
		if len(task.Repeat) == 0 {
			task.Date = now.Format(DateFormat)
		} else {
			task.Date = next
		}
	}

	return nil
}

func writeJson(w http.ResponseWriter, status int, data any) {
	resp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	_, err = w.Write(resp)
	if err != nil {
		log.Printf("Ошибка записи HTTP-ответа: %v", err)
	}
}
