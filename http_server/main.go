package main

import (
	"log"
	"net/http"
)

type inMemoryPlayerStore struct{}

func (i *inMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}
func (i *inMemoryPlayerStore) RecordWin(name string) {}

func main() {
	server := &PlayerServer{&inMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5001", server))
}
