package product

import (
	"context"

	"github.com/steebchen/keskin-api/prisma"
)

func (r *Product) SubCategory(ctx context.Context, obj *prisma.Product) (*prisma.ProductSubCategory, error) {
	if obj.Deleted {
		return nil, nil
	}
	subCategory, err := r.Prisma.Product(prisma.ProductWhereUniqueInput{
		ID: &obj.ID,
	}).SubCategory().Exec(ctx)

	if err != nil {
		return nil, nil
	}

	return subCategory, nil

}
