package service

import (
	"github.com/fbbyqsyea/go-ssh/utils"
)

func Remove(name string) {
	stmt, err := Db.Prepare("delete from " + SSH_TABLE_NAME + " where name=?")
	utils.CheckIfError(err)
	res, err := stmt.Exec(name)
	utils.CheckIfError(err)
	affect, err := res.RowsAffected()
	utils.CheckIfError(err)
	if affect > 0 {
		utils.Info("ssh remove success")
	} else {
		utils.Warning("ssh remove failed")
	}
}
