package helper

import (
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

//SendErrorMail send mail 收件者是Env.maintainer
func SendErrorMail(msg string) {
	fmt.Println("[Helper] send mail")
	m := gomail.NewMessage()
	m.SetHeader("From", viper.GetString("SMTP.from"))
	m.SetHeader("To", viper.GetString("SMTP.to"))
	m.SetHeader("Subject", "iBP test Mail")
	m.SetBody("text/html", msg)

	d := gomail.NewDialer(viper.GetString("SMTP.host"), viper.GetInt("SMTP.port"), viper.GetString("SMTP.user"), viper.GetString("SMTP.pass"))

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
