package main

import (
	"encoding/json"
	"fmt"
	"log"
	// "os"
	"io/ioutil"
	"net/http"

	core "github.com/neetjn/go-cars/core"
)

func main() {

	// go imports all files in the same directory!!!
	// the main func is whats executed :)

	// car := app.CarDto{
	// 	Make:  "Hyundia",
	// 	Model: "Elantra",
	// 	Vin:   "10921341Q41",
	// 	Year:  2010,
	// 	Miles: 76100,
	// }
	// b, _ := json.Marshal(car)
	// fmt.Println(string(b))

	// embedded structs must be initialized by their field name
	// as composite literals

	// truck := app.TruckDto{
	// 	CarDto: app.CarDto{
	// 		Make:  "Toyota",
	// 		Model: "Explorer",
	// 		Vin:   "43133409A53100",
	// 		Year:  1999,
	// 		Miles: 146100,
	// 	},
	// 	Height: 7.1,
	// }
	// b, _ = json.Marshal(truck)
	// fmt.Println(string(b))

	data, _ := ioutil.ReadFile("data/cars.json")
	cars := core.CarCollectionDto{}
	json.Unmarshal(data, &cars)

	data, _ = ioutil.ReadFile("data/trucks.json")
	trucks := core.TruckCollectionDto{}
	json.Unmarshal(data, &trucks)

  // TODO: create http lib for shorthand http functionality
  // that or inject controller into handler, that way we can manage everything from context?

  core.AddResource("/", "root", "", func(w http.ResponseWriter, r *http.Request) {
    sd := core.GetServiceDescription(r)
    processed, _ := json.Marshal(sd)

    switch r.Method {
    case "GET":
      w.Header().Set("Content-Type", "application/json")
      fmt.Fprintf(w, string(processed))
    }
  })

  core.AddResource("/cars/", "cars-collection", "", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
      resp, _ := json.Marshal(cars)
      w.Header().Set("Content-Type", "application/json")
      fmt.Fprintf(w, string(resp))
    case "POST":
      // TODO: left here, figure out how to read request body
      w.Header().Set("Content-Type", "application/json")
      fmt.Fprintf(w, "[]")
    }
  })

  core.AddResource("/trucks/", "trucks-collection", "", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
      resp, _ := json.Marshal(trucks)
      w.Header().Set("Content-Type", "application/json")
      fmt.Fprintf(w, string(resp))
    case "POST":
      // TODO: left here, figure out how to read request body
      w.Header().Set("Content-Type", "application/json")
      fmt.Fprintf(w, "[]")
    }
  })

  core.AddResource("/vehicles/", "vehicles-collection", "", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
      resp, _ := json.Marshal(core.VehicleCollectionDto{
        Cars:   cars,
        Trucks: trucks,
      })
      w.Header().Set("Content-Type", "application/json")
      fmt.Fprintf(w, string(resp))
    case "POST":
      // TODO: left here, figure out how to read request body
      w.Header().Set("Content-Type", "application/json")
      fmt.Fprintf(w, "[]")
    }
  })

	log.Fatal(http.ListenAndServe(":8080", nil))
}
