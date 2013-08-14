package ihtml

type IHtml struct {
	Name string
	Value interface{} 
	Children []IHtml
	Attributes map[string]string
}

func (v *IHtml) IsFragment() bool {
	return v.Name == "" && v.Value == nil && v.Children != nil && v.Attributes == nil
}

func (v *IHtml) IsText() bool {
	return v.Name == "" && v.Value != nil && v.Children == nil && v.Attributes == nil
}

func (v *IHtml) IsElement() bool {
	return v.Name != ""
}

func (v *IHtml) HasAttributes() bool {
	return false
}

func (v *IHtml) HasChildren() bool {
	return len(v.Children) > 0
}

func (v *IHtml) String() string {
	return ""
}


