package product_category

import (
	"context"

	"github.com/steebchen/keskin-api/prisma"
)

func (r *ProductCategory) SubCategories(ctx context.Context, obj *prisma.ProductCategory) ([]*prisma.ProductSubCategory, error) {
	subCg, err := r.Prisma.ProductSubCategories(&prisma.ProductSubCategoriesParams{
		Where: &prisma.ProductSubCategoryWhereInput{
			Category: &prisma.ProductCategoryWhereInput{
				ID: &obj.ID,
			},
		},
	}).Exec(ctx)

	if err != nil {
		return nil, nil
	}

	nodes := []*prisma.ProductSubCategory{}
	for _, cg := range subCg {
		clone := cg
		nodes = append(nodes, &clone)
	}

	return nodes, nil
}
