CREATE TABLE threads (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE posts (
    id UUID PRIMARY KEY,
    thread_id UUID NOT NULL REFERENCES threads (id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    votes INT NOT NULL
);

CREATE TABLE comments (
    id UUID PRIMARY KEY,
    post_id UUID NOT NULL REFERENCES posts (id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    votes INT NOT NULL
);

INSERT INTO "threads" ("id", "title", "description")
VALUES ('a9468f8b-b592-4503-9872-266e43eb3c36', 'golang', 'This is a thread about the go programming language');

INSERT INTO "threads" ("id", "title", "description")
VALUES ('4ce888b9-15be-4ab5-b6dc-836f093aaaef', 'PHP', 'This is a thread about PHP');