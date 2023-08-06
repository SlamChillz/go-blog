---
--- Create database
---
-- DROP DATABASE IF EXISTS blog;

SET timezone = 'Africa/Lagos';

DROP DATABASE IF EXISTS blog;

CREATE DATABASE blog
WITH
	ENCODING = 'UTF8'
	OWNER = mendy
	CONNECTION LIMIT = -1;

---use databse blog
\c blog;

---
---drop table first if it exist
---
DROP TABLE IF EXISTS posts;

---
---choice fields for table
---
DROP TYPE IF EXISTS choices;

CREATE TYPE choices as ENUM ('draft', 'published');

---
--- Create product table
---
CREATE TABLE posts (
	id SERIAL PRIMARY KEY,
	title VARCHAR(250) UNIQUE NOT NULL,
	slug VARCHAR(250) UNIQUE NOT NULL,
	body TEXT NOT NULL,
	published TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	created TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	status choices DEFAULT 'draft'
);

---
---define slug field before insert
---
DROP FUNCTION IF EXISTS generate_slug();

CREATE FUNCTION generate_slug()
RETURNS TRIGGER AS 
$$
BEGIN
	NEW.slug = REPLACE(LOWER(NEW.title), ' ', '-');
	RETURN NEW;
END;
$$ LANGUAGE PLPGSQL;

---
---trigger to execute generated_slug function
---
DROP TRIGGER IF	EXISTS generate_slug_trigger on posts CASCADE;

CREATE TRIGGER generate_slug_trigger
	BEFORE INSERT
	ON posts
	FOR EACH ROW
	EXECUTE PROCEDURE generate_slug();

---
---function to update time at post modification
---
DROP FUNCTION IF EXISTS post_update_time();

CREATE FUNCTION post_update_time()
RETURNS TRIGGER	AS
$$
BEGIN
	NEW.updated = now();
	RETURN NEW;
END;
$$ LANGUAGE PLPGSQL;

---
---trigger to execute post_update_time function
---
DROP TRIGGER IF EXISTS post_update_time_trigger on posts CASCADE;

CREATE TRIGGER post_update_time_trigger
	BEFORE UPDATE
	ON posts
	FOR EACH ROW
	EXECUTE PROCEDURE post_update_time();

---
---populate the posts table with some data
---
INSERT INTO posts (title, body) VALUES ('Software Engineering', 'Software engineering is an engineering-based approach to software development. A software engineer is a person who applies the engineering design process to design, develop, maintain, test, and evaluate computer software.');
INSERT INTO posts (title, body) VALUES ('DevOps Engineer', 'A DevOps engineer is a senior-level technology professional responsible for leading and coordinating the activities of different teams to create and maintain a companys software.');
INSERT INTO posts (title, body) VALUES ('Technical Writer', 'A technical writer is a professional information communicator whose task is to transfer information between two or more parties, through any medium that best facilitates the transfer and comprehension of the information.');
INSERT INTO posts (title, body) VALUES ('Data Analyst', 'A data analyst is a professional responsible for collecting, analyzing, and interpreting large sets of data to identify trends, patterns, and insights that can inform business decisions.');
INSERT INTO posts (title, body) VALUES ('Jazz Music', 'Jazz is a music genre that originated in the African-American communities of New Orleans, Louisiana, in the late 19th and early 20th centuries, with its roots in blues and ragtime. Since the 1920s Jazz Age, it has been recognized as a major form of musical expression in traditional.');
INSERT INTO posts (title, body) VALUES ('Education Technology', 'Educational technology (commonly abbreviated as edutech, or edtech) is the combined use of computer hardware, software, and educational theory and practice to facilitate learning.[1][2] When referred to with its abbreviation, edtech, it often refers to the industry of companies that create educational technology.');
INSERT INTO posts (title, body) VALUES ('Teacher', 'A teacher, also called a schoolteacher or formally an educator, is a person who helps students to acquire knowledge, competence, or virtue, via the practice of teaching.');
INSERT INTO posts (title, body) VALUES ('Song Writer', 'A songwriter is a musician who professionally composes musical compositions or writes lyrics for songs, or both. The writer of the music for a song can be called a composer, although this term tends to be used mainly in the classical music genre and film scoring.');