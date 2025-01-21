package pad

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestPad(t *testing.T) {
    str := "test"
    res := Both(str, " ", 6)
    assert.True(t, res == " test ")
}


func TestPadTooManyChars(t *testing.T) {
    str := "test"
    res := Both(str, " ", 2)
    assert.True(t, res == "test")
}
