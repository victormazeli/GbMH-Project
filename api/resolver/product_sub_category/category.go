package product_sub_category

import (
	"context"

	"github.com/steebchen/keskin-api/prisma"
)

func (r *ProductSubCategory) Category(ctx context.Context, obj *prisma.ProductSubCategory) (*prisma.ProductCategory, error) {

	return r.Prisma.ProductSubCategory(prisma.ProductSubCategoryWhereUniqueInput{
		ID: &obj.ID,
	}).Category().Exec(ctx)
}
