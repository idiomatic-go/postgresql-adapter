CREATE TRIGGER customer_changes
    AFTER INSERT, UPDATE, DELETE
    ON customer
    FOR EACH ROW
    EXECUTE PROCEDURE LogCustomerChanges();