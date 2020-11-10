FROM golang:1.15.4-buster

WORKDIR /go/src/work

COPY go_gallery.go go_gallery.go
COPY view.tmpl view.tmpl

CMD ["go", "run", "go_gallery.go"]
