CREATE TABLE IF NOT EXISTS users (
    user_id       UUID PRIMARY KEY,
    username      VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS scores (
    record_id   SERIAL PRIMARY KEY,
    user_id     UUID NOT NULL,
    WPM         INT NOT NULL,
    accuracy    INT NOT NULL,
    typos       INT NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);