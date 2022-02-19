package helpers

import (
	"context"
	"time"

	mailgun "github.com/mailgun/mailgun-go/v4"
)

func SendInvitationEmail(domain, apiString, sender, subject, recipient, code string) (string, error) {
	// // create an instance of Mailgun Client to send
	mg := mailgun.NewMailgun(domain, apiString)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	templates := mg.ListTemplates(nil)
	var page, results []mailgun.Template
	if templates.Next(ctx, &page) {
		results = append(results, page...)
		for _, r := range results {
			if r.Name == "verifyemail" {
				err := mg.DeleteTemplate(ctx, "verifyemail")
				if err != nil {
					return "", err
				}
			}
		}
	}
	err := mg.CreateTemplate(ctx, &mailgun.Template{
		Name: "verifyemail",
		Version: mailgun.TemplateVersion{
			Template: `
			<!doctype html><html xmlns="http://www.w3.org/1999/xhtml" xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office"><head><title></title><meta http-equiv="X-UA-Compatible" content="IE=edge"><meta http-equiv="Content-Type" content="text/html; charset=UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1"><style type="text/css">#outlook a { padding:0; }
    body { margin:0;padding:0;-webkit-text-size-adjust:100%;-ms-text-size-adjust:100%; }
    table, td { border-collapse:collapse;mso-table-lspace:0pt;mso-table-rspace:0pt; }
    img { border:0;height:auto;line-height:100%; outline:none;text-decoration:none;-ms-interpolation-mode:bicubic; }
    p { display:block;margin:13px 0; }</style><link href="https://fonts.googleapis.com/css?family=Open+Sans:300,400,500,700" rel="stylesheet" type="text/css"><style type="text/css">@import url(https://fonts.googleapis.com/css?family=Open+Sans:300,400,500,700);</style><style type="text/css">@media only screen and (min-width:480px) {
  .mj-column-per-100 { width:100% !important; max-width: 100%; }
}</style><style media="screen and (min-width:480px)">.moz-text-html .mj-column-per-100 { width:100% !important; max-width: 100%; }</style><style type="text/css">@media only screen and (max-width:480px) {
table.mj-full-width-mobile { width: 100% !important; }
td.mj-full-width-mobile { width: auto !important; }
}</style></head><body style="word-spacing:normal;background-color:#fafbfc;"><div style="background-color:#fafbfc;"><div style="margin:0px auto;max-width:600px;"><table align="center" border="0" cellpadding="0" cellspacing="0" role="presentation" style="width:100%;"><tbody><tr><td style="direction:ltr;font-size:0px;padding:20px 0;padding-bottom:20px;padding-top:20px;text-align:center;"><div class="mj-column-per-100 mj-outlook-group-fix" style="font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:middle;width:100%;"><table border="0" cellpadding="0" cellspacing="0" role="presentation" style="vertical-align:middle;" width="100%"><tbody><tr><td align="center" style="font-size:0px;padding:25px;word-break:break-word;"><table border="0" cellpadding="0" cellspacing="0" role="presentation" style="border-collapse:collapse;border-spacing:0px;"><tbody><tr><td style="width:125px;"><img height="auto" src="https://global-uploads.webflow.com/5f059a21d0c1c3278fe69842/5f188b94aebb5983b66610dd_logo-arengu.png" style="border:0;display:block;outline:none;text-decoration:none;height:auto;width:100%;font-size:13px;" width="125"></td></tr></tbody></table></td></tr></tbody></table></div></td></tr></tbody></table></div><div style="background:#ffffff;background-color:#ffffff;margin:0px auto;max-width:600px;"><table align="center" border="0" cellpadding="0" cellspacing="0" role="presentation" style="background:#ffffff;background-color:#ffffff;width:100%;"><tbody><tr><td style="direction:ltr;font-size:0px;padding:20px 0;padding-bottom:20px;padding-top:20px;text-align:center;"><div class="mj-column-per-100 mj-outlook-group-fix" style="font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:middle;width:100%;"><table border="0" cellpadding="0" cellspacing="0" role="presentation" style="vertical-align:middle;" width="100%"><tbody><tr><td align="center" style="font-size:0px;padding:10px 25px;padding-right:25px;padding-left:25px;word-break:break-word;"><div style="font-family:open Sans Helvetica, Arial, sans-serif;font-size:16px;line-height:1;text-align:center;color:#000000;"><span>Hello,</span></div></td></tr><tr><td align="center" style="font-size:0px;padding:10px 25px;padding-right:25px;padding-left:25px;word-break:break-word;"><div style="font-family:open Sans Helvetica, Arial, sans-serif;font-size:16px;line-height:1;text-align:center;color:#000000;">Please use the verification code below on the Organized website:</div></td></tr><tr><td align="center" style="font-size:0px;padding:10px 25px;word-break:break-word;"><div style="font-family:open Sans Helvetica, Arial, sans-serif;font-size:24px;font-weight:bold;line-height:1;text-align:center;color:#000000;">{{.code}}</div></td></tr><tr><td align="center" style="font-size:0px;padding:10px 25px;padding-right:16px;padding-left:25px;word-break:break-word;"><div style="font-family:open Sans Helvetica, Arial, sans-serif;font-size:16px;line-height:1;text-align:center;color:#000000;">If you didn't request this, you can ignore this email or let us know.</div></td></tr><tr><td align="center" style="font-size:0px;padding:10px 25px;padding-right:25px;padding-left:25px;word-break:break-word;"><div style="font-family:open Sans Helvetica, Arial, sans-serif;font-size:16px;line-height:1;text-align:center;color:#000000;">Thanks!<br>Organized team</div></td></tr></tbody></table></div></td></tr></tbody></table></div></div></body></html>
			`,
			Engine: mailgun.TemplateEngineGo,
			Tag:    "v1",
		},
	})
	if err != nil {
		return "", err
	}

	// // gite time for template to show up in the system
	time.Sleep(time.Second * 3)

	message := mg.NewMessage(sender, subject, "")
	message.SetTemplate("verifyemail")
	message.AddRecipient(recipient)

	message.AddVariable("code", code)
	// // send message with a 30 second timeout
	_, id, err := mg.Send(ctx, message)
	if err != nil {
		return "", err
	}
	return id, err
}
