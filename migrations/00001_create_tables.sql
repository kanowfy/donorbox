-- +goose Up
CREATE TYPE project_status AS ENUM (
  'pending',
  'ongoing',
  'rejected',
  'finished',
  'disputed'
);

CREATE TYPE verification_status AS ENUM (
  'unverified',
  'pending',
  'verified'
);

CREATE TYPE notification_type AS ENUM (
  'approved_verification',
  'rejected_verification',
  'approved_project',
  'rejected_project',
  'released_fund_milestone',
  'completed_milestone',
  'refuted_milestone',
  'rejected_proof',
  'approved_proof'
);

CREATE TYPE milestone_status AS ENUM (
  'pending',
  'fund_released',
  'completed',
  'refuted'
);

CREATE TYPE proof_status AS ENUM (
  'pending',
  'rejected',
  'approved'
);

CREATE TABLE IF NOT EXISTS users (
  id bigserial PRIMARY KEY,
  email varchar(255) UNIQUE NOT NULL,
  first_name varchar(64) NOT NULL,
  last_name varchar(64) NOT NULL,
  profile_picture text,
  hashed_password varchar(64) NOT NULL,
  activated boolean NOT NULL DEFAULT false,
  verification_status verification_status NOT NULL DEFAULT 'unverified',
  verification_document_url text,
  created_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS escrow_users (
  id bigserial PRIMARY KEY,
  email varchar(255) UNIQUE NOT NULL,
  hashed_password varchar(64) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS categories (
  id serial PRIMARY KEY,
  name varchar(64) UNIQUE NOT NULL,
  description text NOT NULL,
  cover_picture text NOT NULL
);

CREATE TABLE IF NOT EXISTS projects (
  id bigserial PRIMARY KEY,
  user_id bigint NOT NULL REFERENCES users(id),
  title text NOT NULL,
  description text NOT NULL,
  cover_picture text NOT NULL,
  category_id int NOT NULL REFERENCES categories(id),
  end_date timestamptz NOT NULL,
  receiver_number varchar(64) NOT NULL,
  receiver_name varchar(64) NOT NULL,
  address text NOT NULL,
  district varchar(64) NOT NULL,
  city varchar(64) NOT NULL,
  country varchar(64) NOT NULL,
  status project_status NOT NULL DEFAULT 'pending',
  created_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS milestones (
  id bigserial PRIMARY KEY,
  project_id bigint NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
  title text NOT NULL,
  description text,
  fund_goal bigint NOT NULL,
  current_fund bigint NOT NULL DEFAULT 0,
  bank_description text NOT NULL,
  status milestone_status NOT NULL DEFAULT 'pending',
  created_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS backings (
  id bigserial PRIMARY KEY,
  user_id bigint REFERENCES users(id) ON DELETE CASCADE,
  project_id bigint NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
  amount bigint NOT NULL,
  word_of_support text,
  created_at timestamptz NOT NULL DEFAULT NOW() 
);

CREATE TABLE IF NOT EXISTS escrow_milestone_completions (
  id bigserial PRIMARY KEY,
  milestone_id bigserial NOT NULL REFERENCES milestones(id),
  transfer_amount bigserial NOT NULL,
  transfer_note text,
  transfer_image text,
  created_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_spending_proofs (
  id bigserial PRIMARY KEY,
  milestone_id bigserial NOT NULL REFERENCES milestones(id),
  transfer_image text NOT NULL,
  proof_media_url text NOT NULL,
  description text NOT NULL,
  status proof_status NOT NULL DEFAULT 'pending',
  rejected_cause text,
  created_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS notifications (
  id bigserial PRIMARY KEY,
  user_id bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  notification_type notification_type NOT NULL,
  message text NOT NULL,
  project_id bigint REFERENCES projects(id),
  milestone_id bigint REFERENCES milestones(id),
  is_read boolean NOT NULL DEFAULT FALSE,
  created_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS audit_trails (
  id bigserial PRIMARY KEY,
  user_id bigint, -- either user_id or escrow_id is null, but not both
  escrow_id bigint,
  entity_type text NOT NULL, -- tables that is being manipulated
  entity_id bigserial, -- id of entity in that table that is being manipulated
  operation_type text NOT NULL, -- CREATE, UPDATE, DELETE
  field_name text NOT NULL,
  old_value jsonb,
  new_value jsonb,
  created_at timestamptz NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS audit_trails;
DROP TABLE IF EXISTS backings;
DROP TABLE IF EXISTS notifications;
DROP TABLE IF EXISTS escrow_milestone_completions;
DROP TABLE IF EXISTS user_spending_proofs;
DROP TABLE IF EXISTS milestones;
DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS categories;
-- DROP TABLE IF EXISTS escrow_users;
-- DROP TABLE IF EXISTS users;
-- DROP TYPE verification_status;
DROP TYPE project_status;
DROP TYPE milestone_status;
DROP TYPE notification_type;
DROP TYPE proof_status;