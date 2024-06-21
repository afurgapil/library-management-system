CREATE TABLE IF NOT EXISTS public.employee
(
    employee_id character varying(255) COLLATE pg_catalog."default" NOT NULL,
    employee_mail character varying(255) COLLATE pg_catalog."default" NOT NULL,
    employee_username character varying(255) COLLATE pg_catalog."default" NOT NULL,
    employee_phone_number character varying(20) COLLATE pg_catalog."default" NOT NULL,
    "position" character varying(100) COLLATE pg_catalog."default" NOT NULL,
    employee_password character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT employee_pkey PRIMARY KEY (employee_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.employee
    OWNER to postgres;