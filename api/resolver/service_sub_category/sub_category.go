package service_sub_category

import (
	"github.com/google/wire"
	"github.com/steebchen/keskin-api/prisma"
)

type ServiceSubCategory struct {
	Prisma *prisma.Client
}

func New(client *prisma.Client) *ServiceSubCategory {
	return &ServiceSubCategory{
		Prisma: client,
	}
}

var ProviderSet = wire.NewSet(
	New,
)
