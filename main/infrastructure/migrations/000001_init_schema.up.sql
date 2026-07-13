CREATE TABLE IF NOT EXISTS notes
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT
    );
