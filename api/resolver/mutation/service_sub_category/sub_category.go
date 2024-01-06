package service_sub_category
import (
	"github.com/steebchen/keskin-api/prisma"
)

type ServiceSubCategoryMutation struct {
	Prisma *prisma.Client
}

func New(client *prisma.Client) *ServiceSubCategoryMutation {
	return &ServiceSubCategoryMutation{
		Prisma: client,
	}
}
