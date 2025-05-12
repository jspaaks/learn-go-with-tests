package svg

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"testing"
	"time"
)

func assertLinesIncludeGivenLine(t *testing.T, lines []Line, want Line, s string) {
	t.Helper()
	for _, got := range lines {
		if got == want {
			return
		}
	}
	t.Fatalf("Expected to find the second hand with x2 of %+v and y2 of %+v in the SVG output %v", want.X2, want.Y2, s)
}

func TestSvgWriterMinuteHand(t *testing.T) {
	testcases := []struct {
		minutes  int
		expected Line
	}{
		{0, Line{X1: 150, Y1: 150, X2: 150, Y2: 70}},
		{15, Line{X1: 150, Y1: 150, X2: 230, Y2: 150}},
		{30, Line{X1: 150, Y1: 150, X2: 150, Y2: 230}},
		{45, Line{X1: 150, Y1: 150, X2: 70, Y2: 150}},
	}
	for _, testcase := range testcases {
		desc := fmt.Sprintf("at %d minutes", testcase.minutes)
		t.Run(desc, func(t *testing.T) {
			tm := time.Date(1337, time.January, 1, 0, testcase.minutes, 0, 0, time.UTC)
			b := bytes.Buffer{}
			SvgWriter(&b, tm)
			svg := Svg{}
			xml.Unmarshal(b.Bytes(), &svg)
			assertLinesIncludeGivenLine(t, svg.Lines, testcase.expected, b.String())
		})
	}
}

func TestSvgWriterHourHand(t *testing.T) {
	testcases := []struct {
		hours    int
		expected Line
	}{
		{0, Line{X1: 150, Y1: 150, X2: 150, Y2: 100}},
		{3, Line{X1: 150, Y1: 150, X2: 200, Y2: 150}},
		{6, Line{X1: 150, Y1: 150, X2: 150, Y2: 200}},
		{9, Line{X1: 150, Y1: 150, X2: 100, Y2: 150}},
	}
	for _, testcase := range testcases {
		desc := fmt.Sprintf("at %d hours", testcase.hours)
		t.Run(desc, func(t *testing.T) {
			tm := time.Date(1337, time.January, 1, testcase.hours, 0, 0, 0, time.UTC)
			b := bytes.Buffer{}
			SvgWriter(&b, tm)
			svg := Svg{}
			xml.Unmarshal(b.Bytes(), &svg)
			assertLinesIncludeGivenLine(t, svg.Lines, testcase.expected, b.String())
		})
	}
}

func TestSvgWriterSecondHand(t *testing.T) {
	testcases := []struct {
		seconds  int
		expected Line
	}{
		{0, Line{X1: 150, Y1: 150, X2: 150, Y2: 60}},
		{15, Line{X1: 150, Y1: 150, X2: 240, Y2: 150}},
		{30, Line{X1: 150, Y1: 150, X2: 150, Y2: 240}},
		{45, Line{X1: 150, Y1: 150, X2: 60, Y2: 150}},
	}
	for _, testcase := range testcases {
		desc := fmt.Sprintf("at %d seconds", testcase.seconds)
		t.Run(desc, func(t *testing.T) {
			tm := time.Date(1337, time.January, 1, 0, 0, testcase.seconds, 0, time.UTC)
			b := bytes.Buffer{}
			SvgWriter(&b, tm)
			svg := Svg{}
			xml.Unmarshal(b.Bytes(), &svg)
			assertLinesIncludeGivenLine(t, svg.Lines, testcase.expected, b.String())
		})
	}
}
