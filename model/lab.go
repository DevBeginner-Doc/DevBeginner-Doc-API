package model

import (
	"devbeginner-doc-api/utils"
	"reflect"
)

type Lab struct {
	Uid      int    `json:"uid" db:"uid"`
	Name     string `json:"name" db:"name"`
	Summary  string `json:"summary" db:"summary"`
	College  string `json:"college" db:"college"`
	Position string `json:"position" db:"position"`
	Limit    string `json:"limit" db:"limit"`
	Group    string `json:"group" db:"group"`
	Time     string `json:"time" db:"time"`
	Release  bool   `json:"release" db:"release"`
}

func IsJsonInclude(obj any, col string) bool {
	tagArray := make([]string, 0)
	rel := reflect.TypeOf(obj)
	for i := 0; i < rel.NumField(); i++ {
		tagArray = append(tagArray, rel.Field(i).Tag.Get("json"))
	}
	res := utils.In(col, tagArray)
	return res
}
