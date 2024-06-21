CREATE TABLE IF NOT EXISTS public.book
(
    book_id character varying(255) COLLATE pg_catalog."default" NOT NULL,
    title character varying(255) COLLATE pg_catalog."default" NOT NULL,
    author character varying(255) COLLATE pg_catalog."default" NOT NULL,
    genre character varying(255) COLLATE pg_catalog."default",
    publication_date character varying(255) COLLATE pg_catalog."default",
    publisher character varying(255) COLLATE pg_catalog."default",
    isbn character varying(255) COLLATE pg_catalog."default",
    page_count integer,
    shelf_number character varying(50) COLLATE pg_catalog."default",
    language character varying(255) COLLATE pg_catalog."default",
    donor character varying(255) COLLATE pg_catalog."default",
    CONSTRAINT book_pkey PRIMARY KEY (book_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.book
    OWNER to postgres;