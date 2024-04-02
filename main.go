package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

type challenge struct {
	Data string `json:"data"`
}

var text = "metaverse web3 NFT crypto decentralized meme stock stonk hodl ape GameStop AMC Reddit Robinhood Dogecoin elon tesla Twitter Muskrat quiet quitting great resignation quiet firing layoff recession inflation cost of living supply chain chip shortage climate crisis heat wave drought fire season net zero green energy EV plant-based oat milk cauliflower gnocchi charcuterie grazing board cheugy cringe slay zaddy bussy thirst trap y'all cap no cap fr fr wig go off understood the assignment hot girl walk feral girl summer that's the tweet main character energy unalive sadfishing negging love-bombing gatekeeping cloutlighting sliving going goblin mode crisitunity ambient anxiety"

func main() {
	fmt.Println("Starting RESTful service")

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	}))

	router.Get("/challenges", getChallenge)

	http.ListenAndServe(":8080", router)
}

func getChallenge(w http.ResponseWriter, r *http.Request) {
	res := challenge{
		Data: text,
	}

	render.JSON(w, r, res)
}
