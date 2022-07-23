test:
	gotest -v -coverpkg=./... -coverprofile coverage.out ./... 
	go tool cover -func coverage.out

badge:
	gopherbadger -md="README.md,coverage.out"
