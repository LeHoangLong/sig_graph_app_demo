// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"time"
)

type Material struct {
	ID          string    `json:"Id"`
	Name        string    `json:"Name"`
	Unit        string    `json:"Unit"`
	Quantity    string    `json:"Quantity"`
	CreatedTime time.Time `json:"CreatedTime"`
}

type Peer struct {
	ID    int          `json:"Id"`
	Alias string       `json:"Alias"`
	Keys  []*PublicKey `json:"Keys"`
}

type PublicKey struct {
	ID    int    `json:"Id"`
	Value string `json:"Value"`
}
