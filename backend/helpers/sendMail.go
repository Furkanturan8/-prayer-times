package helpers

import (
	"fmt"
	gomail "gopkg.in/mail.v2"
	"os"
	"strconv"
)

func SendMailToAdmin(from, name, surname, email, message string) error {
	// SMTP sunucusu ve kimlik bilgileri
	smtpServer := os.Getenv("SMTP_SERVER")
	_smtpPort := os.Getenv("SMTP_PORT")
	smtpEmail := os.Getenv("EMAIL_ADDRESS")
	smtpPassword := os.Getenv("EMAIL_PASSWORD")

	smtpPort, err := strconv.Atoi(_smtpPort)
	if err != nil {
		fmt.Errorf("smtp port veri tipi int a çevilmedi", err)
		return err
	}
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
