package entities

import "net/url"

type MockEntities struct {
	Message string   `json:"message"`
	Path    *url.URL `json:"path"`
}
