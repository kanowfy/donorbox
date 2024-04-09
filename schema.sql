CREATE TYPE user_type AS ENUM (
    'regular',
    'escrow'
);

CREATE TYPE backing_status AS ENUM (
	'pending',
	'released',
	'refunded'
);

CREATE TABLE IF NOT EXISTS users (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	username varchar(64) UNIQUE NOT NULL,
	hashed_password text NOT NULL,
	email varchar(255) UNIQUE NOT NULL,
	first_name varchar(64),
	last_name varchar(64),
	profile_picture text,
	activated boolean NOT NULL DEFAULT false,
	user_type user_type NOT NULL DEFAULT 'regular'
);

CREATE TABLE IF NOT EXISTS categories (
	id SERIAL PRIMARY KEY,
	name varchar(64) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS projects (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id uuid NOT NULL REFERENCES users(id),
	category_id integer NOT NULL REFERENCES categories(id),
	title text NOT NULL,
	description text NOT NULL,
	cover_picture text NOT NULL,
	goal_amount decimal(10,2) NOT NULL,
	current_amount decimal(10,2) NOT NULL DEFAULT 0.00,
	country varchar(64) NOT NULL,
	province varchar(64) NOT NULL,
	start_date timestamptz NOT NULL DEFAULT NOW(),
	end_date timestamptz NOT NULL,
	is_active boolean NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS backings (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	project_id uuid NOT NULL REFERENCES projects(id),
	backer_id uuid NOT NULL REFERENCES users(id),
	amount decimal(10,2) NOT NULL,
	backing_date timestamptz NOT NULL DEFAULT NOW(),
	status backing_status NOT NULL DEFAULT 'pending'
);

CREATE TABLE IF NOT EXISTS project_updates (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	project_id uuid NOT NULL REFERENCES projects(id),
	description text NOT NULL,
	update_date timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS project_comments (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	project_id uuid NOT NULL REFERENCES projects(id),
	backer_id uuid NOT NULL REFERENCES users(id),
	parent_comment_id uuid REFERENCES project_comments(id),
	content text NOT NULL,
	commented_at timestamptz NOT NULL DEFAULT NOW()
);
