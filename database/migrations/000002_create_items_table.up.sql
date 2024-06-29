CREATE TABLE items (
                       id VARCHAR(36) NOT NULL PRIMARY KEY,
                       collection_id VARCHAR(36) NOT NULL,
                       name VARCHAR(255) NOT NULL,
                       description TEXT,
                       FOREIGN KEY (collection_id) REFERENCES collections(id) ON DELETE CASCADE
);
