package product

import (
	"context"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *Product) Active(ctx context.Context, obj *prisma.Product) (*bool, error) {
	result := false

	if obj.Deleted {
		return &result, nil
	}

	return &obj.Active, nil
}
