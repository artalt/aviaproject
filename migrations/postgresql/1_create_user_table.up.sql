CREATE TABLE "user"
(
    id         UUID                                                     NOT NULL,
    login      VARCHAR(256)                                             NOT NULL,
    first_name VARCHAR(256)                                             NOT NULL,
    last_name  VARCHAR(256)                                             NOT NULL,
    phone      VARCHAR(64)                                              NOT NULL,
    created_at TIMESTAMP(0) WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP(0) WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
);
