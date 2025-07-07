CREATE TABLE audit_logs
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    flag_id    INT NOT NULL,
    action     VARCHAR(50),
    actor      VARCHAR(100),
    reason     TEXT,
    created_at DATETIME,
    FOREIGN KEY (flag_id) REFERENCES flags (id)
);