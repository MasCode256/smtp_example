package main

import (
	"fmt"
	"net/smtp"
)

// sendEmail отправляет электронное письмо.
func sendEmail(to string, subject string, body string) error {
	fmt.Println("Setting up the SMTP connection...")
	// Настройки SMTP-сервера
	smtpHost := "smtp.mail.ru" // Замените на ваш SMTP-сервер
	smtpPort := ":587"               // Порт SMTP-сервера
	from := "my_address@mail.ru" // Ваш адрес электронной почты
	password := "password" // Ваш пароль

	// Формирование сообщения
	message := []byte("From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\nContent-Type: text/html; charset=\"utf-8\"" + "\n\n" +
		body)

	// Аутентификация
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Отправка письма
	err := smtp.SendMail(smtpHost+smtpPort, auth, from, []string{to}, message)
	return err
}

func main() {
	to := "receiver@mail.ru" // Замените на адрес получателя
	subject := "Hello, world!"
	body := "<h1>This is content of the message.</h1>"

	err := sendEmail(to, subject, body)
	if err != nil {
		fmt.Println("Error sending email:", err)
	} else {
		fmt.Println("Email sent successfully!")
	}
}
