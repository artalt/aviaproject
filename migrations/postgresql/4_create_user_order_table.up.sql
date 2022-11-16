CREATE TABLE user_order (user_id UUID NOT NULL, order_id UUID NOT NULL, PRIMARY KEY(user_id, order_id));
CREATE INDEX IDX_user_order_user_id ON user_order (user_id);
CREATE INDEX IDX_user_order_order_id ON user_order (order_id);
ALTER TABLE user_order ADD CONSTRAINT FK_user_order_user_id FOREIGN KEY (user_id) REFERENCES "user" (id) ON DELETE CASCADE NOT DEFERRABLE INITIALLY IMMEDIATE;
ALTER TABLE user_order ADD CONSTRAINT FK_user_order_order_id FOREIGN KEY (order_id) REFERENCES "order" (id) ON DELETE CASCADE NOT DEFERRABLE INITIALLY IMMEDIATE;
