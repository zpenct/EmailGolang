## Send Email 
This is a simple service to send an email to your emails. Written in Golang using gomail library.
Used in [Sendawa Hackfest 2024](https://github.com/zpenct/sendawa) for more details.

### Prerequisites
- Login in to gmail account
- Enter a security > app password
- Enter your password 
- They will generate a random password 

### Usage
- Clone this repository
- Make a .env
- Enter your Email and password(generated before)
- Check branch main
  
### How to send email?
URL : localhost:8080/send-email
METHOD: POST
BODY:
```json
{
  "to": "Some@gmail.com",
  "subject": "Tes Email Sender Golang",
  "body": "Tess "
}
```

Response :
 200 OK
 ```json
{
    "message": "Email berhasil terkirim"
}
```
Example:
![Example](doc/send-email.jpeg)
