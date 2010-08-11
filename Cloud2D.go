package aida

import . "strconv"
/* This is the ascii representation of a Cloud2d
   It is only useful for i/o, because all fields are parsed into strings
   Therefore it need not be visible outside of this package
   However, the fields must all be exported, because otherwise unmarshal cant parse into them
*/
type cloud2dStr struct {
    Name string "attr"
    Title string "attr"
    ConversionUpperEdgeX string "attr"
    ConversionLowerEdgeX string "attr"
    ConversionUpperEdgeY string "attr"
    ConversionLowerEdgeY string "attr"
    MaxEntries string "attr"
    ConversionBinsX string "attr"
    ConversionBinsY string "attr"
    Path string "attr"
    Annotation Annotation
    Entries2d Entries2d
}


type Cloud2d struct {
    Name string "attr"
    Title string "attr"
    ConversionUpperEdgeX float64
    ConversionLowerEdgeX float64
    ConversionUpperEdgeY float64
    ConversionLowerEdgeY float64
    MaxEntries int
    ConversionBinsX int
    ConversionBinsY int
    Path string
    Annotation Annotation
    Entries2d []Entry2d
}

func (c *cloud2dStr) String() string {
    if c == nil {
        return "nil"
    }
    return c.Name + " " + c.ConversionUpperEdgeY + " " + c.MaxEntries + c.Annotation.String() + c.Entries2d.String()
}

func (cs* cloud2dStr) Convert() *Cloud2d {
    c := new(Cloud2d)
    c.Name = cs.Name
    c.Title = cs.Title
    c.ConversionUpperEdgeX, _ = Atof64(cs.ConversionUpperEdgeX)
    c.ConversionUpperEdgeY, _ = Atof64(cs.ConversionUpperEdgeY)
    c.ConversionLowerEdgeX, _ = Atof64(cs.ConversionLowerEdgeX)
    c.ConversionLowerEdgeY, _ = Atof64(cs.ConversionLowerEdgeY)
    c.MaxEntries, _ = Atoi(cs.MaxEntries)
    c.ConversionBinsX, _ = Atoi(cs.ConversionBinsX)
    c.ConversionBinsY, _ = Atoi(cs.ConversionBinsY)
    c.Path = cs.Path
    c.Annotation = cs.Annotation
    c.Entries2d = make([]Entry2d, len(cs.Entries2d.Entry2d))
    for i:= range(cs.Entries2d.Entry2d) {
        c.Entries2d[i].ValueX, _ = Atof64(cs.Entries2d.Entry2d[i].ValueX)
        c.Entries2d[i].ValueY, _ = Atof64(cs.Entries2d.Entry2d[i].ValueY)
    }
    return c
}
