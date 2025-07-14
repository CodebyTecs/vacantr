CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    telegram_id BIGINT UNIQUE NOT NULL,
    username TEXT
);

CREATE TABLE IF NOT EXISTS vacancies (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL
);

CREATE TABLE user_vacancies (
    user_id BIGINT NOT NULL,
    vacancy_id BIGINT NOT NULL,
    PRIMARY KEY (user_id, vacancy_id),
    FOREIGN KEY (user_id) REFERENCES users(telegram_id),
    FOREIGN KEY (vacancy_id) REFERENCES vacancies(id)
);