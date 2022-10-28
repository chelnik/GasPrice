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

// const fileName = "gas_price.json"

type Response struct {
	Monthly          map[int]float64    // monthly expenses
	PricePerDay      map[string]float64 // average purchase price per day
	PriceDistPerHour map[int]float64    // price distribution per hour
	EntirePeriod     float64            // payment for the entire period
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		sliceByte := workWithData()
		// dobby := new(Response)
		// err := json.Unmarshal(sliceByte, dobby)
		// if err != nil {
		// 	return
		// }
		_, err := w.Write(sliceByte)
		if err != nil {
			log.Println("error in writing bytes", err)
		}
	}).Methods("GET")
	coveredRouter := middlewareHandler(router)
	err := http.ListenAndServe(":8085", coveredRouter)
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
		fmt.Println("middlewareHandler", r.URL.Path)
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
