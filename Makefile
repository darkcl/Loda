.PHONY: build debug build-dmg build-mewn

build-mewn:
	@go build -o build/mewn cmd/mewn/main.go

build-dmg:
	create-dmg ./build/Production/Loda.app

build:
	@rm -Rf ./build/Production/Loda.app/
	@mkdir -p ./build/Production/Loda.app/Contents/MacOS
	@mkdir -p ./build/Production/Loda.app/Contents/Resources
	# Icon files
	# @cp ./assets/appicon.icns ./build/Production/Loda.app/Contents/Resources/
	# Meta Data
	@cp ./meta/Info.plist ./build/Production/Loda.app/Contents/
	@echo "Building GUI"
	@cd ./ui && yarn build
	@echo "Building go-application"
	@mewn build -o build/Production/Loda.app/Contents/MacOS/Loda
	@echo "Completed."

debug:
	@echo "Running in debug mode"
	go run main.go -mode=debug

serve:
	@echo "Starting webpack server"
	@cd ui && yarn serve
	