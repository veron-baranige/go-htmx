run:
	npx tailwindcss -i src/styles/input.css -o dist/output.css
	go run src/server.go

build:
	rm dist/output.css
	npx tailwindcss -i src/styles/input.css -o dist/output.css --minify