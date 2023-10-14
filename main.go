package main

import (
	"fmt"
	"log"

	env "github.com/joho/godotenv"
)

func init() {
	if err := env.Load(); err != nil {
		log.Fatalln("Error : Failed To Load .Env File")
	}
}
func main() {
	store, err := NewPostgresStore("")
	if err != nil {
		log.Fatalln(err)
	}
	if err = store.init(); err != nil {
		log.Fatalln(err)
	}
	server := NewApiServer(store)
	if err := server.Run(); err != nil {
		log.Fatalln("Error : Failed To Start Application !")
	}
	fmt.Println("%+v\n", store)

}
