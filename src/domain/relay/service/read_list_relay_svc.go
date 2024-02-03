package service

import (
	"context"

	"go-firebase/src/domain/relay/model"
)

func (s *RelayService) ReadListRelaySvc(
	ctx context.Context,
) (data map[string]model.RelayState, err error) {
	ref := s.dbClient.NewRef("board1/outputs/digital")

	err = ref.Get(ctx, &data)
	if err != nil {
		return
	}

	return
}
