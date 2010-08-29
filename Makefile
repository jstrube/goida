include $(GOROOT)/src/Make.inc

TARG=aida
GOFILES=\
    Annotation.go \
    Entry2D.go \
    ITuple.go \
    Item.go \
    Cloud2D.go \
    Cloud1D.go \
    H1D.go \
    File.go

include $(GOROOT)/src/Make.pkg

