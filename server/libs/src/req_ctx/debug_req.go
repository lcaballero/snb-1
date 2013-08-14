package req_ctx

import (
	"fmt"
	"strings"
	"strconv"
	"path"
	"sort"
)

type ServerRequestData struct {
	pairs map[string]string
	prefix string
	ctx *RequestContext
}

func NewServerRequestData(rc *RequestContext) *ServerRequestData {
	dd := ServerRequestData{}
	dd.pairs = make(map[string]string)
	dd.ctx = rc

	return &dd
}

func toString(val interface{}) string {
	switch val.(type) {
		case int:
			s, ok := val.(int)
			if ok { return strconv.Itoa(s) } else { return "" }
		case string:
			s, ok := val.(string)
			if ok { return s } else { return "" }
		case int64:
			s, ok := val.(int64)
			if ok { return strconv.FormatInt(s, 10) } else { return "" }
	}
	return ""
}

func (g *ServerRequestData) fill(key string, val interface{}) {
	g.pairs[g.prefix + key] = toString(val)
}

func (rd *ServerRequestData) fillUriInfo() {
	rd.prefix = "Request.URL."
	rd.fill("Scheme", rd.ctx.Request.URL.Scheme)
	rd.fill("Opaque", rd.ctx.Request.URL.Opaque)
	rd.fill("Host", rd.ctx.Request.URL.Host)
	rd.fill("Path", rd.ctx.Request.URL.Path)
	rd.fill("RawQuery", rd.ctx.Request.URL.RawQuery)
	rd.fill("Fragment", rd.ctx.Request.URL.Fragment)
}

func (rd *ServerRequestData) fillQuery() {
	p := "Request.URL.Query()"
	rd.prefix = p
	values := rd.ctx.Request.URL.Query()

	for k, v := range values {
		for i, s := range v {
			n := toString(i)
			rd.fill(p + "[" + n + "]." + k, s)
		}
	}
}

func (rd *ServerRequestData) fillRequestInfo() {
	rd.prefix = "Request."
	rd.fill("Method", rd.ctx.Request.Method)
	rd.fill("Proto", rd.ctx.Request.Proto)
	rd.fill("ProtoMajor", rd.ctx.Request.ProtoMajor)
	rd.fill("ProtoMinor", rd.ctx.Request.ProtoMinor)
	rd.fill("ContentLength", rd.ctx.Request.ContentLength)
	rd.fill("Host", rd.ctx.Request.Host)
	rd.fill("RemoteAddr", rd.ctx.Request.RemoteAddr)
	rd.fill("UserAgent()", rd.ctx.Request.UserAgent())
}

func (rd *ServerRequestData) fillDerivedInfo() {
	rd.prefix = ""
	rd.fill("path.Dir(rd.ctx.Request.URL.Path)", path.Dir(rd.ctx.Request.URL.Path))	
}

func WriteRequestInfo(rd *RequestContext) {
	if rd.AsBool("..debug..") {
		rd.RequestData.writeRequestInfo()
	}	
}

func (sd *ServerRequestData) sortedKeys() []string {

	keys := make([]string, 0)

	for k, _ := range sd.pairs {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}

func (sd *ServerRequestData) writeRequestInfo() {

	sd.fillRequestInfo()
	sd.fillUriInfo()
	sd.fillDerivedInfo()

	rows := []string {"<table>", "<tbody>"}

	keys := sd.sortedKeys()

	for _, k := range keys {
		rows = append(rows, 
			"<tr>",
				"<td>", k, "</td>",
				"<td>", sd.pairs[k], "</td>",
			"</tr>")
	}

	rows = append(rows, "</tbody>", "</table>")

	html := strings.Join(rows, "")

	fmt.Fprintln(sd.ctx.ResponseWriter, html)
}
