include $(GOROOT)/src/Make.$(GOARCH)

TARG=aida
GOFILES=\
    Annotation.go \
    Entry2D.go \
    ITuple.go \
    Item.go \
    Cloud2D.go \
    H1D.go \
    File.go

include $(GOROOT)/src/Make.pkg

