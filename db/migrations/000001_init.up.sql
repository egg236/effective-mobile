CREATE TABLE IF NOT EXISTS records (
                         id SERIAL PRIMARY KEY,
                         service_name VARCHAR(255) NOT NULL,
                         price INTEGER NOT NULL,
                         user_id VARCHAR(255) NOT NULL,
                         start_date varchar(255) NOT NULL,
                         end_date varchar(255)
);