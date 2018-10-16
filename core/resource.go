package app

import (
	"net/http"
)

type HandlerFn func(w http.ResponseWriter, r *http.Request)

type Resource struct {
  Route       string `json:"route"`
  HrefRel     string `json:"rel"`
  Description string `json:"description"`
}

var ResourceCollection = make([]Resource, 0)

// TODO: figure out how to decorate with serialization
// clean up handler logic, figure out service description

func AddResource(route string, rel string, desc string, handler HandlerFn) {
  ResourceCollection = append(ResourceCollection, Resource{
    Route: route,
    HrefRel: rel,
    Description: desc,
  })
  http.HandleFunc(route, handler)
}

func GetServiceDescription(r *http.Request) ServiceDescriptionDto {
  links := make([]LinkDto, 0)
  for _, resource := range ResourceCollection {
    links = append(links, LinkDto{
      Href: r.Host + resource.Route,
      Rel: resource.HrefRel,
    })
  }
  return ServiceDescriptionDto{
    Links: links,
  }
}
