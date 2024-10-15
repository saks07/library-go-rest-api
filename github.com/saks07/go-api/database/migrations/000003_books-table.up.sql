CREATE TABLE public.books
(
    id SERIAL NOT NULL,
    book_title character varying(255) NOT NULL,
    book_available_copies integer,
    CONSTRAINT books_pkey PRIMARY KEY (id)
);

INSERT INTO public.books(book_title, book_available_copies) 
  VALUES 
    ('Book title 1', 5),
    ('Book title 2', 2),
    ('Book title 3', 3),
    ('Book title 4', 5),
    ('Book title 5', 6),
    ('Book title 6', 7),
    ('Book title 7', 9),
    ('Book title 8', 10),
    ('Book title 9', 8),
    ('Book title 10', 4);