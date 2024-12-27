package main

import (
	"log"
	"server/internal/core/model/dao"
	"server/internal/core/model/utils"
	"server/pkg/g"
)

func main() {
	db := g.DB()
	dao.SetDefault(db)
	if err := utils.AutoMigrate(db); err != nil {
		log.Fatalf("db auto migrate error: %s", err)
	}
	if err := utils.AutoInitialData(db); err != nil {
		log.Fatalf("db auto initial data error: %s", err)
	}

	log.Println("DB AutoMigrate")
}
