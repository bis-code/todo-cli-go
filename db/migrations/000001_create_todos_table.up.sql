CREATE TABLE todos
(
    id           SERIAL PRIMARY KEY,
    title        VARCHAR(255) NOT NULL,
    description  TEXT,
    completed    BOOLEAN   DEFAULT FALSE,
    category     VARCHAR(100),
    tags         TEXT[], -- PostgreSQL array type
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP
);