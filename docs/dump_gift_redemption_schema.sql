--
-- PostgreSQL database dump
--

-- Dumped from database version 14.21
-- Dumped by pg_dump version 14.21

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

ALTER TABLE IF EXISTS ONLY public.redemptions DROP CONSTRAINT IF EXISTS redemptions_user_id_fkey;
ALTER TABLE IF EXISTS ONLY public.redemptions DROP CONSTRAINT IF EXISTS redemptions_gift_id_fkey;
ALTER TABLE IF EXISTS ONLY public.ratings DROP CONSTRAINT IF EXISTS ratings_user_id_fkey;
ALTER TABLE IF EXISTS ONLY public.ratings DROP CONSTRAINT IF EXISTS ratings_redemption_id_fkey;
ALTER TABLE IF EXISTS ONLY public.ratings DROP CONSTRAINT IF EXISTS ratings_gift_id_fkey;
DROP INDEX IF EXISTS public.idx_users_email;
DROP INDEX IF EXISTS public.idx_users_deleted_at;
DROP INDEX IF EXISTS public.idx_redemptions_user_id;
DROP INDEX IF EXISTS public.idx_redemptions_gift_id;
DROP INDEX IF EXISTS public.idx_ratings_user_id;
DROP INDEX IF EXISTS public.idx_ratings_gift_id;
DROP INDEX IF EXISTS public.idx_gifts_deleted_at;
DROP INDEX IF EXISTS public.idx_gifts_created_at;
DROP INDEX IF EXISTS public.idx_gifts_avg_rating;
ALTER TABLE IF EXISTS ONLY public.users DROP CONSTRAINT IF EXISTS users_pkey;
ALTER TABLE IF EXISTS ONLY public.users DROP CONSTRAINT IF EXISTS users_email_key;
ALTER TABLE IF EXISTS ONLY public.ratings DROP CONSTRAINT IF EXISTS uq_rating_redemption;
ALTER TABLE IF EXISTS ONLY public.schema_migrations DROP CONSTRAINT IF EXISTS schema_migrations_pkey;
ALTER TABLE IF EXISTS ONLY public.redemptions DROP CONSTRAINT IF EXISTS redemptions_pkey;
ALTER TABLE IF EXISTS ONLY public.ratings DROP CONSTRAINT IF EXISTS ratings_pkey;
ALTER TABLE IF EXISTS ONLY public.gifts DROP CONSTRAINT IF EXISTS gifts_pkey;
ALTER TABLE IF EXISTS public.users ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.redemptions ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.ratings ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.gifts ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE IF EXISTS public.users_id_seq;
DROP TABLE IF EXISTS public.users;
DROP TABLE IF EXISTS public.schema_migrations;
DROP SEQUENCE IF EXISTS public.redemptions_id_seq;
DROP TABLE IF EXISTS public.redemptions;
DROP SEQUENCE IF EXISTS public.ratings_id_seq;
DROP TABLE IF EXISTS public.ratings;
DROP SEQUENCE IF EXISTS public.gifts_id_seq;
DROP TABLE IF EXISTS public.gifts;
SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: gifts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.gifts (
    id integer NOT NULL,
    name character varying(200) NOT NULL,
    description text,
    point integer DEFAULT 0 NOT NULL,
    stock integer DEFAULT 0 NOT NULL,
    image_url character varying(500),
    is_new boolean DEFAULT false NOT NULL,
    is_best_seller boolean DEFAULT false NOT NULL,
    avg_rating numeric(3,2) DEFAULT 0 NOT NULL,
    total_reviews integer DEFAULT 0 NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone
);


--
-- Name: gifts_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.gifts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: gifts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.gifts_id_seq OWNED BY public.gifts.id;


--
-- Name: ratings; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.ratings (
    id integer NOT NULL,
    user_id integer NOT NULL,
    gift_id integer NOT NULL,
    redemption_id integer NOT NULL,
    score numeric(2,1) NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    CONSTRAINT ratings_score_check CHECK (((score >= (1)::numeric) AND (score <= (5)::numeric)))
);


--
-- Name: ratings_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.ratings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: ratings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.ratings_id_seq OWNED BY public.ratings.id;


--
-- Name: redemptions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.redemptions (
    id integer NOT NULL,
    user_id integer NOT NULL,
    gift_id integer NOT NULL,
    quantity integer DEFAULT 1 NOT NULL,
    total_point integer NOT NULL,
    redeemed_at timestamp with time zone DEFAULT now() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


--
-- Name: redemptions_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.redemptions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: redemptions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.redemptions_id_seq OWNED BY public.redemptions.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    email character varying(150) NOT NULL,
    password character varying(255) NOT NULL,
    role character varying(10) DEFAULT 'user'::character varying NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone
);


--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: gifts id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.gifts ALTER COLUMN id SET DEFAULT nextval('public.gifts_id_seq'::regclass);


--
-- Name: ratings id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.ratings ALTER COLUMN id SET DEFAULT nextval('public.ratings_id_seq'::regclass);


