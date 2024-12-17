-- +goose Up
CREATE TYPE report_status AS ENUM (
  'pending',
  'dismissed',
  'resolved'
);

CREATE TABLE IF NOT EXISTS project_reports (
  id bigserial PRIMARY KEY,
  project_id bigint NOT NULL REFERENCES projects(id),
  email varchar(255) UNIQUE NOT NULL,
  full_name varchar(100) NOT NULL,
  phone_number varchar(11) NOT NULL,
  relation text,
  reason text NOT NULL,
  details text NOT NULL,
  status report_status NOT NULL DEFAULT 'pending',
  created_at timestamptz NOT NULL DEFAULT NOW()
);
-- +goose Down
