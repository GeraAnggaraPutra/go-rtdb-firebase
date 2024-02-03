package service

import (
	"context"
	"strconv"
)

func (s *RelayService) UpdateRelayValueSvc(
	ctx context.Context,
	relayNumber, value int,
) (err error) {
	ref := s.dbClient.NewRef("board1/outputs/digital").Child(strconv.Itoa(relayNumber))

	err = ref.Update(ctx, map[string]interface{}{"value": value})
	if err != nil {
		return
	}

	return
}
