package model

import "fmt"

func CreateTodo(name, todo string) error {
	insertQ, err := con.Query("INSERT INTO TODO VALUES(?, ?)", name, todo)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		defer insertQ.Close()
	}
	return nil
}

func DeleteTodo(name string) error {
	insertQ, err := con.Query("DELETE FROM TODO WHERE name=?", name)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		defer insertQ.Close()
	}
	return nil
}
