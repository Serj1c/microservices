package handlers

import (
	"log"
	"net/http"
)

// Products ...
type Products struct {
	l *log.Logger
}

// NewProducts ...
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

}
