CREATE TABLE IF NOT EXISTS public.book_borrow
(
    borrow_id character varying(255) COLLATE pg_catalog."default" NOT NULL,
    student_id character varying(255) COLLATE pg_catalog."default" NOT NULL,
    book_id character varying(255) COLLATE pg_catalog."default" NOT NULL,
    borrow_date character varying(255) COLLATE pg_catalog."default" NOT NULL,
    delivery_date character varying(255) COLLATE pg_catalog."default" NOT NULL,
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
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.book_borrow
    OWNER to postgres;