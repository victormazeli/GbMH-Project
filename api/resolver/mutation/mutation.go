package mutation

import (
	"github.com/google/wire"

	"github.com/steebchen/keskin-api/api/resolver/mutation/appointment"
	"github.com/steebchen/keskin-api/api/resolver/mutation/auth"
	"github.com/steebchen/keskin-api/api/resolver/mutation/branch"
	"github.com/steebchen/keskin-api/api/resolver/mutation/company"
	"github.com/steebchen/keskin-api/api/resolver/mutation/email_template"
	"github.com/steebchen/keskin-api/api/resolver/mutation/favorite"
	"github.com/steebchen/keskin-api/api/resolver/mutation/mailchimp"
	"github.com/steebchen/keskin-api/api/resolver/mutation/news"
	"github.com/steebchen/keskin-api/api/resolver/mutation/notification"
	"github.com/steebchen/keskin-api/api/resolver/mutation/product"
	"github.com/steebchen/keskin-api/api/resolver/mutation/product_category"
	"github.com/steebchen/keskin-api/api/resolver/mutation/product_sub_category"
	"github.com/steebchen/keskin-api/api/resolver/mutation/review"
	"github.com/steebchen/keskin-api/api/resolver/mutation/service"
	"github.com/steebchen/keskin-api/api/resolver/mutation/service_category"
	"github.com/steebchen/keskin-api/api/resolver/mutation/service_sub_category"
	"github.com/steebchen/keskin-api/api/resolver/mutation/user/administrator"
	"github.com/steebchen/keskin-api/api/resolver/mutation/user/customer"
	"github.com/steebchen/keskin-api/api/resolver/mutation/user/employee"
	"github.com/steebchen/keskin-api/api/resolver/mutation/user/manager"
	"github.com/steebchen/keskin-api/api/resolver/mutation/user/viewer"
	"github.com/steebchen/keskin-api/prisma"
)

type Mutation struct {
	Prisma *prisma.Client
	*auth.Auth
	*appointment.Appointment
	*branch.BranchMutation
	*company.CompanyMutation
	*customer.CustomerMutation
	*product_category.ProductCategoryMutation
	*product_sub_category.ProductSubCategoryMutation
	*service_category.ServiceCategoryMutation
	*service_sub_category.ServiceSubCategoryMutation
	*employee.EmployeeMutation
	*product.ProductMutation
	*service.ServiceMutation
	*viewer.ViewerMutation
	*favorite.FavoriteMutation
	*notification.Notification
	*news.NewsMutation
	*email_template.EmailTemplateMutation
	*review.ReviewMutation
	*administrator.AdministratorMutation
	*manager.ManagerMutation
	*mailchimp.MailchimpMutation
}

func New(
	client *prisma.Client,
	auth *auth.Auth,
	appointment *appointment.Appointment,
	branch *branch.BranchMutation,
	company *company.CompanyMutation,
	productCategory *product_category.ProductCategoryMutation,
	productSubCategory *product_sub_category.ProductSubCategoryMutation,
	serviceCategory *service_category.ServiceCategoryMutation,
	serviceSubCategory *service_sub_category.ServiceSubCategoryMutation,
	customer *customer.CustomerMutation,
	employee *employee.EmployeeMutation,
	product *product.ProductMutation,
	service *service.ServiceMutation,
	viewer *viewer.ViewerMutation,
	favorite *favorite.FavoriteMutation,
	notification *notification.Notification,
	news *news.NewsMutation,
	emailTemplate *email_template.EmailTemplateMutation,
	review *review.ReviewMutation,
	administrator *administrator.AdministratorMutation,
	manager *manager.ManagerMutation,
	mailchimp *mailchimp.MailchimpMutation,
) *Mutation {
	return &Mutation{
		client,
		auth,
		appointment,
		branch,
		company,
		customer,
		productCategory,
		productSubCategory,
		serviceCategory,
		serviceSubCategory,
		employee,
		product,
		service,
		viewer,
		favorite,
		notification,
		news,
		emailTemplate,
		review,
		administrator,
		manager,
		mailchimp,
	}
}

var ProviderSet = wire.NewSet(
	New,
	auth.New,
	appointment.New,
	branch.New,
	company.New,
	product_category.New,
	product_sub_category.New,
	service_category.New,
	service_sub_category.New,
	customer.New,
	employee.New,
	product.New,
	service.New,
	viewer.New,
	favorite.New,
	notification.New,
	news.New,
	email_template.New,
	review.New,
	administrator.New,
	manager.New,
	mailchimp.New,
)
