package helpers

import (
	"fmt"
	gomail "gopkg.in/mail.v2"
	"namaz-vakitleri/pkg/config"
)

func SendMailToAdmin(cfg *config.Config, from, name, surname, email, message string) error {
	// SMTP sunucusu ve kimlik bilgileri
	smtpServer := cfg.SMTPServer
	smtpPort := cfg.SMTPPort
	smtpEmail := cfg.EmailAddress
	smtpPassword := cfg.EmailPassword

	if smtpPassword == "" {
		return fmt.Errorf("SMTP şifresi çevresel değişkende bulunamadı")
	}

	// Gönderici ve alıcı bilgileri
	to := smtpEmail

	// E-posta içeriği
	subject := "Namaz Vakitleri Web Sayfasından Bir Mail Var!"
	body := fmt.Sprintf("Ad: %s\nSoyad: %s\nE-posta: %s\nMesaj: %s", name, surname, email, message)

	// E-posta mesajı oluşturma
	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", body)

	// SMTP sunucusuna bağlanma ve e-posta gönderme
	d := gomail.NewDialer(smtpServer, smtpPort, smtpEmail, smtpPassword)
	d.StartTLSPolicy = gomail.MandatoryStartTLS // STARTTLS kullanımı

	// E-posta gönder
	if err := d.DialAndSend(msg); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
