--
-- PostgreSQL database dump
--

-- Dumped from database version 16.1 (Ubuntu 16.1-1.pgdg22.04+1)
-- Dumped by pg_dump version 16.1 (Ubuntu 16.1-1.pgdg22.04+1)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: liked; Type: TABLE; Schema: public; Owner: reddit
--

CREATE TABLE public.liked (
    post_id integer NOT NULL,
    liked boolean,
    no_of_likes integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.liked OWNER TO reddit;

--
-- Name: liked_post_id_seq; Type: SEQUENCE; Schema: public; Owner: reddit
--

CREATE SEQUENCE public.liked_post_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.liked_post_id_seq OWNER TO reddit;

--
-- Name: liked_post_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: reddit
--

ALTER SEQUENCE public.liked_post_id_seq OWNED BY public.liked.post_id;


--
-- Name: post; Type: TABLE; Schema: public; Owner: reddit
--

CREATE TABLE public.post (
    post_id integer NOT NULL,
    title character varying(255) NOT NULL,
    body text NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    subreddit character varying(255) NOT NULL,
    username character varying(255) NOT NULL,
    image_path character varying(255),
    video_path character varying(255)
);


ALTER TABLE public.post OWNER TO reddit;

--
-- Name: post_post_id_seq; Type: SEQUENCE; Schema: public; Owner: reddit
--

CREATE SEQUENCE public.post_post_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.post_post_id_seq OWNER TO reddit;

--
-- Name: post_post_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: reddit
--

ALTER SEQUENCE public.post_post_id_seq OWNED BY public.post.post_id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: reddit
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO reddit;

--
-- Name: user; Type: TABLE; Schema: public; Owner: reddit
--

CREATE TABLE public."user" (
    user_id integer NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(60),
    username character varying(255) NOT NULL
);


ALTER TABLE public."user" OWNER TO reddit;

--
-- Name: user_user_id_seq; Type: SEQUENCE; Schema: public; Owner: reddit
--

CREATE SEQUENCE public.user_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.user_user_id_seq OWNER TO reddit;

--
-- Name: user_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: reddit
--

ALTER SEQUENCE public.user_user_id_seq OWNED BY public."user".user_id;


--
-- Name: liked post_id; Type: DEFAULT; Schema: public; Owner: reddit
--

ALTER TABLE ONLY public.liked ALTER COLUMN post_id SET DEFAULT nextval('public.liked_post_id_seq'::regclass);


--
-- Name: post post_id; Type: DEFAULT; Schema: public; Owner: reddit
--

ALTER TABLE ONLY public.post ALTER COLUMN post_id SET DEFAULT nextval('public.post_post_id_seq'::regclass);


--
-- Name: user user_id; Type: DEFAULT; Schema: public; Owner: reddit
--

ALTER TABLE ONLY public."user" ALTER COLUMN user_id SET DEFAULT nextval('public.user_user_id_seq'::regclass);


--
-- Name: liked liked_pkey; Type: CONSTRAINT; Schema: public; Owner: reddit
--

ALTER TABLE ONLY public.liked
    ADD CONSTRAINT liked_pkey PRIMARY KEY (post_id);


--
-- Name: post post_pkey; Type: CONSTRAINT; Schema: public; Owner: reddit
--

ALTER TABLE ONLY public.post
    ADD CONSTRAINT post_pkey PRIMARY KEY (post_id);


--
-- Name: schema_migration schema_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: reddit
--

ALTER TABLE ONLY public.schema_migration
    ADD CONSTRAINT schema_migration_pkey PRIMARY KEY (version);


--
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: reddit
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (user_id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: reddit
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: user_username_idx; Type: INDEX; Schema: public; Owner: reddit
--

CREATE UNIQUE INDEX user_username_idx ON public."user" USING btree (username);


--
-- Name: post username; Type: FK CONSTRAINT; Schema: public; Owner: reddit
--

ALTER TABLE ONLY public.post
    ADD CONSTRAINT username FOREIGN KEY (username) REFERENCES public."user"(username) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: pg_database_owner
--

GRANT ALL ON SCHEMA public TO reddit;
GRANT ALL ON SCHEMA public TO tsawlergo;


--
-- Name: TABLE schema_migration; Type: ACL; Schema: public; Owner: reddit
--

GRANT ALL ON TABLE public.schema_migration TO tsawlergo;


--
-- Name: TABLE "user"; Type: ACL; Schema: public; Owner: reddit
--

GRANT ALL ON TABLE public."user" TO tsawlergo;


--
-- PostgreSQL database dump complete
--

