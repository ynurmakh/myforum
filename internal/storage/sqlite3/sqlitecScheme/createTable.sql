-- database: /home/student/forum/internal/storage/sqlite3/database.db

-- OK users
-- OK cookies
-- OK posts
-- OK post_category
-- OK categoryes_name
-- OK commemtaries
-- OK post_reacrions
-- OK commentary_reacrions

CREATE TABLE IF NOT EXISTS users (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_lvl INTEGER NOT NULL DEFAULT 1,
    user_email VARCHAR(50) NOT NULL UNIQUE,
    user_nickname VARCHAR(50) NOT NULL UNIQUE,
    hashed_password VARCHAR(64)
);


-- SELECT posts.post_id, posts.post_title, posts.post_content, posts.created_time, users.user_id, users.user_lvl, users.user_email, users.user_nickname
-- FROM posts
-- JOIN users ON posts.user_id = users.user_id
-- WHERE posts.post_id = ?;

CREATE TABLE IF NOT EXISTS cookies (
    cookie VARCHAR(56) NOT NULL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    liveTime INTEGER NOT NULL,
    last_call TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS posts (
    post_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    user_id INTEGER NOT NULL,
    post_title VARCHAR(50) NOT NULL,
    post_content VARCHAR(2048) NOT NULL,
    created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS categories_name (
    category_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    category_name VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS post_category (
    post_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categries_name(category_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS commentaries (
    commentary_id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    commentray_content VARCHAR(2048) NOT NULL,
    created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS post_reactions (
    user_id INTEGER,
    post_id INTEGER,
    reaction INTEGER,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS commentari_reactions (
    user_id INTEGER,
    commentaie_id INTEGER,
    reaction INTEGER,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (commentaie_id) REFERENCES commentaries(commentary_id) ON DELETE CASCADE
);