package aida

// A Tuple is a data structure similar to a table.
// It contains one row of entries per event
// An entry can be a C POD (plain old data) or a Tuple
// (Although nested Tuples are a bit tricky to deal with and should be handled with caution)

// Data is stored in ASCII, so the data is read in two steps
// 1) define containers of the ascii strings and pass to xml unmarshaller to read entries from disk
// 2) define containers of float, double, int, long, ... that contain the numbers and are used in computations.


type ColumnStr struct {
    Name string "attr"
    Type string "attr"
}

type Column struct {
    Name string "attr"
    Title string "attr"
    Path string "attr"
    Annotation Annotation
    Columns []ColumnStr
}

type Tuple struct {
    
}

// The entry type is pretty useless outside of this package
type Entry struct {
    Value string "attr"
}

type ColumnDouble struct {
    
}


type Rows struct {
    Row []float64
}
