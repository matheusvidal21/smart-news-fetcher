CREATE TABLE users (
                       id INT AUTO_INCREMENT PRIMARY KEY,
                       username VARCHAR(100) NOT NULL,
                       email VARCHAR(100) NOT NULL,
                       password TEXT NOT NULL
);


CREATE TABLE sources (
                        id INT AUTO_INCREMENT PRIMARY KEY,
                        name VARCHAR(100) NOT NULL,
                        url TEXT NOT NULL,
                        user_id INT NOT NULL,
                        saved_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE articles (
                        id VARCHAR(36) PRIMARY KEY,
                        title TEXT NOT NULL,
                        description TEXT,
                        content TEXT,
                        link TEXT NOT NULL,
                        pub_date TIMESTAMP,
                        author TEXT,
                        source_id INT NOT NULL,
                        FOREIGN KEY (source_id) REFERENCES sources(id)
);
