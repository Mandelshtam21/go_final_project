package db

import (
	"database/sql"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment,omitempty"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	var id int64

	query := `INSERT INTO scheduler (date, title, comment, repeat) 
	VALUES (:date, :title, :comment, :repeat)`
	res, err := db.Exec(query,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat))
	if err == nil {
		id, err = res.LastInsertId()
	}
	return id, err
}

func Tasks(limit int) ([]*Task, error) {
	var (
		resp *sql.Rows
		err  error
	)
	query := `SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date ASC`
	if limit > 0 {
		query += ` LIMIT ?`
		resp, err = db.Query(query, limit)
	} else {
		resp, err = db.Query(query)
	}
	if err != nil {
		return nil, err
	}

	defer resp.Close()
	tasks := []*Task{}
	for resp.Next() {
		task := &Task{}

		err := resp.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
