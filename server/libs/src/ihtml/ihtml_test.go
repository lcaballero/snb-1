package ihtml_test

import (
    h "ihtml"
    "testing"
)

func Test_Empty_Name(t *testing.T) {
	p := &h.IHtml{}

	actual := p.Name == ""

	if !actual {
		t.Errorf("%v != %v", actual, `""`)
	}
}

func Test_Empty_Value(t *testing.T) {
	p := &h.IHtml{}

	actual := p.Value == nil

	if !actual {
		t.Errorf("%v != %v", actual, `""`)
	}
}

func Test_Empty_Children(t *testing.T) {
	p := &h.IHtml{}

	actual := p.Children == nil

	if !actual {
		t.Errorf("%v != %v", actual, `""`)
	}
}

func Test_Empty_Attributes(t *testing.T) {
	p := &h.IHtml{}

	actual := p.Attributes == nil

	if !actual {
		t.Errorf("%v != %v", actual, `""`)
	}
}

func Test_IsFragment(t *testing.T) {
m
	p := &h.IHtml{ Name:"div" }

	actual := p.IsFragment()

	if !actual {
		t.Errorf("%v != %v", actual, false)
	}
}

func Test_IsText(t *testing.T) {

	p := &h.IHtml{ Value:"text" }

	actual := p.IsText()

	if !actual {
		t.Errorf("%v != %v", actual, false)
	}
}