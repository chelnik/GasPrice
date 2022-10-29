package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/chelnik/GasPriceTask/conf"
	"github.com/gorilla/mux"
)

type Response struct {
	Monthly          map[int]float64    // monthly expenses
	PricePerDay      map[string]float64 // average purchase price per day
	PriceDistPerHour map[int]float64    // price distribution per hour
	EntirePeriod     float64            // payment for the entire period
}

const port = ":8080"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		sliceByte := workWithData()
		_, err := w.Write(sliceByte)
		if err != nil {
			log.Println("error in writing bytes", err)
			http.Error(w, "server-side problems", http.StatusInternalServerError)
		}
	}).Methods("GET")
	coveredRouter := middlewareHandler(router)
	fmt.Printf("listening in http://localhost%s\n", port)
	err := http.ListenAndServe(port, coveredRouter)
	if err != nil {
		log.Println("error in ListenAndServe", err)
	}

}

func workWithData() []byte {
	data := conf.New()
	err := data.GetRequest()
	if err != nil {
		log.Println("error in GetRequest", err)
	}

	everyMonths, err := data.EveryMonths()
	if err != nil {
		log.Println("error in strconv in EveryMonths", err)
	}
	distPerHour, err := data.TakePriceDistPerHour()
	if err != nil {
		log.Println("error in strconv in TakePriceDistPerHour", err)
	}

	response := Response{
		Monthly:          everyMonths,
		PricePerDay:      data.AveragePricePerDay(),
		PriceDistPerHour: distPerHour,
		EntirePeriod:     data.TotalSpent(),
	}

	ourResponce, err := json.Marshal(response)
	if err != nil {
		log.Println("error with Marshal", err)
	}
	// fmt.Println(response.EntirePeriod)
	// fmt.Println(response.PriceDistPerHour)
	// fmt.Println(ourResponce)
	return ourResponce
}

func middlewareHandler(handler http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Printf("[%s] %s %s %s\n\n", r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
	})
}

func testPricePerDay(response *Response) {
	for i, i2 := range response.PricePerDay {
		if "22-08-05" == i {
			fmt.Println(i, i2)
		}
	}
}
