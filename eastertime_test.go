package eastertime

import (
	"testing"
	"time"
)

var fixtureCatholic = []time.Time{
	time.Date(2016, 3, 27, 0, 0, 0, 0, time.Local),
	time.Date(2015, 4, 5, 0, 0, 0, 0, time.Local),
	time.Date(2014, 4, 20, 0, 0, 0, 0, time.Local),
	time.Date(2013, 3, 31, 0, 0, 0, 0, time.Local),
	time.Date(2012, 4, 8, 0, 0, 0, 0, time.Local),
	time.Date(2011, 4, 24, 0, 0, 0, 0, time.Local),
	time.Date(1901, 04, 07, 0, 0, 0, 0, time.Local),
	time.Date(1902, 03, 30, 0, 0, 0, 0, time.Local),
	time.Date(1903, 04, 12, 0, 0, 0, 0, time.Local),
	time.Date(1904, 04, 03, 0, 0, 0, 0, time.Local),
	time.Date(1905, 04, 23, 0, 0, 0, 0, time.Local),
	time.Date(1906, 04, 15, 0, 0, 0, 0, time.Local),
	time.Date(1907, 03, 31, 0, 0, 0, 0, time.Local),
	time.Date(1908, 04, 19, 0, 0, 0, 0, time.Local),
	time.Date(1944, 4, 9, 0, 0, 0, 0, time.Local),
	time.Date(1945, 04, 01, 0, 0, 0, 0, time.Local),
	time.Date(1946, 04, 21, 0, 0, 0, 0, time.Local),
	time.Date(1947, 04, 06, 0, 0, 0, 0, time.Local),
	time.Date(1948, 03, 28, 0, 0, 0, 0, time.Local),
	time.Date(1949, 04, 17, 0, 0, 0, 0, time.Local),
	time.Date(1950, 4, 9, 0, 0, 0, 0, time.Local),
	time.Date(1951, 03, 25, 0, 0, 0, 0, time.Local),
	time.Date(1952, 04, 13, 0, 0, 0, 0, time.Local),
	time.Date(1953, 04, 05, 0, 0, 0, 0, time.Local),
	time.Date(1954, 04, 18, 0, 0, 0, 0, time.Local),
	time.Date(1955, 04, 10, 0, 0, 0, 0, time.Local),
	time.Date(1956, 04, 01, 0, 0, 0, 0, time.Local),
	time.Date(1957, 04, 21, 0, 0, 0, 0, time.Local),
	time.Date(1958, 04, 06, 0, 0, 0, 0, time.Local),
	time.Date(2016, 03, 27, 0, 0, 0, 0, time.Local),
	time.Date(2017, 04, 16, 0, 0, 0, 0, time.Local),
	time.Date(2018, 04, 01, 0, 0, 0, 0, time.Local),
	time.Date(2019, 04, 21, 0, 0, 0, 0, time.Local),
	time.Date(2020, 04, 12, 0, 0, 0, 0, time.Local),
	time.Date(2021, 04, 04, 0, 0, 0, 0, time.Local),
	time.Date(2022, 04, 17, 0, 0, 0, 0, time.Local),
	time.Date(2023, 4, 9, 0, 0, 0, 0, time.Local),
	time.Date(2024, 03, 31, 0, 0, 0, 0, time.Local),
	time.Date(2025, 04, 20, 0, 0, 0, 0, time.Local),
	time.Date(2026, 04, 05, 0, 0, 0, 0, time.Local),
	time.Date(2027, 03, 28, 0, 0, 0, 0, time.Local),
	time.Date(2028, 04, 16, 0, 0, 0, 0, time.Local),
	time.Date(2029, 04, 01, 0, 0, 0, 0, time.Local),
	time.Date(2030, 04, 21, 0, 0, 0, 0, time.Local),
	time.Date(2031, 04, 13, 0, 0, 0, 0, time.Local),
	time.Date(2032, 03, 28, 0, 0, 0, 0, time.Local),
	time.Date(2033, 04, 17, 0, 0, 0, 0, time.Local),
	time.Date(2034, 4, 9, 0, 0, 0, 0, time.Local),
	time.Date(2035, 03, 25, 0, 0, 0, 0, time.Local),
	time.Date(2036, 04, 13, 0, 0, 0, 0, time.Local),
	time.Date(2037, 04, 05, 0, 0, 0, 0, time.Local),
	time.Date(2038, 04, 25, 0, 0, 0, 0, time.Local),
	time.Date(2039, 04, 10, 0, 0, 0, 0, time.Local),
}

var fixtureOrthodox = []time.Time{
	time.Date(2008, 4, 27, 0, 0, 0, 0, time.Local),
	time.Date(2009, 4, 19, 0, 0, 0, 0, time.Local),
	time.Date(2010, 4, 4, 0, 0, 0, 0, time.Local),
	time.Date(2011, 4, 24, 0, 0, 0, 0, time.Local),
}

func TestCatholicByYear(t *testing.T) {
	var res, expectedTime time.Time
	var err error
	// For each year's time in fixtures, compute easter time
	for _, expectedTime = range fixtureCatholic {
		res, _ = CatholicByYear(expectedTime.Year())
		if res != expectedTime {
			t.Errorf("WesternByYear should return %#v, but returns %#v\n", expectedTime, res)
		}
	}

	// Test Error when given year is not greater then 0
	res, err = CatholicByYear(-2)
	if err == nil {
		t.Error("WesternByYear should return an error, but returns null\n")
	}
}

func TestOrthodoxByYear(t *testing.T) {
	var res, expectedTime time.Time
	var err error

	// For each year's time in fixtures, compute easter time
	for _, expectedTime = range fixtureOrthodox {
		res, _ = OrthodoxByYear(expectedTime.Year())
		if res != expectedTime {
			t.Errorf("EasternByYear should return %#v, but returns %#v\n", expectedTime.Format("02-01-2006"), res.Format("02-01-2006"))
		}
	}

	// Test Error when given year is not greater then 326
	res, err = OrthodoxByYear(325)
	if err == nil {
		t.Error("EasternByYear should return an error, but returns null\n")
	}
}
