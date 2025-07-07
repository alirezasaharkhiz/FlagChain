CREATE TABLE flags
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    name       VARCHAR(255) NOT NULL UNIQUE,
    enabled    BOOLEAN DEFAULT FALSE,
    created_at DATETIME,
    updated_at DATETIME
);