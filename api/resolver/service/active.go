package service

import (
	"context"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *Service) Active(ctx context.Context, obj *prisma.Service) (*bool, error) {
	result := false

	if obj.Deleted {
		return &result, nil
	}

	return &obj.Active, nil
}
