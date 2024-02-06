package service

import (
	"context"
)

func (s *RelayService) ReadListRelaySvc(
	ctx context.Context,
) (data map[int64]interface{}, err error) {
	ref := s.dbClient.NewRef("board1/outputs/digital")

	err = ref.Get(ctx, &data)
	if err != nil {
		return
	}

	return
}
