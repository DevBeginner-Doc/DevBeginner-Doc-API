package model

type IdeIndex struct {
	Uid  int    `json:"uid" db:"uid"`
	Rank int    `json:"rank" db:"rank"`
	Name string `json:"name" db:"name"`
}
