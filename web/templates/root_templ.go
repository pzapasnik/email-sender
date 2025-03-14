// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Root(send bool) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><title>email sender</title><script type=\"module\" src=\"/static/lib/htmx.min.js\"></script><script type=\"module\" src=\"/static/lib/tailwindcss@4.min.js\"></script><script type=\"module\" src=\"/static/components/wysiwyg.js\" defer></script></head><style>\n\twysiwyg-editor::part(toolbar) {\n\t\tbackground-color: #4a5568;\n\t}\n</style><body class=\"bg-slate-500\"><div class=\"h-screen flex items-center justify-center\"><div class=\"bg-gray-400 w-2/4 h-9/10 rounded-md shadow-xl p-5\"><form class=\"flex flex-col h-full\" action=\"email/send\" method=\"post\"><div><p class=\"px-2 flex\"><label class=\"my-auto\" for=\"from\">Nadawca:</label> <input class=\"bg-slate-500 flex-auto ml-2 rounded-sm px-2 py-1\" type=\"text\" name=\"from\" value=\"pzapasnik@gmail.com\"></p><p class=\"px-2 my-2 flex\"><label class=\"my-auto\" for=\"to\">Odbiorcy:</label> <input class=\"bg-slate-500 flex-auto ml-2 rounded-sm px-2 py-1\" type=\"text\" name=\"to\" value=\"\"></p></div><div class=\"flex flex-col h-full\"><wysiwyg-editor class=\"h-9/10\"><label class=\"my-2\" slot=\"label\" for=\"message\">Wiadomość:</label><div class=\"flex h-10 bg-slate-500\" slot=\"toolbar\"></div><textarea class=\"bg-slate-500 rounded-sm px-2 py-1 h-full\" slot=\"message\" name=\"message\" placeholder=\"Napisz coś...\"></textarea></wysiwyg-editor><p class=\"my-4 flex justify-end\"><button class=\" bg-cyan-800 hover:bg-cyan-900 rounded-md px-4 py-2\" type=\"submit\">Wyślij</button></p></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if send {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<div>SEND</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "</form></div></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
