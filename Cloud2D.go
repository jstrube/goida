package aida

type Cloud2d struct {
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


func (c *Cloud2d) String() string {
    if c == nil {
        return "nil"
    }
    return c.Name + " " + c.ConversionUpperEdgeY + " " + c.MaxEntries + c.Annotation.String() + c.Entries2d.String()
}

