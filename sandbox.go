package main

import (
	"encoding/json"
	"fmt"
	// "os"
	"io/ioutil"
	app "sandbox/app"
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
	cars := app.CarCollectionDto{}
	json.Unmarshal(data, &cars)
	// fmt.Println(cars)
	for _, e := range cars.Items {
		fmt.Println(e.Model)
	}
}
