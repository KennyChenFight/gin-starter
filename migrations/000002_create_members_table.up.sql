CREATE TABLE IF NOT EXISTS members(
    id UUID PRIMARY KEY NOT NULL,
    email character varying(200) not null,
    password_digest character varying(1000) not null,
    name character varying(255) not null,
    created_at timestamp without time zone not null default current_timestamp,
    updated_at timestamp without time zone not null default current_timestamp,
    deleted_at timestamp without time zone
);