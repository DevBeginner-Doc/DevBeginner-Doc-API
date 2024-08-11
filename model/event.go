package model

type InnerEvent struct {
	Uid       int    `json:"uid" db:"uid"`
	Name      string `json:"name" db:"name"`
	Summary   string `json:"summary" db:"summary"`
	Notes     string `json:"notes" db:"notes"`
	StartTime string `json:"startAt" db:"startAt"`
	Release   bool   `json:"release" db:"release"`
}

type Event struct {
	Name      string `json:"name" db:"name"`
	Platform  string `json:"platform" db:"platform"`
	Link      string `json:"link" db:"link"`
	StartTime string `json:"startAt" db:"startAt"`
	Status    string `json:"status" db:"status"`
}
