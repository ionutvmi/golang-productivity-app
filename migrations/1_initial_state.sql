-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE
    pomodoro_type (
        type_id INTEGER PRIMARY KEY AUTOINCREMENT,
        type_name VARCHAR(255)
    );

INSERT INTO
    pomodoro_type (type_name)
VALUES
    ('work'),
    ('break');

CREATE TABLE
    pomodoro (
        session_id INTEGER PRIMARY KEY AUTOINCREMENT,
        session_type INTEGER,
        start_date DATETIME,
        end_date DATETIME,
        FOREIGN KEY (session_type) REFERENCES session_types (type_id)
    );

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE pomodoro;

DROP TABLE pomodoro_type;