package main

import (
	"goblogart/inits"
	"goblogart/models"
)

func init() {
	inits.LoadEnv()
	inits.MySqlDBInit()
}

func main() {
	inits.DB.AutoMigrate(&models.Post{})
}