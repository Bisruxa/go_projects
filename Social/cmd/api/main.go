package main 

import (
	"log"
	)
func main (){
	cfg :=  config{
			addr: ":3000",
			//addr: env.GetString("ADDR","8080"),
		}
	app := &application{
		config:cfg,
	}
	store :=
	mux := app.mount()
	log.Fatal(app.run(mux))
}