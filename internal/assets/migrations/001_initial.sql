-- +migrate Up

CREATE TABLE kdf
(
    version   bigint NOT NULL,
    algorithm text   NOT NULL,
    bits      bigint NOT NULL,
    n         bigint NOT NULL,
    r         bigint NOT NULL,
    p         bigint NOT NULL
);

CREATE TABLE wallets
(
    id            bigserial primary key,
    wallet_id     character varying(256),
    email         character varying(256) NOT NULL,
    keychain_data text,
    salt          text                   NOT NULL,
    verified      boolean DEFAULT false  NOT NULL
    --  verification_token varchar(255) DEFAULT ''::varchar NOT NULL
);

CREATE TABLE email_tokens
(
    id                 bigserial primary key,
    token              text,
    last_sent_at       timestamp without time zone,
    confirmed          boolean DEFAULT false,
    email              text not null
);

-- insert default kdf versions
INSERT INTO kdf (version, algorithm, bits, n, r, p)
VALUES ('1', 'scrypt', 256, 4096, 8, 1);

-- +migrate Down

DROP TABLE kdf cascade;
DROP TABLE wallets cascade;
