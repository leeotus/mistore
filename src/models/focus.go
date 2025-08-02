package models

type Focus struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	FocusType int    `json:"focustype"`
	FocusImg  string `json:"img"`
	Link      string `json:"link"`
	Sort      int    `json:"sort"`
	Status    int    `json:"status"`
	AddTime   int    `json:"addtime"`
}

func (Focus) TableName() string {
	return "focus"
}
