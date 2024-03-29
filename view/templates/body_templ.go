// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Body() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container\"><div class=\"row\"><div class=\"text-center\"><h2>Access Request App</h2></div></div><div class=\"row\"><div class=\"col-8\"><h2>Active Requests</h2></div><div class=\"col-4\"><h2>Add a request</h2></div><div class=\"row\"><div class=\"col-8\"><table class=\"table\"><thead><tr><th>User</th><th>Provider</th><th>Environment</th><th>Grant</th></tr></thead></table></div><div class=\"col-4\"><form id=\"request-form\"><div class=\"form-group\"><label for=\"provider\">Provider</label> <input type=\"text\" class=\"form-control\" id=\"provider\" name=\"provider\" required></div><div class=\"form-group\"><label for=\"environment\">Environment</label> <input type=\"text\" class=\"form-control\" id=\"environment\" name=\"environment\" required></div><div class=\"form-group\"><label for=\"grant\">Grant</label> <input type=\"text\" class=\"form-control\" id=\"grant\" name=\"grant\" required></div><div class=\"form-group\"><button type=\"submit\" class=\"btn btn-primary\">Submit</button></div></form></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
