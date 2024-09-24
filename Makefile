app = cmd/main.go

.PHONY: run
run:
	go run $(app)

.PHONY: build
build:
	go build -o bin/typing_wizard $(app)
	# GOOS=windows GOARCH=amd64 go build -o release/typing-wizard.exe $(app)
	# GOOS=darwin GOARCH=amd64 go build -o release/typing-wizard_darwin_amd64 $(app)
	# GOOS=darwin GOARCH=arm64 go build -o release/typing-wizard_darwin_arm64 $(app)
