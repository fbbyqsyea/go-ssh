package service

import (
	"os"
	"strconv"

	"github.com/fbbyqsyea/go-ssh/utils"
	"github.com/olekukonko/tablewriter"
)

var SSH_LIST_HEADER = []string{"Name", "Host", "Port", "User", "Password", "Created"}

type SshConnect struct {
	Id       int
	Name     string
	Host     string
	Port     int
	User     string
	Password string
	Created  string
}

func List(isShowPassword bool) {
	rows, err := Db.Query("SELECT * FROM " + SSH_TABLE_NAME + " order by id desc;")
	utils.CheckIfError(err)
	defer rows.Close()
	var data [][]string
	for rows.Next() {
		var name, host, user, password, created string
		var id, port int
		err = rows.Scan(&id, &name, &host, &port, &user, &password, &created)
		utils.CheckIfError(err)
		if !isShowPassword {
			data = append(data, []string{name, host, strconv.Itoa(port), user, "***", created})
		} else {
			data = append(data, []string{name, host, strconv.Itoa(port), user, password, created})
		}
	}
	Output(SSH_LIST_HEADER, data)
}

// list ssh connect by name
func ListName(n string, isShowPassword bool) {
	sc := GetConnectByName(n)
	if !isShowPassword {
		Output(SSH_LIST_HEADER, [][]string{{sc.Name, sc.Host, strconv.Itoa(sc.Port), sc.User, "***", sc.Created}})
	} else {
		Output(SSH_LIST_HEADER, [][]string{{sc.Name, sc.Host, strconv.Itoa(sc.Port), sc.User, sc.Password, sc.Created}})
	}
}

// get ssh connect by name
func GetConnectByName(n string) SshConnect {
	rows, err := Db.Query("SELECT * FROM " + SSH_TABLE_NAME + " WHERE name= \"" + n + "\";")
	utils.CheckIfError(err)
	defer rows.Close()
	var name, host, user, password, created string
	var id, port int
	for rows.Next() {
		err = rows.Scan(&id, &name, &host, &port, &user, &password, &created)
		utils.CheckIfError(err)
	}
	return SshConnect{id, name, host, port, user, password, created}
}

// output
func Output(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader(header)

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
