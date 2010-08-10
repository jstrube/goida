package aida

type Annotation struct {
    Item []Item
}

func (a *Annotation) String() string {
    var result string
    for _, v := range a.Item {
        result += v.String() + " "
    }
    return result
}

