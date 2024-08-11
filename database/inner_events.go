package database

import (
	"devbeginner-doc-api/model"
	"errors"
	"fmt"
)

var InnerEvents *innerEventsDBMethod

type innerEventsDBMethod struct{}

func (*innerEventsDBMethod) Create(m *model.InnerEvent) error {
	sql := "INSERT INTO inner_events(name,summary,notes,startAt,`release`) VALUES(:name,:summary,:notes,:startAt,:release)"
	_, err := DB.NamedExec(sql, m)
	if err != nil {
		return err
	}
	return nil
}

func (*innerEventsDBMethod) Query(isRelease bool) ([]model.InnerEvent, error) {
	var inEvents []model.InnerEvent
	var sql string
	if isRelease {
		sql = "SELECT * FROM inner_events WHERE `release`=TRUE"
	} else {
		sql = "SELECT * FROM inner_events WHERE `release`=FALSE"
	}
	rows, err := DB.Queryx(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		temp := model.InnerEvent{}
		err := rows.StructScan(&temp)
		if err != nil {
			return nil, err
		}
		inEvents = append(inEvents, temp)
	}
	return inEvents, nil
}

func (*innerEventsDBMethod) Delete(uid int) error {
	sql := "DELETE FROM inner_events WHERE `uid`=?"
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

func (*innerEventsDBMethod) Update(uid int, column string, content any) error {
	sql := fmt.Sprintf("UPDATE inner_events SET `%s`=? WHERE `uid`=%d", column, uid)
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
