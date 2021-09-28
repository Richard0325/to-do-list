package todolist

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MariaDao struct {
	DB *sql.DB
}

func (dao MariaDao) GetTasks() (Tasks, error) {
	sqlStr := "SELECT id, title, description, deadline FROM todolist"
	rows, err := dao.DB.Query(sqlStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	defer rows.Close()

	ret := []*Task{}
	for rows.Next() {
		t := Task{}
		err = rows.Scan(&t.ID, &t.Title, &t.Description, &t.Deadline)
		if err != nil {
			fmt.Printf("Scan post row error: %s\n", err.Error())
		}
		ret = append(ret, &t)
	}
	return ret, nil
}
func (dao MariaDao) GetTask(id int) (*Task, error) {
	sqlStr := `SELECT id, title, description, deadline FROM todolist WHERE id = ?`
	stmt, err := dao.DB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("SQL prepare error")
		return nil, err
	}
	defer stmt.Close()
	t := Task{}
	row := stmt.QueryRow(id)
	err = row.Scan(&t.ID, &t.Title, &t.Description, &t.Deadline)
	if err != nil {
		return nil, ErrNotFound
	}
	return &t, nil
}
func (dao MariaDao) AddTask(task *Task) (int, error) {
	if task == nil {
		return -1, ErrInvalidInput
	}
	//if I type "INSERT" instead of insert ,it would cause error
	sqlStr := `insert INTO todolist (title, description, deadline) Values (?, ?, ?)`
	stmt, err := dao.DB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("SQL prepare error")
		return -1, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(task.Title, task.Description, task.Deadline)
	if err != nil {
		fmt.Println("stmt.Exec error")
		return -1, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}
func (dao MariaDao) DeleteTask(id int) error {
	sqlStr := "DELETE FROM todolist WHERE id = ?"
	stmt, err := dao.DB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("stmt prepare error")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}
