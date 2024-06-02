CREATE TABLE sources (
                         id INT AUTO_INCREMENT PRIMARY KEY,
                         name VARCHAR(100) NOT NULL,
                         url TEXT NOT NULL,
                         saved_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE articles (
                        id INT AUTO_INCREMENT PRIMARY KEY,
                        title TEXT NOT NULL,
                        description TEXT,
                        content TEXT,
                        link TEXT NOT NULL,
                        pub_date TIMESTAMP,
                        author TEXT,
                        source_id INT NOT NULL,
                        FOREIGN KEY (source_id) REFERENCES sources(id)
);

CREATE TABLE users (
                       id INT AUTO_INCREMENT PRIMARY KEY,
                       username VARCHAR(100) NOT NULL,
                       email VARCHAR(100) NOT NULL,
                       password TEXT NOT NULL
);
