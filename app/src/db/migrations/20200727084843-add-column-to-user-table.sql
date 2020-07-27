-- +migrate Up
CREATE TABLE public.users
(
    id                 serial       NOT NULL,
    encrypted_password varchar(255) NOT NULL,
    email              varchar(255) NOT NULL,
    created_at         timestamptz  NULL,
    updated_at         timestamptz  NULL,
    deleted            integer      NOT NULL DEFAULT 0,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);
CREATE UNIQUE INDEX uidx_users_email ON public.users USING btree (email, deleted);
-- +migrate Down
DROP TABLE public.users;