package aida

type H1D struct {
    binEdges []float
    binEntries []float
    weights []float
    weight2 []float
}


// func (h *H1D) Append(x float) {
//     h.binEntries.push(x)
// }


func MakeH1D(bins []float) *H1D {
    hist := new(H1D)
    hist.binEdges = bins
    nbins := len(bins)-1
    hist.binEntries = make([]float, nbins)
    hist.weights = make([]float, nbins)
    hist.weight2 = make([]float, nbins)
    return hist
}


