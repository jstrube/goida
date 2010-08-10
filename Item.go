package aida

type Item struct {
    Sticky string "attr"
    Value string "attr"
    Key string "attr"
}

func (i *Item) String() string {
    return i.Sticky + " " + i.Value + " " + i.Key 
}
