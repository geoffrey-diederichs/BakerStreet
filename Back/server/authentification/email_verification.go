package authentification

import (
	"net/smtp"
)

func sendVerificationEmail(to, token string) error {
	from := "98f90e1e5ec36d"
	smtpHost := "sandbox.smtp.mailtrap.io"
	smtpPort := "2525"

	message := []byte("To: " + to + "\r\n" +
		"Subject: Verify your account\r\n\r\n" +
		"Copy and paste this code in the form" + token)

	auth := smtp.PlainAuth("", from, "2cf349ebee1dcb", smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil {
		return err
	}
	return nil
}
