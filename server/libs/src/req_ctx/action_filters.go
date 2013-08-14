package req_ctx

import (
	"fmt"
)

func ParseForm(rc *RequestContext) {
	rc.parseForm();
}

func (rc *RequestContext) parseForm() {

	if !rc.isFormParsed {
		return // Already parsed
	}

	err := rc.Request.ParseForm()

	if err != nil {
		fmt.Println("Parsing form err: ", err)
	} else {
		rc.isFormParsed = true;	
	}	
}