CREATE TABLE posts (
    id SERIAL NOT NULL, 
    title character varying NOT NULL,
    description character varying NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()), 
    updated_at timestamptz NOT NULL DEFAULT (now())
);