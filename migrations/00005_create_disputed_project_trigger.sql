-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION update_project_status()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE projects
    SET status = 'disputed'
    WHERE id = NEW.project_id;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER project_disputed_status_update
AFTER UPDATE OF status ON milestones
FOR EACH ROW
WHEN (NEW.status = 'refuted')
EXECUTE FUNCTION update_project_status();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION update_project_status() CASCADE;
-- +goose StatementEnd
