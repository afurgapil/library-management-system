#!/bin/bash
DATABASE_URL="postgres://postgres:test1234@localhost:5432/librarymanagementsystem_test"
psql -h localhost -p 5432 -U postgres -d librarymanagementsystem_test -w <"$DATABASE_URL" -c "CREATE TABLE IF NOT EXISTS public.book (
        book_id character varying(255) COLLATE pg_catalog.'default' NOT NULL,
        title character varying(255) COLLATE pg_catalog.'default' NOT NULL,
        author character varying(255) COLLATE pg_catalog.'default' NOT NULL,
        genre character varying(255) COLLATE pg_catalog.'default',
        publication_date character varying(255) COLLATE pg_catalog.'default',
        publisher character varying(255) COLLATE pg_catalog.'default',
        isbn character varying(255) COLLATE pg_catalog.'default',
        page_count integer,
        shelf_number character varying(50) COLLATE pg_catalog.'default',
        language character varying(255) COLLATE pg_catalog.'default',
        donor character varying(255) COLLATE pg_catalog.'default',
        CONSTRAINT book_pkey PRIMARY KEY (book_id)
    ) TABLESPACE pg_default;"

psql -h localhost -p 5432 -U postgres -d librarymanagementsystem_test -w <"$DATABASE_URL" -c "CREATE TABLE IF NOT EXISTS public.employee (
        employee_id character varying(255) COLLATE pg_catalog.'default' NOT NULL,
        employee_mail character varying(255) COLLATE pg_catalog.'default' NOT NULL,
        employee_username character varying(255) COLLATE pg_catalog.'default' NOT NULL,
        employee_phone_number character varying(20) COLLATE pg_catalog.'default' NOT NULL,
        'position' character varying(100) COLLATE pg_catalog.'default' NOT NULL,
        employee_password character varying(255) COLLATE pg_catalog.'default' NOT NULL,
        CONSTRAINT employee_pkey PRIMARY KEY (employee_id)
    ) TABLESPACE pg_default;"

psql -h localhost -p 5432 -U postgres -d librarymanagementsystem_test -w <"$DATABASE_URL" -c "CREATE TABLE IF NOT EXISTS public.student (
        student_id character varying(255) COLLATE pg_catalog.'default' NOT NULL,
        student_mail character varying(255) COLLATE pg_catalog.'default' NOT NULL,
        student_password character varying(255) COLLATE pg_catalog.'default' NOT NULL,
        debit integer NOT NULL DEFAULT 0,
        book_limit integer NOT NULL DEFAULT 10,
        is_banned boolean NOT NULL DEFAULT false,
        CONSTRAINT student_pkey PRIMARY KEY (student_id)
    ) TABLESPACE pg_default;"

psql -h localhost -p 5432 -U postgres -d librarymanagementsystem_test -w <"$DATABASE_URL" -c "CREATE TABLE IF NOT EXISTS public.book_borrow (
        borrow_id character varying(255) COLLATE pg_catalog.'default' NOT NULL,
        student_id character varying(255) COLLATE pg_catalog.'default' NOT NULL,
        book_id character varying(255) COLLATE pg_catalog.'default' NOT NULL,
        borrow_date character varying(255) COLLATE pg_catalog.'default' NOT NULL,
        delivery_date character varying(255) COLLATE pg_catalog.'default' NOT NULL,
        is_extended boolean NOT NULL DEFAULT false,
        CONSTRAINT book_borrow_pkey PRIMARY KEY (borrow_id),
        CONSTRAINT fk_book_borrow_boo FOREIGN KEY (book_id)
            REFERENCES public.book (book_id) MATCH SIMPLE
            ON UPDATE CASCADE
            ON DELETE CASCADE
            NOT VALID,
        CONSTRAINT fk_book_borrow_student FOREIGN KEY (student_id)
            REFERENCES public.student (student_id) MATCH SIMPLE
            ON UPDATE CASCADE
            ON DELETE CASCADE
            NOT VALID
    ) TABLESPACE pg_default;"