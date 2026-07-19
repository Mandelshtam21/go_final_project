package db

import (
	"database/sql"
	"fmt"
	"time"
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

func Tasks(limit int, search string) ([]*Task, error) {
	var (
		resp   *sql.Rows
		err    error
		dbDate string
	)

	date, parseErr := time.Parse("02.01.2006", search)
	if parseErr == nil {
		dbDate = date.Format("20060102")
	}
	if search == "" {
		query := `SELECT id, date, title, comment, repeat 
			FROM scheduler 
			ORDER BY date ASC`
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
	} else if parseErr == nil {
		query := `SELECT * FROM scheduler 
			WHERE date = :date 
			LIMIT :limit`
		resp, err = db.Query(query,
			sql.Named("date", dbDate),
			sql.Named("limit", limit))
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
	} else {
		searchValue := "%" + search + "%"
		query := `SELECT * FROM scheduler 
			WHERE title 
			LIKE :search OR comment LIKE :search 
			ORDER BY date 
			LIMIT :limit`
		resp, err = db.Query(query,
			sql.Named("search", searchValue),
			sql.Named("limit", limit))
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
}

func GetTask(id string) (*Task, error) {
	task := &Task{}
	query := `SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?`
	resp := db.QueryRow(query, id)
	err := resp.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func UpdateTask(task *Task) error {
	query := `UPDATE scheduler SET 
				date = :date, 
				title = :title, 
				comment = :comment, 
				repeat = :repeat
				WHERE id = :id`
	res, err := db.Exec(query,
		sql.Named("id", task.ID),
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat))
	if err != nil {
		return err
	}
	// метод RowsAffected() возвращает количество записей к которым
	// была применена SQL команда
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}
	return nil
}

func DeleteTask(id string) error {
	query := `DELETE FROM scheduler WHERE id = :id`
	resp, err := db.Exec(query,
		sql.Named("id", id))
	if err != nil {
		return err
	}

	count, err := resp.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for deleting task`)
	}
	return nil
}

func UpdateDate(next string, id string) error {
	query := `UPDATE scheduler SET date = :date WHERE id = :id`
	resp, err := db.Exec(query,
		sql.Named("date", next),
		sql.Named("id", id))
	if err != nil {
		return err
	}

	count, err := resp.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}
	return nil
}
