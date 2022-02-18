ZIP := zip
GO := GO
MAKE := make

archive: dependencies build zip
	@echo "#  executing " $@
	echo 'Builded and zipped'
	@echo

build: 
	@echo "#  executing " $@
	GOOS=linux GOARCH=amd64 $(GO) build -o main main.go
	@echo

dependencies:
	@echo "#  executing " $@
	$(GO) get github.com/aws/aws-lambda-go/lambda
	@echo

zip:
	@echo "#  executing " $@
	$(ZIP) archive.zip main
	@echo
