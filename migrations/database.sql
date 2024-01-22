CREATE TABLE IF NOT EXISTS `attempts`
(
    id      INTEGER PRIMARY KEY AUTOINCREMENT,
    ip      TEXT    NOT NULL,
    attempt JSONB   NOT NULL,
    success BOOLEAN NOT NULL,
    date    DATE    NOT NULL
);