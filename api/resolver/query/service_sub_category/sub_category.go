package service_sub_category

import (
	"github.com/google/wire"
	"github.com/steebchen/keskin-api/prisma"
)

type ServiceSubCategoryQuery struct {
	Prisma *prisma.Client
}

func New(client *prisma.Client) *ServiceSubCategoryQuery {
	return &ServiceSubCategoryQuery{
		Prisma: client,
	}
}

var ProviderSet = wire.NewSet(
	New,
)
