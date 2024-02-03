package model

type UpdateRelayRequest struct {
	Value int `json:"value" validate:"required"`
}