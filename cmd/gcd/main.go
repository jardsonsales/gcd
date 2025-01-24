package main

import (
    "time"
    "fmt"

    "github.com/jardsonsales/gcd/internal/calendar"
)

func main() {

    year, month, _ := time.Now().Date()
    currentMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)

    calendar_render := calendar.Mount{Date: currentMonth}
    calendar_render.CLI()

    fmt.Println()

    calendar_render = calendar.Mount{Date: currentMonth.AddDate(0,1,0)}
    calendar_render.CLI()

    fmt.Println()

    calendar_render = calendar.Mount{Date: currentMonth.AddDate(0,2,0)}
    calendar_render.CLI()
}
