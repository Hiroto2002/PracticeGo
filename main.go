package main

import (
	"practiceGo/app/src/infrastructure"
    "os"
)

func main() {
    
    env := os.Getenv("ENV")
	db := infrastructure.NewDB()
    if env == "development" {
        db.ResetDatabase()
        db.InsertSeedData()
    }else{
        db.AutoMigrate()
    }
	http := infrastructure.NewHttp()
	r := infrastructure.NewRouting(db, http)
	r.Run()
}
