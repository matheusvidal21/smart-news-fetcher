-- +goose Down
ALTER TABLE articles DROP FOREIGN KEY articles_ibfk_1;

DROP TABLE IF EXISTS sources;
DROP TABLE IF EXISTS articles;