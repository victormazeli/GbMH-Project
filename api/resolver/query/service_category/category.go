package service_category

import (
	"github.com/google/wire"
	"github.com/steebchen/keskin-api/prisma"
)

type ServiceCategoryQuery struct {
	Prisma *prisma.Client
}

func New(client *prisma.Client) *ServiceCategoryQuery {
	return &ServiceCategoryQuery{
		Prisma: client,
	}
}

var ProviderSet = wire.NewSet(
	New,
)
