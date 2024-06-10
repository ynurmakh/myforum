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
    user_id INTEGER PRIMARY KEY NOT NULL ,
    user_lvl INTEGER NOT NULL DEFAULT 1,
    user_email VARCHAR(50) NOT NULL UNIQUE,
    user_nickname VARCHAR(50) NOT NULL UNIQUE,
    hashed_password VARCHAR(64)
    );
CREATE TABLE IF NOT EXISTS cookies (
    cookie VARCHAR(56) NOT NULL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    deadTime TIMESTAMP NOT NULL
    );

CREATE TABLE IF NOT EXISTS posts (
    post_id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL,
    post_title VARCHAR(50),
    post_content VARCHAR(2048),
    cteated_time CURRENT_TIMESTAMP NOT NULL
    );

CREATE TABLE IF NOT EXISTS categries_name (
    category_id INTEGER PRIMARY KEY,
    category_name VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS posts_category (
    post_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL
    );
CREATE TABLE IF NOT EXISTS commetaries (
    commentary_id INTEGER PRIMARY KEY,
    post_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    commentray_content VARCHAR(2048),
    created_time CURRENT_TIMESTAMP
    );
CREATE TABLE IF NOT EXISTS post_reactions(
    user_id INTEGER,
    post_id INTEGER,
    reaction INTEGER
);
CREATE TABLE IF NOT EXISTS commentari_reactions(
    user_id INTEGER,
    commentaie_id INTEGER,
    reaction INTEGER
);