CREATE TABLE Problems (
    problem_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    difficulty VARCHAR(10) CHECK (difficulty IN ('Easy', 'Medium', 'Hard')) NOT NULL,
    description TEXT
);

CREATE TABLE Users (
    user_id INT PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Solved_Problems (
    id SERIAL PRIMARY KEY,
    user_id INT,
    problem_id INT,
    solved_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, problem_id),
    FOREIGN KEY (user_id) REFERENCES Users(user_id),
    FOREIGN KEY (problem_id) REFERENCES Problems(problem_id)
);


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
