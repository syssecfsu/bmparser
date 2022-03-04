build:
	go build -o ./bmparser main.go
	go build -o ./dib_viewer viewer/viewer.go
clean:
	rm bmparser dib_viewer
