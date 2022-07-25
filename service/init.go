package service

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	"github.com/fbbyqsyea/go-ssh/utils"
)

var (
	Db *sql.DB
)

const (
	SSH_TABLE_CREATE_SQL = `
	CREATE TABLE IF NOT EXISTS "ssh_connect_details" (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"name" VARCHAR(32) NOT NULL DEFAULT "",
		"host" VARCHAR(32) NOT NULL DEFAULT "",
		"port" SMALLINTNT(5) NOT NULL DEFAULT 0,
		"user" VARCHAR(32) NOT NULL DEFAULT "",
		"password" VARCHAR(255) NOT NULL DEFAULT "",
		"created" TIMESTAMP default (datetime('now', 'localtime')),
		UNIQUE (name)
	);`
	SSH_TABLE_NAME = "ssh_connect_details"
)

// run pre check
// if not installed sqlite3, remind to install first
func PreCheck() {
	res := utils.ExecCommandWithResult("sqlite3 --version")
	if strings.Contains(res, "command not found") {
		utils.Error(fmt.Errorf("sqlite3 is not installed, please install sqlite3 first"))
		os.Exit(1)
	}
}

// init sqlite3 connect and database
func InitSqlite3ConnectAndDatabase() {
	// init sqlite3 connect
	Db = initSqlite3Connect()
	// init sqlite3 database
	initDatabse()
}

// init sqlite3 connect
func initSqlite3Connect() *sql.DB {
	homeDir, err := os.UserHomeDir()
	utils.CheckIfError(err)
	db, err := sql.Open("sqlite3", homeDir+"/.sshDB.db")
	utils.CheckIfError(err)
	return db
}

// init sqlite3 database
func initDatabse() {
	if !databaseInited() {
		_, err := Db.Exec(SSH_TABLE_CREATE_SQL)
		utils.CheckIfError(err)
	}
}

// check database is inited
func databaseInited() bool {
	rows, err := Db.Query("SELECT count(*) as count FROM sqlite_master WHERE type=\"table\" AND name = \"" + SSH_TABLE_NAME + "\";")
	utils.CheckIfError(err)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err = rows.Scan(&count)
		utils.CheckIfError(err)
	}
	return count > 0
}

// close db
func CloseDB() {
	Db.Close()
}
