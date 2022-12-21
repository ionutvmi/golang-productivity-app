-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE pomodoro_sessions (
    session_id INT,
    created_date DATETIME,
    modified_date DATETIME
);
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE pomodoro_sessions;