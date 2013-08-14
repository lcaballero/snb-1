package req_ctx

import (
	"net/http"
	"fmt"
	"regexp"
	//"strings"
	//"strconv"
)

type RequestContext struct {
	ResponseWriter http.ResponseWriter
	Request *http.Request
	ContextItems map[string]interface{}
	RequestData *ServerRequestData
	Templates *HtmlTemplateCache

	RouteValues interface{}
	Preferrences interface{}
	Session interface{}
	User interface{}

	isFormParsed bool
}

type actionMap struct {
	pattern string
	regexp *regexp.Regexp
	actions []func(*RequestContext)
	staticDir string
}

type ContextWrapper struct {
	actionMaps []actionMap
	post []func(*RequestContext)
	pre []func(*RequestContext)
}

func NewRequestContext(w http.ResponseWriter, r *http.Request) *RequestContext {
	rc := RequestContext {}
	rc.ResponseWriter = w
	rc.Request = r
	rc.ContextItems = make(map[string]interface{})
	rc.RequestData = NewServerRequestData(&rc)
	rc.isFormParsed = false
	return &rc
}

func NewContextWrapper() *ContextWrapper {
	cw := ContextWrapper {}
	cw.actionMaps = make([]actionMap, 0)
	cw.post = make([]func(*RequestContext), 0)
	cw.pre = make([]func(*RequestContext), 0)
	return &cw
}

func (cw *ContextWrapper) HandleFunc(key string, fn ...func(*RequestContext)) {
	re := regexp.MustCompile(key)
	cw.actionMaps = append(cw.actionMaps, actionMap { key, re, fn, "" })
}

func (cw *ContextWrapper) OnActionExecuted(fn ...func(*RequestContext)) {
	cw.post = append(cw.post, fn...)
}

func (cw *ContextWrapper) OnActionExecuting(fn ...func(*RequestContext)) {
	cw.pre = append(cw.pre, fn...)
}

func findAction(path string, actions []actionMap) (a *actionMap, found bool) {
	for _, a := range actions {
		found = a.regexp.MatchString(path)
		if found {
			return &a, true
		}
	}
	return nil, false
}

func (cw *RequestContext) Get(key string) (val string) {

	cw.parseForm()

	val = cw.Request.FormValue(key)

	return val
}

func (cw *RequestContext) AsBool(key string) bool {
	val := cw.Get(key)

	b := val == "1" || val == "on" || val == "y" || val == "true"

	return b
}

func (cw *ContextWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	rc := NewRequestContext(w, r)
	ToTextHtml(rc)

	path := r.URL.Path

	am, ok := findAction(path, cw.actionMaps)

	if ok {
		fmt.Println("Path: ", path, "Matched Action: ", am.regexp)
	} else {
		fmt.Println("Could not find aciton for path: ", path)
	}

	for _, fn := range cw.pre {
		fn(rc)
	}

	if ok {
		for _, fn := range am.actions {
			fn(rc)
		}
	}

	for _, fn := range cw.post {
		fn(rc)
	}
}

