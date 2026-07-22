CREATE TABLE IF NOT EXISTS scheduler (
    id INT PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT '', 
    title TEXT NOT NULL, 
    comment TEXT NOT NULL DEFAULT '', 
    repeat VARCHAR(128) NOT NULL DEFAULT '', 
);

CREATE INDEX IF NOT EXISTS scheduler_date ON scheduler (date);