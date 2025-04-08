CREATE OR REPLACE TRIGGER `adjust_remaining_count` BEFORE
INSERT ON store_orders FOR EACH ROW BEGIN IF (
    SELECT s_c.remaining
    FROM store_coupons s_c
    WHERE s_c.id = NEW.coupon_id
      AND s_c.status = 'active'
      AND DATE(s_c.exp) > CURDATE()
  ) > 0 THEN
UPDATE store_coupons
SET remaining = remaining - 1
WHERE store_coupons.id = NEW.coupon_id;
ELSE IF (
  SELECT s_c.status
  FROM store_coupons s_c
  WHERE s_c.id = NEW.coupon_id
) = 'inactive' THEN SIGNAL SQLSTATE '45000'
SET MESSAGE_TEXT = 'Coupon is not active';
END IF;
IF (
  SELECT DATE(s_c.exp)
  FROM store_coupons s_c
  WHERE s_c.id = NEW.coupon_id
) < CURDATE() THEN SIGNAL SQLSTATE '45000'
SET MESSAGE_TEXT = 'Coupon is expired';
END IF;
IF (
  SELECT s_c.remaining
  FROM store_coupons s_c
  WHERE s_c.id = NEW.coupon_id
) <= 0 THEN SIGNAL SQLSTATE '45000'
SET MESSAGE_TEXT = 'Coupon has no use left';
END IF;
END IF;
END;
