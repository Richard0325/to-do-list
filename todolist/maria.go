package todolist

import (
	"database/sql"
	"fmt"
	"strconv"
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
		var id int
		err = rows.Scan(&id, &t.Title, &t.Description, &t.Deadline)
		idStr := strconv.Itoa(id)
		t.ID = idStr
		if err != nil {
			fmt.Printf("Scan post row error: %s\n", err.Error())
		}
		ret = append(ret, &t)
	}
	return ret, nil
}
func (dao MariaDao) GetTask(idStr string) (*Task, error) {
	sqlStr := `SELECT id, title, description, deadline FROM todolist WHERE id = ?`
	stmt, err := dao.DB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("SQL prepare error")
		return nil, err
	}
	defer stmt.Close()
	t := Task{}
	id, _ := strconv.Atoi(idStr)
	row := stmt.QueryRow(id)
	err = row.Scan(&t.ID, &t.Title, &t.Description, &t.Deadline)
	if err != nil {
		return nil, ErrNotFound
	}
	return &t, nil
}
func (dao MariaDao) AddTask(task *Task) (string, error) {
	var count int
	err := dao.DB.QueryRow("SELECT COUNT(*) FROM todolist").Scan(&count)
	if err != nil{
		fmt.Println("query counts error")
		return "-1", err
	}
	if count >= ListSize{
		return "-1", ErrNoSpace
	}
	if task == nil {
		return "-1", ErrInvalidInput
	}
	//if I type "INSERT" instead of insert ,it would cause error
	sqlStr := `insert INTO todolist (title, description, deadline) Values (?, ?, ?)`
	stmt, err := dao.DB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("SQL prepare error")
		return "-1", err
	}
	defer stmt.Close()
	result, err := stmt.Exec(task.Title, task.Description, task.Deadline)
	if err != nil {
		fmt.Println("stmt.Exec error")
		return "-1", err
	}
	id, err := result.LastInsertId()
	idStr:= strconv.Itoa(int(id))
	return idStr, err
}
func (dao MariaDao) DeleteTask(idStr string) error {
	sqlStr := "DELETE FROM todolist WHERE id = ?"
	stmt, err := dao.DB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("stmt prepare error")
		return err
	}
	defer stmt.Close()
	id, _ := strconv.Atoi(idStr)
	_, err = stmt.Exec(id)
	return err
}
