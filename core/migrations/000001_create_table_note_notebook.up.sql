CREATE TABLE IF NOT EXISTS notebook (
    id VARCHAR PRIMARY KEY,
    name TEXT NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS note (
    id VARCHAR PRIMARY KEY,
    notebook_id VARCHAR NOT NULL,
    content TEXT NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (notebook_id) REFERENCES notebook(id)
);