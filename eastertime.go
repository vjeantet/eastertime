package eastertime

import (
  "time"
	"errors"
)

/**
 * return UTC time.Time for midnight on Easter of a given year
 * use Delambre and Butcher's algorithm Western Gregorian dates
 *
 * @param int year The year as a number greater than 0
 * @return time.Time The easter date as a time.Time
 * @return errors.Error Error if exists, else nil
 **/
func Byyear(year int) (time.Time, error) {
	var a, b, c, d, e, r int

  if year < 0 {
  	return  time.Now(), errors.New("year have to be greater than 0")
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
