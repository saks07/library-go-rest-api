CREATE TABLE public.borrowed_books
(
    id SERIAL NOT NULL,
    book_id integer NOT NULL,
    user_id integer NOT NULL,
    borrow_date timestamp without time zone NOT NULL,
    return_date timestamp without time zone,
    CONSTRAINT borrowed_books_pkey PRIMARY KEY (id),
    CONSTRAINT borrowed_books_unique_user_id_book_id UNIQUE (user_id, book_id),
    CONSTRAINT borrowed_books_fkey_book_id FOREIGN KEY (book_id)
        REFERENCES public.books (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT borrowed_books_fkey_user_id FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);

INSERT INTO public.borrowed_books(book_id, user_id, borrow_date, return_date) 
  VALUES 
    ('5', '1', '2024-10-08 08:45:00', '2024-10-18 10:45:00'),
    ('7', '1', '2024-10-08 08:45:00', NULL),
    ('3', '2', '2024-10-05 08:45:00', NULL),
    ('2', '2', '2024-10-05 08:45:00', '2024-10-15 12:00:00'),
    ('5', '2', '2024-10-08 09:30:00', '2024-10-15 12:00:00'),
    ('7', '2', '2024-10-08 09:30:00', NULL),
    ('5', '3', '2024-10-12 13:05:00', NULL),
    ('6', '3', '2024-10-12 13:05:00', NULL),
    ('9', '3', '2024-10-12 13:05:00', NULL),
    ('5', '5', '2024-10-02 11:15:00', '2024-10-25 07:15:00'),
    ('3', '5', '2024-10-02 11:15:00', '2024-10-22 16:45:00'),
    ('1', '5', '2024-10-02 11:15:00', NULL),
    ('2', '5', '2024-10-09 15:30:00', '2024-10-25 07:15:00');