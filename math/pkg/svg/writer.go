package svg

import (
	"encoding/xml"
	"fmt"
	"io"
	"time"

	"github.com/jspaaks/learn-go-with-tests/math/pkg/unitclock"
)

type Svg struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Lines   []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func Writer(w io.Writer, t time.Time) {
	const (
		svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
		<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
		<svg xmlns="http://www.w3.org/2000/svg"
		     width="100%"
		     height="100%"
		     viewBox="0 0 300 300"
		     version="2.0">`
		bezel  = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
		svgEnd = `</svg>`

		cx               float64 = 150
		cy               float64 = 150
		radiusHourHand   float64 = 50
		radiusMinuteHand float64 = 80
		radiusSecondHand float64 = 90
	)

	hh := unitclock.HourHand(t).FlipVerticalDirection().Scale(radiusHourHand).Translate(cx, cy)
	mh := unitclock.MinuteHand(t).FlipVerticalDirection().Scale(radiusMinuteHand).Translate(cx, cy)
	sh := unitclock.SecondHand(t).FlipVerticalDirection().Scale(radiusSecondHand).Translate(cx, cy)

	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	io.WriteString(w, hourHandTag(hh))
	io.WriteString(w, minuteHandTag(mh))
	io.WriteString(w, secondHandTag(sh))
	io.WriteString(w, svgEnd)
}

func handTag(p unitclock.Point, color string, strokeWidth int) string {
	return fmt.Sprintf(
		`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:%s;stroke-width:%dpx;"/>`,
		p.X, p.Y, color, strokeWidth)
}

func hourHandTag(p unitclock.Point) string {
	return handTag(p, "#000", 7)
}

func minuteHandTag(p unitclock.Point) string {
	return handTag(p, "#000", 5)
}

func secondHandTag(p unitclock.Point) string {
	return handTag(p, "#F00", 3)
}
