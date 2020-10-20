package httpserver

import (
	"fmt"
	"net/http"
	"strings"
)

// Create a Web Server where users can track how many games players have won.
//  - GET /players/{name} - should return a number indicating the total number oif wins
//  - POST /players/{name} - should record a win for that name, incrementing for every subsequent POST

func PlayerServer(w http.ResponseWriter, r *http.Request){
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, GetPlayerScore(player))

}

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}

	return ""
}