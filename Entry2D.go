package aida

import (
    "strconv"
    "fmt"
)

type Entry2dStr struct {
    ValueX string "attr"
    ValueY string "attr"
}


type Entry2d struct {
    ValueX float64
    ValueY float64
}

/*
    This is an example of strind to float conversion.
    Ideally, the aida file would be parsed somewhere in a function,
    then converted to numbers,
    then the string representation dies with the local scope and can be
    garbage-collected.
*/
func (e *Entry2dStr) Convert() Entry2d {
    x, _ := strconv.Atof64(e.ValueX)
    y, _ := strconv.Atof64(e.ValueY)
    return Entry2d{x, y}
}

func (x *Entry2dStr) String() string {
    var e = x.Convert()
    return fmt.Sprint(e.ValueX) + " " + fmt.Sprint(e.ValueY)
}

type Entries2d struct {
    Entry2d []Entry2dStr
}

func (e *Entries2d) String() string {
    var result string
    for _, v := range e.Entry2d {
        result += v.String() + "\n"
    }
    return result
}

