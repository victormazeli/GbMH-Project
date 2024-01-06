package product

import (
	"context"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *Product) Category(ctx context.Context, obj *prisma.Product) (*prisma.ProductCategory, error) {
	if obj.Deleted {
		return nil, nil
	}
	category, err := r.Prisma.Product(prisma.ProductWhereUniqueInput{
		ID: &obj.ID,
	}).Category().Exec(ctx)

	if err != nil {
		return nil, nil
	}

	return category, nil

}
