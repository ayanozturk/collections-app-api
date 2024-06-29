CREATE TABLE items (
                       id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
                       collection_id int NOT NULL,
                       name VARCHAR(255) NOT NULL,
                       description TEXT,
                       FOREIGN KEY (collection_id) REFERENCES collections(id) ON DELETE CASCADE
);
