package main

import (
	"bytes"
	"strings"
	"text/template"
)

var ginTemplate = `
{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}

// 这里定义 handler interface
type {{.ServiceType}}HTTPHandler interface {
{{- range .Methods}}
    {{.Name}}(context.Context, *{{.Request}}) (*{{.Reply}}, error)
{{- end}}
}

// Register{{.ServiceType}}HTTPHandler define http router handle by gin. 
// 注册路由 handler
func Register{{.ServiceType}}HTTPHandler(g *gin.RouterGroup, srv {{.ServiceType}}HTTPHandler) {
    {{- range .Methods}}
    g.{{.Method}}("{{.Path}}", _{{$svrType}}_{{.Name}}{{.Num}}_HTTP_Handler(srv))
    {{- end}}
}

// 定义 handler
// 遍历之前解析到所有 rpc 方法信息
{{range .Methods}}
func _{{$svrType}}_{{.Name}}{{.Num}}_HTTP_Handler(srv {{$svrType}}HTTPHandler) func(c *gin.Context) {
    return func(c *gin.Context) {
        var (
			err error
            in  = new({{.Request}})
            out = new({{.Reply}})
            ctx = context.TODO()
        )

       {{- if .HasBody}}
		if err = c.ShouldBind(in{{.Body}}); err != nil {
            c.AbortWithStatusJSON(400, gin.H{"err": err.Error()})
            return
        }
		
		{{- if not (eq .Body "")}}
		if err = c.ShouldBindQuery(in}); err != nil {
            c.AbortWithStatusJSON(400, gin.H{"err": err.Error()})
            return
        }
		{{- end}}
		{{- else}}
		if err = c.ShouldBindQuery(in); err != nil {
            c.AbortWithStatusJSON(400, gin.H{"err": err.Error()})
            return
        }
		{{- end}}

        // execute
        out, err = srv.{{.Name}}(ctx, in)
        if err != nil {
            c.AbortWithStatusJSON(500, gin.H{"err": err.Error()})
            return
        }
        
        c.JSON(200, out)
    }
}
{{end}}
// Client defines call remote server client and implement selector
type Client interface{
  Call(ctx context.Context, req, rsp interface{}) error
}

// {{.ServiceType}}HTTPClient defines call {{.ServiceType}}Server client
type {{.ServiceType}}HTTPClient interface {
{{- range .Methods}}
    {{.Name}}(context.Context, *{{.Request}}) (*{{.Reply}}, error)
{{- end}}
}

// {{.ServiceType}}HTTPClientImpl implement {{.ServiceType}}HTTPClient
type {{.ServiceType}}HTTPClientImpl struct {
	cli Client
}

func New{{.ServiceType}}HTTPClient(cli Client) {{.ServiceType}}HTTPClient {
	return &{{.ServiceType}}HTTPClientImpl{
		cli: cli,
	}
}

{{range .Methods}}
func (c *{{$svrType}}HTTPClientImpl) {{.Name}}(ctx context.Context, req *{{.Request}})(resp *{{.Reply}} ,err error) {
	resp = new({{.Reply}})
	err = c.cli.Call(ctx, req, resp)

	return
}
{{end}}
`

type serviceDesc struct {
	ServiceType string
	ServiceName string
	Metadata    string
	Methods     []*methodDesc
	MethodSets  map[string]*methodDesc
}

type methodDesc struct {
	// method
	Name    string
	Num     int
	Request string
	Reply   string
	// http_rule
	Path         string
	Method       string
	HasVars      bool
	HasBody      bool
	Body         string
	ResponseBody string
}

func (s *serviceDesc) execute() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("http_server").Parse(strings.TrimSpace(ginTemplate))
	if err != nil {
		panic(err)
	}
	if err = tmpl.Execute(buf, s); err != nil {
		panic(err)
	}

	return strings.Trim(buf.String(), "\r\n")
}
