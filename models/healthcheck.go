package models

type Healthcheck struct {
	Id   uint `gorm:"primaryKey"`
	Name string
}
