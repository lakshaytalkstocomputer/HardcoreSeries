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
}

type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodPost:
		p.processWin(w)
	case http.MethodGet:
		p.showScore(w, r)

	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request){
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := p.Store.GetPlayerScore(player)

	if score == 0{
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter){
	w.WriteHeader(http.StatusAccepted)
}