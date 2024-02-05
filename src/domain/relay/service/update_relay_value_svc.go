package service

import (
	"context"
	"log"
	"strconv"

	"firebase.google.com/go/v4/db"
)

type RelayData struct {
	Value int `json:"value"`
}

func (s *RelayService) UpdateRelayValueSvc(
	ctx context.Context,
	relayNumber, value int,
) (err error) {
	ref := s.dbClient.NewRef("board1/outputs/digital").Child(strconv.Itoa(relayNumber))

	updateRelayValue := func(tn db.TransactionNode) (interface{}, error) {
		var relayData RelayData
		if err := tn.Unmarshal(&relayData); err != nil {
			return nil, err
		}

		relayData.Value = value

		return relayData, nil
	}

	if err := ref.Transaction(ctx, updateRelayValue); err != nil {
		log.Printf("failed update relay value with err %v", err)
		return err
	}

	return
}
