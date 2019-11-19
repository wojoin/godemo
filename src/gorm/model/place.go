package model

type Place struct {
	ID     int    `gorm:"primary_key"`
	Name   string `gorm:"name"`
	Town   Town
	TownId int `gorm:"ForeignKey:id"`
}

type Town struct {
	ID   int    `gorm:"primary_key"`
	Name string `gorm:"name"`
}
