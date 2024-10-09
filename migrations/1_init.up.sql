CREATE TABLE IF NOT EXISTS songs
(
    id           SERIAL PRIMARY KEY, -- Уникальный идентификатор записи
    group_name   TEXT NOT NULL,      -- Название группы
    song         TEXT NOT NULL,      -- Название песни
    release_date DATE NOT NULL,      -- Дата выхода песни
    text         TEXT NOT NULL,      -- Текст песни
    link         TEXT NOT NULL,      -- Ссылка на песню
    created_at   TIMESTAMPTZ DEFAULT NOW()
);