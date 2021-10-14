--
-- PostgreSQL database dump
--

-- Dumped from database version 12.3 (Debian 12.3-1.pgdg100+1)
-- Dumped by pg_dump version 12.7 (Ubuntu 12.7-0ubuntu0.20.10.1)

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
-- Name: cache; Type: TABLE; Schema: public; Owner: cachet
--

CREATE TABLE public.cache (
    key character varying(255) NOT NULL,
    value text NOT NULL,
    expiration integer NOT NULL
);


ALTER TABLE public.cache OWNER TO cachet;

--
-- Name: component_groups; Type: TABLE; Schema: public; Owner: cachet
--

CREATE TABLE public.component_groups (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    "order" integer DEFAULT 0 NOT NULL,
    collapsed integer DEFAULT 1 NOT NULL
);


ALTER TABLE public.component_groups OWNER TO cachet;

--
-- Name: component_groups_id_seq; Type: SEQUENCE; Schema: public; Owner: cachet
--

CREATE SEQUENCE public.component_groups_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.component_groups_id_seq OWNER TO cachet;

--
-- Name: component_groups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cachet
--

ALTER SEQUENCE public.component_groups_id_seq OWNED BY public.component_groups.id;


--
-- Name: component_tag; Type: TABLE; Schema: public; Owner: cachet
--

CREATE TABLE public.component_tag (
    id integer NOT NULL,
    component_id integer NOT NULL,
    tag_id integer NOT NULL
);


ALTER TABLE public.component_tag OWNER TO cachet;

--
-- Name: component_tag_id_seq; Type: SEQUENCE; Schema: public; Owner: cachet
--

CREATE SEQUENCE public.component_tag_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.component_tag_id_seq OWNER TO cachet;

--
-- Name: component_tag_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cachet
--

ALTER SEQUENCE public.component_tag_id_seq OWNED BY public.component_tag.id;


--
-- Name: components; Type: TABLE; Schema: public; Owner: cachet
--

CREATE TABLE public.components (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description text NOT NULL,
    link text NOT NULL,
    status integer NOT NULL,
    "order" integer NOT NULL,
    group_id integer NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone,
    enabled boolean DEFAULT true NOT NULL
);


ALTER TABLE public.components OWNER TO cachet;

--
-- Name: components_id_seq; Type: SEQUENCE; Schema: public; Owner: cachet
--

CREATE SEQUENCE public.components_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.components_id_seq OWNER TO cachet;

--
-- Name: components_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cachet
--

ALTER SEQUENCE public.components_id_seq OWNED BY public.components.id;


--
-- Name: failed_jobs; Type: TABLE; Schema: public; Owner: cachet
--

CREATE TABLE public.failed_jobs (
    id integer NOT NULL,
    connection text NOT NULL,
    queue text NOT NULL,
    payload text NOT NULL,
    failed_at timestamp(0) without time zone NOT NULL
);


ALTER TABLE public.failed_jobs OWNER TO cachet;

--
-- Name: failed_jobs_id_seq; Type: SEQUENCE; Schema: public; Owner: cachet
--

CREATE SEQUENCE public.failed_jobs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.failed_jobs_id_seq OWNER TO cachet;

--
-- Name: failed_jobs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cachet
--

ALTER SEQUENCE public.failed_jobs_id_seq OWNED BY public.failed_jobs.id;


--
-- Name: incident_templates; Type: TABLE; Schema: public; Owner: cachet
--

CREATE TABLE public.incident_templates (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    slug character varying(255) NOT NULL,
    template text NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone
);


ALTER TABLE public.incident_templates OWNER TO cachet;

--
-- Name: incident_templates_id_seq; Type: SEQUENCE; Schema: public; Owner: cachet
--

CREATE SEQUENCE public.incident_templates_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.incident_templates_id_seq OWNER TO cachet;

--
-- Name: incident_templates_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cachet
--

ALTER SEQUENCE public.incident_templates_id_seq OWNED BY public.incident_templates.id;


--
-- Name: incidents; Type: TABLE; Schema: public; Owner: cachet
--

