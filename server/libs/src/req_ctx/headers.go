package req_ctx

func ToTextHtml(rc *RequestContext) {
	h := rc.ResponseWriter.Header()
	h["Content-Type"] = []string { "text/html", "charset-utf-8", }
}

func ToTextCss(rc *RequestContext) {
	h := rc.ResponseWriter.Header()
	h["Content-Type"] = []string { "text/css", "charset-utf-8",	}
}


func ToTextJavascript(rc *RequestContext) {
	h := rc.ResponseWriter.Header()
	h["Content-Type"] = []string { "text/javascript", "charset-utf-8", }
}
