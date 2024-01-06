package product_category

import (
	"context"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *ProductCategory) Products(ctx context.Context, obj *prisma.ProductCategory) ([]*prisma.Product, error) {
	products, err := r.Prisma.Products(&prisma.ProductsParams{
		Where: &prisma.ProductWhereInput{
			Category: &prisma.ProductCategoryWhereInput{
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
