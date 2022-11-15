CREATE TABLE ticket
(
    id              UUID                                                     NOT NULL,
    flight_id       UUID                                                     NOT NULL,
    order_id        UUID                           DEFAULT NULL,
    status          VARCHAR(64)                                              NOT NULL,
    type            VARCHAR(64)                                              NOT NULL,
    pass_first_name VARCHAR(256)                   DEFAULT NULL,
    pass_last_name  VARCHAR(256)                   DEFAULT NULL,
    seat            VARCHAR(64)                                              NOT NULL,
    luggage         INT                                                      NOT NULL,
    price           INT                                                      NOT NULL,
    can_return      BOOLEAN                                                  NOT NULL,
    created_at      TIMESTAMP(0) WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at      TIMESTAMP(0) WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
);
CREATE INDEX IDX_ticket_flight_id ON ticket (flight_id);
CREATE INDEX IDX_ticket_order_id ON ticket (order_id);
ALTER TABLE ticket ADD CONSTRAINT FK_ticket_flight_id FOREIGN KEY (flight_id) REFERENCES flight (id) NOT DEFERRABLE INITIALLY IMMEDIATE;
ALTER TABLE ticket ADD CONSTRAINT FK_ticket_order_id FOREIGN KEY (order_id) REFERENCES "order" (id) NOT DEFERRABLE INITIALLY IMMEDIATE;
