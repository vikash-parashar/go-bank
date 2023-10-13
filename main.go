package main

import (
	"log"

	env "github.com/joho/godotenv"
)

func init() {
	if err := env.Load(); err != nil {
		log.Fatalln("Error : Failed To Load .Env File")
	}
}
func main() {
	s := NewApiServer()
	if err := s.Run(); err != nil {
		log.Fatalln("Error : Failed To Start Application !")
	}

}
