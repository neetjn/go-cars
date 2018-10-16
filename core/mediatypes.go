package app

type LinkDto struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

type ServiceDescriptionDto struct {
	Links           []LinkDto `json:"links"`
	HealthEndpoints []LinkDto `json:"healthEndpoints"`
}

type CarDto struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Vin   string `json:"vin"`
	Year  int    `json:"year"`
	Miles int    `json:"miles"`
}

type CarCollectionDto struct {
	Items []CarDto `json:"items"`
}

type TruckDto struct {
	CarDto

	Height float32 `json:"height"`
}

type TruckCollectionDto struct {
	Items []TruckDto `json:"items"`
}

type VehicleCollectionDto struct {
	Cars   CarCollectionDto   `json:"cars"`
	Trucks TruckCollectionDto `json:"trucks"`
}
