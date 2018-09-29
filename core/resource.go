package app

import (
	"net/http"
)

type Resource struct {
  Route       string `json:"route"`
  HrefRel     string `json:"rel"`
  Description string `json:"description"`
}

var ResourceCollection = make([]Resource, 0)

// TODO: figure out how to decorate with serialization
// clean up handler logic, figure out service description

func AddResource(route string, rel string, desc string, handler func) {
  ResourceCollection = append(ResourceCollection, Resource{
    Route: route,
    HrefRel: rel,
    Description: desc
  })
  http.HandleFunc(route, handler)
}
