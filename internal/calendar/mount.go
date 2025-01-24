package calendar

import (
    "fmt"
    "log"
    "time"

    "github.com/jardsonsales/gcd/internal/pad"
)

type Mount struct {
    Date time.Time
}

func (m Mount) CLI() {
    holiday := &Holiday{Year: 2025} // find a way to parse only holidays of the affected years
    err := holiday.Parse(); if err != nil {
        err, ok := err.(*ErrHolidayFileNotFound); if ok {
            fmt.Printf("\033[38;2;255;0;02mNote:\033[m Cannot find the holidays file at: %s\n\n", err.Path)
        } else {
            log.Fatal(err)
        }
    }

    todayDate := time.Now()

    year, month, _ := m.Date.Date()
    location := m.Date.Location()

    fmt.Println(pad.Both(m.Date.Month().String(), " ", 27))

    fmt.Println(header)

    firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, location)

    weekdayIndex := int(firstOfMonth.Weekday())

    for i := 0; i < weekdayIndex; i++ {
        fmt.Print("    ")
    }

    daysInMonth := DaysIn(month, year)
    for day := 1; day <= daysInMonth; day++ {
        currentDay := fmt.Sprintf("%d-%02d-%02d",year,month,day)
        if (currentDay == todayDate.Format(time.DateOnly)) {
            if(isHoliday(year, currentDay)) {
                fmt.Printf("\033[48;2;200;150;02m%3d \033[m", day)
            } else {
                fmt.Printf("\033[42m%3d \033[m", day)
            }
        } else {
            if(isHoliday(year, currentDay)) {
                fmt.Printf("\033[48;2;180;0;02m%3d \033[m", day)
            } else {
                fmt.Printf("%3d ", day)
            }
        }
        if (day+weekdayIndex)%7 == 0 {
            fmt.Println()
        }
    }

    holidaysFromActualMonth := holidaysFromMonth(year, fmt.Sprintf("%02d", month))

    if(len(holidaysFromActualMonth) > 0) {
        fmt.Println()
    }

    for _, h := range holidaysFromActualMonth {
        fmt.Printf("%s: %s ", h.Date[8:], h.Name)
    }

    fmt.Println()
}
