-- database: /home/student/forum/internal/storage/sqlite3/database.db
CREATE TABLE IF NOT EXISTS users (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_lvl INTEGER NOT NULL DEFAULT 1,
    user_email VARCHAR(50) NOT NULL UNIQUE,
    user_nickname VARCHAR(50) NOT NULL UNIQUE,
    hashed_password VARCHAR(64)
);
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
    categories_id TEXT NOT NULL DEFAULT '[-1]', 
    liked_ids TEXT not null default '[]', 
    disliked_ids TEXT not null default '[]',
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS categories_name (
    category_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    category_name VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE "commentaries" (
    commentary_id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    commentray_content VARCHAR(2048) NOT NULL,
    liked_ids TEXT NOT NULL DEFAULT '[]',
    disliked_ids TEXT NOT NULL DEFAULT '[]',
    created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);


UPDATE cookies
SET last_call = datetime(CURRENT_TIMESTAMP) , livetime = 5
WHERE cookies.cookie = 'cooki'