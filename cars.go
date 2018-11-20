package main

import (
	"encoding/json"
	"fmt"
	"log"
  "io/ioutil"
	"net/http"

	core "github.com/neetjn/go-cars/core"
  utils "github.com/neetjn/go-cars/utils"
  // zap "go.uber.org/zap"
)

// TODO: add error utilities to construct error payload responses for easy consumption

func main() {

	data, _ := ioutil.ReadFile("data/cars.json")
	var cars *core.CarCollectionDto = &core.CarCollectionDto{}
	json.Unmarshal(data, &cars)

	data, _ = ioutil.ReadFile("data/trucks.json")
	var trucks *core.TruckCollectionDto = &core.TruckCollectionDto{}
	json.Unmarshal(data, &trucks)

  core.AddResource("/", "root", "", func(w http.ResponseWriter, r *http.Request) {
    sd := core.GetServiceDescription(r)
    processed, _ := json.Marshal(sd)

    switch r.Method {
    case http.MethodGet:
      w.Header().Set("Content-Type", "application/json")
      fmt.Fprintf(w, string(processed))
    }
  })

  core.AddResource("/car/", "car", "", func(w http.ResponseWriter, r *http.Request) {
    vin := r.URL.Query().Get("vin")
    if !utils.ValidateVin(vin) {
      http.Error(w, "Invalid Vin Provided", http.StatusInternalServerError)
      return
    }
    switch r.Method {
    case http.MethodGet:
      for _, car := range cars.Items {
        if car.Vin == vin {
          resp, _ := json.Marshal(car)
          w.Header().Set("Content-Type", "application/json")
          fmt.Fprintf(w, string(resp))
          return
        }
      }
      http.Error(w, "Car Does Not Exist", http.StatusInternalServerError)
    }
  })

  core.AddResource("/cars/", "cars-collection", "", func(w http.ResponseWriter, r *http.Request) {
    // TODO: add check for existing car by unique vin on post
    switch r.Method {
    case http.MethodGet:
      resp, _ := json.Marshal(cars)
      w.Header().Set("Content-Type", "application/json")
      fmt.Fprintf(w, string(resp))
    case http.MethodPost:
      body, _ := ioutil.ReadAll(r.Body)
      defer r.Body.Close()
      var car *core.CarDto = &core.CarDto{}
      err := json.Unmarshal(body, *car)
      if err != nil {
        cars.Items = append(cars.Items, *car)
        w.WriteHeader(http.StatusCreated)
      } else {
        http.Error(w, err.Error(), http.StatusInternalServerError)
      }
    }
  })

  core.AddResource("/trucks/", "trucks-collection", "", func(w http.ResponseWriter, r *http.Request) {
    // TODO: add check for existing truck by unique vin on post
    switch r.Method {
    case http.MethodGet:
      resp, _ := json.Marshal(trucks)
      w.Header().Set("Content-Type", "application/json")
      fmt.Fprintf(w, string(resp))
    case http.MethodPost:
      body, _ := ioutil.ReadAll(r.Body)
      defer r.Body.Close()
      var truck *core.TruckDto = &core.TruckDto{}
      err := json.Unmarshal(body, *truck)
      if err != nil {
        trucks.Items = append(trucks.Items, *truck)
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
        Cars:   *cars,
        Trucks: *trucks,
      })
      w.Header().Set("Content-Type", "application/json")
      fmt.Fprintf(w, string(resp))
    }
  })

	log.Fatal(http.ListenAndServe(":8080", nil))
}