--
-- Name: redemptions id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.redemptions ALTER COLUMN id SET DEFAULT nextval('public.redemptions_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: gifts; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.gifts (id, name, description, point, stock, image_url, is_new, is_best_seller, avg_rating, total_reviews, created_at, updated_at, deleted_at) VALUES (3, 'Xiaomi Redmi Note 12 Pro', '6.67 inch AMOLED, 200MP Camera, 5000mAh battery, 67W HyperCharge', 150000, 15, 'https://via.placeholder.com/400x400?text=Redmi+Note+12', false, true, 3.80, 95, '2026-02-17 20:51:41.870161+00', '2026-02-17 20:51:41.870161+00', NULL);
INSERT INTO public.gifts (id, name, description, point, stock, image_url, is_new, is_best_seller, avg_rating, total_reviews, created_at, updated_at, deleted_at) VALUES (4, 'JBL Flip 6 Bluetooth Speaker', 'Portable waterproof speaker, 12 hours battery life, PartyBoost compatible', 80000, 20, 'https://via.placeholder.com/400x400?text=JBL+Flip+6', false, false, 4.10, 210, '2026-02-17 20:51:41.870161+00', '2026-02-17 20:51:41.870161+00', NULL);
INSERT INTO public.gifts (id, name, description, point, stock, image_url, is_new, is_best_seller, avg_rating, total_reviews, created_at, updated_at, deleted_at) VALUES (5, 'Sony WH-1000XM5 Headphones', 'Industry-leading noise canceling, 30-hour battery, multipoint connection', 420000, 0, 'https://via.placeholder.com/400x400?text=Sony+WH1000XM5', false, false, 4.90, 540, '2026-02-17 20:51:41.870161+00', '2026-02-17 20:51:41.870161+00', NULL);
INSERT INTO public.gifts (id, name, description, point, stock, image_url, is_new, is_best_seller, avg_rating, total_reviews, created_at, updated_at, deleted_at) VALUES (2, 'Apple AirPods Pro 2nd Gen', 'Active Noise Cancellation, Adaptive Transparency, Personalized Spatial Audio with dynamic head tracking', 350000, 2, 'https://via.placeholder.com/400x400?text=AirPods+Pro', false, true, 4.70, 320, '2026-02-17 20:51:41.870161+00', '2026-02-17 20:56:48.308289+00', NULL);
INSERT INTO public.gifts (id, name, description, point, stock, image_url, is_new, is_best_seller, avg_rating, total_reviews, created_at, updated_at, deleted_at) VALUES (7, 'Mechanical Keyboard Keychron K2', 'TKL wireless mechanical keyboard, hot-swappable, RGB backlight', 180000, 8, 'https://via.placeholder.com/400x400?text=Keychron+K2', true, false, 0.00, 0, '2026-02-17 20:57:57.822985+00', '2026-02-17 20:57:57.822985+00', NULL);
INSERT INTO public.gifts (id, name, description, point, stock, image_url, is_new, is_best_seller, avg_rating, total_reviews, created_at, updated_at, deleted_at) VALUES (8, 'Mechanical Keyboard Keychron K2', 'TKL wireless mechanical keyboard, hot-swappable, RGB backlight', 180000, 8, 'https://via.placeholder.com/400x400?text=Keychron+K2', true, false, 0.00, 0, '2026-02-17 20:58:40.469246+00', '2026-02-17 20:58:40.469246+00', NULL);
INSERT INTO public.gifts (id, name, description, point, stock, image_url, is_new, is_best_seller, avg_rating, total_reviews, created_at, updated_at, deleted_at) VALUES (1, 'Samsung Galaxy S9 Midnight Black 4/64 GB', 'Updated description - Ukuran layar 6.2 inci, Super AMOLED', 210000, 15, 'https://via.placeholder.com/400x400?text=Galaxy+S9', true, true, 3.70, 2, '2026-02-17 20:51:41.870161+00', '2026-02-17 20:59:15.923472+00', NULL);
INSERT INTO public.gifts (id, name, description, point, stock, image_url, is_new, is_best_seller, avg_rating, total_reviews, created_at, updated_at, deleted_at) VALUES (6, 'Logitech MX Master 3S Mouse', '8K DPI sensor, MagSpeed scroll wheel, quiet clicks, USB-C charging', 120000, 8, 'https://via.placeholder.com/400x400?text=MX+Master+3S', true, false, 4.60, 185, '2026-02-17 20:51:41.870161+00', '2026-02-17 20:51:41.870161+00', '2026-02-17 20:59:29.680161+00');


--
-- Data for Name: ratings; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.ratings (id, user_id, gift_id, redemption_id, score, created_at, updated_at) VALUES (1, 2, 1, 1, 3.6, '2026-02-17 20:52:59.623286+00', '2026-02-17 20:52:59.623286+00');
INSERT INTO public.ratings (id, user_id, gift_id, redemption_id, score, created_at, updated_at) VALUES (2, 4, 1, 2, 3.8, '2026-02-17 20:54:40.264597+00', '2026-02-17 20:54:40.264597+00');


--
-- Data for Name: redemptions; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.redemptions (id, user_id, gift_id, quantity, total_point, redeemed_at, created_at) VALUES (1, 2, 1, 1, 200000, '2026-02-17 20:52:54.340763+00', '2026-02-17 20:52:54.340944+00');
INSERT INTO public.redemptions (id, user_id, gift_id, quantity, total_point, redeemed_at, created_at) VALUES (2, 4, 1, 1, 200000, '2026-02-17 20:54:37.291731+00', '2026-02-17 20:54:37.291755+00');
INSERT INTO public.redemptions (id, user_id, gift_id, quantity, total_point, redeemed_at, created_at) VALUES (3, 1, 2, 3, 1050000, '2026-02-17 20:56:48.309369+00', '2026-02-17 20:56:48.309406+00');


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.schema_migrations (version, dirty) VALUES (3, false);


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.users (id, name, email, password, role, created_at, updated_at, deleted_at) VALUES (1, 'Admin gift-redemption', 'admin@gift-redemption.com', '$2a$10$NblMkRiUEwlZ.s.TZKmBsOsR/NOaySvlXpGeq/0naqNrNt19dcC46', 'admin', '2026-02-17 20:51:41.858925+00', '2026-02-17 20:51:41.858925+00', NULL);
INSERT INTO public.users (id, name, email, password, role, created_at, updated_at, deleted_at) VALUES (4, 'Jane Doe', 'jane@example.com', '$2a$10$Huoyl71bZE2VzAW9GoWXaOwSWCznKK/RXeILxKL551XJmUFP3D0.q', 'user', '2026-02-17 20:54:14.063494+00', '2026-02-17 20:54:14.063494+00', NULL);
INSERT INTO public.users (id, name, email, password, role, created_at, updated_at, deleted_at) VALUES (5, 'Janet Doe', 'janet@example.com', '$2a$10$IDQvhVsWXkl34a/l.gVV4ecD0XpHXnegX9neiE7JOSS9ndjqdzQli', 'user', '2026-02-17 20:55:37.397691+00', '2026-02-17 20:55:37.397691+00', NULL);
INSERT INTO public.users (id, name, email, password, role, created_at, updated_at, deleted_at) VALUES (2, 'John Doe Updated', 'john@example.com', '$2a$10$X/.DcRmDctyhzW0MXVrvFO4BDSAhU5JJGn8lM8pVDx5tsnJwG/ybO', 'user', '2026-02-17 20:51:41.858925+00', '2026-02-17 20:55:26.899497+00', '2026-02-17 20:55:58.156227+00');


--
-- Name: gifts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.gifts_id_seq', 8, true);


--
-- Name: ratings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.ratings_id_seq', 2, true);


--
-- Name: redemptions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.redemptions_id_seq', 3, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.users_id_seq', 5, true);


