package service_category

import "github.com/steebchen/keskin-api/prisma"

type ServiceCategoryMutation struct {
	Prisma *prisma.Client
}

func New(client *prisma.Client) *ServiceCategoryMutation {
	return &ServiceCategoryMutation{
		Prisma: client,
	}
}
