package calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDaysIn(t *testing.T) {
    year, month, _ := time.Date(2025,1,1,0,0,0,0,time.Local).Date()
    res := DaysIn(month, year)
    assert.Equal(t, res, 31)


    year, month, _ = time.Date(2025,6,1,0,0,0,0,time.Local).Date()
    res = DaysIn(month, year)
    assert.Equal(t, res, 30)


    year, month, _ = time.Date(2025,2,1,0,0,0,0,time.Local).Date()
    res = DaysIn(month, year)
    assert.Equal(t, res, 28)


    year, month, _ = time.Date(2024,2,1,0,0,0,0,time.Local).Date()
    res = DaysIn(month, year)
    assert.Equal(t, res, 29)
}

func TestIsLeapYear(t *testing.T) {
    res := isLeapYear(2024)
    assert.True(t, res)

    res = isLeapYear(2025)
    assert.False(t, res)
}
