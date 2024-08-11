package database

import (
	"devbeginner-doc-api/model"
	"errors"
	"fmt"
)

var IdeIndex *indexDBMethod

type indexDBMethod struct{}

func (*indexDBMethod) Create(m *model.IdeIndex) error {
	sql := "INSERT INTO ide_index(`rank`,`name`) VALUES(:rank,:name)"
	_, err := DB.NamedExec(sql, m)
	if err != nil {
		return err
	}
	return nil
}

func (*indexDBMethod) Query() ([]model.IdeIndex, error) {
	var mides []model.IdeIndex
	sql := "SELECT * FROM ide_index"
	rows, err := DB.Queryx(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		mide := model.IdeIndex{}
		err := rows.StructScan(&mide)
		if err != nil {
			return nil, err
		}
		mides = append(mides, mide)
	}
	return mides, nil
}

func (*indexDBMethod) Delete(uid int) error {
	sql := "DELETE FROM ide_index WHERE `uid`=?"
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

func (*indexDBMethod) Update(uid int, column string, content any) error {
	sql := fmt.Sprintf("UPDATE ide_index SET `%s`=? WHERE `uid`=%d", column, uid)
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
