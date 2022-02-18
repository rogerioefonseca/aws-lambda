ZIP := zip

archive: build zip
	echo 'Builded and zipped'

build:
	GOOS=linux GOARCH=amd64 go build -o main main.go

zip:
	$(ZIP) archive.zip main