--
-- Name: gifts gifts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.gifts
    ADD CONSTRAINT gifts_pkey PRIMARY KEY (id);


--
-- Name: ratings ratings_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.ratings
    ADD CONSTRAINT ratings_pkey PRIMARY KEY (id);


--
-- Name: redemptions redemptions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.redemptions
    ADD CONSTRAINT redemptions_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: ratings uq_rating_redemption; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.ratings
    ADD CONSTRAINT uq_rating_redemption UNIQUE (redemption_id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_gifts_avg_rating; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_gifts_avg_rating ON public.gifts USING btree (avg_rating) WHERE (deleted_at IS NULL);


--
-- Name: idx_gifts_created_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_gifts_created_at ON public.gifts USING btree (created_at) WHERE (deleted_at IS NULL);


--
-- Name: idx_gifts_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_gifts_deleted_at ON public.gifts USING btree (deleted_at);


--
-- Name: idx_ratings_gift_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_ratings_gift_id ON public.ratings USING btree (gift_id);


--
-- Name: idx_ratings_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_ratings_user_id ON public.ratings USING btree (user_id);


--
-- Name: idx_redemptions_gift_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_redemptions_gift_id ON public.redemptions USING btree (gift_id);


--
-- Name: idx_redemptions_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_redemptions_user_id ON public.redemptions USING btree (user_id);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- Name: idx_users_email; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_users_email ON public.users USING btree (email) WHERE (deleted_at IS NULL);


--
-- Name: ratings ratings_gift_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.ratings
    ADD CONSTRAINT ratings_gift_id_fkey FOREIGN KEY (gift_id) REFERENCES public.gifts(id);


--
-- Name: ratings ratings_redemption_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.ratings
    ADD CONSTRAINT ratings_redemption_id_fkey FOREIGN KEY (redemption_id) REFERENCES public.redemptions(id);


--
-- Name: ratings ratings_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.ratings
    ADD CONSTRAINT ratings_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: redemptions redemptions_gift_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.redemptions
    ADD CONSTRAINT redemptions_gift_id_fkey FOREIGN KEY (gift_id) REFERENCES public.gifts(id);


--
-- Name: redemptions redemptions_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.redemptions
    ADD CONSTRAINT redemptions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--


