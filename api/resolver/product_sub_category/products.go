package product_sub_category

import (
	"context"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *ProductSubCategory) Products(ctx context.Context, obj *prisma.ProductSubCategory) ([]*prisma.Product, error) {

	products, err := r.Prisma.Products(&prisma.ProductsParams{
		Where: &prisma.ProductWhereInput{
			SubCategory: &prisma.ProductSubCategoryWhereInput{
				ID: &obj.ID,
			},
		},
	}).Exec(ctx)

	if err != nil {
		return nil, nil
	}

	var nodes []*prisma.Product
	for _, product := range products {
		clone := product
		nodes = append(nodes, &clone)
	}

	return nodes, nil
}
