// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"time"
)

type Material struct {
	ID                     int       `json:"Id"`
	NodeID                 string    `json:"NodeId"`
	Name                   string    `json:"Name"`
	Unit                   string    `json:"Unit"`
	Quantity               string    `json:"Quantity"`
	CreatedTime            time.Time `json:"CreatedTime"`
	OwnerPublicKey         string    `json:"OwnerPublicKey"`
	PreviousNodesHashedIds []*string `json:"PreviousNodesHashedIds"`
	NextNodesHashedIds     []*string `json:"NextNodesHashedIds"`
}

type MaterialGraph struct {
	MainMaterial     *Material   `json:"MainMaterial"`
	RelatedMaterials []*Material `json:"RelatedMaterials"`
}

type Peer struct {
	ID         int          `json:"Id"`
	Alias      string       `json:"Alias"`
	PublicKeys []*PublicKey `json:"PublicKeys"`
}

type PublicKey struct {
	ID    int    `json:"Id"`
	Value string `json:"Value"`
}

type ReceiveMaterialRequestRequest struct {
	TransferMaterial *Material   `json:"TransferMaterial"`
	ExposedMaterials []*Material `json:"ExposedMaterials"`
	TransferTime     time.Time   `json:"TransferTime"`
	SenderPublicKey  string      `json:"SenderPublicKey"`
}

type ReceiveMaterialRequestResponse struct {
	Accepted bool                           `json:"Accepted"`
	Request  *ReceiveMaterialRequestRequest `json:"Request"`
}
