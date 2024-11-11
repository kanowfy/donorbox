-- +goose Up
CREATE TYPE project_status AS ENUM (
  'pending',
  'ongoing',
  'rejected',
  'finished'
);

CREATE TABLE IF NOT EXISTS users (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  email varchar(255) UNIQUE NOT NULL,
  first_name varchar(64) NOT NULL,
  last_name varchar(64) NOT NULL,
  profile_picture text,
  hashed_password varchar(64) NOT NULL,
  activated boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS escrow_users (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  email varchar(255) UNIQUE NOT NULL,
  hashed_password varchar(64) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS categories (
  id SERIAL PRIMARY KEY,
  name varchar(64) UNIQUE NOT NULL,
  description text NOT NULL,
  cover_picture text NOT NULL
);

CREATE TABLE IF NOT EXISTS projects (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id uuid NOT NULL REFERENCES users(id),
  title text NOT NULL,
  description text NOT NULL,
  cover_picture text NOT NULL,
  category_id int NOT NULL REFERENCES categories(id),
  total_fund bigint NOT NULL DEFAULT 0,
  start_date timestamptz NOT NULL DEFAULT NOW(),
  end_date timestamptz NOT NULL,
  receiver_number varchar(64) NOT NULL,
  receiver_name varchar(64) NOT NULL,
  address text NOT NULL,
  district varchar(64) NOT NULL,
  city varchar(64) NOT NULL,
  country varchar(64) NOT NULL,
  status project_status DEFAULT 'pending'
);

CREATE TABLE IF NOT EXISTS milestones (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  project_id uuid NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
  title text NOT NULL,
  description text,
  fund_goal bigint NOT NULL,
  current_fund bigint NOT NULL DEFAULT 0,
  bank_description text NOT NULL,
  completed boolean NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS project_updates (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  project_id uuid NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
  attachment_photo text,
  description text NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS backings (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id uuid REFERENCES users(id),
  project_id uuid NOT NULL REFERENCES projects(id),
  amount bigint NOT NULL,
  word_of_support text,
  created_at timestamptz NOT NULL DEFAULT NOW() 
);

CREATE TABLE IF NOT EXISTS certificates (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  escrow_user_id uuid NOT NULL REFERENCES escrow_users(id),
  user_id uuid NOT NULL REFERENCES users(id),
  milestone_id uuid NOT NULL references milestones(id),
  verified bool DEFAULT false,
  verified_at timestamptz,
  created_at timestamptz NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TYPE project_status;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS escrow_users;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS project_updates;
DROP TABLE IF EXISTS milestones;
DROP TABLE IF EXISTS backings;
DROP TABLE IF EXISTS certificates;
