-- +goose Up
CREATE TYPE project_status AS ENUM (
  'pending',
  'ongoing',
  'rejected',
  'finished'
);

CREATE TYPE verification_status AS ENUM (
  'unverified',
  'pending',
  'verified'
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
  completed boolean NOT NULL DEFAULT FALSE,
  created_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS project_updates (
  id bigserial PRIMARY KEY,
  project_id bigint NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
  attachment_photo text,
  description text NOT NULL,
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

CREATE TABLE IF NOT EXISTS milestone_completions (
  id bigserial PRIMARY KEY,
  milestone_id bigserial NOT NULL REFERENCES milestones(id),
  transfer_amount bigserial NOT NULL,
  transfer_note text,
  transfer_image text,
  completed_at timestamptz NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS backings;
DROP TABLE IF EXISTS milestone_completions;
DROP TABLE IF EXISTS milestones;
DROP TABLE IF EXISTS project_updates;
DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS categories;
--DROP TABLE IF EXISTS escrow_users;
DROP TABLE IF EXISTS users;
DROP TYPE verification_status;
DROP TYPE project_status;
