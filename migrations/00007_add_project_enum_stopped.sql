-- +goose Up
ALTER TYPE project_status ADD VALUE 'stopped';

-- +goose Down
