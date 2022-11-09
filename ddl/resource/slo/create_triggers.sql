CREATE TRIGGER msre.slo_entry_changes
    AFTER INSERT, UPDATE, DELETE
    ON msre.slo_entry
    FOR EACH ROW
    EXECUTE PROCEDURE msre.LogSLOEntryChanges();