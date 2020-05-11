package template2

import "testing"

func TestTxtTemplate(t *testing.T) {
	err := TxtTemplate()
	if err != nil {
		t.Error(err)
	}
}
func TestTxtTemplateToParam(t *testing.T) {
	err := TxtTemplateToParam()
	if err != nil {
		t.Error(err)
	}
}
func TestTxtTemplateByStruct(t *testing.T) {
	err := TxtTemplateByStruct()
	if err != nil {
		t.Error(err)
	}
}
