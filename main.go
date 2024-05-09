package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	justarouter "github.com/marcotheo/justarouter"
)

type UserInfo struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	Id   string `json:"id"`
}

func pokemonRoutes(subRouter justarouter.SubRouter) {
	subRouter.POST("/info", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No Pokemon Info Yet")
	})

	subRouter.GET("/list", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No Pokemon List Yet because I'm tired")
	})
}

func userRoutes(subRouter justarouter.SubRouter) {
	subRouter.GET("/{userId}", func(w http.ResponseWriter, r *http.Request) {
		val := r.PathValue("userId")

		me := UserInfo{
			Name: "marco",
			Age:  "25",
			Id:   val,
		}

		b, err := json.Marshal(me)

		if err != nil {
			fmt.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(b))
	})
}

func main() {
	router := justarouter.CreateRouter()

	router.AddSubRoutes("/pokemon", pokemonRoutes)
	router.AddSubRoutes("/user", userRoutes)

	router.POST("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "App is Healthy")
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router.Mux,
	}

	fmt.Println("Server running at port :8080")

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
		return
	}
}
