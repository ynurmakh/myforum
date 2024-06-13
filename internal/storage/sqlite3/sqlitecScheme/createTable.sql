-- database: /home/student/forum/internal/storage/sqlite3/database.db



-- сортировка по категориям

-- SELECT post_id, categories_id
-- FROM posts, json_each(posts.categories_id)
-- WHERE json_each.value in (1,2,3); "1,2,3,4"



-- Получение 1го поста сджойнам

-- SELECT posts.post_id, posts.post_title, posts.post_content, posts.created_time, users.user_id, users.user_lvl, users.user_email, users.user_nickname
-- FROM posts
-- JOIN users ON posts.user_id = users.user_id
-- WHERE posts.post_id = ?;


-- filter

-- WITH filtered_posts AS (
--     SELECT 
--         p.post_id,
--         p.user_id,
--         p.post_title,
--         p.post_content,
--         p.created_time,
--         c.category_id,
--         c.category_name,
--         ROW_NUMBER() OVER (PARTITION BY p.post_id ORDER BY c.category_id) AS row_num
--     FROM 
--         posts p
--     JOIN 
--         json_each(p.categories_id) AS j ON CAST(j.value AS INTEGER) IN (1,2)
--     JOIN 
--         categories_name c ON j.value = c.category_id
-- )
-- SELECT 
--     post_id,
--     user_id,
--     post_title,
--     post_content,
--     created_time,
--     category_id,
--     category_name
-- FROM 
--     filtered_posts
-- WHERE 
--     row_num = 1;




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