CREATE TABLE public.users
(
    id SERIAL NOT NULL,
    username character varying(50) NOT NULL,
    email character varying(50) NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_email_unique UNIQUE (email),
    CONSTRAINT users_username_unique UNIQUE (username)
);

INSERT INTO public.users(username, email) 
  VALUES 
    ('user1', 'user1@mail.com'),
    ('user2', 'user2@mail.com'),
    ('user3', 'user3@mail.com'),
    ('user4', 'user4@mail.com'),
    ('user5', 'user5@mail.com');