# Go WhatsApp Business API Wrapper
Currently supports template messages. I wrote this to be easily extensible. Pull requests welcome.

```golang
wa := NewWhatsApp(phoneNumberID, authToken)

msg := wa.NewTemplateMessage(
    "my_template",
    "193891938482",
    "en",
)
msg.Header().AddText("Verification Code")
msg.Body().AddText("111-111")

msg.Send()
```