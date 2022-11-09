CREATE TRIGGER msre.customer_changes
    AFTER INSERT, UPDATE, DELETE
    ON msre.customer
    FOR EACH ROW
    EXECUTE PROCEDURE msre.LogCustomerChanges();