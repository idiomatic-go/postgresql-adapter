CREATE TRIGGER slo_entry_changes
    AFTER INSERT, UPDATE, DELETE
    ON slo_entry
    FOR EACH ROW
    EXECUTE PROCEDURE LogSLOEntryChanges();