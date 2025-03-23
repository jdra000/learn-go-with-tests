package clockface_test

import (
	"math"
	"testing"
	"time"
    ."github.com/jdra000/learn-go-with-tests/clockface"
)

// These tests are based on an unitary circle.
// First I test the radians conversion
// Second I test the x,y position

func TestSecondsInRadians(t *testing.T){
    cases := []struct {
        time time.Time 
        angle float64
    }{
        {simpleTime(0, 0, 30), math.Pi},
        {simpleTime(0, 0, 0), 0},
        {simpleTime(0, 0, 45), 3 * (math.Pi/2)},
        {simpleTime(0, 0, 7), 7 * (math.Pi/30)},
    }
    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T){
            got := SecondsInRadians(c.time)
            if !roughlyEqualFloat(got, c.angle){
                t.Fatalf("wanted %v radians got %v", c.angle, got)
            }
        })
    }
}
func TestSecondHandPoint(t *testing.T){
    cases := []struct {
        time time.Time 
        point Point
    }{
        {simpleTime(0, 0, 30), Point{0, -1}},
        {simpleTime(0, 0, 45), Point{-1, 0}},
    }
    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T){
            got := SecondHandPoint(c.time)
            if !roughlyEqualPoint(got, c.point){
                t.Fatalf("wanted %v point got %v", c.point, got)
            }
        })
    }
}

func TestMinutesInRadians(t *testing.T){
    cases := []struct{
        time time.Time
        angle float64
    }{
        {simpleTime(0, 30, 0), math.Pi},
        {simpleTime(0, 0, 7), 7 * (math.Pi / (30*60))}, 
    } 

    for _, c := range cases {
       t.Run(testName(c.time), func(t *testing.T) {
            got := MinutesInRadians(c.time)
            if !roughlyEqualFloat(got, c.angle){
                t.Fatalf("wanted %v radians got %v", c.angle, got)
            }
       })
    }
}
func TestMinuteHandPoint(t *testing.T){
    cases := []struct{
        time time.Time 
        point Point
    }{
        {simpleTime(0, 30, 0), Point{0, -1}},
        {simpleTime(0, 45, 0), Point{-1, 0}},
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            got := MinuteHandPoint(c.time)
            if !roughlyEqualPoint(got, c.point) {
                t.Fatalf("wanted %v point got %v", c.point, got)
            } 
        })
    }

}

func TestHoursInRadians(t *testing.T){
    cases := []struct{
        time time.Time 
        angle float64
    }{
        {simpleTime(6, 0, 0), math.Pi},
        {simpleTime(0, 0, 0), 0},
        {simpleTime(21, 0, 0), math.Pi * 1.5},
        {simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            got := HoursInRadians(c.time)
            if !roughlyEqualFloat(got, c.angle) {
                t.Fatalf("wanted %v radians got %v", c.angle, got)
            }
        })
    }
}
func TestHourHandPoint(t *testing.T){
    cases := []struct{
        time time.Time 
        point Point
    }{
        {simpleTime(6, 0, 0), Point{0, -1}},
        {simpleTime(21, 0, 0), Point{-1, 0}},
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            got := HourHandPoint(c.time)
            if !roughlyEqualPoint(got, c.point){
                t.Fatalf("wanted %v point got %v", c.point, got)
            }
        })
    }
}

func roughlyEqualFloat(a, b float64) bool {
    const equalityThreshold = 1e-7
    return math.Abs(a - b) < equalityThreshold
}
// Define equality between two Points
// they'll work if the X and Y elements are within 0.00000001 of each other
func roughlyEqualPoint(a, b Point) bool {
    return roughlyEqualFloat(a.X, b.X) && roughlyEqualFloat(a.Y, b.Y)
}

// Helper functions for tests
func simpleTime(hours, minutes, seconds int) time.Time {
    return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
    return t.Format("15:04:05")
}
