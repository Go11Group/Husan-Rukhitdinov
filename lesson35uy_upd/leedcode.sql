CREATE TABLE Problems (
    problem_id uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    title VARCHAR(255),
    difficulty VARCHAR(10) CHECK (difficulty IN ('Easy', 'Medium', 'Hard')) NOT NULL,
    description TEXT
);

CREATE TABLE Users (
    user_id uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    username VARCHAR(50),
    email VARCHAR(100),
    password_hash VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE solved_problems (
    id uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    user_id uuid not null REFERENCES users(user_id),
    problem_id uuid NOT NULL REFERENCES Problems(problem_id),
    solved_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   
);
drop table users;


INSERT INTO Problems (title, difficulty, description) VALUES
('Two Sum', 'Easy', 'Given an array of integers, return indices of the two numbers such that they add up to a specific target.'),
('Longest Substring Without Repeating Characters', 'Medium', 'Given a string, find the length of the longest substring without repeating characters.'),
('Median of Two Sorted Arrays', 'Hard', 'Given two sorted arrays, find the median of the two sorted arrays.');

INSERT INTO Users (username, email, password_hash) VALUES
('alice', 'alice@example.com', 'hashedpassword1'),
('bob', 'bob@example.com', 'hashedpassword2'),
('carol', 'carol@example.com', 'hashedpassword3');

INSERT INTO Solved_Problems (user_id, problem_id) VALUES
(1, 1),
(1, 2),
(2, 1),
(2, 3),
(3, 2);
