# Run templ generation in watch mode
templ:
	go tool templ generate --watch --proxy="http://localhost:8090" --open-browser=false -v

# Watch Tailwind CSS changes
tailwind:
	tailwindcss -i ./cmd/service/sources/css/input.css -o ./cmd/service/assets/css/styles.css --watch

# Start development server with all watchers
dev:
	go tool air
