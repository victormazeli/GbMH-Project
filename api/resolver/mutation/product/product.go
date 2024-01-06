package product

import (
	"github.com/steebchen/keskin-api/prisma"
)

type ProductMutation struct {
	Prisma *prisma.Client
}

var allowedTypes = []prisma.UserType{
	prisma.UserTypeManager,
	prisma.UserTypeAdministrator,
}

func New(client *prisma.Client) *ProductMutation {
	return &ProductMutation{
		Prisma: client,
	}
}
