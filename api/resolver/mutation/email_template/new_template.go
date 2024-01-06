package email_template

import (
	"context"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *EmailTemplateMutation) NewEmailTemplate(
	ctx context.Context,
) (string, error) {

	_, err := r.Prisma.CreateEmailTemplate(prisma.EmailTemplateCreateInput{
		Name: "appointmentBooked",
		Title: prisma.LocalizedStringCreateOneInput{
			Create: &prisma.LocalizedStringCreateInput{
				De: prisma.Str("Termin gebucht"),
				En: prisma.Str("Appointment Booked"),
				Tr: prisma.Str("Randevu Alındı"),
			},
		},

		Content: prisma.LocalizedStringCreateOneInput{
			Create: &prisma.LocalizedStringCreateInput{
				De: prisma.Str("{{salutation}} {{lastName}},\n\nWir freuen uns, Ihnen mitteilen zu können, dass ein Kunde einen Termin für unsere Dienstleistungen vereinbart hat. Die Details lauten wie folgt:\n\n- Kundenname: {{customerName}}\n- Termin Datum: {{appointmentDate}}\n- Terminzeit: {{appointmentTime}}\n\nBitte treffen Sie die notwendigen Vorbereitungen, um die angeforderten Dienstleistungen zu erbringen.\n\nUm diesen Termin anzunehmen und zu verwalten, melden Sie sich bitte in Ihrem Dashboard an. Ihre prompte Aufmerksamkeit in dieser Angelegenheit wird sehr geschätzt.\n\nMit freundlichen Grüßen, {{companyName}}-Team"),
				En: prisma.Str("{{salutation}} {{firstName}},\n\nWe are pleased to inform you that a customer has booked an appointment for our services. The details are as follows:\n\n- Customer Name: {{customerName}}\n- Appointment Date: {{appointmentDate}}\n- Appointment Time: {{appointmentTime}}\n\nPlease make the necessary preparations to provide the requested services.\n\nTo accept and manage this appointment, kindly log in to your dashboard. Your prompt attention to this matter is greatly appreciated.\n\nBest regards,\n{{companyName}}"),
				Tr: prisma.Str("Merhaba {{firstname}} {{salutation}},\n\nSizi bilgilendirmekten memnuniyet duyarız ki, bir müşteri hizmetlerimiz için bir randevu almıştır. Detaylar şu şekildedir:\n\n- Müşteri Adı: {{customerName}}\n- Randevu Tarihi: {{appointmentDate}}\n- Randevu Saati: {{appointmentTime}}\n\nLütfen istenen hizmetleri sağlamak için gerekli hazırlıkları yapın.\n\nBu randevuyu kabul etmek ve yönetmek için lütfen dashboard'unuza giriş yapın. Bu konudaki hızlı tepkiniz çok takdir edilmektedir.\n\nSaygılarımla,\n{{companyName}}"),
			},
		},
	}).Exec(ctx)
	if err != nil {
		return "", err
	}
	output := "completed"
	return output, nil

}
