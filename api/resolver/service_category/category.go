package service_category

import (
	"github.com/google/wire"
	"github.com/steebchen/keskin-api/prisma"
)

type ServiceCategory struct {
	Prisma *prisma.Client
}

func New(client *prisma.Client) *ServiceCategory {
	return &ServiceCategory{
		Prisma: client,
	}
}

var ProviderSet = wire.NewSet(
	New,
)
