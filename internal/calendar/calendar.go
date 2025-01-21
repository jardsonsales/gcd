package calendar

import (
    "time"
)

var header string = "Dom Seg Ter Qua Qui Sex Sab"

// Função para determinar o número de dias em um mês específico
func DaysIn(month time.Month, year int) int {
    switch month {
    case time.February:
        if isLeapYear(year) {
            return 29
        }
        return 28
    case time.April, time.June, time.September, time.November:
        return 30
    default:
        return 31
    }
}

// Função para verificar se um ano é bissexto
func isLeapYear(year int) bool {
    if year%4 == 0 {
        if year%100 == 0 {
            return year%400 == 0
        }
        return true
    }
    return false
}
