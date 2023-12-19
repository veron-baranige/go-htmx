run:
	go run src/server.go

build:
	rm dist/output.css
	npx tailwindcss -i static/index.css -o static/tailwind.css --minify