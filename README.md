golang-examples
===============

My go language examples

Pie chart generator
includes:
* testPieChartSvg.go -- Main file for testing the pie chart generator
* src/chart/pie.go -- Actual brute force pie chart generator tha could be made way more configurable. Just wanting to learn some go.
* testPieChartSvgOut.svg -- Sample output from the generator
	![alt SVG output](testPieChartSvgOut.png)
* src/util/util.go -- includes GetFractionalYear, StringToFile, PopulateTemplate, PopulateTemplateWithFuncMap, checkError, Inc, IncFloat, WordCount utility functions
* src/util/util_test.go -- unit tests for util functions, run with: go test -v util


