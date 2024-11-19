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
migrate/up:
	@echo 'Running database up migration...'
	@goose postgres "postgres://postgres:a@localhost:5432/donorbox" up -dir=./migrations/
migrate/down:
	@echo 'Running database down migration...'
	@goose postgres "postgres://postgres:a@localhost:5432/donorbox" down -dir=./migrations/
