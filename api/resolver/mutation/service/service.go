package service

import (
	"github.com/steebchen/keskin-api/prisma"
)

type ServiceMutation struct {
	Prisma *prisma.Client
}

var allowedTypes = []prisma.UserType{
	prisma.UserTypeManager,
	prisma.UserTypeAdministrator,
}

func New(client *prisma.Client) *ServiceMutation {
	return &ServiceMutation{
		Prisma: client,
	}
}
