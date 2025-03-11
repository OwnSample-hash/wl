DELIMITER //
DROP TRIGGER IF EXISTS `prevent_update_final_price`//
CREATE TRIGGER `prevent_update_final_price`
    BEFORE UPDATE ON store_orders
    FOR EACH ROW
    BEGIN
      SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = 'Final price cannot be updated';
    END //
DELIMITER ;
