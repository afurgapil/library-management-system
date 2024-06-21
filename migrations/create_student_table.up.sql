CREATE TABLE IF NOT EXISTS public.student
(
    student_id character varying(255) COLLATE pg_catalog."default" NOT NULL,
    student_mail character varying(255) COLLATE pg_catalog."default" NOT NULL,
    student_password character varying(255) COLLATE pg_catalog."default" NOT NULL,
    debit integer NOT NULL DEFAULT 0,
    book_limit integer NOT NULL DEFAULT 10,
    is_banned boolean NOT NULL DEFAULT false,
    CONSTRAINT student_pkey PRIMARY KEY (student_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.student
    OWNER to postgres;