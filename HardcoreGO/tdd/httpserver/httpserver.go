package httpserver

import (
	"fmt"
	"net/http"
	"strings"
)

// Create a Web Server where users can track how many games players have won.
//  - GET /players/{name} - should return a number indicating the total number oif wins
//  - POST /players/{name} - should record a win for that name, incrementing for every subsequent POST

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request){
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)

	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string){

	score := p.Store.GetPlayerScore(player)

	if score == 0{
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string){

	p.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}