CREATE TABLE IF NOT EXISTS article (
    id INT AUTO_INCREMENT PRIMARY KEY,
    create_time DATETIME NOT NULL,
    article_type_id INT NOT NULL,
    article_sender_type_id INT NOT NULL
);