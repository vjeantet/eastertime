// eastertime provides functions to get midnight date time of easter day
// for Catholic and Orthodox on Gregorian Calendar
// web github/vjeantet/eastertime

package eastertime

import (
	"errors"
	"math"
	"time"
)

// CatholicByYear returns time.Time for midnight on Catholic Easter of a given year
// use Delambre and Butcher's and return time based on Gregorian calendar
//
// @param int year The year as a number greater than 0
// @return time.Time The easter date as a time.Time
// @return errors.Error Error if exists, else nil
func CatholicByYear(year int) (time.Time, error) {
	var a, b, c, d, e, r int

	if year < 0 {
		return time.Now(), errors.New("year have to be greater than 0")
	}

	a = year % 19
	if year >= 1583 {
		var f, g, h, i, k, l, m int
		b = year / 100
		c = year % 100
		d = b / 4
		e = b % 4
		f = (b + 8) / 25
		g = (b - f + 1) / 3
		h = (19*a + b - d - g + 15) % 30
		i = c / 4
		k = c % 4
		l = (32 + 2*e + 2*i - h - k) % 7
		m = (a + 11*h + 22*l) / 451
		r = 22 + h + l - 7*m
	} else {
		b = year % 7
		c = year % 4
		d = (19*a + 15) % 30
		e = (2*c + 4*b - d + 34) % 7
		r = 22 + d + e
	}

	return time.Date(year, time.March, r, 0, 0, 0, 0, time.Local), nil
}

// OrthodoxByYear returns time.Time for midnight on Orthodox Easter of a given year
// use Meeus Julian algorithm and return time based on Gregorian calendar
//
// @param int year The year as a number greater than 325
// @return time.Time The easter date as a time.Time
// @return errors.Error Error if exists, else nil
func OrthodoxByYear(year int) (time.Time, error) {

	if year < 326 {
		return time.Now(), errors.New("year have to be greater than 325")
	}

	var a, b, c, d, e int
	var month time.Month
	var day float64

	a = year % 4
	b = year % 7
	c = year % 19
	d = (19*c + 15) % 30
	e = (2*a + 4*b - d + 34) % 7
	month = time.Month((d + e + 114) / 31)
	day = math.Floor(float64((d+e+114)%31 + 1))
	day = day + 13

	return time.Date(year, month, int(day), 0, 0, 0, 0, time.Local), nil
}
