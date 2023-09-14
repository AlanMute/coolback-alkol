-- For Courses --

CREATE OR REPLACE PROCEDURE add_course(
    course_name VARCHAR(255),
    course_description TEXT,
    course_name_folder VARCHAR(255)
) LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO courses (name, description, name_folder)
    VALUES (course_name, course_description, course_name_folder);
END;
$$;


CREATE OR REPLACE PROCEDURE delete_course(course_id INT) LANGUAGE plpgsql AS $$
BEGIN
    DELETE FROM courses WHERE id = course_id;
END;
$$;


CREATE OR REPLACE PROCEDURE update_course(
    course_id INT,
    new_name VARCHAR(255),
    new_description TEXT,
    new_name_folder VARCHAR(255)
) LANGUAGE plpgsql AS $$
BEGIN
    UPDATE courses
    SET
        name = new_name,
        description = new_description,
        name_folder = new_name_folder
    WHERE id = course_id;
END;
$$;

-----------------------

-- For Modules --

CREATE OR REPLACE PROCEDURE add_module(
    module_name VARCHAR(255),
    module_description TEXT,
    course_id INT,
    module_name_folder VARCHAR(255)
) LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO modules (name, description, courses_ID, name_folder)
    VALUES (module_name, module_description, course_id, module_name_folder);
END;
$$;


CREATE OR REPLACE PROCEDURE delete_module(module_id INT) LANGUAGE plpgsql AS $$
BEGIN
    DELETE FROM modules WHERE id = module_id;
END;
$$;


CREATE OR REPLACE PROCEDURE update_module(
    module_id INT,
    new_name VARCHAR(255),
    new_description TEXT,
    new_name_folder VARCHAR(255)
) LANGUAGE plpgsql AS $$
BEGIN
    UPDATE modules
    SET
        name = new_name,
        description = new_description,
        name_folder = new_name_folder
    WHERE id = module_id;
END;
$$;

--------------------------

-- For lessons --


CREATE OR REPLACE PROCEDURE add_lesson(
    lesson_name VARCHAR(255),
    lesson_description TEXT,
    module_id INT,
    lesson_name_folder VARCHAR(255)
) LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO lessons (name, description, modules_ID, name_folder)
    VALUES (lesson_name, lesson_description, module_id, lesson_name_folder);
END;
$$;


CREATE OR REPLACE PROCEDURE delete_lesson(lesson_id INT) LANGUAGE plpgsql AS $$
BEGIN
    DELETE FROM lessons WHERE id = lesson_id;
END;
$$;


CREATE OR REPLACE PROCEDURE update_lesson(
    lesson_id INT,
    new_name VARCHAR(255),
    new_description TEXT,
    new_name_folder VARCHAR(255)
) LANGUAGE plpgsql AS $$
BEGIN
    UPDATE lessons
    SET
        name = new_name,
        description = new_description,
        name_folder = new_name_folder
    WHERE id = lesson_id;
END;
$$;
