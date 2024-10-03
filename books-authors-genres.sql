--
-- Name: books; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.books (
    id integer NOT NULL,
    title character varying(512),
    author_id integer,
    publication_year integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    slug character varying(512),
    description text
);



--
-- Name: books books_author_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_author_id_fkey FOREIGN KEY (author_id) REFERENCES public.authors(id) ON UPDATE CASCADE ON DELETE CASCADE;

