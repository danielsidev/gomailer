# gomailer
### Last version: **v1.0.5**
 This is a module to send email with your email provider or Gmail.
 
 ## Requirements to Gmail

 >  First, It's necessary to config your gmail. Configure 2step verfication and register a second email and a phone to recovery your data.

 >  Second, access this link: **[AppPassword](https://myaccount.google.com/apppasswords)** to create a password app(**this is so important!**). We need this appPassword.

 ## Excute

 Create a new project, add this dependecy(**check the last version**):
 ```
 go get github.com/danielsidev/gomailer@v1.0.5

 ```

 ### About Envs

 Wen need 2 envs:

 - **EMAIL_SENDER_FROM** this is sender
	
 - **EMAIL_SENDER_PASSWORD** this is sender's password (In the Gmail usecase, is the appPassword created)  

 **So, export these envs**

 ```
 export EMAIL_SENDER_FROM="my_email_sender"
 ```
 ```
 export EMAIL_SENDER_PASSWORD="my_email_sender_password"
 ```

 ### Use Case  Example
- **my_project**
   - **main.go**

 ```
package main

import (
	"log"

	"github.com/danielsidev/gomailer"
)

func main() {
	css := "p{ font-size: 16px;}"
    /*
    It's necessary a public image hosted.
    Can you create a github page from repo and make upload your image, like this example below:
    logo := https://danielsidev.github.io/repo-images/my_image.png    
    */
	logo := "your full image path" 
	var data gomailer.DataEmail = gomailer.DataEmail{
		EmailTo:      "your recipient",
		EmailBody:    "your html body",
		Username:     "your recipient name",
		SenderName:   "Who sends the email",
		EmailSubject: "Subject",
        SmtpHost:     "smtp.gmail.com",  // default gmail: "smtp.gmail.com"
        SmtpPort:     "587",            // default gmail: "587" // Porta TLS/STARTTLS
		Css:          &css,            // optional
		Logo:         &logo,          // optional
	}
	erro := gomailer.SendEmail(data)
	if erro != nil {
		log.Printf("Ops! Houston, we have a problem::ERROR: %v", erro)
		return
	}
	log.Println("Email stn with success!")

}

 ```
