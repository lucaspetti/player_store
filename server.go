package main

import (
	"net/http"
	"fmt"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, GetPlayerScore(player))
}

func GetPlayerScore(name string) string {
	switch name {
	case "Pippen":
		return "20"
	case "Jordan":
		return "45"
	default:
		return ""
	}
}
