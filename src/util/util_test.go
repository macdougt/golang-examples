package util_test

// Run with: go test -v util

import (
	"testing"
	"text/template"
	"time"
	. "util"
)

func TestInc(t *testing.T) {
	var v int
	var incVal int

	v = 32
	incVal = 3

	var resVal = Inc(v, incVal)
	if resVal != 35 {
		t.Error("Expected 35, got ", resVal)
	}
}

func TestIncFloat(t *testing.T) {
	var v float32
	var incVal float32

	v = 32.05
	incVal = 3

	var resVal = IncFloat(v, incVal)
	if resVal != 35.05 {
		t.Error("Expected 35.05, got ", resVal)
	}
}

func TestGetFractionalYear(t *testing.T) {
	// July 1st 1997, approximately half through 1997
	var year uint16
	var month time.Month
	var day uint8

	year = 1997
	month = 7
	day = 1

	var resVal = GetFractionalYear(year, month, day)
	if resVal != 1997.4987 {
		t.Error("Expected 1997.4987, got ", resVal)
	}
}

type AnimValSet struct {
	FromX float32
	ToX   float32
	FromY float32
	ToY   float32
}

func TestPopulateTemplateWithFuncMap(t *testing.T) {

	currAnimVals := AnimValSet{20, 40, 30, 50}

	var testTemplate = `<circle cx="{{.AnimVals.FromX}}" cy="{{.AnimVals.FromY}}" {{inc .AnimVals.FromY -20}}  {{.AnimVals.ToX}},{{inc .AnimVals.ToY -20}}`
	var templateWithResolution = `<circle cx="20" cy="30" 10  40,30`
	var resVal = PopulateTemplateWithFuncMap("testTemplate", testTemplate, map[string]interface{}{"AnimVals": currAnimVals}, template.FuncMap{"inc": IncFloat})
	if resVal != templateWithResolution {
		t.Error("Expected ", templateWithResolution, ", got ", resVal)
	}
}

func TestPopulateTemplate(t *testing.T) {

	currAnimVals := AnimValSet{20, 40, 30, 50}

	var testTemplate = `<circle cx="{{.AnimVals.FromX}}" cy="{{.AnimVals.FromY}}" {{.AnimVals.FromY}}  {{.AnimVals.ToX}},{{.AnimVals.ToY}}`
	var templateWithResolution = `<circle cx="20" cy="30" 30  40,50`
	var resVal = PopulateTemplate("testTemplate", testTemplate, map[string]interface{}{"AnimVals": currAnimVals})

	if resVal != templateWithResolution {
		t.Error("Expected ", templateWithResolution, ", got ", resVal)
	}
}
