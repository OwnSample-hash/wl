CREATE OR REPLACE TRIGGER `calc_final_price_insert`
AFTER
INSERT ON store_orders FOR EACH ROW BEGIN
SET @discount_c = (
    SELECT discount
    FROM store_coupons
    WHERE id = NEW.coupon_id
  );
SET @discount_p = (
    SELECT discount
    FROM store_products
    WHERE id = NEW.product_id
  );
IF NEW.is_lifetime THEN
SET @final_price = (
    (
      SELECT price
      FROM store_products
      WHERE id = NEW.product_id
    ) - @discount_c
  ) * (100 - @discount_p) / 100;
UPDATE store_orders
SET final_price = @final_price
WHERE id = NEW.id;
ELSE
SET @final_price = (
    (
      SELECT price_per_month
      FROM store_products
      WHERE id = NEW.product_id
    ) - @discount_c
  ) * (100 - @discount_p) / 100;
UPDATE store_orders
SET final_price = @final_price
WHERE id = NEW.id;
END IF;
END;
