CREATE TABLE courses (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(255) NOT NULL,
                         description TEXT NOT NULL,
                         age_group VARCHAR(20) NOT NULL,
                         price_per_class DECIMAL(10,2) NOT NULL,
                         duration INT NOT NULL
);

INSERT INTO courses (name, description, age_group, price_per_class, duration) VALUES
                        ('Баскетбол для начинающих', 'Базовые навыки и техника.', '6-10 лет', 500.00, 10),
                        ('Йога для детей', 'Гибкость и координация.', '8-12 лет', 450.00, 12),
                        ('Футбол для подростков', 'Улучшение техники и подготовки.', '13-17 лет', 600.00, 15),
                        ('Плавание для взрослых', 'Обучение плаванию.', '18+ лет', 700.00, 20),
                                                                                  ('Теннис для детей', 'Основы тенниса.', '6-15 лет', 550.00, 12);

CREATE TABLE trainers (
                          id SERIAL PRIMARY KEY,
                          full_name VARCHAR(255) NOT NULL,
                          email VARCHAR(255) NOT NULL UNIQUE,
                          phone_number VARCHAR(20) NOT NULL,
                          specialization INT NOT NULL,
                          address TEXT NOT NULL,
                          FOREIGN KEY (specialization) REFERENCES courses(id) ON DELETE CASCADE
);

INSERT INTO trainers (full_name, email, phone_number, specialization, address) VALUES
                        ('Иван Петров', 'ivan.petrov@example.com', '+7 900 123 45 67', 1, 'Москва, ул. Ленина, д. 10'),
                        ('Ольга Смирнова', 'olga.smirnova@example.com', '+7 901 234 56 78', 2, 'СПб, ул. Невский пр., д. 20'),
                        ('Алексей Иванов', 'alex.ivanov@example.com', '+7 902 345 67 89', 3, 'Екатеринбург, ул. Мира, д. 5'),
                        ('Наталья Кузнецова', 'nat.kuz@example.com', '+7 903 456 78 90', 4, 'Казань, ул. Баумана, д. 12'),
                        ('Дмитрий Соколов', 'd.sokolov@example.com', '+7 904 567 89 01', 5, 'Новосибирск, ул. Красный проспект, д. 30');

CREATE TABLE lessons (
                         id SERIAL PRIMARY KEY,
                         course_id INT NOT NULL,
                         trainer_id INT NOT NULL,
                         date DATE NOT NULL,
                         start_time TIME NOT NULL,
                         end_time TIME NOT NULL,
                         FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE,
                         FOREIGN KEY (trainer_id) REFERENCES trainers(id) ON DELETE CASCADE
);

INSERT INTO lessons (course_id, trainer_id, date, start_time, end_time) VALUES
                        (1, 1, '2025-04-01', '10:00:00', '11:30:00'),
                        (2, 2, '2025-04-02', '14:00:00', '15:30:00'),
                        (3, 3, '2025-04-03', '16:00:00', '17:30:00'),
                        (4, 4, '2025-04-04', '18:00:00', '19:30:00'),
                        (5, 5, '2025-04-05', '09:00:00', '10:30:00');

CREATE TABLE students (
                          id SERIAL PRIMARY KEY,
                          full_name VARCHAR(255) NOT NULL,
                          phone_number VARCHAR(20) NOT NULL,
                          age INT NOT NULL,
                          registration_date DATE NOT NULL,
                          course_id INT NOT NULL,
                          FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE
);

INSERT INTO students (full_name, phone_number, age, registration_date, course_id) VALUES
                        ('Алексей Смирнов', '+7 900 123 45 67', 10, '2025-03-01', 1),
                        ('Мария Иванова', '+7 901 234 56 78', 12, '2025-03-02', 2),
                        ('Иван Петров', '+7 902 345 67 89', 15, '2025-03-03', 3),
                        ('Екатерина Соколова', '+7 903 456 78 90', 8, '2025-03-04', 4),
                        ('Дмитрий Кузнецов', '+7 904 567 89 01', 17, '2025-03-05', 5);

CREATE TABLE payments (
                          id SERIAL PRIMARY KEY,
                          student_id INT NOT NULL,
                          course_id INT NOT NULL,
                          attended_lessons INT NOT NULL,
                          paid_lessons INT NOT NULL,
                          FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE,
                          FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE
);

INSERT INTO payments (student_id, course_id, attended_lessons, paid_lessons) VALUES
                        (1, 1, 5, 10),
                        (2, 2, 8, 10),
                        (3, 3, 3, 5),
                        (4, 4, 7, 8),
                        (5, 5, 10, 10);
