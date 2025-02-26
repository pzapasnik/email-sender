// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Root() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><title></title><script src=\"/static/lib/htmx.min.js\"></script><script src=\"/static/lib/tailwindcss@4.min.js\"></script><script src=\"/static/components/wysiwyg.js\" defer></script></head><body class=\"bg-slate-500\"><div class=\"h-screen flex items-center justify-center\"><div class=\"bg-gray-400 min-w-2/4 min-h-9/10 rounded-md shadow-xl p-5\"><form class=\"flex flex-col\"><p class=\"px-2\"><label for=\"from\">Nadawca:</label></p><p class=\"px-2\"><input class=\"bg-slate-500\" type=\"text\" name=\"from\" value=\"pzapasnik@gmail.com\"></p><p class=\"px-2\"><label for=\"to\">Odbiorcy:</label></p><p class=\"px-2\"><input class=\"bg-slate-500\" type=\"text\" name=\"to\" value=\"\"></p><p><wysiwyg-editor></wysiwyg-editor></p></form></div></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
