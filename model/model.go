package model

type Fruit struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Sugar  int    `json:"sugar"`
}
