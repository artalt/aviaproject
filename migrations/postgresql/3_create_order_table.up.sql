CREATE TABLE "order"
(
    id           UUID                                                     NOT NULL,
    status       VARCHAR(64)                                              NOT NULL,
    payment_link VARCHAR(4096)                                            NOT NULL,
    created_at   TIMESTAMP(0) WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at   TIMESTAMP(0) WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
);
