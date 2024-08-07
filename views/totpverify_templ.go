// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func MakeTOTPVerifyPage(email string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html><head><title>Two-Factor Authentication</title><style>\n    body {\n      font-family: Arial, sans-serif;\n      display: flex;\n      justify-content: center;\n      align-items: center;\n      height: 100vh;\n      margin: 0;\n      background-color: #f0f0f0;\n    }\n\n    .container {\n      background-color: white;\n      padding: 2rem;\n      border-radius: 8px;\n      box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);\n      text-align: center;\n      max-width: 400px;\n      width: 100%;\n    }\n\n    h1 {\n      color: #333;\n    }\n\n    input[type=\"text\"] {\n      width: 100%;\n      padding: 0.5rem;\n      margin: 1rem 0;\n      border: 1px solid #ddd;\n      border-radius: 4px;\n    }\n\n    button {\n      background-color: #007bff;\n      color: white;\n      border: none;\n      padding: 0.5rem 1rem;\n      border-radius: 4px;\n      cursor: pointer;\n    }\n\n    button:hover {\n      background-color: #0056b3;\n    }\n  </style></head><body><div class=\"container\"><h1>Two-Factor Authentication</h1><p>Enter the code from your authenticator app:</p><form action=\"/totp-verify\" method=\"POST\"><input type=\"text\" name=\"otp_code\" placeholder=\"Enter TOTP code\" required> <input type=\"hidden\" name=\"account_name\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(email))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" required> <button type=\"submit\">Verify</button></form></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
