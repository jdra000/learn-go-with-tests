// Package clockface provides functions that calculate the position of 
// the hands of an analogue clock.

package clockface

import (
	"math"
	"time"
)

// A Point is a two-dimensional Cartesian coordinate
type Point struct {
	X float64 
	Y float64 
}
// SecondHandPoint is the unit vector of the second hand at time 't',
// represented a Point.
func SecondHandPoint(t time.Time) Point {
    return angleToPoint(SecondsInRadians(t))
}
// SecondsInRadians returns the angle of the second hand from 12 o'clock in radians.
func SecondsInRadians(t time.Time) float64 {
    return math.Pi/(30/float64(t.Second()))
}

// MinuteHandPoint is the unit vector of the minute hand at time 't',
// represented a Point.
func MinuteHandPoint(t time.Time) Point {
    return angleToPoint(MinutesInRadians(t))
}
// MinutesInRadians returns the angle of the minute hand from 12 o'clock in radians.
func MinutesInRadians(t time.Time) float64 {
    // For every second the minute hand will move 1/60th of the angle the second hand moves
    return (SecondsInRadians(t) / 60) + (math.Pi/(30/float64(t.Minute()))) 
}

// HourHandPoint is the unit vector of the hour hand at time 't',
// represented a Point.
func HourHandPoint(t time.Time) Point {
    return angleToPoint(HoursInRadians(t))
}
// HourInRadians returns the angle of the hour hand from 12 o'clock in radians.
func HoursInRadians(t time.Time) float64 {
    // For every minute the hour hand will move 1/12th of the angle of the minute hand moves
    return (MinutesInRadians(t) / 12) + (math.Pi / (6 / float64(t.Hour() % 12)))  
}

// Helper functions
func angleToPoint(angle float64) Point {
    x := math.Sin(angle)
    y := math.Cos(angle)
    return Point{x, y}
}
