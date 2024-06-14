CREATE TABLE Users (
    user_id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    birthday TIMESTAMP NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);



CREATE TABLE Courses (
    course_id UUID PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at BEGINT DEFAULT 0
);

CREATE TABLE Lessons (
    lesson_id UUID PRIMARY KEY,
    course_id UUID NOT NULL,
    title VARCHAR(100) NOT NULL,
    content TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0,
    FOREIGN KEY (course_id) REFERENCES Courses(course_id)
);


CREATE TABLE Enrollments (
    enrollment_id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    course_id UUID NOT NULL,
    enrollment_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0,
    FOREIGN KEY (user_id) REFERENCES Users(user_id),
    FOREIGN KEY (course_id) REFERENCES Courses(course_id)
);

ALTER TABLE Enrollments
ADD CONSTRAINT enrollments_user_id_fkey
FOREIGN KEY (user_id)
REFERENCES Users(user_id)
ON DELETE CASCADE;







INSERT INTO Users (user_id, name, email, birthday, password, created_at, updated_at, deleted_at)
VALUES
(gen_random_uuid(), 'User1', 'user1@example.com', '1990-01-01', 'password1', now(), now(), NULL),
(gen_random_uuid(), 'User2', 'user2@example.com', '1991-02-01', 'password2', now(), now(), NULL),
(gen_random_uuid(), 'User50', 'user50@example.com', '1990-02-19', 'password50', now(), now(), NULL);


INSERT INTO Courses (course_id, title, description, created_at, updated_at, deleted_at)
VALUES
(gen_random_uuid(), 'Course1', 'Description for Course1', now(), now(), NULL),
(gen_random_uuid(), 'Course2', 'Description for Course2', now(), now(), NULL),
(gen_random_uuid(), 'Course50', 'Description for Course50', now(), now(), NULL);


INSERT INTO Lessons (lesson_id, course_id, title, content, created_at, updated_at, deleted_at)
VALUES
('c4c4c6ab-12ac-4213-bd88-e2bd5cbaba6d','8e1e7b88-c621-43ba-b94d-08694ac7a5c8', 'Lesson1', 'Content for Lesson1', now(), now(), NULL),
('18c88b31-b4a9-4d36-894a-8854fbb089d4', '8e775104-8058-4fab-a7e5-afa761e104b1','Lesson2', 'Content for Lesson2', now(), now(), NULL),
('af1f92d3-ae44-4022-903c-8af929b83dea', '52f65263-fb1c-4075-8215-bc66ebe6c389', 'Lesson50', 'Content for Lesson50', now(), now(), NULL);


INSERT INTO Enrollments (enrollment_id, user_id, course_id, enrollment_date, created_at, updated_at, deleted_at)
VALUES
(gen_random_uuid(), 'c4c4c6ab-12ac-4213-bd88-e2bd5cbaba6d', '8e1e7b88-c621-43ba-b94d-08694ac7a5c8', now(), now(), now(), NULL),
(gen_random_uuid(),'18c88b31-b4a9-4d36-894a-8854fbb089d4' , '8e775104-8058-4fab-a7e5-afa761e104b1', now(), now(), now(), NULL),
(gen_random_uuid(), 'af1f92d3-ae44-4022-903c-8af929b83dea', '52f65263-fb1c-4075-8215-bc66ebe6c389', now(), now(), now(), NULL);
