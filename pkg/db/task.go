package db

import (
	"database/sql"
	"fmt"
)

const DateFormat = "20060102"

type Task struct {
	Id      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	var id int64

	query := "INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)"
	///@todo ExecContext()
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
	tasks := []*Task{}

	query := "SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date DESC LIMIT :limit"
	rows, err := db.Query(query, sql.Named("limit", limit))
	if err != nil {
		return []*Task{}, err
	}
	defer rows.Close()

	for rows.Next() {
		task := Task{}
		err := rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return []*Task{}, err
		}

		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
		return []*Task{}, err
	}

	return tasks, nil
}

func GetTask(id string) (*Task, error) {
	query := "SELECT id, date, title, comment, repeat FROM scheduler WHERE id = :id"

	row := db.QueryRow(query, sql.Named("id", id))

	task := Task{}
	err := row.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return &Task{}, err
	}

	return &task, nil
}

func UpdateTask(task *Task) error {
	query := "UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id"

	res, err := db.Exec(query,
		sql.Named("id", task.Id),
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat))

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}
	return nil
}

func UpdateDateTask(next string, id string) error {
	query := "UPDATE scheduler SET date = :date WHERE id = :id"

	res, err := db.Exec(query,
		sql.Named("id", id),
		sql.Named("date", next))

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating date task`)
	}
	return nil
}

func DeleteTask(id string) error {
	query := "DELETE FROM scheduler WHERE id = :id"

	_, err := db.Exec(query, sql.Named("id", id))
	if err != nil {
		return err
	}

	return nil
}
