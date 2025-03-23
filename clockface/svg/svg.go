package svg

import (
    cf "github.com/jdra000/learn-go-with-tests/clockface"
    "io"
    "time"
    "fmt"
)

const (
   secondHandLength = 90
   minuteHandLength = 80
   hourHandLength = 50
   clockCentreX = 150
   clockCentreY = 150
)

// SVGWriter writes an svg representation of an analogue clock, showing the time t, to the writer w.
func SVGWriter(w io.Writer, t time.Time){
   io.WriteString(w, svgStart)
   io.WriteString(w, bezel)
   SecondHand(w, t)
   MinuteHand(w, t)
   HourHand(w, t)
   io.WriteString(w, svgEnd)
}

func SecondHand(w io.Writer, t time.Time) {
    p := makeHand(cf.SecondHandPoint(t), secondHandLength)
    fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}
func MinuteHand(w io.Writer, t time.Time) {
    p := makeHand(cf.MinuteHandPoint(t), minuteHandLength)
    fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}
func HourHand(w io.Writer, t time.Time) {
    p := makeHand(cf.HourHandPoint(t), hourHandLength)
    fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y) 
}
func makeHand(p cf.Point, length float64) cf.Point {
    p = cf.Point{X: p.X * length, Y: p.Y * length}                  // scale
    p = cf.Point{X: p.X, Y: -p.Y}                                   // flip
    return cf.Point{X: p.X + clockCentreX, Y: p.Y + clockCentreY}   // translate
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`

