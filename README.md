# Trying to get forms to repopulate after a failed attempt using cookies
---

## Here are some handy templates for vscode if you are interested in making your own router like this
{

	"router main app": {
		"prefix": "router-app",
		"body": [
			"// App is main router for the web app. it is the entrypoint",
			"type App struct {",
			"",
			"}",
			"func (h *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {",
			"",
			"    return",
			"}"
	],
		"description": "Creates  for web app routing"
	},
	"router struct": {
		"prefix": "router-struct",
		"body": [
			"type $1 struct {",
			"",
			"}",
			"func (h *$1) ServeHTTP(w http.ResponseWriter, r *http.Request) {",
			"",
			"    return",
			"}"
	],
		"description": "Creates struct and servehttp method"
	},
	"router handlerfunc method": {
		"prefix": "router-handlefunc",
		"body": [
			"func (*$1) $2() http.Handler {",
			"    // code if you want",
			"    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {",
			"       $0",
			"       return",
			"    })",
			"}"
		],
		"description": "Creates  for web app routing"
	},
	"router NotFound and MethodNotAllowed func definitions" : {
		"prefix": "router-errorfuncs",
		"body": [
			"// NotFound sends not found error to w",
			"func NotFound(w http.ResponseWriter) {",
			"	http.Error(w, \"404 not found!!!\", http.StatusNotFound)",
			"}",
			"",
			"// MethodNotAllowed sends MethodNotAllowed error to w",
			"func MethodNotAllowed(w http.ResponseWriter) {",
			"	http.Error(w, \"Method not allowed\", http.StatusMethodNotAllowed)",
			"}",
		],
		"description": "NotFound and MethodNotAllowed func definitions"
	}
}