CREATE TABLE users (
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    username VARCHAR(20) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)  ENGINE=INNODB

CREATE TABLE posts (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    subtitle VARCHAR(120) NOT NULL,
    author_id INT NOT NULL,
    FOREIGN KEY (author_id)
        REFERENCES users (id)
        ON DELETE CASCADE,
    content TEXT NOT NULL,
    votes INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)  ENGINE=INNODB

CREATE TABLE topics (
	id	INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    topic VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)  ENGINE=INNODB

CREATE TABLE post_topics (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    topic_id int NOT NULL,
    FOREIGN KEY (topic_id)
        REFERENCES topics (id)
        ON DELETE CASCADE,

    post_id INT NOT NULL,
    FOREIGN KEY (post_id)
        REFERENCES posts (id)
        ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)  ENGINE=INNODB;

CREATE TABLE post_comments (
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id)
        REFERENCES users (id)
        ON DELETE CASCADE,
	post_id INT NOT NULL,
    FOREIGN KEY (post_id)
        REFERENCES posts (id)
        ON DELETE CASCADE,
	comment TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)

CREATE TABLE post_votes (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    post_id INT NOT NULL,
    FOREIGN KEY (post_id)
        REFERENCES posts (id)
        ON DELETE CASCADE,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id)
        REFERENCES users (id)
        ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)

select * from posts;

insert into post_votes (post_id) values ()

insert into topics (topic) 
values 
("Blockchain"),
("Carrer"),
("Cloud"),
("Databases"),
("Data Science"),
("Developer tools"),
("DevOps & SRE"),
("DevRel"),
("Fundamentls"),
("Gaming"),
("Mobile Development"),
("Programming language"),
("Security"),
("Testing"),
("Web Development")