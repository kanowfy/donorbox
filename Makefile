run/api:
	@echo 'Running cmd/api...'
	@go run ./cmd/api
run/frontend:
	@echo 'Running frontend...'
	@npm run dev --prefix frontend
watch/css:
	@(cd frontend && npx tailwindcss -i ./src/input.css -o ./src/output.css --watch)
