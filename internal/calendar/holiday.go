package calendar

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var HolidaysList = make(map[int][]holidayDay)

type Holiday struct{
    Year int
}

type holidayDay struct {
    Date string `json:"date"`
    Name string `json:"name"`
}

type ErrHolidayFileNotFound struct{ Path string }
func (e *ErrHolidayFileNotFound) Error() string {
    return fmt.Sprintf("File for holidays not found in: %s", e.Path)
}

func (h *Holiday) Parse() (error) {
    _, ok := HolidaysList[h.Year]; if ok {
        return nil
    }

    homeDir, err := os.UserHomeDir(); if err != nil {
        log.Fatalln(err)
    }
    holidayFile := filepath.Join(homeDir, ".config", "gcd", "holidays", fmt.Sprintf("%d.json", h.Year))
    _, err = os.Stat(holidayFile); if err != nil {
        return &ErrHolidayFileNotFound{ Path: holidayFile }
    }
    holidayJson, err := os.ReadFile(holidayFile); if err != nil {
        log.Fatalln(err)
    }

    var holidaysResult []holidayDay
    json.Unmarshal(holidayJson, &holidaysResult)

    for _, v := range holidaysResult {
        HolidaysList[h.Year] = append(HolidaysList[h.Year], v)
    }

    return nil
}

func isHoliday(year int, date string) bool {
    _, ok := HolidaysList[year]; if ok {
        for _, v := range HolidaysList[year] {
            if(date == v.Date) {
                return true
            }
        }
    }
    return false
}

func holidaysFromMonth(year int, month string) []holidayDay {
    var res []holidayDay
    _, ok := HolidaysList[year]; if ok {
        for _, v := range HolidaysList[year] {
            if(strings.HasPrefix(v.Date, fmt.Sprintf("%d-%s", year, month))) {
                res = append(res, v)
            }
        }
    }
    return res
}
