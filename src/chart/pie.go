package chart

// This chart package ()i.e. directory) needs to be accessible for go files to use it
// the src directory that it belongs to, must have its parent in the GOPATH env var

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"text/template"
)

const Pi = 3.14159

type section struct {
	colourString string
	angle        float64
}

func RoundToInt(x float64) int {
	intVal := int(x)
	intermed := x - float64(intVal)
	if intermed >= 0.5 {
		intVal++
	}
	return intVal
}

const svgTemplate = `<svg xmlns="http://www.w3.org/2000/svg" width="160" height="170">
<defs>
<radialGradient id="greenGradient" fx="5%" fy="5%" r="65%" spreadMethod="pad"><stop offset="0%" stop-color="#00ee00" stop-opacity="1"/><stop offset="100%" stop-color="#006600" stop-opacity="1" /></radialGradient>
<radialGradient id="redGradient" fx="5%" fy="5%" r="65%" spreadMethod="pad"><stop offset="0%" stop-color="#ee0000" stop-opacity="1"/><stop offset="100%" stop-color="#660000" stop-opacity="1" /></radialGradient>
<radialGradient id="greyGradient" fx="5%" fy="5%" r="65%" spreadMethod="pad"><stop offset="0%" stop-color="#dedede" stop-opacity="1"/><stop offset="100%" stop-color="#5e5e5e" stop-opacity="1" /></radialGradient>
</defs><g>
<text x="20" y="18">{{.Title}}</text> 
{{.Circle}}{{range .Paths}}{{.}}{{end}}
{{.Tooltip}}
{{if .Paths}}<circle cx="80" cy="90" r="1" stroke="black" stroke-width="1"/>{{end}}
</g></svg>`

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func PieDraw(input []float64, title string) string {

	// Green
	greenSection := section{}
	greenSection.colourString = "green"

	// Red
	redSection := section{}
	redSection.colourString = "red"

	// Grey
	greySection := section{}
	greySection.colourString = "grey"

	colours := []section{greenSection, redSection, greySection}

	strokeWidth := 3
	cx := float64(80)
	cy := float64(90)
	radius := float64(60)

	var total float64

	// Calculate the total value
	for _, value := range input {
		total += value
	}

	drawCircle := false

	var sections []section
	var circleColour string

	for index2, value2 := range input {

		fl := 2.0 * value2 / total

		angleVal := fl * Pi

		// Are any of the angles more than 180 (pi radians)
		if angleVal > Pi {
			drawCircle = true // and don't push the angle
			circleColour = colours[index2].colourString

		} else {
			if angleVal > 0 {
				colours[index2].angle = angleVal
				sections = append(sections, colours[index2])
			}
		}
	}

	var paths []string
	circleContent := ""

	// First the circle if necessary
	if drawCircle {
		circleContent = fmt.Sprintf("<circle cx=\"%v\" cy=\"%v\" r=\"%v\" stroke=\"black\" stroke-width=\"%v\" style=\"fill:url(#%sGradient)\"/>", cx, cy, radius, strokeWidth, circleColour)
	}

	var startAngle, endAngle float64

	for _, value3 := range sections {
		startAngle = endAngle
		endAngle = startAngle + value3.angle

		x1 := RoundToInt(cx + radius*math.Cos(startAngle))
		y1 := RoundToInt(cy + radius*math.Sin(startAngle))

		x2 := RoundToInt(cx + radius*math.Cos(endAngle))
		y2 := RoundToInt(cy + radius*math.Sin(endAngle))

		paths = append(paths, fmt.Sprintf("<path d=\"M%v,%v  L%v,%v A%v,%v 0 0,1 %v,%v z\" stroke-width=\"%d\" stroke=\"#000000\" style=\"stroke-linejoin:bevel;fill:url(#%vGradient)\"/>", cx, cy, x1, y1, radius, radius, x2, y2, strokeWidth, value3.colourString))
	}

	// Set the tooltip
	tooltip := "<title>Green: " + fmt.Sprintf("%v", input[0]) + "\nRed: " + fmt.Sprintf("%v", input[1]) + "\nGrey: " + fmt.Sprintf("%v", input[2]) + "</title>"

	// Populate the template
	var err error
	t := template.New("svgTemplate")
	t, err = t.Parse(svgTemplate)
	checkError(err)

	buff := bytes.NewBufferString("")
	t.Execute(buff, map[string]interface{}{"Circle": circleContent, "Paths": paths, "Tooltip": tooltip, "Title": title})

	return buff.String()
}
