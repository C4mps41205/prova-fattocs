DROP TABLE IF EXISTS public.tasks;
CREATE TABLE public.tasks
(
    id                 SERIAL PRIMARY KEY,
    name               VARCHAR(255) UNIQUE NOT NULL,
    cost               NUMERIC(10, 2)      NOT NULL,
    deadline           DATE                NOT NULL,
    presentation_order INTEGER UNIQUE      NOT NULL
);
