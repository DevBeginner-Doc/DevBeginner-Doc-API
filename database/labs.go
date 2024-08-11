package database

import (
	"devbeginner-doc-api/model"
	"errors"
	"fmt"
)

var Labs *labsDBMethod

type labsDBMethod struct{}

func (*labsDBMethod) Create(m *model.Lab) error {
	sql := "INSERT INTO labs(name,summary,college,position,`limit`,`group`,`time`,`release`) VALUES(:name,:summary,:college,:position,:limit,:group,:time,:release)"
	_, err := DB.NamedExec(sql, m)
	if err != nil {
		return err
	}
	return nil
}

func (*labsDBMethod) Query(isRelease bool) ([]model.Lab, error) {
	var mlabs []model.Lab
	var sql string
	if isRelease {
		sql = "SELECT * FROM labs WHERE `release`=TRUE"
	} else {
		sql = "SELECT * FROM labs WHERE `release`=FALSE"
	}
	rows, err := DB.Queryx(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		mlab := model.Lab{}
		err := rows.StructScan(&mlab)
		if err != nil {
			return nil, err
		}
		mlabs = append(mlabs, mlab)
	}
	return mlabs, nil
}

func (*labsDBMethod) Delete(uid int) error {
	sql := "DELETE FROM labs WHERE `uid`=?"
	res, err := DB.Exec(sql, uid)
	cnt, _ := res.RowsAffected()
	if err != nil {
		return err
	}
	if cnt == 0 {
		return errors.New("无数据行被修改")
	}
	return nil
}

func (*labsDBMethod) Update(uid int, column string, content any) error {
	sql := fmt.Sprintf("UPDATE labs SET `%s`=? WHERE `uid`=%d", column, uid)
	res, err := DB.Exec(sql, content)
	cnt, _ := res.RowsAffected()
	if err != nil {
		return err
	}
	if cnt == 0 {
		return errors.New("无数据行被修改")
	}
	return nil
}
