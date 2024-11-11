run/api:
	@echo 'Running server...'
	@go run .
run/client:
	@echo 'Running frontend...'
	@npm run dev --prefix frontend/client
run/dashboard:
	@echo 'Running escrow dashboard...'
	@npm run dev --prefix frontend/escrow_dashboard
watch/css/client:
	@(cd frontend/client && npx tailwindcss -i ./src/input.css -o ./src/output.css --watch)
watch/css/dashboard:
	@(cd frontend/escrow_dashboard && npx tailwindcss -i ./src/input.css -o ./src/output.css --watch)
