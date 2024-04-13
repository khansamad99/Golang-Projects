package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())

}

func getSession() *mgo.Session {

	s, err := mgo.Dial("mongodb://localhost:27107")

	if err != nil {
		panic(err)
	}

	return s
}
