ALTER TABLE ticket DROP CONSTRAINT FK_ticket_flight_id;
ALTER TABLE ticket DROP CONSTRAINT FK_ticket_order_id;
DROP TABLE ticket;
