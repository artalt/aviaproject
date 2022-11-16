CREATE TABLE flight
(
    id                  UUID                                                     NOT NULL,
    number              VARCHAR(256)                                             NOT NULL,
    departure           VARCHAR(256)                                             NOT NULL,
    arrival             VARCHAR(256)                                             NOT NULL,
    departure_date_time TIMESTAMP(0) WITHOUT TIME ZONE                           NOT NULL,
    arrival_date_time   TIMESTAMP(0) WITHOUT TIME ZONE                           NOT NULL,
    created_at          TIMESTAMP(0) WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at          TIMESTAMP(0) WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
);
