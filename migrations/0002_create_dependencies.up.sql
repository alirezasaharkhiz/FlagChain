CREATE TABLE dependencies
(
    id            INT AUTO_INCREMENT PRIMARY KEY,
    flag_id       INT NOT NULL,
    depends_on_id INT NOT NULL,
    created_at    DATETIME,
    updated_at    DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (flag_id) REFERENCES flags (id),
    FOREIGN KEY (depends_on_id) REFERENCES flags (id)
);