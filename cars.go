package main

import (
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"

	core "github.com/neetjn/go-cars/core"
)

func main() {

	data, _ := ioutil.ReadFile("data/cars.json")
	cars := core.CarCollectionDto{}
	json.Unmarshal(data, cars)

	data, _ = ioutil.ReadFile("data/trucks.json")
	trucks := core.TruckCollectionDto{}
	json.Unmarshal(data, trucks)

  // TODO: create http lib for shorthand http functionality
  // that or inject controller into handler, that way we can manage everything from context?

  core.AddResource("/", "root", "", func(w http.ResponseWriter, r *http.Request) {
    sd := core.GetServiceDescription(r)
    processed, _ := json.Marshal(sd)

    switch r.Method {
    case http.MethodGet:
      w.Header().Set("Content-Type", "application/json")
      fmt.Fprintf(w, string(processed))
    }
  })

  core.AddResource("/cars/", "cars-collection", "", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
      resp, _ := json.Marshal(cars)
      w.Header().Set("Content-Type", "application/json")
      fmt.Fprintf(w, string(resp))
    case http.MethodPost:
      body, _ := ioutil.ReadAll(r.Body)
      defer r.Body.Close()
      car := core.CarDto{}
      err := json.Unmarshal(body, car)
      fmt.Println(err.Error())
      if err != nil {
        cars.Items = append(cars.Items, car)
        w.WriteHeader(http.StatusCreated)
      } else {
        http.Error(w, err.Error(), http.StatusInternalServerError)
      }
    }
  })

  core.AddResource("/trucks/", "trucks-collection", "", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
      resp, _ := json.Marshal(trucks)
      w.Header().Set("Content-Type", "application/json")
      fmt.Fprintf(w, string(resp))
    case http.MethodPost:
      body, _ := ioutil.ReadAll(r.Body)
      defer r.Body.Close()
      truck := core.TruckDto{}
      err := json.Unmarshal(body, truck)
      if err != nil {
        trucks.Items = append(trucks.Items, truck)
        w.WriteHeader(http.StatusCreated)
      } else {
        http.Error(w, err.Error(), http.StatusInternalServerError)
      }
    }
  })

  core.AddResource("/vehicles/", "vehicles-collection", "", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
      resp, _ := json.Marshal(core.VehicleCollectionDto{
        Cars:   cars,
        Trucks: trucks,
      })
      w.Header().Set("Content-Type", "application/json")
      fmt.Fprintf(w, string(resp))
    case http.MethodPost:
      // TODO: left here, figure out how to read request body
      w.Header().Set("Content-Type", "application/json")
      fmt.Fprintf(w, "[]")
    }
  })

	log.Fatal(http.ListenAndServe(":8080", nil))
}
