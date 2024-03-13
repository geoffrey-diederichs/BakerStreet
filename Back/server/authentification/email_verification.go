package authentification

import (
	"net/smtp"

	"go.uber.org/zap"
)

func sendVerificationEmail(to, token string) error {
	from := "from@example.com"
	smtpHost := "sandbox.smtp.mailtrap.io"
	smtpPort := "2525"
	username := "98f90e1e5ec36d" // Use your Mailtrap username
	password := "2cf349ebee1dcb" // Use your Mailtrap password

	message := []byte("To: " + to + "\r\n" +
		"Subject: Verify your account\r\n\r\n" +
		"Copy and paste this code in the form : " + token)

	auth := smtp.PlainAuth("", username, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil {
		logger.Error("Error sending email: ", zap.Error(err))
		return err
	}
	logger.Info("Email sent successfully")
	return nil
}