CREATE TABLE public.incidents (
    id integer NOT NULL,
    component_id integer DEFAULT 0 NOT NULL,
    name character varying(255) NOT NULL,
    status integer NOT NULL,
    message text NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone,
    scheduled_at timestamp(0) without time zone,
    visible boolean DEFAULT true NOT NULL
);


ALTER TABLE public.incidents OWNER TO cachet;

--
-- Name: incidents_id_seq; Type: SEQUENCE; Schema: public; Owner: cachet
--

CREATE SEQUENCE public.incidents_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.incidents_id_seq OWNER TO cachet;

--
-- Name: incidents_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cachet
--

ALTER SEQUENCE public.incidents_id_seq OWNED BY public.incidents.id;


--
-- Name: invites; Type: TABLE; Schema: public; Owner: cachet
--

CREATE TABLE public.invites (
    id integer NOT NULL,
    code character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    claimed_at timestamp(0) without time zone,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


ALTER TABLE public.invites OWNER TO cachet;

--
-- Name: invites_id_seq; Type: SEQUENCE; Schema: public; Owner: cachet
--

CREATE SEQUENCE public.invites_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.invites_id_seq OWNER TO cachet;

--
-- Name: invites_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cachet
--

ALTER SEQUENCE public.invites_id_seq OWNED BY public.invites.id;


--
-- Name: jobs; Type: TABLE; Schema: public; Owner: cachet
--

CREATE TABLE public.jobs (
    id bigint NOT NULL,
    queue character varying(255) NOT NULL,
    payload text NOT NULL,
    attempts smallint NOT NULL,
    reserved smallint NOT NULL,
    reserved_at integer,
    available_at integer NOT NULL,
    created_at integer NOT NULL
);


ALTER TABLE public.jobs OWNER TO cachet;

--
-- Name: jobs_id_seq; Type: SEQUENCE; Schema: public; Owner: cachet
--

CREATE SEQUENCE public.jobs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.jobs_id_seq OWNER TO cachet;

--
-- Name: jobs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cachet
--

ALTER SEQUENCE public.jobs_id_seq OWNED BY public.jobs.id;


--
-- Name: metric_points; Type: TABLE; Schema: public; Owner: cachet
--

CREATE TABLE public.metric_points (
    id integer NOT NULL,
    metric_id integer NOT NULL,
    value numeric(15,3) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    counter integer DEFAULT 1 NOT NULL
);


ALTER TABLE public.metric_points OWNER TO cachet;

--
-- Name: metric_points_id_seq; Type: SEQUENCE; Schema: public; Owner: cachet
--

CREATE SEQUENCE public.metric_points_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.metric_points_id_seq OWNER TO cachet;

--
-- Name: metric_points_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cachet
--

ALTER SEQUENCE public.metric_points_id_seq OWNED BY public.metric_points.id;


--
-- Name: metrics; Type: TABLE; Schema: public; Owner: cachet
--

CREATE TABLE public.metrics (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    suffix character varying(255) NOT NULL,
    description text NOT NULL,
    default_value numeric(10,3) NOT NULL,
    calc_type smallint NOT NULL,
    display_chart boolean DEFAULT true NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    places integer DEFAULT 2 NOT NULL,
    default_view smallint DEFAULT '1'::smallint NOT NULL,
    threshold integer DEFAULT 5 NOT NULL,
    "order" smallint DEFAULT '0'::smallint NOT NULL
);


ALTER TABLE public.metrics OWNER TO cachet;

--
-- Name: metrics_id_seq; Type: SEQUENCE; Schema: public; Owner: cachet
--

CREATE SEQUENCE public.metrics_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.metrics_id_seq OWNER TO cachet;

--
-- Name: metrics_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cachet
--

ALTER SEQUENCE public.metrics_id_seq OWNED BY public.metrics.id;


--
-- Name: migrations; Type: TABLE; Schema: public; Owner: cachet
--

CREATE TABLE public.migrations (
    migration character varying(255) NOT NULL,
    batch integer NOT NULL
);


ALTER TABLE public.migrations OWNER TO cachet;

--
-- Name: sessions; Type: TABLE; Schema: public; Owner: cachet
--

CREATE TABLE public.sessions (
    id character varying(255) NOT NULL,
    user_id integer,
    ip_address character varying(45),
    user_agent text,
    payload text NOT NULL,
    last_activity integer NOT NULL
);


ALTER TABLE public.sessions OWNER TO cachet;

--
-- Name: settings; Type: TABLE; Schema: public; Owner: cachet
--

CREATE TABLE public.settings (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    value text NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


ALTER TABLE public.settings OWNER TO cachet;

--
-- Name: settings_id_seq; Type: SEQUENCE; Schema: public; Owner: cachet
--

CREATE SEQUENCE public.settings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.settings_id_seq OWNER TO cachet;

--
-- Name: settings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cachet
--

ALTER SEQUENCE public.settings_id_seq OWNED BY public.settings.id;


--
-- Name: subscribers; Type: TABLE; Schema: public; Owner: cachet
--

CREATE TABLE public.subscribers (
    id integer NOT NULL,
    email character varying(255) NOT NULL,
    verify_code character varying(255) NOT NULL,
    verified_at timestamp(0) without time zone,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    global boolean DEFAULT true NOT NULL
);


ALTER TABLE public.subscribers OWNER TO cachet;

--
-- Name: subscribers_id_seq; Type: SEQUENCE; Schema: public; Owner: cachet
--

CREATE SEQUENCE public.subscribers_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.subscribers_id_seq OWNER TO cachet;

--
-- Name: subscribers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cachet
--

ALTER SEQUENCE public.subscribers_id_seq OWNED BY public.subscribers.id;


--
-- Name: subscriptions; Type: TABLE; Schema: public; Owner: cachet
--

CREATE TABLE public.subscriptions (
    id integer NOT NULL,
    subscriber_id integer NOT NULL,
    component_id integer NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


ALTER TABLE public.subscriptions OWNER TO cachet;

--
-- Name: subscriptions_id_seq; Type: SEQUENCE; Schema: public; Owner: cachet
--

CREATE SEQUENCE public.subscriptions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.subscriptions_id_seq OWNER TO cachet;

--
-- Name: subscriptions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cachet
--

ALTER SEQUENCE public.subscriptions_id_seq OWNED BY public.subscriptions.id;


--
-- Name: tags; Type: TABLE; Schema: public; Owner: cachet
--

CREATE TABLE public.tags (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    slug character varying(255) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


ALTER TABLE public.tags OWNER TO cachet;

--
-- Name: tags_id_seq; Type: SEQUENCE; Schema: public; Owner: cachet
--

CREATE SEQUENCE public.tags_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tags_id_seq OWNER TO cachet;

--
-- Name: tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cachet
--

ALTER SEQUENCE public.tags_id_seq OWNED BY public.tags.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: cachet
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    remember_token character varying(100),
    email character varying(255) NOT NULL,
    api_key character varying(255) NOT NULL,
    active boolean DEFAULT true NOT NULL,
    level smallint DEFAULT '2'::smallint NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    google_2fa_secret character varying(255)
);


ALTER TABLE public.users OWNER TO cachet;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: cachet
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO cachet;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cachet
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: component_groups id; Type: DEFAULT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.component_groups ALTER COLUMN id SET DEFAULT nextval('public.component_groups_id_seq'::regclass);


--
-- Name: component_tag id; Type: DEFAULT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.component_tag ALTER COLUMN id SET DEFAULT nextval('public.component_tag_id_seq'::regclass);


--
-- Name: components id; Type: DEFAULT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.components ALTER COLUMN id SET DEFAULT nextval('public.components_id_seq'::regclass);


--
-- Name: failed_jobs id; Type: DEFAULT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.failed_jobs ALTER COLUMN id SET DEFAULT nextval('public.failed_jobs_id_seq'::regclass);


--
-- Name: incident_templates id; Type: DEFAULT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.incident_templates ALTER COLUMN id SET DEFAULT nextval('public.incident_templates_id_seq'::regclass);


--
-- Name: incidents id; Type: DEFAULT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.incidents ALTER COLUMN id SET DEFAULT nextval('public.incidents_id_seq'::regclass);


--
-- Name: invites id; Type: DEFAULT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.invites ALTER COLUMN id SET DEFAULT nextval('public.invites_id_seq'::regclass);


--
-- Name: jobs id; Type: DEFAULT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.jobs ALTER COLUMN id SET DEFAULT nextval('public.jobs_id_seq'::regclass);


--
-- Name: metric_points id; Type: DEFAULT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.metric_points ALTER COLUMN id SET DEFAULT nextval('public.metric_points_id_seq'::regclass);


--
-- Name: metrics id; Type: DEFAULT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.metrics ALTER COLUMN id SET DEFAULT nextval('public.metrics_id_seq'::regclass);


--
-- Name: settings id; Type: DEFAULT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.settings ALTER COLUMN id SET DEFAULT nextval('public.settings_id_seq'::regclass);


--
-- Name: subscribers id; Type: DEFAULT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.subscribers ALTER COLUMN id SET DEFAULT nextval('public.subscribers_id_seq'::regclass);


--
-- Name: subscriptions id; Type: DEFAULT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.subscriptions ALTER COLUMN id SET DEFAULT nextval('public.subscriptions_id_seq'::regclass);


--
-- Name: tags id; Type: DEFAULT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.tags ALTER COLUMN id SET DEFAULT nextval('public.tags_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: cache; Type: TABLE DATA; Schema: public; Owner: cachet
--

COPY public.cache (key, value, expiration) FROM stdin;
\.


--
-- Data for Name: component_groups; Type: TABLE DATA; Schema: public; Owner: cachet
--

COPY public.component_groups (id, name, created_at, updated_at, "order", collapsed) FROM stdin;
1	group1	2021-08-15 09:24:41	2021-08-15 09:24:41	0	0
2	group2	2021-08-15 09:24:47	2021-08-15 09:24:47	0	0
\.


--
-- Data for Name: component_tag; Type: TABLE DATA; Schema: public; Owner: cachet
--

COPY public.component_tag (id, component_id, tag_id) FROM stdin;
1	1	1
2	2	1
3	3	1
4	4	1
5	5	1
\.


--
-- Data for Name: components; Type: TABLE DATA; Schema: public; Owner: cachet
--

COPY public.components (id, name, description, link, status, "order", group_id, created_at, updated_at, deleted_at, enabled) FROM stdin;
1	triple1			1	0	0	2021-08-15 09:25:16	2021-08-15 09:25:16	\N	t
2	triple1			1	0	1	2021-08-15 09:25:23	2021-08-15 09:25:23	\N	t
3	triple1			1	0	2	2021-08-15 09:25:29	2021-08-15 09:25:29	\N	t
4	component1			1	0	0	2021-08-15 09:26:04	2021-08-15 09:26:04	\N	t
5	component2			1	0	1	2021-08-15 09:26:15	2021-08-15 09:26:15	\N	t
\.


--
-- Data for Name: failed_jobs; Type: TABLE DATA; Schema: public; Owner: cachet
--

COPY public.failed_jobs (id, connection, queue, payload, failed_at) FROM stdin;
\.


--
-- Data for Name: incident_templates; Type: TABLE DATA; Schema: public; Owner: cachet
--

COPY public.incident_templates (id, name, slug, template, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: incidents; Type: TABLE DATA; Schema: public; Owner: cachet
--

COPY public.incidents (id, component_id, name, status, message, created_at, updated_at, deleted_at, scheduled_at, visible) FROM stdin;
\.


--
-- Data for Name: invites; Type: TABLE DATA; Schema: public; Owner: cachet
--

COPY public.invites (id, code, email, claimed_at, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: jobs; Type: TABLE DATA; Schema: public; Owner: cachet
--

COPY public.jobs (id, queue, payload, attempts, reserved, reserved_at, available_at, created_at) FROM stdin;
\.


--
-- Data for Name: metric_points; Type: TABLE DATA; Schema: public; Owner: cachet
--

COPY public.metric_points (id, metric_id, value, created_at, updated_at, counter) FROM stdin;
\.


--
-- Data for Name: metrics; Type: TABLE DATA; Schema: public; Owner: cachet
--

COPY public.metrics (id, name, suffix, description, default_value, calc_type, display_chart, created_at, updated_at, places, default_view, threshold, "order") FROM stdin;
\.


--
-- Data for Name: migrations; Type: TABLE DATA; Schema: public; Owner: cachet
--

COPY public.migrations (migration, batch) FROM stdin;
2015_01_05_201324_CreateComponentGroupsTable	1
2015_01_05_201444_CreateComponentsTable	1
2015_01_05_202446_CreateIncidentTemplatesTable	1
2015_01_05_202609_CreateIncidentsTable	1
2015_01_05_202730_CreateMetricPointsTable	1
2015_01_05_202826_CreateMetricsTable	1
2015_01_05_203014_CreateSettingsTable	1
2015_01_05_203235_CreateSubscribersTable	1
2015_01_05_203341_CreateUsersTable	1
2015_01_09_083419_AlterTableUsersAdd2FA	1
2015_01_16_083825_CreateTagsTable	1
2015_01_16_084030_CreateComponentTagTable	1
2015_02_28_214642_UpdateIncidentsAddScheduledAt	1
2015_05_19_214534_AlterTableComponentGroupsAddOrder	1
2015_05_20_073041_AlterTableIncidentsAddVisibileColumn	1
2015_05_24_210939_create_jobs_table	1
2015_05_24_210948_create_failed_jobs_table	1
2015_06_10_122216_AlterTableComponentsDropUserIdColumn	1
2015_06_10_122229_AlterTableIncidentsDropUserIdColumn	1
2015_08_02_120436_AlterTableSubscribersRemoveDeletedAt	1
2015_08_13_214123_AlterTableMetricsAddDecimalPlacesColumn	1
2015_10_31_211944_CreateInvitesTable	1
2015_11_03_211049_AlterTableComponentsAddEnabledColumn	1
2015_12_26_162258_AlterTableMetricsAddDefaultViewColumn	1
2016_01_09_141852_CreateSubscriptionsTable	1
2016_01_29_154937_AlterTableComponentGroupsAddCollapsedColumn	1
2016_02_18_085210_AlterTableMetricPointsChangeValueColumn	1
2016_03_01_174858_AlterTableMetricPointsAddCounterColumn	1
2016_03_10_144613_AlterTableComponentGroupsMakeColumnInteger	1
2016_04_05_142933_create_sessions_table	1
2016_04_29_061916_AlterTableSubscribersAddGlobalColumn	1
2016_06_02_075012_AlterTableMetricsAddOrderColumn	1
2016_06_05_091615_create_cache_table	1
\.


--
-- Data for Name: sessions; Type: TABLE DATA; Schema: public; Owner: cachet
--

COPY public.sessions (id, user_id, ip_address, user_agent, payload, last_activity) FROM stdin;
\.


--
-- Data for Name: settings; Type: TABLE DATA; Schema: public; Owner: cachet
--

COPY public.settings (id, name, value, created_at, updated_at) FROM stdin;
1	app_name	fake	2021-08-15 09:19:43	2021-08-15 09:19:43
2	app_domain	http://localhost:8000	2021-08-15 09:19:43	2021-08-15 09:19:43
3	app_timezone	Europe/Paris	2021-08-15 09:19:43	2021-08-15 09:19:43
4	app_locale	en	2021-08-15 09:19:43	2021-08-15 09:19:43
5	show_support	1	2021-08-15 09:19:43	2021-08-15 09:19:43
6	app_incident_days	7	2021-08-15 09:19:43	2021-08-15 09:19:43
\.


--
-- Data for Name: subscribers; Type: TABLE DATA; Schema: public; Owner: cachet
--

COPY public.subscribers (id, email, verify_code, verified_at, created_at, updated_at, global) FROM stdin;
\.


--
-- Data for Name: subscriptions; Type: TABLE DATA; Schema: public; Owner: cachet
--

COPY public.subscriptions (id, subscriber_id, component_id, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: tags; Type: TABLE DATA; Schema: public; Owner: cachet
--

COPY public.tags (id, name, slug, created_at, updated_at) FROM stdin;
1			2021-08-15 09:25:16	2021-08-15 09:25:16
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: cachet
--

COPY public.users (id, username, password, remember_token, email, api_key, active, level, created_at, updated_at, google_2fa_secret) FROM stdin;
1	admin	$2y$10$R0xSDAuOaKNJmrNuEdjf..iZleTUKe5qu.xiuDTtgbYfC3519bl.y	\N	admin@fake.com	8I0tUhh4WUZIL7eyFUMX	t	1	2021-08-15 09:19:43	2021-08-15 09:19:43	\N
\.


--
-- Name: component_groups_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cachet
--

SELECT pg_catalog.setval('public.component_groups_id_seq', 2, true);


--
-- Name: component_tag_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cachet
--

SELECT pg_catalog.setval('public.component_tag_id_seq', 5, true);


--
-- Name: components_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cachet
--

SELECT pg_catalog.setval('public.components_id_seq', 5, true);


--
-- Name: failed_jobs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cachet
--

SELECT pg_catalog.setval('public.failed_jobs_id_seq', 1, false);


--
-- Name: incident_templates_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cachet
--

SELECT pg_catalog.setval('public.incident_templates_id_seq', 1, false);


--
-- Name: incidents_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cachet
--

SELECT pg_catalog.setval('public.incidents_id_seq', 1, false);


--
-- Name: invites_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cachet
--

SELECT pg_catalog.setval('public.invites_id_seq', 1, false);


--
-- Name: jobs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cachet
--

SELECT pg_catalog.setval('public.jobs_id_seq', 1, false);


--
-- Name: metric_points_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cachet
--

SELECT pg_catalog.setval('public.metric_points_id_seq', 1, false);


--
-- Name: metrics_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cachet
--

SELECT pg_catalog.setval('public.metrics_id_seq', 1, false);


--
-- Name: settings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cachet
--

SELECT pg_catalog.setval('public.settings_id_seq', 6, true);


--
-- Name: subscribers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cachet
--

SELECT pg_catalog.setval('public.subscribers_id_seq', 1, false);


--
-- Name: subscriptions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cachet
--

SELECT pg_catalog.setval('public.subscriptions_id_seq', 1, false);


--
-- Name: tags_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cachet
--

SELECT pg_catalog.setval('public.tags_id_seq', 1, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cachet
--

SELECT pg_catalog.setval('public.users_id_seq', 1, true);


--
-- Name: cache cache_key_unique; Type: CONSTRAINT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.cache
    ADD CONSTRAINT cache_key_unique UNIQUE (key);


--
-- Name: component_groups component_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.component_groups
    ADD CONSTRAINT component_groups_pkey PRIMARY KEY (id);


--
-- Name: component_tag component_tag_pkey; Type: CONSTRAINT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.component_tag
    ADD CONSTRAINT component_tag_pkey PRIMARY KEY (id);


--
-- Name: components components_pkey; Type: CONSTRAINT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.components
    ADD CONSTRAINT components_pkey PRIMARY KEY (id);


--
-- Name: failed_jobs failed_jobs_pkey; Type: CONSTRAINT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.failed_jobs
    ADD CONSTRAINT failed_jobs_pkey PRIMARY KEY (id);


--
-- Name: incident_templates incident_templates_pkey; Type: CONSTRAINT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.incident_templates
    ADD CONSTRAINT incident_templates_pkey PRIMARY KEY (id);


--
-- Name: incidents incidents_pkey; Type: CONSTRAINT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.incidents
    ADD CONSTRAINT incidents_pkey PRIMARY KEY (id);


--
-- Name: invites invites_code_unique; Type: CONSTRAINT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.invites
    ADD CONSTRAINT invites_code_unique UNIQUE (code);


--
-- Name: invites invites_pkey; Type: CONSTRAINT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.invites
    ADD CONSTRAINT invites_pkey PRIMARY KEY (id);


--
-- Name: jobs jobs_pkey; Type: CONSTRAINT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.jobs
    ADD CONSTRAINT jobs_pkey PRIMARY KEY (id);


--
-- Name: metric_points metric_points_pkey; Type: CONSTRAINT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.metric_points
    ADD CONSTRAINT metric_points_pkey PRIMARY KEY (id);


--
-- Name: metrics metrics_pkey; Type: CONSTRAINT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.metrics
    ADD CONSTRAINT metrics_pkey PRIMARY KEY (id);


--
-- Name: sessions sessions_id_unique; Type: CONSTRAINT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_id_unique UNIQUE (id);


--
-- Name: settings settings_pkey; Type: CONSTRAINT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.settings
    ADD CONSTRAINT settings_pkey PRIMARY KEY (id);


--
-- Name: subscribers subscribers_email_unique; Type: CONSTRAINT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.subscribers
    ADD CONSTRAINT subscribers_email_unique UNIQUE (email);


--
-- Name: subscribers subscribers_pkey; Type: CONSTRAINT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.subscribers
    ADD CONSTRAINT subscribers_pkey PRIMARY KEY (id);


--
-- Name: subscriptions subscriptions_pkey; Type: CONSTRAINT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.subscriptions
    ADD CONSTRAINT subscriptions_pkey PRIMARY KEY (id);


--
-- Name: tags tags_name_slug_unique; Type: CONSTRAINT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_name_slug_unique UNIQUE (name, slug);


--
-- Name: tags tags_pkey; Type: CONSTRAINT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (id);


--
-- Name: users users_api_key_unique; Type: CONSTRAINT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_api_key_unique UNIQUE (api_key);


--
-- Name: users users_email_unique; Type: CONSTRAINT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_unique UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_unique; Type: CONSTRAINT; Schema: public; Owner: cachet
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_unique UNIQUE (username);


--
-- Name: component_groups_order_index; Type: INDEX; Schema: public; Owner: cachet
--

CREATE INDEX component_groups_order_index ON public.component_groups USING btree ("order");


--
-- Name: component_tag_component_id_index; Type: INDEX; Schema: public; Owner: cachet
--

CREATE INDEX component_tag_component_id_index ON public.component_tag USING btree (component_id);


--
-- Name: component_tag_tag_id_index; Type: INDEX; Schema: public; Owner: cachet
--

CREATE INDEX component_tag_tag_id_index ON public.component_tag USING btree (tag_id);


--
-- Name: components_group_id_index; Type: INDEX; Schema: public; Owner: cachet
--

CREATE INDEX components_group_id_index ON public.components USING btree (group_id);


--
-- Name: components_order_index; Type: INDEX; Schema: public; Owner: cachet
--

CREATE INDEX components_order_index ON public.components USING btree ("order");


--
-- Name: components_status_index; Type: INDEX; Schema: public; Owner: cachet
--

CREATE INDEX components_status_index ON public.components USING btree (status);


--
-- Name: incidents_component_id_index; Type: INDEX; Schema: public; Owner: cachet
--

CREATE INDEX incidents_component_id_index ON public.incidents USING btree (component_id);


--
-- Name: incidents_status_index; Type: INDEX; Schema: public; Owner: cachet
--

CREATE INDEX incidents_status_index ON public.incidents USING btree (status);


--
-- Name: incidents_visible_index; Type: INDEX; Schema: public; Owner: cachet
--

CREATE INDEX incidents_visible_index ON public.incidents USING btree (visible);


--
-- Name: metric_points_metric_id_index; Type: INDEX; Schema: public; Owner: cachet
--

CREATE INDEX metric_points_metric_id_index ON public.metric_points USING btree (metric_id);


--
-- Name: metrics_display_chart_index; Type: INDEX; Schema: public; Owner: cachet
--

CREATE INDEX metrics_display_chart_index ON public.metrics USING btree (display_chart);


--
-- Name: subscriptions_component_id_index; Type: INDEX; Schema: public; Owner: cachet
--

CREATE INDEX subscriptions_component_id_index ON public.subscriptions USING btree (component_id);


--
-- Name: subscriptions_subscriber_id_index; Type: INDEX; Schema: public; Owner: cachet
--

CREATE INDEX subscriptions_subscriber_id_index ON public.subscriptions USING btree (subscriber_id);


--
-- Name: users_active_index; Type: INDEX; Schema: public; Owner: cachet
--

CREATE INDEX users_active_index ON public.users USING btree (active);


--
-- Name: users_remember_token_index; Type: INDEX; Schema: public; Owner: cachet
--

CREATE INDEX users_remember_token_index ON public.users USING btree (remember_token);


--
-- PostgreSQL database dump complete
--

