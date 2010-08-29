package aida

import . "strconv"

type entry1dstr struct {
    ValueX string "attr"
}

type Entry1d struct {
    ValueX float64 "attr"
}

/* This is the ascii representation of a Cloud1d
   It is only useful for i/o, because all fields are parsed into strings
   Therefore it need not be visible outside of this package
   However, the fields must all be exported, because otherwise unmarshal cant parse into them
*/


type Entries1dStr []entry1dstr

type cloud1dStr struct {
    Name string "attr"
    Title string "attr"
    ConversionUpperEdge string "attr"
    ConversionLowerEdge string "attr"
    MaxEntries string "attr"
    ConversionBins string "attr"
    Path string "attr"
    Annotation Annotation
    Entries1d Entries1dStr
}

type Cloud1d struct {
    Name string "attr"
    Title string "attr"
    ConversionUpperEdge float64
    ConversionLowerEdge float64
    MaxEntries int
    ConversionBins int
    Path string
    Annotation Annotation
    Entries1d []Entry1d
}

func (e Entries1dStr) String() string {
    var x string
    for i := range(e) {
        x += e[i].ValueX
    }
    return x
}

func (c *cloud1dStr) String() string {
    if c == nil {
        return "nil"
    }
    return c.Name + " " + c.ConversionUpperEdge + " " + c.MaxEntries + c.Annotation.String() + c.Entries1d.String()
}

func (cs* cloud1dStr) Convert() *Cloud1d {
    c := new(Cloud1d)
    c.Name = cs.Name
    c.Title = cs.Title
    c.ConversionUpperEdge, _ = Atof64(cs.ConversionUpperEdge)
    c.ConversionLowerEdge, _ = Atof64(cs.ConversionLowerEdge)
    c.MaxEntries, _ = Atoi(cs.MaxEntries)
    c.ConversionBins, _ = Atoi(cs.ConversionBins)
    c.Path = cs.Path
    c.Annotation = cs.Annotation
    c.Entries1d = make([]Entry1d, len(cs.Entries1d))
    for i:= range(cs.Entries1d) {
        c.Entries1d[i].ValueX, _ = Atof64(cs.Entries1d[i].ValueX)
    }
    return c
}
