DELIMITER //
DROP TRIGGER IF EXISTS `create_uuid`//
CREATE TRIGGER `create_uuid`
    AFTER UPDATE ON store_orders
    FOR EACH ROW
    BEGIN
      IF NEW.status = 'completed' THEN
        SET @uuid = UUID(); -- TODO: replace UUID with UUID_v7() once it is available
        IF NEW.is_lifetime THEN
          INSERT INTO store_licenses (uuid, valid_till) VALUES (@uuid, 0);
        ELSE
          INSERT INTO store_licenses (uuid, valid_till) VALUES (@uuid, DATE_ADD(CURDATE(), INTERVAL 1 MONTH));
        END IF;
        UPDATE store_orders SET license_id = (SELECT id FROM store_licenses WHERE uuid = @uuid), completed_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
      END IF;
    END //
DELIMITER ;
