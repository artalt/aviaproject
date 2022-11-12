ALTER TABLE user_order DROP CONSTRAINT FK_user_order_user_id;
ALTER TABLE user_order DROP CONSTRAINT FK_user_order_order_id;
DROP TABLE "order";
DROP TABLE user_order;
