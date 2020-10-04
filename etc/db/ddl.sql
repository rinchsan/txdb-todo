CREATE TABLE users (
    id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    username VARCHAR(191) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
INSERT INTO users (id, username) VALUES
	(1, 'John'),
	(2, 'Charles'),
	(3, 'Herbert')
;

CREATE TABLE todos (
    id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    title VARCHAR(191) NOT NULL,
    detail TEXT NOT NULL,
    author_user_id BIGINT(20) UNSIGNED NOT NULL,
    due_date DATE NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT todos_author_user_id FOREIGN KEY (author_user_id) REFERENCES users (id)
);

CREATE TABLE todos_users (
    todo_id BIGINT(20) UNSIGNED NOT NULL,
    user_id BIGINT(20) UNSIGNED NOT NULL,
    PRIMARY KEY (todo_id, user_id),
    CONSTRAINT todos_users_todo_id FOREIGN KEY (todo_id) REFERENCES todos (id),
    CONSTRAINT todos_users_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);
