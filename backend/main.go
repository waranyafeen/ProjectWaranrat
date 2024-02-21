package main

import (
	"github.com/waranyafeen/Projectwaranrat/routers"
	"github.com/waranyafeen/Projectwaranrat/entity"
)

func main() {
	entity.SetupDatabase("WaranratDB")
	entity.SetupData(entity.DB())
	route := routers.SetupRouter()
	
	routers.InitRouter(route)
	route.Run()

}

