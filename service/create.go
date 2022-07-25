package service

import (
	"fmt"
	"os"

	"github.com/fbbyqsyea/go-ssh/utils"
)

func Create(name, host, user, password string, port int) {
	// check name is exists in table
	if NameExists(name) {
		utils.Error(fmt.Errorf("name %s is exists", name))
		os.Exit(1)
	}

	// insert into table
	insert(name, host, user, password, port)
	utils.Info("ssh create success")
}

// check name is exists
func NameExists(name string) bool {
	rows, err := Db.Query("SELECT count(*) as count FROM " + SSH_TABLE_NAME + " WHERE name= \"" + name + "\";")
	utils.CheckIfError(err)
	defer rows.Close()
	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		utils.CheckIfError(err)
	}
	return count > 0
}

func insert(name, host, user, password string, port int) int64 {
	stmt, err := Db.Prepare("INSERT INTO " + SSH_TABLE_NAME + "(name, host, port, user, password)  values(?, ?, ?, ?, ?)")
	utils.CheckIfError(err)
	res, err := stmt.Exec(name, host, port, user, password)
	utils.CheckIfError(err)
	id, err := res.LastInsertId() //返回新增的id号
	utils.CheckIfError(err)
	return id
}
