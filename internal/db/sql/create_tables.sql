CREATE TABLE courses (
	id serial NOT NULL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	description TEXT,
	name_folder VARCHAR(255)
);


CREATE TABLE modules (
	id serial NOT NULL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	description TEXT,
	courses_ID integer NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
	name_folder VARCHAR(255)
);


CREATE TABLE lessons (
	id serial NOT NULL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	description TEXT,
	modules_ID integer NOT NULL REFERENCES modules(id) ON DELETE CASCADE,
	name_folder VARCHAR(255)
);