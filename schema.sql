CREATE TYPE user_type AS ENUM (
    'regular',
    'escrow'
);

CREATE TYPE project_status AS ENUM (
	'ongoing',
	'ended',
	'completed_payout',
	'completed_refund'
);

CREATE TYPE backing_status AS ENUM (
	'pending',
	'released',
	'refunded'
);

CREATE TYPE transaction_type AS ENUM (
	'backing',
	'payout',
	'refund'
);

CREATE TYPE transaction_status AS ENUM (
	'pending',
	'completed',
	'failed'
);

CREATE TYPE card_brand AS ENUM (
	'VISA',
	'MASTERCARD'
);

CREATE TABLE IF NOT EXISTS cards (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	token char(29) UNIQUE NOT NULL,
	card_owner_name varchar(255) NOT NULL,
	last_four_digits char(4) NOT NULL,
	brand card_brand NOT NULL,
	created_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS users (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	email varchar(255) UNIQUE NOT NULL,
	hashed_password text NOT NULL,
	first_name varchar(64) NOT NULL,
	last_name varchar(64) NOT NULL,
	profile_picture text,
	activated boolean NOT NULL DEFAULT false,
	user_type user_type NOT NULL DEFAULT 'regular',
	created_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS escrow_users (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	email varchar(255) UNIQUE NOT NULL,
	hashed_password text NOT NULL,
	user_type user_type NOT NULL DEFAULT 'escrow',
	card_id uuid REFERENCES cards(id),
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
	category_id integer NOT NULL REFERENCES categories(id),
	title text NOT NULL,
	description text NOT NULL,
	cover_picture text NOT NULL,
	goal_amount bigint NOT NULL,
	current_amount bigint NOT NULL DEFAULT 0,
	country varchar(64) NOT NULL,
	province varchar(64) NOT NULL,
	card_id uuid REFERENCES cards(id),
	start_date timestamptz NOT NULL DEFAULT NOW(),
	end_date timestamptz NOT NULL,
	status project_status NOT NULL DEFAULT 'ongoing'
);

CREATE TABLE IF NOT EXISTS backings (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	project_id uuid NOT NULL REFERENCES projects(id),
	backer_id uuid REFERENCES users(id),
	amount bigint NOT NULL,
	word_of_support text,
	status backing_status NOT NULL DEFAULT 'pending',
	created_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS project_updates (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	project_id uuid NOT NULL REFERENCES projects(id),
	attachment_photo text,
	description text NOT NULL,
	created_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS transactions (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	project_id uuid NOT NULL REFERENCES projects(id),
	transaction_type transaction_type NOT NULL,
	amount bigint NOT NULL,
	initiator_card_id uuid NOT NULL REFERENCES cards(id),
	recipient_card_id uuid NOT NULL REFERENCES cards(id),
	status transaction_status NOT NULL DEFAULT 'pending',
	created_at timestamptz NOT NULL DEFAULT NOW()
);
