--
-- PostgreSQL database dump
--

-- Dumped from database version 11.6
-- Dumped by pg_dump version 11.6

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

--
-- Name: postgis; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS postgis WITH SCHEMA public;


--
-- Name: EXTENSION postgis; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION postgis IS 'PostGIS geometry, geography, and raster spatial types and functions';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: admins; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.admins (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text NOT NULL,
    email text NOT NULL,
    password text NOT NULL,
    auth_token text,
    image text,
    is_active boolean NOT NULL
);


--
-- Name: admins_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.admins_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: admins_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.admins_id_seq OWNED BY public.admins.id;


--
-- Name: driver_document_uploads; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.driver_document_uploads (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    doc_id integer NOT NULL,
    driver_id integer NOT NULL,
    image text NOT NULL,
    is_active boolean
);


--
-- Name: driver_document_uploads_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.driver_document_uploads_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: driver_document_uploads_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.driver_document_uploads_id_seq OWNED BY public.driver_document_uploads.id;


--
-- Name: driver_documents; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.driver_documents (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    operator_id integer NOT NULL,
    name text NOT NULL,
    is_active boolean
);


--
-- Name: driver_documents_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.driver_documents_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: driver_documents_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.driver_documents_id_seq OWNED BY public.driver_documents.id;


--
-- Name: drivers; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.drivers (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text NOT NULL,
    dial_code bigint NOT NULL,
    mobile_number text NOT NULL,
    operator_id integer NOT NULL,
    license_number text,
    vehicle_name text NOT NULL,
    vehicle_type_id integer NOT NULL,
    vehicle_brand text NOT NULL,
    vehicle_model text,
    vehicle_color text NOT NULL,
    vehicle_number text NOT NULL,
    vehicle_image text,
    auth_token text,
    driver_image text,
    fcm_id text,
    is_profile_completed boolean,
    is_online boolean,
    is_ride boolean,
    is_active boolean,
    latlng public.geometry,
    fleet_id integer DEFAULT 0 NOT NULL,
    balance double precision DEFAULT 0,
    outgo_pending double precision DEFAULT 0
);


--
-- Name: drivers_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.drivers_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: drivers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.drivers_id_seq OWNED BY public.drivers.id;


--
-- Name: fares; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.fares (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    vehicle_type_id integer NOT NULL,
    operator_id integer NOT NULL,
    base_fare numeric NOT NULL,
    minimum_fare numeric NOT NULL,
    waiting_time_limit numeric NOT NULL,
    waiting_fee numeric NOT NULL,
    cancellation_time_limit numeric NOT NULL,
    cancellation_fee numeric NOT NULL,
    duration_fare numeric NOT NULL,
    distance_fare numeric NOT NULL,
    tax numeric NOT NULL,
    traffic_factor numeric NOT NULL,
    is_active boolean
);


--
-- Name: fares_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.fares_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: fares_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.fares_id_seq OWNED BY public.fares.id;


--
-- Name: fleets_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.fleets_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: fleets; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.fleets (
    id integer DEFAULT nextval('public.fleets_id_seq'::regclass) NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    operator_id integer NOT NULL,
    name text,
    is_active boolean,
    email text,
    password text,
    balance double precision DEFAULT 0,
    outgo_pending double precision DEFAULT 0
);


--
-- Name: operators; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.operators (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text NOT NULL,
    location_name text NOT NULL,
    email text NOT NULL,
    password text NOT NULL,
    platform_commission numeric NOT NULL,
    operator_commission numeric NOT NULL,
    driver_work_time integer,
    driver_rest_time integer,
    currency text,
    auth_token text,
    is_active boolean,
    polygon public.geometry,
    refer_amount numeric,
    refer_type integer,
    balance double precision DEFAULT 0,
    outgo_pending double precision DEFAULT 0
);


--
-- Name: operators_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.operators_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: operators_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.operators_id_seq OWNED BY public.operators.id;


--
-- Name: otps; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.otps (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    dial_code bigint NOT NULL,
    country_code text NOT NULL,
    mobile_number text NOT NULL,
    otp text NOT NULL,
    is_used boolean
);


--
-- Name: otps_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.otps_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: otps_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.otps_id_seq OWNED BY public.otps.id;


--
-- Name: passengers; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.passengers (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text NOT NULL,
    dial_code bigint NOT NULL,
    country_code text NOT NULL,
    mobile_number text NOT NULL,
    auth_token text,
    image text,
    fcm_id text,
    is_active boolean,
    referral_code text,
    referred_by text,
    wallet_balance numeric,
    balance double precision DEFAULT 0,
    income_pending double precision DEFAULT 0
);


--
-- Name: passengers_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.passengers_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: passengers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.passengers_id_seq OWNED BY public.passengers.id;


--
-- Name: rides; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.rides (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    passenger_id integer NOT NULL,
    operator_id integer NOT NULL,
    zone_id integer,
    vehicle_type_id integer NOT NULL,
    driver_id integer,
    pickup_point text,
    pickup_location text,
    pickup_latitude numeric NOT NULL,
    pickup_longitude numeric NOT NULL,
    drop_location text,
    drop_latitude numeric NOT NULL,
    drop_longitude numeric NOT NULL,
    ride_date_time timestamp with time zone NOT NULL,
    ride_driver_arrived_time timestamp with time zone NOT NULL,
    ride_start_time timestamp with time zone NOT NULL,
    ride_end_time timestamp with time zone NOT NULL,
    ride_type bigint NOT NULL,
    is_ride_later boolean NOT NULL,
    distance numeric,
    duration numeric,
    duration_readable text,
    fare_id integer,
    zone_fare_id integer,
    distance_fare numeric,
    duration_fare numeric,
    waiting_fare numeric,
    cancellation_fee numeric,
    tax numeric,
    is_paid boolean,
    transaction_id text,
    total_fare numeric,
    passenger_rating numeric,
    driver_rating numeric,
    passenger_review text,
    driver_review text,
    ride_status bigint NOT NULL,
    is_active boolean,
    is_multi_stop boolean,
    payment_verified integer DEFAULT 0,
    fleet_id integer DEFAULT 0
);


--
-- Name: pendings_driver; Type: VIEW; Schema: public; Owner: -
--

CREATE VIEW public.pendings_driver AS
 SELECT r.driver_id AS id,
    (sum(r.total_fare) * 0.85) AS income_pending
   FROM public.rides r
  WHERE (r.payment_verified = 0)
  GROUP BY r.driver_id;


--
-- Name: pendings_fleet; Type: VIEW; Schema: public; Owner: -
--

CREATE VIEW public.pendings_fleet AS
 SELECT r.fleet_id AS id,
    sum(((r.total_fare * (o.platform_commission - o.operator_commission)) / (100)::numeric)) AS income_pending
   FROM (public.rides r
     LEFT JOIN public.operators o ON ((o.id = r.operator_id)))
  WHERE (r.payment_verified = 0)
  GROUP BY r.fleet_id;


--
-- Name: pendings_operator; Type: VIEW; Schema: public; Owner: -
--

CREATE VIEW public.pendings_operator AS
 SELECT r.operator_id AS id,
    sum(
        CASE
            WHEN (r.fleet_id <> 0) THEN ((r.total_fare * o.operator_commission) / (100)::numeric)
            WHEN (r.fleet_id = 0) THEN (r.total_fare * ((o.platform_commission / (100)::numeric) + 0.05))
            ELSE NULL::numeric
        END) AS income_pending
   FROM (public.rides r
     LEFT JOIN public.operators o ON ((o.id = r.operator_id)))
  WHERE (r.payment_verified = 0)
  GROUP BY r.operator_id;


--
-- Name: pendings_passenger; Type: VIEW; Schema: public; Owner: -
--

CREATE VIEW public.pendings_passenger AS
 SELECT r.passenger_id AS id,
    sum(r.total_fare) AS outgo_pending
   FROM public.rides r
  WHERE (r.payment_verified = 0)
  GROUP BY r.passenger_id;


--
-- Name: pickup_points; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.pickup_points (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text NOT NULL,
    zone_id integer NOT NULL,
    is_active boolean
);


--
-- Name: pickup_points_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.pickup_points_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: pickup_points_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.pickup_points_id_seq OWNED BY public.pickup_points.id;


--
-- Name: ride_event_logs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.ride_event_logs (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    ride_id integer NOT NULL,
    ride_status bigint NOT NULL,
    message text NOT NULL,
    is_active boolean
);


--
-- Name: ride_event_logs_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.ride_event_logs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: ride_event_logs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.ride_event_logs_id_seq OWNED BY public.ride_event_logs.id;


--
-- Name: ride_locations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.ride_locations (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    ride_id integer NOT NULL,
    "time" timestamp with time zone NOT NULL,
    is_active boolean,
    latlng public.geometry
);


--
-- Name: ride_locations_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.ride_locations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: ride_locations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.ride_locations_id_seq OWNED BY public.ride_locations.id;


--
-- Name: ride_messages; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.ride_messages (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    ride_id integer NOT NULL,
    message text,
    "from" integer,
    is_active boolean
);


--
-- Name: ride_messages_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.ride_messages_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: ride_messages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.ride_messages_id_seq OWNED BY public.ride_messages.id;


--
-- Name: ride_stops; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.ride_stops (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    ride_id integer NOT NULL,
    location text,
    latitude numeric NOT NULL,
    longitude numeric NOT NULL,
    is_reached boolean,
    is_active boolean
);


--
-- Name: ride_stops_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.ride_stops_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: ride_stops_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.ride_stops_id_seq OWNED BY public.ride_stops.id;


--
-- Name: rides_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.rides_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: rides_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.rides_id_seq OWNED BY public.rides.id;


--
-- Name: sent_ride_requests; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.sent_ride_requests (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    driver_id integer NOT NULL,
    ride_id integer NOT NULL,
    is_active boolean
);


--
-- Name: sent_ride_requests_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.sent_ride_requests_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: sent_ride_requests_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.sent_ride_requests_id_seq OWNED BY public.sent_ride_requests.id;


--
-- Name: vehicle_categories; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.vehicle_categories (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text NOT NULL,
    description text NOT NULL,
    is_active boolean
);


--
-- Name: vehicle_categories_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.vehicle_categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: vehicle_categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.vehicle_categories_id_seq OWNED BY public.vehicle_categories.id;


--
-- Name: vehicle_types; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.vehicle_types (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text NOT NULL,
    image text NOT NULL,
    vehicle_category_id integer NOT NULL,
    description text NOT NULL,
    image_active text NOT NULL,
    seat_capacity bigint NOT NULL,
    is_active boolean
);


--
-- Name: vehicle_types_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.vehicle_types_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: vehicle_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.vehicle_types_id_seq OWNED BY public.vehicle_types.id;


--
-- Name: zone_fares; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.zone_fares (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    vehicle_type_id integer NOT NULL,
    zone_id integer NOT NULL,
    base_fare numeric NOT NULL,
    minimum_fare numeric NOT NULL,
    waiting_time_limit numeric NOT NULL,
    waiting_fee numeric NOT NULL,
    cancellation_time_limit numeric NOT NULL,
    cancellation_fee numeric NOT NULL,
    duration_fare numeric NOT NULL,
    distance_fare numeric NOT NULL,
    tax numeric NOT NULL,
    traffic_factor numeric NOT NULL,
    is_active boolean
);


--
-- Name: zone_fares_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.zone_fares_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: zone_fares_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.zone_fares_id_seq OWNED BY public.zone_fares.id;


--
-- Name: zones; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.zones (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text NOT NULL,
    operator_id integer NOT NULL,
    is_active boolean,
    polygon public.geometry
);


--
-- Name: zones_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.zones_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: zones_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.zones_id_seq OWNED BY public.zones.id;


--
-- Name: admins id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.admins ALTER COLUMN id SET DEFAULT nextval('public.admins_id_seq'::regclass);


--
-- Name: driver_document_uploads id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.driver_document_uploads ALTER COLUMN id SET DEFAULT nextval('public.driver_document_uploads_id_seq'::regclass);


--
-- Name: driver_documents id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.driver_documents ALTER COLUMN id SET DEFAULT nextval('public.driver_documents_id_seq'::regclass);


--
-- Name: drivers id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.drivers ALTER COLUMN id SET DEFAULT nextval('public.drivers_id_seq'::regclass);


--
-- Name: fares id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.fares ALTER COLUMN id SET DEFAULT nextval('public.fares_id_seq'::regclass);


--
-- Name: operators id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.operators ALTER COLUMN id SET DEFAULT nextval('public.operators_id_seq'::regclass);


--
-- Name: otps id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.otps ALTER COLUMN id SET DEFAULT nextval('public.otps_id_seq'::regclass);


--
-- Name: passengers id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.passengers ALTER COLUMN id SET DEFAULT nextval('public.passengers_id_seq'::regclass);


--
-- Name: pickup_points id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.pickup_points ALTER COLUMN id SET DEFAULT nextval('public.pickup_points_id_seq'::regclass);


--
-- Name: ride_event_logs id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.ride_event_logs ALTER COLUMN id SET DEFAULT nextval('public.ride_event_logs_id_seq'::regclass);


--
-- Name: ride_locations id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.ride_locations ALTER COLUMN id SET DEFAULT nextval('public.ride_locations_id_seq'::regclass);


--
-- Name: ride_messages id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.ride_messages ALTER COLUMN id SET DEFAULT nextval('public.ride_messages_id_seq'::regclass);


--
-- Name: ride_stops id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.ride_stops ALTER COLUMN id SET DEFAULT nextval('public.ride_stops_id_seq'::regclass);


--
-- Name: rides id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.rides ALTER COLUMN id SET DEFAULT nextval('public.rides_id_seq'::regclass);


--
-- Name: sent_ride_requests id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sent_ride_requests ALTER COLUMN id SET DEFAULT nextval('public.sent_ride_requests_id_seq'::regclass);


--
-- Name: vehicle_categories id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.vehicle_categories ALTER COLUMN id SET DEFAULT nextval('public.vehicle_categories_id_seq'::regclass);


--
-- Name: vehicle_types id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.vehicle_types ALTER COLUMN id SET DEFAULT nextval('public.vehicle_types_id_seq'::regclass);


--
-- Name: zone_fares id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.zone_fares ALTER COLUMN id SET DEFAULT nextval('public.zone_fares_id_seq'::regclass);


--
-- Name: zones id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.zones ALTER COLUMN id SET DEFAULT nextval('public.zones_id_seq'::regclass);


--
-- Data for Name: admins; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.admins (id, created_at, updated_at, deleted_at, name, email, password, auth_token, image, is_active) FROM stdin;
1	2020-02-02 17:00:08.455054+00	2020-02-02 17:00:08.455054+00	\N	Taasai Admin	admin@taasai.com	$2a$14$PHtm0kilvlUjNy6ALoajFO3oXp83.tzVIKsHBXOU5i15V41l6LGdK	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX3R5cGUiOiJhZG1pbiIsImV4cCI6MTU4MzQ4NTkwOSwiaXNzIjoidGFhc2FpIn0.O8ufVnk4nQXpJKvLhqXXEuAuV5_hBg3YPojjea3rKgI		t
\.


--
-- Data for Name: driver_document_uploads; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.driver_document_uploads (id, created_at, updated_at, deleted_at, doc_id, driver_id, image, is_active) FROM stdin;
1	2020-02-02 18:55:19.945106+00	2020-02-02 18:55:19.945106+00	\N	1	1	1580669719468988913_image_cropper_1580669715096.jpg	t
2	2020-02-02 18:55:29.234415+00	2020-02-02 18:55:29.234415+00	\N	2	1	1580669728758730133_image_cropper_1580669724199.jpg	t
3	2020-02-02 18:55:40.3798+00	2020-02-02 18:55:40.3798+00	\N	3	1	1580669739904255636_image_cropper_1580669734626.jpg	t
4	2020-02-02 18:55:52.219072+00	2020-02-02 18:55:52.219072+00	\N	4	1	1580669751743051832_image_cropper_1580669746182.jpg	t
5	2020-02-03 03:22:05.943217+00	2020-02-03 03:22:05.943217+00	\N	1	2	1580700125458543656_image_cropper_1580700122073.jpg	t
6	2020-02-03 03:22:18.74031+00	2020-02-03 03:22:18.74031+00	\N	2	2	1580700138254777500_image_cropper_1580700132084.jpg	t
7	2020-02-03 03:22:28.656166+00	2020-02-03 03:22:28.656166+00	\N	3	2	1580700148171322419_image_cropper_1580700143692.jpg	t
8	2020-02-03 03:22:37.937201+00	2020-02-03 03:22:37.937201+00	\N	4	2	1580700157451501632_image_cropper_1580700153817.jpg	t
9	2020-02-03 07:55:43.545679+00	2020-02-03 07:55:43.545679+00	\N	1	3	1580716543054180313_image_cropper_1580716538573.jpg	t
10	2020-02-03 07:55:52.524791+00	2020-02-03 07:55:52.524791+00	\N	2	3	1580716552033204307_image_cropper_1580716547260.jpg	t
11	2020-02-03 07:56:02.49699+00	2020-02-03 07:56:02.49699+00	\N	3	3	1580716562005603141_image_cropper_1580716557751.jpg	t
12	2020-02-03 07:56:12.671245+00	2020-02-03 07:56:12.671245+00	\N	4	3	1580716572180444711_image_cropper_1580716567776.jpg	t
13	2020-02-03 10:47:54.192798+00	2020-02-03 10:47:54.192798+00	\N	5	4	1580726873717793761_image_cropper_1580726865192.jpg	t
14	2020-02-03 15:50:50.437734+00	2020-02-03 15:50:50.437734+00	\N	1	13	1580745049950095362_image_cropper_1580745041864.jpg	t
15	2020-02-03 16:31:24.834657+00	2020-02-03 16:31:24.834657+00	\N	1	16	1580747484333460288_image_cropper_1580747478603.jpg	t
16	2020-02-03 16:31:35.711037+00	2020-02-03 16:31:35.711037+00	\N	2	16	1580747495209950314_image_cropper_1580747489129.jpg	t
17	2020-02-03 16:31:50.37012+00	2020-02-03 16:31:50.37012+00	\N	3	16	1580747509879733421_image_cropper_1580747504368.jpg	t
18	2020-02-03 16:32:05.447831+00	2020-02-03 16:32:05.447831+00	\N	4	16	1580747524956570879_image_cropper_1580747514076.jpg	t
19	2020-02-03 16:37:20.642934+00	2020-02-03 16:37:20.642934+00	\N	1	17	1580747840151313322_image_cropper_1580747830464.jpg	t
20	2020-02-03 16:37:36.222708+00	2020-02-03 16:37:36.222708+00	\N	2	17	1580747855732411162_image_cropper_1580747847054.jpg	t
21	2020-02-03 16:37:56.926046+00	2020-02-03 16:37:56.926046+00	\N	3	17	1580747876436210271_image_cropper_1580747871047.jpg	t
22	2020-02-03 16:38:11.358448+00	2020-02-03 16:38:11.358448+00	\N	4	17	1580747890868101589_image_cropper_1580747882744.jpg	t
23	2020-02-04 14:36:58.380166+00	2020-02-04 14:36:58.380166+00	\N	5	18	1580827017910615617_image_cropper_1580827011114.jpg	t
24	2020-02-05 03:21:01.148825+00	2020-02-05 03:21:01.148825+00	\N	5	20	1580872860661767661_image_cropper_1580872855754.jpg	t
25	2020-02-05 04:09:35.03043+00	2020-02-05 04:09:35.03043+00	\N	5	22	1580875774555153084_image_cropper_1580875770208.jpg	t
26	2020-02-05 08:18:27.168983+00	2020-02-05 08:18:27.168983+00	\N	5	25	1580890706677667749_image_cropper_1580890693890.jpg	f
27	2020-02-05 08:20:32.030265+00	2020-02-05 08:20:32.030265+00	\N	5	25	1580890831538564127_image_cropper_1580890823286.jpg	t
28	2020-02-05 08:42:14.639279+00	2020-02-05 08:42:14.639279+00	\N	5	27	1580892134158696933_image_cropper_1580892110267.jpg	t
29	2020-02-05 09:01:34.068355+00	2020-02-05 09:01:34.068355+00	\N	5	28	1580893293588762908_image_cropper_1580893287283.jpg	t
30	2020-02-05 09:27:24.673123+00	2020-02-05 09:27:24.673123+00	\N	5	31	1580894844189531704_image_cropper_1580894839830.jpg	t
31	2020-02-05 09:41:48.622379+00	2020-02-05 09:41:48.622379+00	\N	5	32	1580895708148180048_image_cropper_1580895701998.jpg	t
32	2020-02-05 10:40:51.747071+00	2020-02-05 10:40:51.747071+00	\N	5	33	1580899251273847170_image_cropper_1580899245621.jpg	t
33	2020-02-07 03:21:18.903226+00	2020-02-07 03:21:18.903226+00	\N	5	34	1581045678433272261_image_cropper_1581045671998.jpg	t
34	2020-02-07 06:48:48.342014+00	2020-02-07 06:48:48.342014+00	\N	5	35	1581058127868034896_image_cropper_1581058122583.jpg	t
35	2020-02-16 13:20:19.329354+00	2020-02-16 13:20:19.329354+00	\N	5	36	1581859218828278787_image_cropper_1581859214535.jpg	t
36	2020-02-17 11:24:55.654025+00	2020-02-17 11:24:55.654025+00	\N	6	37	1581938695157289764_image_cropper_1581938689547.jpg	t
37	2020-02-19 04:13:47.493537+00	2020-02-19 04:13:47.493537+00	\N	6	38	1582085627003693302_image_cropper_1582085622408.jpg	t
\.


--
-- Data for Name: driver_documents; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.driver_documents (id, created_at, updated_at, deleted_at, operator_id, name, is_active) FROM stdin;
1	2020-02-02 17:03:43.733525+00	2020-02-02 17:03:43.733525+00	\N	1	Driver License Front	f
2	2020-02-02 17:03:44.132597+00	2020-02-02 17:03:44.132597+00	\N	1	Driver License Back	f
3	2020-02-02 17:03:44.526982+00	2020-02-02 17:03:44.526982+00	\N	1	Vehicle Registration Certificate	f
4	2020-02-02 17:03:44.92127+00	2020-02-02 17:03:44.92127+00	\N	1	Vehicle Insurance	f
5	2020-02-02 20:46:51.492201+00	2020-02-02 20:46:51.492201+00	\N	2	Driving License	f
6	2020-02-16 16:43:59.128123+00	2020-02-16 16:43:59.128123+00	\N	3	Passport	f
\.


--
-- Data for Name: drivers; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.drivers (id, created_at, updated_at, deleted_at, name, dial_code, mobile_number, operator_id, license_number, vehicle_name, vehicle_type_id, vehicle_brand, vehicle_model, vehicle_color, vehicle_number, vehicle_image, auth_token, driver_image, fcm_id, is_profile_completed, is_online, is_ride, is_active, latlng, fleet_id, balance, outgo_pending) FROM stdin;
1	2020-02-02 18:55:06.548098+00	2020-02-02 18:55:06.548098+00	\N	Praveen	91	8364727485	1		Kwid	1	Taasai Go	2001	Red	TN 41 AN 1837	public/vehicle/1580669705878696245_image_cropper_1580669680510.jpg	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjowLCJ1c2VyX3R5cGUiOiJkcml2ZXIiLCJleHAiOjE1ODM0MjE3MTEsImlzcyI6InRhYXNhaSJ9.Jyi7bLJJHxiD_I-Vj0dyW7d2ltGM6aQBUiFhWCsoso4	public/driver/1580669705878692277_image_cropper_1580669656988.jpg		t	f	f	t	010100000089CE328B5054254059C0046EDD3F5340	0	24	0
54	2020-02-29 21:49:42.376462+00	2020-02-29 21:49:42.376462+00	\N	Michael	381	616582159	2	TH1551	VIP	1	NK	Wagon	White	9981	public/driver/1583012981886596660_tenor.gif		public/driver/1583012981886828839_tenor.gif		f	f	f	f	\N	0	0	0
4	2020-02-03 10:47:21.963781+00	2020-02-03 10:47:21.963781+00	\N	Praveen	91	2747484848	2	MH15278888	Kwid	1	Taasai Go	2001	Black	TN 41 AT 1737	public/vehicle/1580726841296229381_image_cropper_1580726812211.jpg	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjowLCJ1c2VyX3R5cGUiOiJkcml2ZXIiLCJleHAiOjE1ODM0MjE3MTEsImlzcyI6InRhYXNhaSJ9.Jyi7bLJJHxiD_I-Vj0dyW7d2ltGM6aQBUiFhWCsoso4	public/driver/1580726841296226418_image_cropper_1580726788349.jpg		t	f	f	t	\N	0	0	0
2	2020-02-03 03:21:15.936594+00	2020-02-03 03:21:15.936594+00	\N	Praveen	91	9485748485	1		Kwid	1	Taasai Go	2001	white	TN 41 AR 2837	public/vehicle/1580700075263741703_image_cropper_1580699997617.jpg	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjowLCJ1c2VyX3R5cGUiOiJkcml2ZXIiLCJleHAiOjE1ODM0MjE3MTEsImlzcyI6InRhYXNhaSJ9.Jyi7bLJJHxiD_I-Vj0dyW7d2ltGM6aQBUiFhWCsoso4	public/driver/1580700075263737692_image_cropper_1580699981553.jpg		t	f	f	f	\N	0	0	0
9	2020-02-03 12:17:14.695274+00	2020-02-03 12:17:14.695274+00	\N	Praveen	91	8765432176	1	MH282893939	Asta	1	Taasai Go	2001	Ehite	TN 41 AR 8383	public/vehicle/1580732233984685528_image_cropper_1580732206896.jpg	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjowLCJ1c2VyX3R5cGUiOiJkcml2ZXIiLCJleHAiOjE1ODM0MjE3MTEsImlzcyI6InRhYXNhaSJ9.Jyi7bLJJHxiD_I-Vj0dyW7d2ltGM6aQBUiFhWCsoso4	public/driver/1580732233984682525_image_cropper_1580732190727.jpg		f	f	f	f	\N	0	0	0
10	2020-02-03 12:18:41.157762+00	2020-02-03 12:18:41.157762+00	\N	Praveen	91	8375858384	1	MH39399449	Asta	1	Taasai Go	2001	Black	TN 41 AR 2838	public/vehicle/1580732320446653924_image_cropper_1580732295571.jpg	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjowLCJ1c2VyX3R5cGUiOiJkcml2ZXIiLCJleHAiOjE1ODM0MjE3MTEsImlzcyI6InRhYXNhaSJ9.Jyi7bLJJHxiD_I-Vj0dyW7d2ltGM6aQBUiFhWCsoso4	public/driver/1580732320446650591_image_cropper_1580732278458.jpg		f	f	f	f	\N	0	0	0
5	2020-02-03 11:01:05.416364+00	2020-02-03 11:01:05.416364+00	\N	Praveen	91	8508008472	1	MH2838399449	Kwid	1	Taasai Go	2001	White	TN 41 AD 2833	public/vehicle/1580727664719052164_image_cropper_1580727640964.jpg	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjowLCJ1c2VyX3R5cGUiOiJkcml2ZXIiLCJleHAiOjE1ODM0MjE3MTEsImlzcyI6InRhYXNhaSJ9.Jyi7bLJJHxiD_I-Vj0dyW7d2ltGM6aQBUiFhWCsoso4	public/driver/1580727664719049649_image_cropper_1580727615437.jpg		f	f	f	f	\N	0	0	0
6	2020-02-03 11:22:36.179155+00	2020-02-03 11:22:36.179155+00	\N	Praveen	91	0987654321	1	MH1737282	Asta	1	Taasai Go	2001	black	TN 41 AC 2827	public/vehicle/1580728955508907343_image_cropper_1580728927425.jpg	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjowLCJ1c2VyX3R5cGUiOiJkcml2ZXIiLCJleHAiOjE1ODM0MjE3MTEsImlzcyI6InRhYXNhaSJ9.Jyi7bLJJHxiD_I-Vj0dyW7d2ltGM6aQBUiFhWCsoso4	public/driver/1580728955508904198_image_cropper_1580728910021.jpg		f	f	f	f	\N	0	0	0
8	2020-02-03 12:14:13.929073+00	2020-02-03 12:14:13.929073+00	\N	Praveen	91	8072981626	1	MH166737383	Asta	1	Taasai Go	2001	Red	TN 41 AC 1627	public/vehicle/1580732053218723648_image_cropper_1580732001229.jpg	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjowLCJ1c2VyX3R5cGUiOiJkcml2ZXIiLCJleHAiOjE1ODM0MjE3MTEsImlzcyI6InRhYXNhaSJ9.Jyi7bLJJHxiD_I-Vj0dyW7d2ltGM6aQBUiFhWCsoso4	public/driver/1580732053218721152_image_cropper_1580731983541.jpg		f	f	f	f	\N	0	0	0
11	2020-02-03 12:22:31.515737+00	2020-02-03 12:22:31.515737+00	\N	Praveen	91	9965498366	1	MY3889494	Asta	1	Taasai Go	2001	Black	TN 41 AR 2622	public/vehicle/1580732550804322496_image_cropper_1580732523399.jpg	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjowLCJ1c2VyX3R5cGUiOiJkcml2ZXIiLCJleHAiOjE1ODM0MjE3MTEsImlzcyI6InRhYXNhaSJ9.Jyi7bLJJHxiD_I-Vj0dyW7d2ltGM6aQBUiFhWCsoso4	public/driver/1580732550804319366_image_cropper_1580732504241.jpg		f	f	f	f	\N	0	0	0
12	2020-02-03 12:31:38.028221+00	2020-02-03 12:31:38.028221+00	\N	Praveen	91	9965498344	1	MH73378338	Asta	3	Taasai Premium	2001	Red	TN 41 BA 3829	public/vehicle/1580733097347042917_image_cropper_1580733069509.jpg	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjowLCJ1c2VyX3R5cGUiOiJkcml2ZXIiLCJleHAiOjE1ODM0MjE3MTEsImlzcyI6InRhYXNhaSJ9.Jyi7bLJJHxiD_I-Vj0dyW7d2ltGM6aQBUiFhWCsoso4	public/driver/1580733097347039653_image_cropper_1580733057093.jpg		f	f	f	f	\N	0	0	0
13	2020-02-03 12:33:59.439191+00	2020-02-03 12:33:59.439191+00	\N	Praveen	91	9876543215	1	MH46787543	Asra	1	Taasai Go	2009	Red	TN 41 AR 5677	public/vehicle/1580733238771374132_image_cropper_1580733188307.jpg	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjowLCJ1c2VyX3R5cGUiOiJkcml2ZXIiLCJleHAiOjE1ODM0MjE3MTEsImlzcyI6InRhYXNhaSJ9.Jyi7bLJJHxiD_I-Vj0dyW7d2ltGM6aQBUiFhWCsoso4	public/driver/1580733238771372243_image_cropper_1580733160134.jpg		f	f	f	f	\N	0	0	0
7	2020-02-03 12:11:48.277894+00	2020-02-26 13:11:40.063617+00	\N	Praveen	91	7654321098	1	MH833848483	Asta	1	Taasai Go	2001	White	TN 41 AC 2733	public/vehicle/1580731907566283572_image_cropper_1580729008464.jpg	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjowLCJ1c2VyX3R5cGUiOiJkcml2ZXIiLCJleHAiOjE1ODM0MjE3MTEsImlzcyI6InRhYXNhaSJ9.Jyi7bLJJHxiD_I-Vj0dyW7d2ltGM6aQBUiFhWCsoso4	public/driver/1580731907566281006_image_cropper_1580729008464.jpg		f	f	f	f	\N	0	0	0
3	2020-02-03 07:55:32.024651+00	2020-02-03 07:55:32.024651+00	\N	Praveen	91	9344664559	1	MH12274648484	Asta	1	Taasai Go	2001	White	TN 41 AR 2846	public/vehicle/1580716531333919662_image_cropper_1580716485190.jpg	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozLCJ1c2VyX3R5cGUiOiJkcml2ZXIiLCJleHAiOjE1ODQ3ODA5MjUsImlzcyI6InRhYXNhaSJ9.xPbt9nZhVTvPfv9XL-6zKay898ldZBFesyv_Aq2gBRQ	public/driver/1580716531333915499_image_cropper_1580716452086.jpg		t	f	f	t	0101000000A9328CBB415425400C3D62F4DC3F5340	0	0	0
\.


--
-- Data for Name: fares; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.fares (id, created_at, updated_at, deleted_at, vehicle_type_id, operator_id, base_fare, minimum_fare, waiting_time_limit, waiting_fee, cancellation_time_limit, cancellation_fee, duration_fare, distance_fare, tax, traffic_factor, is_active) FROM stdin;
2	2020-02-03 07:50:43.838008+00	2020-02-03 07:50:43.838008+00	\N	2	1	10	15	5	10	2	10	1	2	10	15	t
3	2020-02-03 07:53:40.634401+00	2020-02-03 07:53:40.634401+00	\N	3	1	50	75	5	5	5	30	2	5	20	25	t
4	2020-02-03 17:44:53.568945+00	2020-02-03 17:44:53.568945+00	\N	1	2	2.5	5	5	5	5	5	1	1	0	10	t
1	2020-02-02 17:07:10.470814+00	2020-02-02 17:07:10.470814+00	\N	1	1	20	20	2	2	2	20	20	20	10	20	t
5	2020-02-16 16:45:08.309225+00	2020-02-16 16:45:08.309225+00	\N	1	3	4	4	4	3	4	4	4	1	0	5	t
\.


--
-- Data for Name: fleets; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.fleets (id, created_at, updated_at, deleted_at, operator_id, name, is_active, email, password, balance, outgo_pending) FROM stdin;
29	2020-03-01 20:41:08.087411+00	2020-03-01 20:41:08.087411+00	\N	3	aaaaaaaaaaaaaaa	f	admin1541@taasai.com	$2a$14$qjlaQZLBZk5qFYrW/V2wB.z0Ndiyw58ih9N1f3UrI8nvuv6p7Gbw6	0	0
30	2020-03-02 11:00:09.838911+00	2020-03-02 11:00:09.838911+00	\N	2	slough	f	slough@taasai.com	$2a$14$lLpaf/SGw2ooVJwItx7mzeAU0oZ12U1I2WSkuhuGwcqXZ7vhkChI2	0	0
25	2020-03-01 20:38:57.23527+00	2020-03-01 21:04:53.950402+00	\N	3	bbbbbbbbb	t	admin@taasai.com	$2a$14$A8ibDQdVJjRPwuRw2HFX1eJMObzB5h0/5HalXXTtDxxUiKiO7Eefa	0	0
26	2020-03-01 20:39:29.05134+00	2020-03-01 20:41:52.723428+00	\N	3	aaaaaaaaaaaaaaa	t	admin1@taasai.com	$2a$14$8hipSth0XHrwAl3RYSAVXO1i9yOxIRI2Cq/3eD3FAWdRRZfNaEvv2	0	0
27	2020-03-01 20:39:54.139353+00	2020-03-01 20:39:54.139353+00	\N	3	aaaaaaaaaaaaaaa	f	admin15@taasai.com	$2a$14$8DmPgKeJVyLWm6VjyaPNje1sgqGxFTwJmAq6N6oWhEqR19COMGLuK	0	0
28	2020-03-01 20:40:54.627489+00	2020-03-01 20:40:54.627489+00	\N	3	aaaaaaaaaaaaaaa	t	admin151@taasai.com	$2a$14$PAjqLriF99wH9WR022BMkOMfb46Ijh96YkYNb0/wxn8CNlIBiRgmO	0	0
24	2020-03-01 20:36:38.231267+00	2020-03-01 20:36:38.231267+00	\N	3	aaaaaaaaaaaaaaa	t	admi111@taasai.com	$2a$14$qtjU6cJzHYE2OzxFR4xgwewLy1oKWQ3XdQD3Hubs2ZiUvRy0uldJK	112	0
\.


--
-- Data for Name: operators; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.operators (id, created_at, updated_at, deleted_at, name, location_name, email, password, platform_commission, operator_commission, driver_work_time, driver_rest_time, currency, auth_token, is_active, polygon, refer_amount, refer_type, balance, outgo_pending) FROM stdin;
2	\N	\N	\N	Taasai London	London	london@taasai.com	$2a$14$1GVkI05hQYZc15brlpsp3.z.Bvxq/XJjbD3adjgD5cFLNNJDj/v2e	15.000000	15.000000	8	6	£	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VyX3R5cGUiOiJvcGVyYXRvciIsImV4cCI6MTU4MDg0OTg3OSwiaXNzIjoidGFhc2FpIn0.0vm9zdL3kZsIZvhhlf4_yYMZVnKXi-RP8dXQkeczdmg	t	0103000000010000000F00000054E0641BB8DB4940B75F3E59315CD5BF58ACE122F7DA4940C9CA2F833122C13FD6C8AEB48CCA494061E124CD1FD3D23FF1D8CF6229A44940DE54A4C2D842C43F66A4DE5339A1494074982F2FC03EBEBFE5F21FD26FA94940E50E9BC8CC05D8BF0B7BDAE1AFA749408A39083A5AD5DABFBE6A65C22FAD4940D15B3CBCE7C0DFBF9468C9E369C549400C3B8C497FAFEABF1EC4CE143AD349406A865451BC4AE9BF849F38807ED149405F29CB10C7BAE5BF5D88D51F61D249403A3C84F1D338E3BFFF5D9F39EBD34940E1CE85915E54E1BF2FC03E3A75D549403FABCC94D6DFDEBF54E0641BB8DB4940B75F3E59315CD5BF	\N	\N	\N	\N
3	\N	\N	\N	Poland	Warsaw	poland@taasai.com	$2a$14$4tNh6RY4zQ48REYkUOLeAOdxUiyndrcONCqfrP5YOz1JSyITM/Q8u	10.000000	10.000000	5	6	£	\N	t	010300000001000000090000001EC539EAE8284A401310937021033540DE3B6A4C88254A4033C34659BF1D35407A53910A631D4A402B33A5F5B728354056F2B1BB40154A40C07971E2AB1D3540276C3F19E3114A40131093702103354056F2B1BB40154A4065A6B4FE96E834407A53910A631D4A405AF5B9DA8ADD3440DE3B6A4C88254A405265187783E834401EC539EAE8284A401310937021033540	12.000000	\N	\N	\N
1	\N	\N	\N	Pollachi Taxi Services	Pollachi	pollachi@taasai.com	$2a$14$fzYh32FtOXEgoxVDz8S9LeMZZsk1p9wOVBsIlkEjCg/Q04kghvmBm	10.000000	5.000000	8	8	$	\N	t	0103000000010000000D0000009DBAF2599E8F254061DD7877644053403B54539275782540809C306134465340D57B2AA73D452540D7F6764B72485340185B0872501A2540CBF3E0EEAC4553409EB47059850D25406D7022FAB54153403370404B57082540B3075A81213D5340B859BC5818122540009013268C37534097749483D93425406AD95A5F243353409A417C60C76F2540A5C002983231534086E3F90CA88725407F130A11703953409DBAF2599E8F254061DD7877644053409DBAF2599E8F254061DD7877644053409DBAF2599E8F254061DD787764405340	\N	\N	89	\N
\.


--
-- Data for Name: otps; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.otps (id, created_at, updated_at, deleted_at, dial_code, country_code, mobile_number, otp, is_used) FROM stdin;
1	2020-02-02 17:38:10.500297+00	2020-02-02 17:38:10.500297+00	\N	91	IN	7344664559	6374	t
2	2020-02-02 18:33:30.177351+00	2020-02-02 18:33:30.177351+00	\N	91	IN	9344336645	6218	f
3	2020-02-02 18:46:24.949915+00	2020-02-02 18:46:24.949915+00	\N	91	IN	9334686478	4088	f
4	2020-02-02 18:49:03.621444+00	2020-02-02 18:49:03.621444+00	\N	91	IN	8364727485	1891	t
5	2020-02-02 18:49:18.817215+00	2020-02-02 18:49:18.817215+00	\N	91	IN	8364727485	4596	t
6	2020-02-02 18:51:23.634538+00	2020-02-02 18:51:23.634538+00	\N	91	IN	8364727485	4108	t
7	2020-02-02 18:54:11.656984+00	2020-02-02 18:54:11.656984+00	\N	91	IN	8364727485	2642	t
8	2020-02-02 18:57:37.023625+00	2020-02-02 18:57:37.023625+00	\N	91	IN	9876543210	6674	t
9	2020-02-03 03:09:08.058529+00	2020-02-03 03:09:08.058529+00	\N	91	IN	9366847588	6374	t
10	2020-02-03 03:10:19.89536+00	2020-02-03 03:10:19.89536+00	\N	91	IN	9366847588	2326	t
11	2020-02-03 03:10:48.381083+00	2020-02-03 03:10:48.381083+00	\N	91	IN	9366847588	8434	t
12	2020-02-03 03:11:26.447407+00	2020-02-03 03:11:26.447407+00	\N	91	IN	9366847588	3513	t
13	2020-02-03 03:12:25.018595+00	2020-02-03 03:12:25.018595+00	\N	91	IN	9366847588	7414	t
14	2020-02-03 03:12:53.820767+00	2020-02-03 03:12:53.820767+00	\N	91	IN	9366847588	6218	t
15	2020-02-03 03:13:11.100311+00	2020-02-03 03:13:11.100311+00	\N	91	IN	9366847588	4088	t
16	2020-02-03 03:14:12.545727+00	2020-02-03 03:14:12.545727+00	\N	91	IN	9366847588	1891	t
17	2020-02-03 03:16:30.772265+00	2020-02-03 03:16:30.772265+00	\N	91	IN	9366847588	4596	f
18	2020-02-03 03:18:10.7875+00	2020-02-03 03:18:10.7875+00	\N	91	IN	9485748485	4108	t
19	2020-02-03 03:19:38.614353+00	2020-02-03 03:19:38.614353+00	\N	91	IN	9485748485	2642	t
20	2020-02-03 03:21:49.47955+00	2020-02-03 03:21:49.47955+00	\N	91	IN	9485748485	6674	t
21	2020-02-03 05:24:42.956933+00	2020-02-03 05:24:42.956933+00	\N	91	IN	9876543219	1293	t
23	2020-02-03 05:25:21.326364+00	2020-02-03 05:25:21.326364+00	\N	91	IN	8765432198	1628	f
24	2020-02-03 05:27:56.732508+00	2020-02-03 05:27:56.732508+00	\N	91	IN	8646965368	9167	f
22	2020-02-03 05:25:01.771202+00	2020-02-03 05:25:01.771202+00	\N	91	IN	9876543219	4927	t
25	2020-02-03 06:07:15.414137+00	2020-02-03 06:07:15.414137+00	\N	91	IN	9876543219	1753	t
27	2020-02-03 07:03:47.531603+00	2020-02-03 07:03:47.531603+00	\N	91	IN	8364727485	9290	t
26	2020-02-03 06:58:12.128933+00	2020-02-03 06:58:12.128933+00	\N	91	IN	9876543219	6389	t
30	2020-02-03 07:18:44.383217+00	2020-02-03 07:18:44.383217+00	\N	91	IN	8344664459	8582	f
31	2020-02-03 07:20:48.551918+00	2020-02-03 07:20:48.551918+00	\N	91	IN	8765432187	1711	f
32	2020-02-03 07:24:43.872085+00	2020-02-03 07:24:43.872085+00	\N	91	IN	8364727485	4079	t
29	2020-02-03 07:17:19.590417+00	2020-02-03 07:17:19.590417+00	\N	91	IN	9344664559	6982	t
33	2020-02-03 07:45:37.208449+00	2020-02-03 07:45:37.208449+00	\N	91	IN	9344664559	9289	t
34	2020-02-03 07:53:26.343358+00	2020-02-03 07:53:26.343358+00	\N	91	IN	9344664559	2162	t
35	2020-02-03 07:54:08.715242+00	2020-02-03 07:54:08.715242+00	\N	91	IN	9344664559	6156	t
36	2020-02-03 09:24:25.29432+00	2020-02-03 09:24:25.29432+00	\N	91	IN	0000000000	1980	f
37	2020-02-03 09:28:43.567723+00	2020-02-03 09:28:43.567723+00	\N	91	IN	000000000/	5760	f
38	2020-02-03 09:39:54.813859+00	2020-02-03 09:39:54.813859+00	\N	91	IN	9999999999	8215	f
39	2020-02-03 09:45:53.857217+00	2020-02-03 09:45:53.857217+00	\N	91	IN	9999776678	1049	f
40	2020-02-03 09:48:56.250235+00	2020-02-03 09:48:56.250235+00	\N	91	IN	0000063637	9024	f
41	2020-02-03 09:50:01.000385+00	2020-02-03 09:50:01.000385+00	\N	91	IN	9344667890	3412	f
42	2020-02-03 09:53:23.146703+00	2020-02-03 09:53:23.146703+00	\N	91	IN	8374748585	7763	f
43	2020-02-03 09:56:34.137349+00	2020-02-03 09:56:34.137349+00	\N	91	IN	9377447833	8581	f
44	2020-02-03 10:02:43.516774+00	2020-02-03 10:02:43.516774+00	\N	91	IN	8484848888	6638	f
45	2020-02-03 10:03:42.126568+00	2020-02-03 10:03:42.126568+00	\N	91	IN	8374838484	4348	f
46	2020-02-03 10:04:45.459389+00	2020-02-03 10:04:45.459389+00	\N	91	IN	8384848484	7692	f
47	2020-02-03 10:05:41.470327+00	2020-02-03 10:05:41.470327+00	\N	91	IN	9494948484	2086	f
48	2020-02-03 10:08:34.272144+00	2020-02-03 10:08:34.272144+00	\N	91	IN	9377374848	6622	f
49	2020-02-03 10:09:26.985971+00	2020-02-03 10:09:26.985971+00	\N	91	IN	8374748474	7615	f
50	2020-02-03 10:15:59.382251+00	2020-02-03 10:15:59.382251+00	\N	91	IN	8484848585	2455	f
51	2020-02-03 10:18:07.419025+00	2020-02-03 10:18:07.419025+00	\N	91	IN	1234567891	4127	f
52	2020-02-03 10:21:38.883821+00	2020-02-03 10:21:38.883821+00	\N	91	IN	8508742734	5433	f
53	2020-02-03 10:39:32.702422+00	2020-02-03 10:39:32.702422+00	\N	91	IN	9876598765	6374	f
54	2020-02-03 10:46:23.743499+00	2020-02-03 10:46:23.743499+00	\N	91	IN	2747484848	6374	t
56	2020-02-03 10:51:29.843518+00	2020-02-03 10:51:29.843518+00	\N	91	IN	9237474844	8434	f
57	2020-02-03 11:00:12.505119+00	2020-02-03 11:00:12.505119+00	\N	91	IN	8508008472	6374	t
58	2020-02-03 11:02:25.892698+00	2020-02-03 11:02:25.892698+00	\N	91	IN	49438484848	2326	f
28	2020-02-03 07:14:29.995898+00	2020-02-03 07:14:29.995898+00	\N	91	IN	9876543219	6260	t
60	2020-02-03 11:21:44.749202+00	2020-02-03 11:21:44.749202+00	\N	91	IN	0987654321	2326	t
61	2020-02-03 11:23:22.821366+00	2020-02-03 11:23:22.821366+00	\N	91	IN	7654321098	8434	t
62	2020-02-03 12:12:57.239354+00	2020-02-03 12:12:57.239354+00	\N	91	IN	8072981626	6374	t
63	2020-02-03 12:16:24.113294+00	2020-02-03 12:16:24.113294+00	\N	91	IN	8765432176	2326	t
64	2020-02-03 12:17:53.51465+00	2020-02-03 12:17:53.51465+00	\N	91	IN	8375858384	8434	t
65	2020-02-03 12:21:39.527807+00	2020-02-03 12:21:39.527807+00	\N	91	IN	9965498366	3513	t
66	2020-02-03 12:27:50.517406+00	2020-02-03 12:27:50.517406+00	\N	91	IN	9484848484	7414	f
67	2020-02-03 12:28:44.63403+00	2020-02-03 12:28:44.63403+00	\N	91	IN	9876543216	6218	f
68	2020-02-03 12:30:54.174346+00	2020-02-03 12:30:54.174346+00	\N	91	IN	9965498344	4088	t
55	2020-02-03 10:48:53.024856+00	2020-02-03 10:48:53.024856+00	\N	91	IN	9876543215	2326	t
69	2020-02-03 12:32:32.825936+00	2020-02-03 12:32:32.825936+00	\N	91	IN	9876543215	1891	t
70	2020-02-03 16:10:42.81437+00	2020-02-03 16:10:42.81437+00	\N	91	IN	987653636367	4596	t
71	2020-02-03 16:17:04.935557+00	2020-02-03 16:17:04.935557+00	\N	91	IN	93838383883	4108	t
72	2020-02-03 16:30:21.448354+00	2020-02-03 16:30:21.448354+00	\N	91	IN	9494949484	2642	t
73	2020-02-03 16:34:35.143788+00	2020-02-03 16:34:35.143788+00	\N	91	IN	986657876667	6674	f
74	2020-02-03 16:35:25.418087+00	2020-02-03 16:35:25.418087+00	\N	91	IN	977567886677	1293	t
75	2020-02-03 17:23:30.816546+00	2020-02-03 17:23:30.816546+00	\N	91	IN	123456789	9167	t
76	2020-02-03 18:23:44.333837+00	2020-02-03 18:23:44.333837+00	\N	91	IN	123456789	6982	t
77	2020-02-03 18:53:04.972209+00	2020-02-03 18:53:04.972209+00	\N	91	IN	12345678	4079	t
78	2020-02-03 19:12:22.58726+00	2020-02-03 19:12:22.58726+00	\N	91	IN	12345678	5994	f
79	2020-02-03 22:34:18.913782+00	2020-02-03 22:34:18.913782+00	\N	44	UK	12345678	5841	f
80	2020-02-03 22:35:10.941984+00	2020-02-03 22:35:10.941984+00	\N	91	IN	123456677	8877	f
81	2020-02-04 03:53:58.185462+00	2020-02-04 03:53:58.185462+00	\N	91	IN	98383838383	1186	t
82	2020-02-04 05:51:11.191372+00	2020-02-04 05:51:11.191372+00	\N	91	IN	98737373772	9289	t
83	2020-02-04 13:10:15.955193+00	2020-02-04 13:10:15.955193+00	\N	91	IN	7777788888	2277	f
84	2020-02-04 14:29:08.029286+00	2020-02-04 14:29:08.029286+00	\N	91	IN	9377448844	3543	f
85	2020-02-04 14:31:06.666258+00	2020-02-04 14:31:06.666258+00	\N	91	IN	9344770025	2162	t
86	2020-02-04 15:09:56.750202+00	2020-02-04 15:09:56.750202+00	\N	91	IN	9374627384	8169	f
87	2020-02-04 15:13:53.3322+00	2020-02-04 15:13:53.3322+00	\N	91	IN	9383838473	4743	f
88	2020-02-04 15:15:14.17952+00	2020-02-04 15:15:14.17952+00	\N	91	IN	8474748383	6156	f
89	2020-02-04 15:16:10.070259+00	2020-02-04 15:16:10.070259+00	\N	91	IN	8393939383	5316	t
90	2020-02-04 15:24:46.643664+00	2020-02-04 15:24:46.643664+00	\N	91	IN	5566676655	6679	f
91	2020-02-04 15:34:12.876726+00	2020-02-04 15:34:12.876726+00	\N	91	IN	6855677678	3739	f
92	2020-02-04 15:35:08.285582+00	2020-02-04 15:35:08.285582+00	\N	91	IN	7867788765	8626	f
93	2020-02-04 15:37:08.099719+00	2020-02-04 15:37:08.099719+00	\N	91	IN	7777666677	1980	f
94	2020-02-04 15:45:26.134159+00	2020-02-04 15:45:26.134159+00	\N	91	IN	6789999988	5760	f
95	2020-02-04 15:47:29.079076+00	2020-02-04 15:47:29.079076+00	\N	91	IN	5566666677	8215	f
96	2020-02-04 15:49:26.119352+00	2020-02-04 15:49:26.119352+00	\N	91	IN	4567776666	1049	f
97	2020-02-04 15:50:53.959415+00	2020-02-04 15:50:53.959415+00	\N	91	IN	4666777777	9024	f
98	2020-02-04 15:51:20.305884+00	2020-02-04 15:51:20.305884+00	\N	91	IN	4678998876	3412	f
99	2020-02-04 15:54:45.619953+00	2020-02-04 15:54:45.619953+00	\N	91	IN	4667777765	7763	f
100	2020-02-04 15:56:04.608443+00	2020-02-04 15:56:04.608443+00	\N	91	IN	4567888986	8581	f
101	2020-02-04 15:57:02.500548+00	2020-02-04 15:57:02.500548+00	\N	91	IN	4568986533	6638	f
102	2020-02-04 16:02:17.578463+00	2020-02-04 16:02:17.578463+00	\N	91	IN	4678975467	4348	f
103	2020-02-04 16:03:46.939387+00	2020-02-04 16:03:46.939387+00	\N	91	IN	6677896436	7692	f
107	2020-02-05 03:19:54.627163+00	2020-02-05 03:19:54.627163+00	\N	91	IN	7964688644	2455	t
108	2020-02-05 03:23:56.164238+00	2020-02-05 03:23:56.164238+00	\N	91	IN	9754257788	4127	f
109	2020-02-05 03:24:12.856689+00	2020-02-05 03:24:12.856689+00	\N	91	IN	83399393939	5433	f
110	2020-02-05 03:25:05.712625+00	2020-02-05 03:25:05.712625+00	\N	91	IN	9344722893	2934	f
104	2020-02-04 16:05:12.670025+00	2020-02-04 16:05:12.670025+00	\N	91	IN	6889546797	2086	f
105	2020-02-04 16:12:24.55333+00	2020-02-04 16:12:24.55333+00	\N	91	IN	6789668895	6622	f
106	2020-02-05 03:06:23.51327+00	2020-02-05 03:06:23.51327+00	\N	91	IN	9678966675	7615	f
111	2020-02-05 03:26:47.05914+00	2020-02-05 03:26:47.05914+00	\N	91	IN	83829292822	6374	t
112	2020-02-05 03:29:11.001707+00	2020-02-05 03:29:11.001707+00	\N	91	IN	8754347	2326	f
113	2020-02-05 03:33:50.632527+00	2020-02-05 03:33:50.632527+00	\N	91	IN	6654433455	8434	f
114	2020-02-05 03:38:44.277234+00	2020-02-05 03:38:44.277234+00	\N	91	IN	5466776	6374	f
115	2020-02-05 03:41:14.390046+00	2020-02-05 03:41:14.390046+00	\N	91	IN	4335665444	2326	f
116	2020-02-05 03:43:34.230132+00	2020-02-05 03:43:34.230132+00	\N	91	IN	9384849494	6374	f
117	2020-02-05 03:46:49.693749+00	2020-02-05 03:46:49.693749+00	\N	91	IN	6776665555	2326	f
118	2020-02-05 03:47:18.850064+00	2020-02-05 03:47:18.850064+00	\N	91	IN	7776665555	8434	f
119	2020-02-05 03:56:22.099027+00	2020-02-05 03:56:22.099027+00	\N	91	IN	7475757575	6374	f
120	2020-02-05 04:00:28.86866+00	2020-02-05 04:00:28.86866+00	\N	91	IN	8595949	2326	f
121	2020-02-05 04:04:19.252615+00	2020-02-05 04:04:19.252615+00	\N	91	IN	9876456789	8434	f
122	2020-02-05 04:06:41.398498+00	2020-02-05 04:06:41.398498+00	\N	91	IN	9495949494	3513	t
123	2020-02-05 04:17:40.777333+00	2020-02-05 04:17:40.777333+00	\N	91	IN	9876446865	7414	f
124	2020-02-05 04:18:18.837716+00	2020-02-05 04:18:18.837716+00	\N	91	IN	7899988777	6218	f
125	2020-02-05 04:19:42.228306+00	2020-02-05 04:19:42.228306+00	\N	91	IN	9876543329	4088	f
127	2020-02-05 04:21:21.350438+00	2020-02-05 04:21:21.350438+00	\N	91	IN	8977986678	4596	f
128	2020-02-05 04:22:37.566677+00	2020-02-05 04:22:37.566677+00	\N	91	IN	8678908656	4108	f
129	2020-02-05 04:23:35.926805+00	2020-02-05 04:23:35.926805+00	\N	91	IN	78765666	2642	f
130	2020-02-05 04:24:11.558674+00	2020-02-05 04:24:11.558674+00	\N	91	IN	7088678877	6674	f
131	2020-02-05 04:28:24.559843+00	2020-02-05 04:28:24.559843+00	\N	91	IN	74474848484	1293	f
132	2020-02-05 04:29:59.308681+00	2020-02-05 04:29:59.308681+00	\N	91	IN	9876688899	4927	t
133	2020-02-05 04:30:31.786026+00	2020-02-05 04:30:31.786026+00	\N	91	IN	9876688899	1628	t
134	2020-02-05 04:36:10.510121+00	2020-02-05 04:36:10.510121+00	\N	91	IN	9876688899	9167	f
135	2020-02-05 04:37:01.118393+00	2020-02-05 04:37:01.118393+00	\N	91	IN	8474747474	1753	t
136	2020-02-05 04:37:08.491356+00	2020-02-05 04:37:08.491356+00	\N	91	IN	8474747474	6389	t
137	2020-02-05 04:38:58.995676+00	2020-02-05 04:38:58.995676+00	\N	91	IN	8474747474	6374	t
138	2020-02-05 04:40:20.679204+00	2020-02-05 04:40:20.679204+00	\N	91	IN	8474747474	2326	t
139	2020-02-05 04:42:02.731156+00	2020-02-05 04:42:02.731156+00	\N	91	IN	8474747474	8434	t
140	2020-02-05 04:52:14.388028+00	2020-02-05 04:52:14.388028+00	\N	91	IN	9344664559	6374	t
141	2020-02-05 07:03:02.110481+00	2020-02-05 07:03:02.110481+00	\N	91	IN	987646373738	6374	f
59	2020-02-03 11:12:49.954252+00	2020-02-03 11:12:49.954252+00	\N	91	IN	9876543219	6374	t
142	2020-02-05 07:26:27.058268+00	2020-02-05 07:26:27.058268+00	\N	91	IN	9876543219	6374	f
143	2020-02-05 07:34:36.832294+00	2020-02-05 07:34:36.832294+00	\N	91	IN	9586447747	2326	f
144	2020-02-05 07:40:40.081721+00	2020-02-05 07:40:40.081721+00	\N	91	IN	9775567878	8434	f
145	2020-02-05 07:46:31.692622+00	2020-02-05 07:46:31.692622+00	\N	91	IN	850774737373	3513	f
146	2020-02-05 07:51:54.312379+00	2020-02-05 07:51:54.312379+00	\N	91	IN	9344664551	7414	t
147	2020-02-05 07:51:59.754086+00	2020-02-05 07:51:59.754086+00	\N	91	IN	9344664551	6218	t
148	2020-02-05 07:52:23.116342+00	2020-02-05 07:52:23.116342+00	\N	91	IN	9344664551	4088	t
149	2020-02-05 07:53:11.451374+00	2020-02-05 07:53:11.451374+00	\N	91	IN	9344664551	1891	t
150	2020-02-05 07:55:01.597692+00	2020-02-05 07:55:01.597692+00	\N	91	IN	9344664551	4596	t
151	2020-02-05 07:56:51.361438+00	2020-02-05 07:56:51.361438+00	\N	91	IN	9344664551	4108	f
152	2020-02-05 07:59:49.786445+00	2020-02-05 07:59:49.786445+00	\N	91	IN	9809878909	2642	t
153	2020-02-05 08:02:18.641396+00	2020-02-05 08:02:18.641396+00	\N	91	IN	9809878909	6674	f
154	2020-02-05 08:04:27.162051+00	2020-02-05 08:04:27.162051+00	\N	91	IN	9090909090	1293	t
155	2020-02-05 08:06:32.181242+00	2020-02-05 08:06:32.181242+00	\N	91	IN	9090909090	4927	t
156	2020-02-05 08:07:34.55553+00	2020-02-05 08:07:34.55553+00	\N	91	IN	9090909090	1628	t
157	2020-02-05 08:10:01.076239+00	2020-02-05 08:10:01.076239+00	\N	91	IN	9090909090	9167	t
158	2020-02-05 08:10:35.319314+00	2020-02-05 08:10:35.319314+00	\N	91	IN	9090909090	1753	f
159	2020-02-05 08:11:45.244803+00	2020-02-05 08:11:45.244803+00	\N	91	IN	9878987898	6389	t
160	2020-02-05 08:12:35.109597+00	2020-02-05 08:12:35.109597+00	\N	91	IN	9878987898	9290	f
161	2020-02-05 08:12:58.83849+00	2020-02-05 08:12:58.83849+00	\N	91	IN	984848484	6260	t
162	2020-02-05 08:14:03.298515+00	2020-02-05 08:14:03.298515+00	\N	91	IN	984848484	6982	t
126	2020-02-05 04:20:14.406699+00	2020-02-05 04:20:14.406699+00	\N	91	IN	9876543217	1891	t
163	2020-02-05 08:26:34.810548+00	2020-02-05 08:26:34.810548+00	\N	91	IN	9876543217	8582	t
164	2020-02-05 08:28:39.401383+00	2020-02-05 08:28:39.401383+00	\N	91	IN	9876543217	1711	f
165	2020-02-05 08:31:42.651806+00	2020-02-05 08:31:42.651806+00	\N	91	IN	9878987678	4079	t
166	2020-02-05 08:38:52.482165+00	2020-02-05 08:38:52.482165+00	\N	91	IN	9098909890	5994	t
168	2020-02-05 08:40:20.129805+00	2020-02-05 08:40:20.129805+00	\N	91	IN	8679876986	9289	t
169	2020-02-05 08:40:41.073896+00	2020-02-05 08:40:41.073896+00	\N	91	IN	8679876986	2277	t
170	2020-02-05 08:40:53.22288+00	2020-02-05 08:40:53.22288+00	\N	91	IN	8679876986	3543	t
171	2020-02-05 08:49:40.44322+00	2020-02-05 08:49:40.44322+00	\N	91	IN	9876543210	2162	t
172	2020-02-05 09:00:27.016602+00	2020-02-05 09:00:27.016602+00	\N	91	IN	8909890989	6156	t
173	2020-02-05 09:04:27.267916+00	2020-02-05 09:04:27.267916+00	\N	91	IN	8787878787	5316	t
174	2020-02-05 09:07:29.980579+00	2020-02-05 09:07:29.980579+00	\N	91	IN	8787878787	6679	t
175	2020-02-05 09:09:12.444141+00	2020-02-05 09:09:12.444141+00	\N	91	IN	8787878787	3739	t
176	2020-02-05 09:10:13.793269+00	2020-02-05 09:10:13.793269+00	\N	91	IN	8787878787	8626	t
177	2020-02-05 09:11:30.561699+00	2020-02-05 09:11:30.561699+00	\N	91	IN	8787878787	1980	t
178	2020-02-05 09:11:59.518967+00	2020-02-05 09:11:59.518967+00	\N	91	IN	8787878787	5760	t
179	2020-02-05 09:13:11.523137+00	2020-02-05 09:13:11.523137+00	\N	91	IN	8787878787	8215	t
181	2020-02-05 09:15:59.682755+00	2020-02-05 09:15:59.682755+00	\N	91	IN	9344667899	9024	f
182	2020-02-05 09:18:35.200868+00	2020-02-05 09:18:35.200868+00	\N	91	IN	8089808980	3412	t
183	2020-02-05 09:24:53.339325+00	2020-02-05 09:24:53.339325+00	\N	91	IN	9898989898	7763	t
167	2020-02-05 08:39:18.802191+00	2020-02-05 08:39:18.802191+00	\N	91	IN	9098909890	5841	t
184	2020-02-05 09:26:30.944957+00	2020-02-05 09:26:30.944957+00	\N	91	IN	9098909890	8581	t
185	2020-02-05 09:40:45.348396+00	2020-02-05 09:40:45.348396+00	\N	91	IN	8072981623	6638	t
186	2020-02-05 09:47:29.188559+00	2020-02-05 09:47:29.188559+00	\N	91	IN	8090898768	4348	t
187	2020-02-05 09:48:34.886957+00	2020-02-05 09:48:34.886957+00	\N	91	IN	8090898768	7692	t
188	2020-02-05 09:48:55.388473+00	2020-02-05 09:48:55.388473+00	\N	91	IN	8090898768	2086	t
189	2020-02-05 09:51:12.990136+00	2020-02-05 09:51:12.990136+00	\N	91	IN	8090898768	6622	f
190	2020-02-05 09:52:13.023217+00	2020-02-05 09:52:13.023217+00	\N	91	IN	9798964689	7615	f
191	2020-02-05 09:59:07.099617+00	2020-02-05 09:59:07.099617+00	\N	91	IN	8901235678	5433	f
192	2020-02-05 10:01:36.534871+00	2020-02-05 10:01:36.534871+00	\N	91	IN	9898909898	2337	f
193	2020-02-05 10:04:25.219674+00	2020-02-05 10:04:25.219674+00	\N	91	IN	9876567809	4695	f
194	2020-02-05 10:10:18.758617+00	2020-02-05 10:10:18.758617+00	\N	91	IN	9890989094	4918	f
197	2020-02-05 10:27:02.578719+00	2020-02-05 10:27:02.578719+00	\N	91	IN	9898909543	8818	t
198	2020-02-05 10:29:03.075952+00	2020-02-05 10:29:03.075952+00	\N	91	IN	9898909543	3214	t
199	2020-02-05 10:34:47.299406+00	2020-02-05 10:34:47.299406+00	\N	91	IN	9898909543	1001	t
180	2020-02-05 09:14:18.408647+00	2020-02-05 09:14:18.408647+00	\N	91	IN	8787878787	1049	t
203	2020-02-05 12:02:19.400588+00	2020-02-05 12:02:19.400588+00	\N	91	IN	8787868678	3453	f
204	2020-02-05 12:10:02.448061+00	2020-02-05 12:10:02.448061+00	\N	91	IN	7676767676	3299	f
205	2020-02-06 06:39:03.362981+00	2020-02-06 06:39:03.362981+00	\N	91	IN	9890989098	9498	f
206	2020-02-06 06:40:20.143156+00	2020-02-06 06:40:20.143156+00	\N	91	IN	9890989093	7438	f
208	2020-02-06 06:45:29.96184+00	2020-02-06 06:45:29.96184+00	\N	91	IN	8989899098	5340	f
195	2020-02-05 10:11:55.392353+00	2020-02-05 10:11:55.392353+00	\N	91	IN	8764789064	4049	f
196	2020-02-05 10:18:26.995376+00	2020-02-05 10:18:26.995376+00	\N	91	IN	8642567809	6390	f
200	2020-02-05 10:34:54.718112+00	2020-02-05 10:34:54.718112+00	\N	91	IN	9898909543	4987	t
201	2020-02-05 11:54:13.748969+00	2020-02-05 11:54:13.748969+00	\N	91	IN	98765443214	5419	f
202	2020-02-05 12:00:27.090503+00	2020-02-05 12:00:27.090503+00	\N	91	IN	8787878787	7124	f
207	2020-02-06 06:42:02.946649+00	2020-02-06 06:42:02.946649+00	\N	91	IN	9876546789	8090	f
209	2020-02-06 10:22:33.126059+00	2020-02-06 10:22:33.126059+00	\N	91	IN	9879879767	1521	t
210	2020-02-07 03:06:16.491612+00	2020-02-07 03:06:16.491612+00	\N	91	IN	9809654324	6695	t
211	2020-02-07 03:11:56.276788+00	2020-02-07 03:11:56.276788+00	\N	91	IN	9809654324	2205	f
212	2020-02-07 03:16:17.844829+00	2020-02-07 03:16:17.844829+00	\N	91	IN	9090989099	6532	f
213	2020-02-07 03:20:30.660328+00	2020-02-07 03:20:30.660328+00	\N	91	IN	4689975678	3937	t
214	2020-02-07 03:22:57.261321+00	2020-02-07 03:22:57.261321+00	\N	91	IN	9344664559	5040	t
215	2020-02-07 03:23:02.446479+00	2020-02-07 03:23:02.446479+00	\N	91	IN	9344664559	2441	t
216	2020-02-07 06:47:21.274141+00	2020-02-07 06:47:21.274141+00	\N	91	IN	9898989845	7651	t
217	2020-02-07 06:50:16.817142+00	2020-02-07 06:50:16.817142+00	\N	91	IN	9344664559	7422	t
218	2020-02-11 04:09:42.12181+00	2020-02-11 04:09:42.12181+00	\N	91	IN	9344664559	6982	t
219	2020-02-12 09:05:48.156537+00	2020-02-12 09:05:48.156537+00	\N	91	IN	8765432245	1753	t
220	2020-02-13 03:29:00.881768+00	2020-02-13 03:29:00.881768+00	\N	91	IN	9344664559	6389	t
221	2020-02-13 03:32:13.128082+00	2020-02-13 03:32:13.128082+00	\N	91	IN	9344664559	9290	t
222	2020-02-13 06:54:05.1853+00	2020-02-13 06:54:05.1853+00	\N	91	IN	9344664559	6260	t
224	2020-02-13 07:15:00.312102+00	2020-02-13 07:15:00.312102+00	\N	91	IN	9965498366	2326	t
225	2020-02-13 07:16:53.55497+00	2020-02-13 07:16:53.55497+00	\N	91	IN	9965498366	8434	t
226	2020-02-13 07:19:59.413876+00	2020-02-13 07:19:59.413876+00	\N	91	IN	9965498366	3513	t
227	2020-02-13 07:21:42.023799+00	2020-02-13 07:21:42.023799+00	\N	91	IN	9965498366	7414	t
228	2020-02-13 07:23:25.071945+00	2020-02-13 07:23:25.071945+00	\N	91	IN	9965498366	6218	t
229	2020-02-13 07:26:51.860256+00	2020-02-13 07:26:51.860256+00	\N	91	IN	9965498366	4088	t
230	2020-02-13 07:55:57.912078+00	2020-02-13 07:55:57.912078+00	\N	91	IN	9965498366	1293	t
232	2020-02-13 08:00:54.862298+00	2020-02-13 08:00:54.862298+00	\N	91	IN	9965498377	1628	f
231	2020-02-13 07:58:47.821451+00	2020-02-13 07:58:47.821451+00	\N	91	IN	9965498366	4927	t
233	2020-02-13 08:01:06.714553+00	2020-02-13 08:01:06.714553+00	\N	91	IN	9965498366	9167	f
234	2020-02-13 19:08:04.18506+00	2020-02-13 19:08:04.18506+00	\N	44	UK	7729286321	6374	t
235	2020-02-16 13:16:26.180875+00	2020-02-16 13:16:26.180875+00	\N	91	IN	123123	4088	t
236	2020-02-16 14:09:25.397721+00	2020-02-16 14:09:25.397721+00	\N	91	IN	235671232312	1891	t
237	2020-02-16 15:25:42.220666+00	2020-02-16 15:25:42.220666+00	\N	91	IN	123456879	4927	t
238	2020-02-16 17:08:34.96226+00	2020-02-16 17:08:34.96226+00	\N	91	IN	12342134	6260	t
239	2020-02-16 20:52:15.129989+00	2020-02-16 20:52:15.129989+00	\N	91	IN	126566777	2162	t
240	2020-02-17 11:23:18.282672+00	2020-02-17 11:23:18.282672+00	\N	91	IN	8526783789	3739	t
241	2020-02-19 03:30:44.174055+00	2020-02-19 03:30:44.174055+00	\N	1	US	8163769637	8626	t
242	2020-02-19 04:10:58.782341+00	2020-02-19 04:10:58.782341+00	\N	1	US	8163769637	1980	t
223	2020-02-13 07:06:48.291447+00	2020-02-13 07:06:48.291447+00	\N	91	IN	9344664559	6374	t
243	2020-02-20 08:55:16.506418+00	2020-02-20 08:55:16.506418+00	\N	91	IN	9344664559	5760	t
244	2020-02-20 19:46:47.014629+00	2020-02-20 19:46:47.014629+00	\N	91	IN	245685264	4348	f
245	2020-02-20 20:28:54.258958+00	2020-02-20 20:28:54.258958+00	\N	44	UK	07421346254	7692	f
246	2020-02-20 20:29:32.550934+00	2020-02-20 20:29:32.550934+00	\N	44	UK	12345678901	2086	t
247	2020-02-20 20:32:39.022759+00	2020-02-20 20:32:39.022759+00	\N	44	UK	12345678901	6622	f
248	2020-02-20 20:35:45.915756+00	2020-02-20 20:35:45.915756+00	\N	44	UK	7729286333	7615	f
249	2020-02-20 20:58:12.262467+00	2020-02-20 20:58:12.262467+00	\N	91	IN	7729286333	2455	f
250	2020-02-27 05:28:08.320856+00	2020-02-27 05:28:08.320856+00	\N	91	IN	9344669951	6374	f
\.


--
-- Data for Name: passengers; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.passengers (id, created_at, updated_at, deleted_at, name, dial_code, country_code, mobile_number, auth_token, image, fcm_id, is_active, referral_code, referred_by, wallet_balance, balance, income_pending) FROM stdin;
3	2020-02-03 17:23:39.297791+00	2020-02-03 17:23:39.297791+00	\N	mike	91	IN	123456789	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozLCJ1c2VyX3R5cGUiOiJwYXNzZW5nZXIifQ.ndbLjHbirFtERkQAo-RYjouEP-anAPcXWCv-vDBZtN0		eT7rnU7nUes:APA91bFfSwbz7iNZMHdXsIYEXumRqSJSmHDdl6Q8TettIQcpsbmQ0Z_ctyJQ5OvVerhVHPaCLXqsSRUeD3EsjUFchmolMAhrrURYL0lpzXCK8Eb7ahesMG7uyIEiQKxMcTlIRZjIuD6q	t	\N	\N	\N	101	\N
7	2020-02-12 09:06:50.93725+00	2020-02-12 09:06:50.93725+00	\N	praveen	91	IN	8765432245	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo3LCJ1c2VyX3R5cGUiOiJwYXNzZW5nZXIiLCJleHAiOjE1ODQwOTA0MTEsImlzcyI6Im9ucmlkZSJ9.Ua6VJCfyjjBm7F7tfnEmISJeSizXif15vjKQA7_sQh8		e1fD-Yj1m30:APA91bGiaWTIjUO3FpCCMQMWHy1V_eXTHuv0x_2nSSgBV9t8inD4dMeuja5yUvPTLy4sGajW1oSPNFDbbBi7tYRGnWpMPwac4rFO8yL5-m8IhRRMdgPOdCFYJFO2xE1wbkejnECb4ySd	t	\N	\N	\N	\N	\N
1	2020-02-02 17:38:19.974761+00	2020-02-02 17:38:19.974761+00	\N	Praveen	91	IN	7344664559	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX3R5cGUiOiJwYXNzZW5nZXIiLCJleHAiOjE1ODMyNTcxMDAsImlzcyI6Im9ucmlkZSJ9.Og6AdQbJfV0COjiiAputMP4CUL7YlxQjxuxcMkgYX-w		fpGo46M3mu0:APA91bFBEOICWPfX1gfF20f8mWneaPVrrKlz-2yDlolhW2IrroYP-OvR8GClUtxDX-5tOCEu6okjQEAljCsNxLZNjwok6idyB1YKVLkbGdAssRyEUqiKmBxUQcp_pnl-Qwya01L24tNF	t	\N	\N	280	\N	\N
13	2020-02-16 20:52:48.059564+00	2020-02-16 20:52:48.059564+00	\N	mike	91	IN	126566777	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMywidXNlcl90eXBlIjoicGFzc2VuZ2VyIiwiZXhwIjoxNTg0NDc4MzY4LCJpc3MiOiJvbnJpZGUifQ.8D2EwofqZ055pgpEbDJlsSw0MNNQfmqKTYVZ1HkLIIg		fbw3aSx8NHE:APA91bHouDjWuAjIrstVtsesR7osMkiU9iPWxNug5qoj6nWEQY1M87eeX31FhvtAoQlECNcJ1C5dSRhxCSziEQJA5q1Z_uaYNvgnsimWNSqpFdLm7Qr7JdR_Gkd9EQhCt9jtXFa9pq9x	t	13MGEGM	mike	0	\N	\N
5	2020-02-04 05:51:24.455728+00	2020-02-04 05:51:24.455728+00	\N	Test Passenger	91	IN	98737373772	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJ1c2VyX3R5cGUiOiJwYXNzZW5nZXIiLCJleHAiOjE1ODMzODc0ODQsImlzcyI6Im9ucmlkZSJ9.MbU8mhu-yLpiKZmNxXLZgJQ-XysmTAhp4ErgOpwfZII		eZB0yaqZPKs:APA91bG4MprDuPJCXbG_zFhHqREo7Iq9Gqj6Nfs-_rnhSylD5RZ0JTokjRrxCAz0RAgXrcTd5ydYD1zU3NwysmsQV2N125sF3boyUZEFjkNG86H_s87n2KXSxcBFDLL51WtB08dENodu	t	\N	\N	\N	\N	\N
2	2020-02-02 18:57:55.696994+00	2020-02-02 18:57:55.696994+00	\N	John Michael	91	IN	9876543210	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VyX3R5cGUiOiJwYXNzZW5nZXIifQ.duDaELxbo1sGBcfH3sQwCgeTvvHu1lqBy0bUP1EqmRY		fJnt8LBZeN8:APA91bHY-CwrtKkng7QVcq5Kgft1Ld3wpd3MlWl7SaAGpccY9kkjbWe3SuU0topTCN_kBEy3wWX6vWNFwmoialThWHRWK6MaNM1mJrxW7oy5hOQzDyyrdaoG3h0x680E1XRKhrKeFtNx	t	\N	\N	100	\N	\N
4	2020-02-04 03:54:09.404192+00	2020-02-04 03:54:09.404192+00	\N	Test Passenger	91	IN	98383838383	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0LCJ1c2VyX3R5cGUiOiJwYXNzZW5nZXIiLCJleHAiOjE1ODMzODA0NDksImlzcyI6Im9ucmlkZSJ9.Kcei_T2-VAVEkzz7_lMzOjH92HXExfHc2zsonP3EK-w		e2u2IRjNb0Q:APA91bHu8WncgrSe0y28Ld5m5hCXMawCR2dyfuBEF_A9Xx2X6e9rqHebciOVh3swJ1n2A4EMmINlRy7pAwU4fdGq3dh25nm1YuacCG6t3QVfT5aeNAuAV2qPUsH4G--m6vP-jFxgRJR6	t	\N	\N	\N	\N	\N
6	2020-02-06 10:22:42.391664+00	2020-02-06 10:22:42.391664+00	\N	sdfddfsdf	91	IN	9879879767	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VyX3R5cGUiOiJwYXNzZW5nZXIiLCJleHAiOjE1ODM1NzY1NjIsImlzcyI6Im9ucmlkZSJ9.qfKqSoo7i3vPbfUuupnfrmVsypf9DPdnwuwE8UtOwPM		dIJ6gYqBC2I:APA91bHNY4nyz0S2Tlh5LoMug3xXpqpp02rFpSA6UsEj0j-UKrFsraW98I5MrI_nUjlcppkBU1-IASxeVNapFLaPjUFCOJUGGqm-_RZChiAKSBkD6_gCb7D5-i-y6jNApSA7r0tnPdU9	t	6NIWDY	\N	\N	\N	\N
8	2020-02-13 07:27:12.216839+00	2020-02-13 07:27:12.216839+00	\N	Praveen	91	IN	9965498366	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo4LCJ1c2VyX3R5cGUiOiJwYXNzZW5nZXIifQ.Sn6KMhWrB2ZP85FxZ6ZQ_KYqMxt4KAc4LyLitcpufEg		fAj82MsgR_w:APA91bH1v1xVEgcTXgFN3V7ieQPMNUXlZuR0vkhVwmJwkB5rZWLzxEJAxn6J5GKASoAk0GDLbO30RXmZ2pnW6r9VPgKDx-qjfcS9_KFGNk4Rvv5_dza2fJX31gD6y0V0omr7OMzShltl	t	8PGATL	MYN30	0	\N	\N
12	2020-02-16 17:09:00.743859+00	2020-02-16 17:09:00.743859+00	\N	Ihar Passenger	91	IN	12342134	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMiwidXNlcl90eXBlIjoicGFzc2VuZ2VyIiwiZXhwIjoxNTg0NDY0OTQxLCJpc3MiOiJvbnJpZGUifQ.O3NXySCuWDYfQEPHPtmIPXhqfC6oFxremeCd8xSevJo		cCphwNV_RL8:APA91bFiQtHbdzQgYaAjaL76xu__B75fvk5DQtZBA8Z1p1BTZa_tAsr8NSg06hV4ObjA-t3gySn7Gi3SZpFrvJLvDhE7Ib3YeZkomc4RNlFtxmIwQvDW4LJkQII3-4XkgXsN6iCudoe2	t	12UQDIW	21234	0	\N	\N
9	2020-02-13 19:08:17.351657+00	2020-02-13 19:08:17.351657+00	\N	mike	44	UK	7729286321	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo5LCJ1c2VyX3R5cGUiOiJwYXNzZW5nZXIiLCJleHAiOjE1ODQyMTI4OTgsImlzcyI6Im9ucmlkZSJ9.yY3NksXIFikIiWi7-TG55mI4NwwQe7PzkG09kCNLoBQ		fk80a4cXnb4:APA91bFedZ_EMSKkz8lQmC9kBX6DHB3ZAlGqjohog7B2VHGooT8d8yZIgNeYtH_otgsTwB8Z155yAHaYsd6kSyWq9O3i9JZ-JQweXb2V-NG50QVWdMlXatMkPf8gZL-jE3B16I0hEf8W	t	9MWJGS		0	\N	\N
10	2020-02-16 14:10:07.736326+00	2020-02-16 14:10:07.736326+00	\N	qqqqqq	91	IN	235671232312	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMCwidXNlcl90eXBlIjoicGFzc2VuZ2VyIiwiZXhwIjoxNTg0NDU0MjA4LCJpc3MiOiJvbnJpZGUifQ.bcgTi--pok8I6AW15dQqs1nDiF_rY-cyHhqGjXqzCNY		da57pGHa-wY:APA91bH4AwiE4FAw_5DbU_qwTGpkbGh8u_Wifco3VtGI2pU9BvXMLmE1OOl393aQgOUQoSmBxLLH-cO_804jT0ps_geY2rzyWHiMgpnfXmB0LjMTg0PTuUZ4IBDfSM8KVWGGQmhDt1gg	t	10GATLM	qqqqqqq	0	\N	\N
11	2020-02-16 15:25:56.912968+00	2020-02-16 15:25:56.912968+00	\N	qqqq	91	IN	123456879	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMSwidXNlcl90eXBlIjoicGFzc2VuZ2VyIiwiZXhwIjoxNTg0NDU4NzU3LCJpc3MiOiJvbnJpZGUifQ.advF1-bXcmhDIp_OCb9zie5PHy8U0mxCsKHOlRR0ru0			t	11DYLUM	12321	0	\N	\N
\.


--
-- Data for Name: pickup_points; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.pickup_points (id, created_at, updated_at, deleted_at, name, zone_id, is_active) FROM stdin;
1	2020-02-02 17:09:13.196607+00	2020-02-02 17:09:13.196607+00	\N	Front Gate	1	t
2	2020-02-02 17:09:13.594523+00	2020-02-02 17:09:13.594523+00	\N	Back Gate	1	t
3	2020-02-02 20:47:50.761096+00	2020-02-02 20:47:50.761096+00	\N	Airport Level 1	2	t
4	2020-02-03 12:23:09.337529+00	2020-02-03 12:23:09.337529+00	\N	Gate 1	3	t
5	2020-02-03 12:23:09.74234+00	2020-02-03 12:23:09.74234+00	\N	Gate 2	3	t
6	2020-02-03 12:23:10.147556+00	2020-02-03 12:23:10.147556+00	\N	Gate 3	3	t
7	2020-02-03 12:23:10.552003+00	2020-02-03 12:23:10.552003+00	\N	Gate 4	3	t
\.


--
-- Data for Name: ride_event_logs; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.ride_event_logs (id, created_at, updated_at, deleted_at, ride_id, ride_status, message, is_active) FROM stdin;
1	2020-02-02 18:58:22.959852+00	2020-02-02 18:58:22.959852+00	\N	1	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
2	2020-02-02 18:58:39.417065+00	2020-02-02 18:58:39.417065+00	\N	1	1	Driver Assigned For Ride	f
3	2020-02-02 19:00:19.464446+00	2020-02-02 19:00:19.464446+00	\N	1	2	Driver Arrived At PickUp Location	f
4	2020-02-02 19:00:23.285138+00	2020-02-02 19:00:23.285138+00	\N	1	3	Ride Started	f
5	2020-02-02 19:00:27.96907+00	2020-02-02 19:00:27.96907+00	\N	1	4	Ride Completed	f
6	2020-02-03 07:36:14.881854+00	2020-02-03 07:36:14.881854+00	\N	2	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
7	2020-02-03 07:36:23.74531+00	2020-02-03 07:36:23.74531+00	\N	2	1	Driver Assigned For Ride	f
8	2020-02-03 07:40:58.03116+00	2020-02-03 07:40:58.03116+00	\N	2	6	Ride Cancelled	f
9	2020-02-03 08:36:46.480627+00	2020-02-03 08:36:46.480627+00	\N	3	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
10	2020-02-03 08:37:17.6262+00	2020-02-03 08:37:17.6262+00	\N	3	1	Driver Assigned For Ride	f
11	2020-02-03 08:38:30.928772+00	2020-02-03 08:38:30.928772+00	\N	3	6	Ride Cancelled  By Driver	f
12	2020-02-03 08:38:32.076351+00	2020-02-03 08:38:32.076351+00	\N	3	0	Ride status changed to waiting & operator started new driver search	f
13	2020-02-03 08:38:34.450967+00	2020-02-03 08:38:34.450967+00	\N	3	6	Ride Cancelled  By Driver	f
14	2020-02-03 08:38:35.591881+00	2020-02-03 08:38:35.591881+00	\N	3	0	Ride status changed to waiting & operator started new driver search	f
15	2020-02-03 08:38:57.790007+00	2020-02-03 08:38:57.790007+00	\N	3	5	Driver unavailable	f
16	2020-02-03 16:51:03.143929+00	2020-02-03 16:51:03.143929+00	\N	4	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
17	2020-02-03 16:51:13.393968+00	2020-02-03 16:51:13.393968+00	\N	4	1	Driver Assigned For Ride	f
18	2020-02-03 16:51:32.216072+00	2020-02-03 16:51:32.216072+00	\N	4	6	Ride Cancelled  By Driver	f
19	2020-02-03 16:51:33.402999+00	2020-02-03 16:51:33.402999+00	\N	4	0	Ride status changed to waiting & operator started new driver search	f
20	2020-02-03 16:51:35.331661+00	2020-02-03 16:51:35.331661+00	\N	4	6	Ride Cancelled  By Driver	f
21	2020-02-03 16:51:36.516583+00	2020-02-03 16:51:36.516583+00	\N	4	0	Ride status changed to waiting & operator started new driver search	f
22	2020-02-03 16:51:59.180491+00	2020-02-03 16:51:59.180491+00	\N	4	5	Driver unavailable	f
23	2020-02-05 08:53:51.188688+00	2020-02-05 08:53:51.188688+00	\N	5	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
24	2020-02-05 08:54:17.038744+00	2020-02-05 08:54:17.038744+00	\N	5	5	Driver unavailable	f
25	2020-02-05 13:52:42.887511+00	2020-02-05 13:52:42.887511+00	\N	6	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
26	2020-02-05 13:53:08.702929+00	2020-02-05 13:53:08.702929+00	\N	6	5	Driver unavailable	f
27	2020-02-07 10:47:13.174655+00	2020-02-07 10:47:13.174655+00	\N	7	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
28	2020-02-07 10:47:39.065242+00	2020-02-07 10:47:39.065242+00	\N	7	5	Driver unavailable	f
29	2020-02-07 10:49:55.17491+00	2020-02-07 10:49:55.17491+00	\N	8	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
30	2020-02-07 10:50:21.036417+00	2020-02-07 10:50:21.036417+00	\N	8	5	Driver unavailable	f
31	2020-02-07 11:37:41.166487+00	2020-02-07 11:37:41.166487+00	\N	9	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
32	2020-02-07 11:37:46.919491+00	2020-02-07 11:37:46.919491+00	\N	9	1	Driver Assigned For Ride	f
33	2020-02-07 11:38:37.742963+00	2020-02-07 11:38:37.742963+00	\N	9	6	Ride Cancelled	f
34	2020-02-07 11:39:08.723894+00	2020-02-07 11:39:08.723894+00	\N	10	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
35	2020-02-07 11:39:43.194468+00	2020-02-07 11:39:43.194468+00	\N	10	1	Driver Assigned For Ride	f
36	2020-02-07 11:39:51.896856+00	2020-02-07 11:39:51.896856+00	\N	10	2	Driver Arrived At PickUp Location	f
37	2020-02-07 11:39:56.377626+00	2020-02-07 11:39:56.377626+00	\N	10	3	Ride Started	f
38	2020-02-07 11:40:03.422689+00	2020-02-07 11:40:03.422689+00	\N	10	4	Ride Completed	f
39	2020-02-07 11:42:49.000755+00	2020-02-07 11:42:49.000755+00	\N	11	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
40	2020-02-07 11:43:19.466878+00	2020-02-07 11:43:19.466878+00	\N	11	1	Driver Assigned For Ride	f
41	2020-02-07 11:43:24.741166+00	2020-02-07 11:43:24.741166+00	\N	11	2	Driver Arrived At PickUp Location	f
42	2020-02-07 11:43:27.950048+00	2020-02-07 11:43:27.950048+00	\N	11	3	Ride Started	f
43	2020-02-07 11:47:31.127741+00	2020-02-07 11:47:31.127741+00	\N	12	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
44	2020-02-07 11:47:36.716361+00	2020-02-07 11:47:36.716361+00	\N	12	1	Driver Assigned For Ride	f
45	2020-02-07 11:47:47.122337+00	2020-02-07 11:47:47.122337+00	\N	12	2	Driver Arrived At PickUp Location	f
46	2020-02-07 11:47:50.644865+00	2020-02-07 11:47:50.644865+00	\N	12	3	Ride Started	f
47	2020-02-07 11:48:04.419419+00	2020-02-07 11:48:04.419419+00	\N	12	4	Ride Completed	f
48	2020-02-07 11:49:40.409014+00	2020-02-07 11:49:40.409014+00	\N	13	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
49	2020-02-07 11:50:10.23649+00	2020-02-07 11:50:10.23649+00	\N	13	7	Driver Assigned For Ride	f
50	2020-02-07 11:52:59.892581+00	2020-02-07 11:52:59.892581+00	\N	14	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
51	2020-02-07 11:53:31.147772+00	2020-02-07 11:53:31.147772+00	\N	14	1	Driver Assigned For Ride	f
52	2020-02-07 11:53:37.513043+00	2020-02-07 11:53:37.513043+00	\N	14	2	Driver Arrived At PickUp Location	f
53	2020-02-07 11:53:44.23803+00	2020-02-07 11:53:44.23803+00	\N	14	3	Ride Started	f
54	2020-02-07 11:53:49.692591+00	2020-02-07 11:53:49.692591+00	\N	14	4	Ride Completed	f
55	2020-02-07 11:56:50.360485+00	2020-02-07 11:56:50.360485+00	\N	15	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
56	2020-02-07 11:57:22.271333+00	2020-02-07 11:57:22.271333+00	\N	15	1	Driver Assigned For Ride	f
57	2020-02-07 11:57:26.606229+00	2020-02-07 11:57:26.606229+00	\N	15	2	Driver Arrived At PickUp Location	f
58	2020-02-07 11:57:31.277005+00	2020-02-07 11:57:31.277005+00	\N	15	3	Ride Started	f
59	2020-02-07 11:57:38.443685+00	2020-02-07 11:57:38.443685+00	\N	15	4	Ride Completed	f
60	2020-02-11 04:10:31.885388+00	2020-02-11 04:10:31.885388+00	\N	16	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
61	2020-02-11 04:11:02.567925+00	2020-02-11 04:11:02.567925+00	\N	16	1	Driver Assigned For Ride	f
62	2020-02-11 04:47:54.201642+00	2020-02-11 04:47:54.201642+00	\N	16	6	Ride Cancelled	f
63	2020-02-11 04:48:22.225739+00	2020-02-11 04:48:22.225739+00	\N	17	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
64	2020-02-11 04:48:29.209314+00	2020-02-11 04:48:29.209314+00	\N	17	1	Driver Assigned For Ride	f
65	2020-02-11 04:50:39.593273+00	2020-02-11 04:50:39.593273+00	\N	17	6	Ride Cancelled	f
66	2020-02-11 04:52:17.606739+00	2020-02-11 04:52:17.606739+00	\N	18	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
67	2020-02-11 04:52:50.136181+00	2020-02-11 04:52:50.136181+00	\N	18	1	Driver Assigned For Ride	f
68	2020-02-11 06:11:48.423247+00	2020-02-11 06:11:48.423247+00	\N	18	6	Ride Cancelled	f
69	2020-02-11 06:15:07.412471+00	2020-02-11 06:15:07.412471+00	\N	19	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
70	2020-02-11 06:15:38.081973+00	2020-02-11 06:15:38.081973+00	\N	19	1	Driver Assigned For Ride	f
71	2020-02-11 06:28:11.476351+00	2020-02-11 06:28:11.476351+00	\N	19	6	Ride Cancelled	f
72	2020-02-11 06:28:45.567676+00	2020-02-11 06:28:45.567676+00	\N	20	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
73	2020-02-11 06:29:17.133654+00	2020-02-11 06:29:17.133654+00	\N	20	1	Driver Assigned For Ride	f
74	2020-02-11 06:31:22.357564+00	2020-02-11 06:31:22.357564+00	\N	20	6	Ride Cancelled	f
75	2020-02-11 06:32:30.559223+00	2020-02-11 06:32:30.559223+00	\N	21	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
76	2020-02-11 06:32:36.268212+00	2020-02-11 06:32:36.268212+00	\N	21	1	Driver Assigned For Ride	f
77	2020-02-11 06:41:52.627868+00	2020-02-11 06:41:52.627868+00	\N	21	6	Ride Cancelled	f
78	2020-02-11 06:43:08.635024+00	2020-02-11 06:43:08.635024+00	\N	22	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
79	2020-02-11 06:43:39.8967+00	2020-02-11 06:43:39.8967+00	\N	22	1	Driver Assigned For Ride	f
80	2020-02-11 07:06:24.702532+00	2020-02-11 07:06:24.702532+00	\N	22	6	Ride Cancelled	f
83	2020-02-11 07:08:12.083232+00	2020-02-11 07:08:12.083232+00	\N	23	6	Ride Cancelled	f
86	2020-02-11 07:11:29.972336+00	2020-02-11 07:11:29.972336+00	\N	25	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
88	2020-02-11 07:16:34.199656+00	2020-02-11 07:16:34.199656+00	\N	25	6	Ride Cancelled	f
89	2020-02-11 07:17:15.05974+00	2020-02-11 07:17:15.05974+00	\N	26	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
92	2020-02-11 07:25:04.174559+00	2020-02-11 07:25:04.174559+00	\N	27	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
96	2020-02-11 07:31:58.469622+00	2020-02-11 07:31:58.469622+00	\N	28	1	Driver Assigned For Ride	f
81	2020-02-11 07:07:21.174294+00	2020-02-11 07:07:21.174294+00	\N	23	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
82	2020-02-11 07:07:26.459091+00	2020-02-11 07:07:26.459091+00	\N	23	1	Driver Assigned For Ride	f
84	2020-02-11 07:08:57.643797+00	2020-02-11 07:08:57.643797+00	\N	24	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
85	2020-02-11 07:09:02.802824+00	2020-02-11 07:09:02.802824+00	\N	24	1	Driver Assigned For Ride	f
87	2020-02-11 07:12:04.815634+00	2020-02-11 07:12:04.815634+00	\N	25	1	Driver Assigned For Ride	f
90	2020-02-11 07:17:20.779936+00	2020-02-11 07:17:20.779936+00	\N	26	1	Driver Assigned For Ride	f
91	2020-02-11 07:22:25.333892+00	2020-02-11 07:22:25.333892+00	\N	26	6	Ride Cancelled	f
93	2020-02-11 07:25:35.038219+00	2020-02-11 07:25:35.038219+00	\N	27	1	Driver Assigned For Ride	f
94	2020-02-11 07:30:59.26416+00	2020-02-11 07:30:59.26416+00	\N	27	6	Ride Cancelled	f
95	2020-02-11 07:31:52.746772+00	2020-02-11 07:31:52.746772+00	\N	28	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
97	2020-02-11 07:35:22.965608+00	2020-02-11 07:35:22.965608+00	\N	28	6	Ride Cancelled	f
98	2020-02-11 07:36:43.457315+00	2020-02-11 07:36:43.457315+00	\N	29	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
99	2020-02-11 07:37:13.953599+00	2020-02-11 07:37:13.953599+00	\N	29	1	Driver Assigned For Ride	f
100	2020-02-11 07:38:37.61491+00	2020-02-11 07:38:37.61491+00	\N	29	6	Ride Cancelled	f
101	2020-02-11 07:39:36.755869+00	2020-02-11 07:39:36.755869+00	\N	30	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
102	2020-02-11 07:40:07.997262+00	2020-02-11 07:40:07.997262+00	\N	30	1	Driver Assigned For Ride	f
103	2020-02-11 09:42:33.494038+00	2020-02-11 09:42:33.494038+00	\N	31	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
104	2020-02-11 09:43:06.134842+00	2020-02-11 09:43:06.134842+00	\N	31	1	Driver Assigned For Ride	f
105	2020-02-11 10:23:51.880389+00	2020-02-11 10:23:51.880389+00	\N	32	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
106	2020-02-11 10:23:57.872944+00	2020-02-11 10:23:57.872944+00	\N	32	1	Driver Assigned For Ride	f
107	2020-02-11 10:25:47.082736+00	2020-02-11 10:25:47.082736+00	\N	33	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
108	2020-02-11 10:25:53.719647+00	2020-02-11 10:25:53.719647+00	\N	33	1	Driver Assigned For Ride	f
109	2020-02-11 10:29:33.064622+00	2020-02-11 10:29:33.064622+00	\N	34	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
110	2020-02-11 10:29:41.277514+00	2020-02-11 10:29:41.277514+00	\N	34	1	Driver Assigned For Ride	f
111	2020-02-11 10:35:44.544111+00	2020-02-11 10:35:44.544111+00	\N	35	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
112	2020-02-11 10:35:51.22732+00	2020-02-11 10:35:51.22732+00	\N	35	1	Driver Assigned For Ride	f
113	2020-02-11 10:35:54.500508+00	2020-02-11 10:35:54.500508+00	\N	35	2	Driver Arrived At PickUp Location	f
114	2020-02-11 10:48:22.826296+00	2020-02-11 10:48:22.826296+00	\N	36	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
115	2020-02-11 10:48:54.33098+00	2020-02-11 10:48:54.33098+00	\N	36	1	Driver Assigned For Ride	f
116	2020-02-11 10:50:10.313446+00	2020-02-11 10:50:10.313446+00	\N	36	6	Ride Cancelled	f
117	2020-02-11 10:51:50.166984+00	2020-02-11 10:51:50.166984+00	\N	37	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
118	2020-02-11 10:51:57.622676+00	2020-02-11 10:51:57.622676+00	\N	37	1	Driver Assigned For Ride	f
119	2020-02-11 10:59:28.575776+00	2020-02-11 10:59:28.575776+00	\N	38	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
120	2020-02-11 10:59:54.418725+00	2020-02-11 10:59:54.418725+00	\N	38	5	Driver unavailable	f
121	2020-02-11 11:00:40.556572+00	2020-02-11 11:00:40.556572+00	\N	39	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
122	2020-02-11 11:01:11.113216+00	2020-02-11 11:01:11.113216+00	\N	39	1	Driver Assigned For Ride	f
123	2020-02-16 17:13:19.083195+00	2020-02-16 17:13:19.083195+00	\N	40	0	Ride Booking Accepted By Poland Operator	f
124	2020-02-16 17:13:45.005415+00	2020-02-16 17:13:45.005415+00	\N	40	5	Driver unavailable	f
125	2020-02-16 17:18:14.515396+00	2020-02-16 17:18:14.515396+00	\N	41	0	Ride Booking Accepted By Poland Operator	f
126	2020-02-16 17:18:40.441683+00	2020-02-16 17:18:40.441683+00	\N	41	5	Driver unavailable	f
127	2020-02-20 08:57:09.564753+00	2020-02-20 08:57:09.564753+00	\N	42	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
128	2020-02-20 08:57:41.517266+00	2020-02-20 08:57:41.517266+00	\N	42	1	Driver Assigned For Ride	f
129	2020-02-20 08:59:52.593246+00	2020-02-20 08:59:52.593246+00	\N	42	2	Driver Arrived At PickUp Location	f
130	2020-02-20 08:59:58.032113+00	2020-02-20 08:59:58.032113+00	\N	42	3	Ride Started	f
131	2020-02-20 09:00:05.632462+00	2020-02-20 09:00:05.632462+00	\N	42	4	Ride Completed	f
132	2020-02-21 08:10:57.455174+00	2020-02-21 08:10:57.455174+00	\N	43	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
133	2020-02-21 08:11:03.219985+00	2020-02-21 08:11:03.219985+00	\N	43	1	Driver Assigned For Ride	f
134	2020-02-21 08:12:32.249051+00	2020-02-21 08:12:32.249051+00	\N	43	6	Ride Cancelled  By Driver	f
135	2020-02-21 08:12:33.408354+00	2020-02-21 08:12:33.408354+00	\N	43	0	Ride status changed to waiting & operator started new driver search	f
136	2020-02-21 08:12:59.185846+00	2020-02-21 08:12:59.185846+00	\N	43	5	Driver unavailable	f
137	2020-02-21 08:14:40.722641+00	2020-02-21 08:14:40.722641+00	\N	44	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
138	2020-02-21 08:14:45.674399+00	2020-02-21 08:14:45.674399+00	\N	44	1	Driver Assigned For Ride	f
139	2020-02-21 08:18:35.473923+00	2020-02-21 08:18:35.473923+00	\N	44	2	Driver Arrived At PickUp Location	f
140	2020-02-21 08:18:42.821868+00	2020-02-21 08:18:42.821868+00	\N	44	3	Ride Started	f
141	2020-02-21 08:19:43.47551+00	2020-02-21 08:19:43.47551+00	\N	44	4	Ride Completed	f
142	2020-02-21 08:37:55.9136+00	2020-02-21 08:37:55.9136+00	\N	45	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
143	2020-02-21 08:38:25.461311+00	2020-02-21 08:38:25.461311+00	\N	45	1	Driver Assigned For Ride	f
144	2020-02-21 08:38:36.584309+00	2020-02-21 08:38:36.584309+00	\N	45	6	Ride Cancelled  By Driver	f
145	2020-02-21 08:38:37.750615+00	2020-02-21 08:38:37.750615+00	\N	45	0	Ride status changed to waiting & operator started new driver search	f
146	2020-02-21 08:38:47.043134+00	2020-02-21 08:38:47.043134+00	\N	45	5	Driver unavailable	f
147	2020-02-21 08:41:11.944458+00	2020-02-21 08:41:11.944458+00	\N	46	0	Ride Booking Accepted By Pollachi Taxi Services Operator	f
148	2020-02-21 08:41:41.658846+00	2020-02-21 08:41:41.658846+00	\N	46	1	Driver Assigned For Ride	f
149	2020-02-21 08:41:59.663505+00	2020-02-21 08:41:59.663505+00	\N	46	6	Ride Cancelled  By Driver	f
150	2020-02-21 08:42:00.835634+00	2020-02-21 08:42:00.835634+00	\N	46	0	Ride status changed to waiting & operator started new driver search	f
151	2020-02-21 08:42:03.052837+00	2020-02-21 08:42:03.052837+00	\N	46	5	Driver unavailable	f
\.


--
-- Data for Name: ride_locations; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.ride_locations (id, created_at, updated_at, deleted_at, ride_id, "time", is_active, latlng) FROM stdin;
\.


--
-- Data for Name: ride_messages; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.ride_messages (id, created_at, updated_at, deleted_at, ride_id, message, "from", is_active) FROM stdin;
1	2020-02-11 09:04:48.952988+00	2020-02-11 09:04:48.952988+00	\N	30	hi	0	t
2	2020-02-11 09:06:06.061162+00	2020-02-11 09:06:06.061162+00	\N	30	hi	0	t
3	2020-02-11 09:21:20.641202+00	2020-02-11 09:21:20.641202+00	\N	30	hi	0	t
4	2020-02-11 09:22:20.248235+00	2020-02-11 09:22:20.248235+00	\N	30	hi	0	t
5	2020-02-11 09:22:28.061436+00	2020-02-11 09:22:28.061436+00	\N	30	hi	0	t
6	2020-02-11 09:22:51.238569+00	2020-02-11 09:22:51.238569+00	\N	30	hi	0	t
7	2020-02-11 09:23:36.569396+00	2020-02-11 09:23:36.569396+00	\N	30	hi	0	t
8	2020-02-11 09:25:04.166428+00	2020-02-11 09:25:04.166428+00	\N	30	hi	0	t
9	2020-02-11 09:25:14.116491+00	2020-02-11 09:25:14.116491+00	\N	30	htrdc	0	t
10	2020-02-11 09:25:45.129718+00	2020-02-11 09:25:45.129718+00	\N	30	hikjgg	0	t
11	2020-02-11 09:28:42.835534+00	2020-02-11 09:28:42.835534+00	\N	30	hi	0	t
12	2020-02-11 09:29:07.314386+00	2020-02-11 09:29:07.314386+00	\N	30	hi	0	t
13	2020-02-11 09:29:18.288629+00	2020-02-11 09:29:18.288629+00	\N	30	hello	0	t
14	2020-02-11 09:30:35.864401+00	2020-02-11 09:30:35.864401+00	\N	30	ho	0	t
15	2020-02-11 09:31:07.276278+00	2020-02-11 09:31:07.276278+00	\N	30	hello	0	t
16	2020-02-11 09:35:57.049617+00	2020-02-11 09:35:57.049617+00	\N	30	hi	0	t
17	2020-02-11 09:36:11.205054+00	2020-02-11 09:36:11.205054+00	\N	30	gsshdhjd	0	t
18	2020-02-11 09:37:42.472454+00	2020-02-11 09:37:42.472454+00	\N	30	jfjfjfjfjfjfjf	0	t
19	2020-02-11 09:37:50.388083+00	2020-02-11 09:37:50.388083+00	\N	30	djdjfjfjfj	0	t
20	2020-02-11 09:43:28.827092+00	2020-02-11 09:43:28.827092+00	\N	31	Hi	0	t
21	2020-02-11 09:43:40.875574+00	2020-02-11 09:43:40.875574+00	\N	31	hi	0	t
22	2020-02-11 09:43:48.290555+00	2020-02-11 09:43:48.290555+00	\N	31	hi	0	t
23	2020-02-11 09:43:54.11641+00	2020-02-11 09:43:54.11641+00	\N	31	hi	0	t
24	2020-02-11 09:43:58.896728+00	2020-02-11 09:43:58.896728+00	\N	31	jfjfj	0	t
25	2020-02-11 09:44:03.582734+00	2020-02-11 09:44:03.582734+00	\N	31	jfjdkc	0	t
26	2020-02-11 09:44:08.556872+00	2020-02-11 09:44:08.556872+00	\N	31	fjfkkck	0	t
27	2020-02-11 09:44:16.876371+00	2020-02-11 09:44:16.876371+00	\N	31	fjfjfi	0	t
28	2020-02-11 09:44:48.74579+00	2020-02-11 09:44:48.74579+00	\N	31	xhfjcjfk	0	t
29	2020-02-11 09:44:54.038909+00	2020-02-11 09:44:54.038909+00	\N	31	hffjjcjc	0	t
30	2020-02-11 09:44:59.769638+00	2020-02-11 09:44:59.769638+00	\N	31	hffjfj	0	t
31	2020-02-11 09:45:04.847602+00	2020-02-11 09:45:04.847602+00	\N	31	hcfjjcjcjc	0	t
32	2020-02-11 09:45:10.296364+00	2020-02-11 09:45:10.296364+00	\N	31	fjjcjcjfkf	0	t
33	2020-02-11 09:45:20.263179+00	2020-02-11 09:45:20.263179+00	\N	31	jffjjficig	0	t
34	2020-02-11 10:30:09.052241+00	2020-02-11 10:30:09.052241+00	\N	34	hi	1	t
35	2020-02-11 10:30:41.932459+00	2020-02-11 10:30:41.932459+00	\N	34	hi	0	t
36	2020-02-11 10:36:13.70735+00	2020-02-11 10:36:13.70735+00	\N	35	hi	0	t
37	2020-02-11 10:37:53.817291+00	2020-02-11 10:37:53.817291+00	\N	35	hi	0	t
38	2020-02-11 10:52:10.015939+00	2020-02-11 10:52:10.015939+00	\N	37	hi	0	t
39	2020-02-11 10:52:24.461122+00	2020-02-11 10:52:24.461122+00	\N	37	hi	1	t
40	2020-02-11 10:52:46.602411+00	2020-02-11 10:52:46.602411+00	\N	37	jjgjhgj	1	t
41	2020-02-11 10:52:55.270678+00	2020-02-11 10:52:55.270678+00	\N	37	hi	0	t
42	2020-02-11 10:53:22.2446+00	2020-02-11 10:53:22.2446+00	\N	37	ghjgjgc	1	t
43	2020-02-11 10:53:33.60139+00	2020-02-11 10:53:33.60139+00	\N	37	ghjjj	0	t
44	2020-02-11 10:53:47.403634+00	2020-02-11 10:53:47.403634+00	\N	37	l;nah lol nm	1	t
45	2020-02-11 10:53:55.732584+00	2020-02-11 10:53:55.732584+00	\N	37	gghh	0	t
46	2020-02-11 10:58:24.494907+00	2020-02-11 10:58:24.494907+00	\N	37	vcbnvbnvbnvb	1	t
47	2020-02-11 10:58:34.333149+00	2020-02-11 10:58:34.333149+00	\N	37	ghjjnn	0	t
48	2020-02-11 11:01:26.050498+00	2020-02-11 11:01:26.050498+00	\N	39	hi	0	t
49	2020-02-11 11:01:38.623212+00	2020-02-11 11:01:38.623212+00	\N	39	hi	0	t
50	2020-02-11 11:01:50.300323+00	2020-02-11 11:01:50.300323+00	\N	39	kih	1	t
51	2020-02-11 11:02:22.45436+00	2020-02-11 11:02:22.45436+00	\N	39	gdfdfg	1	t
52	2020-02-11 11:02:30.575751+00	2020-02-11 11:02:30.575751+00	\N	39	fdgdfgdfgdf	1	t
53	2020-02-11 11:02:48.083189+00	2020-02-11 11:02:48.083189+00	\N	39	hhhh	0	t
54	2020-02-11 11:03:09.165058+00	2020-02-11 11:03:09.165058+00	\N	39	vxcvxcvxc	1	t
55	2020-02-11 11:04:03.14251+00	2020-02-11 11:04:03.14251+00	\N	39	ggggggggjcufufufufufucuucucucucucuvucuvuvuvuvuvuvuvuvucuuvuvuvuvuvuvuvuguugucuguguvuuvuvuvuguguguigigigigigigigigiigigigigigigiigigig	0	t
56	2020-02-11 11:04:23.710721+00	2020-02-11 11:04:23.710721+00	\N	39	gujii	0	t
57	2020-02-11 11:04:31.440592+00	2020-02-11 11:04:31.440592+00	\N	39	fdgdfgdfgdfgdf	1	t
58	2020-02-11 11:05:29.028761+00	2020-02-11 11:05:29.028761+00	\N	39	jcjcjcjc	0	t
59	2020-02-11 11:05:40.392737+00	2020-02-11 11:05:40.392737+00	\N	39	][‘p=[;=/.=[p/=[p/[=/=[/[=	1	t
60	2020-02-20 08:58:16.272523+00	2020-02-20 08:58:16.272523+00	\N	42	hi	0	t
61	2020-02-20 08:58:26.287109+00	2020-02-20 08:58:26.287109+00	\N	42	hello	1	t
62	2020-02-20 08:58:42.04598+00	2020-02-20 08:58:42.04598+00	\N	42	where are you now?	1	t
63	2020-02-20 08:58:57.718122+00	2020-02-20 08:58:57.718122+00	\N	42	I'm on the way 	0	t
64	2020-02-20 08:59:08.840709+00	2020-02-20 08:59:08.840709+00	\N	42	I will reach in 5 minutes	0	t
65	2020-02-20 08:59:18.927023+00	2020-02-20 08:59:18.927023+00	\N	42	ok thanks	1	t
66	2020-02-20 08:59:37.564659+00	2020-02-20 08:59:37.564659+00	\N	42	thanks	0	t
\.


--
-- Data for Name: ride_stops; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.ride_stops (id, created_at, updated_at, deleted_at, ride_id, location, latitude, longitude, is_reached, is_active) FROM stdin;
1	2020-02-07 10:49:54.782848+00	2020-02-07 10:49:54.782848+00	\N	8	Aymmapalayam Mill Bus Stop, SH78A, Thalakkarai, Tamil Nadu 642005, India	10.674765780824403	76.96036107838154	f	t
2	2020-02-07 11:37:40.77475+00	2020-02-07 11:37:40.77475+00	\N	9	78, Palakkad - Pollachi Rd, Palaghat, Pollachi, Tamil Nadu 642001, India	10.668961052536083	76.98045048862696	f	t
3	2020-02-07 11:39:08.332428+00	2020-02-07 11:39:08.332428+00	\N	10	Kannan Coir, Palakkad - Pollachi Rd, Nallur, Tamil Nadu 642004, India	10.667527150947588	76.98564123362303	f	t
4	2020-02-07 11:42:48.609984+00	2020-02-07 11:42:48.609984+00	\N	11	101, Palakkad - Pollachi Rd, Bodipalayam village, Tamil Nadu 642001, India	10.669167307125733	76.98000557720661	f	t
5	2020-02-07 11:47:30.737229+00	2020-02-07 11:47:30.737229+00	\N	12	101, Palakkad - Pollachi Rd, Bodipalayam village, Tamil Nadu 642001, India	10.669663833175981	76.97951439768076	f	t
6	2020-02-07 11:49:40.018834+00	2020-02-07 11:49:40.018834+00	\N	13	101, Palakkad - Pollachi Rd, Bodipalayam village, Tamil Nadu 642001, India	10.668916902174448	76.98008939623833	f	t
7	2020-02-07 11:52:59.500787+00	2020-02-07 11:52:59.500787+00	\N	14	78, Palakkad - Pollachi Rd, Palaghat, Pollachi, Tamil Nadu 642001, India	10.669024971705019	76.98047764599323	f	t
8	2020-02-07 11:56:49.969513+00	2020-02-07 11:56:49.969513+00	\N	15	11, Krishna Anaicut Rd, Bodipalayam village, Tamil Nadu 642005, India	10.66941968874959	76.97926729917526	f	t
9	2020-02-11 07:07:20.791044+00	2020-02-11 07:07:20.791044+00	\N	23	Pollachi Main Rd, Jeeva Nagar, Tamil Nadu 642002, India	10.67807071651588	77.00412098318338	f	t
10	2020-02-11 07:08:57.262345+00	2020-02-11 07:08:57.262345+00	\N	24	S.S Kovil Street, Puliampatti, Pollachi, Tamil Nadu 642001, India	10.657068229143961	77.01024312525988	f	t
11	2020-02-11 07:11:29.590697+00	2020-02-11 07:11:29.590697+00	\N	25	Raja Mill Rd, Vinayagar Kovil, Pollachi, Tamil Nadu 642001, India	10.660227393207643	77.00739562511444	f	t
12	2020-02-11 07:17:14.678229+00	2020-02-11 07:17:14.678229+00	\N	26	Pollachi Main Rd, Palaghat, Pollachi, Tamil Nadu 642001, India	10.660902187456616	77.00808227062225	f	t
13	2020-02-11 07:25:03.793843+00	2020-02-11 07:25:03.793843+00	\N	27	6/2, street, Krishnaswamy Nagar, Palaghat, Pollachi, Tamil Nadu 642001, India	10.661295926573825	77.00550734996796	f	t
14	2020-02-11 07:31:52.365581+00	2020-02-11 07:31:52.365581+00	\N	28	32, Bazaar St, Puliampatti, Pollachi, Tamil Nadu 642001, India	10.656994093173084	77.01011370867491	f	t
15	2020-02-11 07:36:43.07618+00	2020-02-11 07:36:43.07618+00	\N	29	T.N.S.T.C Depot, Venkatasa Colony, Pollachi, Tamil Nadu 642001, India	10.669646041281078	77.00753811746836	f	t
16	2020-02-11 07:39:36.373968+00	2020-02-11 07:39:36.373968+00	\N	30	Warehouse Bus Stop, SH19, Kannappan Nagar, Pollachi, Tamil Nadu 642006, India	10.652973260571228	76.99560765177011	f	t
17	2020-02-11 09:42:33.109716+00	2020-02-11 09:42:33.109716+00	\N	31	48, Municipal Office Rd, near A.T.S.C Theatre., Palaghat, Pollachi, Tamil Nadu 642001, India	10.661715035769063	77.00102135539055	f	t
18	2020-02-11 10:23:51.498121+00	2020-02-11 10:23:51.498121+00	\N	32	Kesav Vidya Mandhir, Nallur, Tamil Nadu 642001, India	10.66761512258618	76.98803309351206	f	t
19	2020-02-11 10:25:46.700424+00	2020-02-11 10:25:46.700424+00	\N	33	Andy Gounder St, Kettimalanpudur, Pollachi, Tamil Nadu 642001, India	10.659309103904308	77.00650412589312	f	t
20	2020-02-11 10:29:32.681544+00	2020-02-11 10:29:32.681544+00	\N	34	1, Gowri nagar Panickampatti road, T . Kottampatty, Mahalakshmi Nagar, Pollachi, Tamil Nadu 642002, India	10.673385281643688	77.0199554041028	f	t
21	2020-02-11 10:35:44.152079+00	2020-02-11 10:35:44.152079+00	\N	35	26, Vadugapalayam, Tamil Nadu 642001, India	10.668731734169974	76.99141468852758	f	t
22	2020-02-11 10:48:22.434613+00	2020-02-11 10:48:22.434613+00	\N	36	Bharathi Street, Arumugam Nagar, Pollachi, Tamil Nadu 642002, India	10.67065523545213	77.00952798128128	f	t
23	2020-02-11 10:51:49.782901+00	2020-02-11 10:51:49.782901+00	\N	37	54/7, Ramakrishna Puram St, Mahalingapuram, Tamil Nadu 642001, India	10.66044057299776	77.01233390718699	f	t
24	2020-02-20 08:57:09.181669+00	2020-02-20 08:57:09.181669+00	\N	42	Raja Mill Rd, Vinayagar Kovil, Pollachi, Tamil Nadu 642001, India	10.66049988110447	77.00749453157187	f	t
25	2020-02-21 08:10:57.023052+00	2020-02-21 08:10:57.023052+00	\N	43	16/¹9, Alagappa Layout Venkatesa Colony, Venkatasa Colony, Pollachi, Tamil Nadu 642001, India	10.664596725346517	77.00406465679407	f	t
26	2020-02-21 08:14:40.332347+00	2020-02-21 08:14:40.332347+00	\N	44	16/¹9, Alagappa Layout Venkatesa Colony, Venkatasa Colony, Pollachi, Tamil Nadu 642001, India	10.664596725346517	77.00406465679407	f	t
27	2020-02-21 08:37:55.531671+00	2020-02-21 08:37:55.531671+00	\N	45	96, Palaghat Rd, near Nallamuthu Gounder Mahalingam College, Palaghat, Pollachi, Tamil Nadu 642001, India	10.66314830577764	76.99749860912561	f	t
28	2020-02-21 08:41:11.554048+00	2020-02-21 08:41:11.554048+00	\N	46	96, Palaghat Rd, near Nallamuthu Gounder Mahalingam College, Palaghat, Pollachi, Tamil Nadu 642001, India	10.66314830577764	76.99749860912561	f	t
\.


--
-- Data for Name: rides; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.rides (id, created_at, updated_at, deleted_at, passenger_id, operator_id, zone_id, vehicle_type_id, driver_id, pickup_point, pickup_location, pickup_latitude, pickup_longitude, drop_location, drop_latitude, drop_longitude, ride_date_time, ride_driver_arrived_time, ride_start_time, ride_end_time, ride_type, is_ride_later, distance, duration, duration_readable, fare_id, zone_fare_id, distance_fare, duration_fare, waiting_fare, cancellation_fee, tax, is_paid, transaction_id, total_fare, passenger_rating, driver_rating, passenger_review, driver_review, ride_status, is_active, is_multi_stop, payment_verified, fleet_id) FROM stdin;
9	2020-02-07 11:37:40.376031+00	2020-02-07 11:37:40.376031+00	\N	2	1	3	1	3		135, Palakkad - Pollachi Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.664574979334827	76.99791703373194	78, Palakkad - Pollachi Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.668961052536083	76.98045048862696	2020-02-07 11:37:39.874323+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			6	f	t	0	0
2	2020-02-03 07:36:14.487795+00	2020-02-03 07:36:14.487795+00	\N	2	1	0	1	1		Palakkad - Pollachi Road, 642005, Zamin Muthur, Coimbatore, Tamil Nadu, India	10.669666469012164	76.97609022259712	Unnamed Road, Zamin Uthukuli, 642005, Zamin Uthukuli, Coimbatore, Tamil Nadu, India	10.66064551540637	76.98665913194418	2020-02-03 07:36:14.119222+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	0	0	0	0	0	0	f		0	0	0			6	f	\N	0	0
4	2020-02-03 16:51:02.748751+00	2020-02-03 16:51:02.748751+00	\N	2	1	0	2	17		Palakkad - Koduvayur - Meenakshipuram Highway, 678534, Moolathara, Palakkad, Kerala, India	10.63224583915019	76.85758419334888	Palakkad - Koduvayur - Meenakshipuram Highway, 642103, Meenakshipuram, Coimbatore, Tamil Nadu, India	10.631275730744782	76.871832087636	2020-02-03 16:51:02.346429+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		2	0	0	0	0	0	0	f		0	0	0			5	f	\N	0	0
5	2020-02-05 08:53:50.796299+00	2020-02-05 08:53:50.796299+00	\N	2	1	3	2	0		135, Palakkad - Pollachi Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.664446809628231	76.9977879524231	Pollachi Main Road, Puliampatti, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.6578254034929	77.00971841812134	2020-02-05 08:53:50.434434+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		2	4	0	0	0	0	0	f		0	0	0			5	f	\N	0	0
3	2020-02-03 08:36:46.098086+00	2020-02-03 08:36:46.098086+00	\N	2	1	0	1	3		Palaghat Road, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.666572640558998	76.98821481317282	Pollachi Main Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.66134304343825	77.00774129480124	2020-02-03 08:36:45.604154+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	0	0	0	0	0	0	f		0	0	0			5	f	\N	0	0
12	2020-02-07 11:47:30.344156+00	2020-02-07 11:47:30.344156+00	\N	2	1	3	1	3		135, Palakkad - Pollachi Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.664583875430708	76.99798811227083	101, Palakkad - Pollachi Road, 642001, Bodipalayam village, Coimbatore, Tamil Nadu, India	10.669663833175981	76.97951439768076	2020-02-07 11:47:29.935114+00	2020-02-07 11:47:46.635236+00	2020-02-07 11:47:50.156765+00	2020-02-07 11:48:05.00655+00	0	f	0	0.25	14.849785348s	1	3	0	0.25	0	0	2.03	f		22.28	0	0			4	f	t	0	0
6	2020-02-05 13:52:42.506426+00	2020-02-05 13:52:42.506426+00	\N	2	1	0	3	0		Palakkad - Koduvayur - Meenakshipuram Highway, 678534, Moolathara, Palakkad, Kerala, India	10.635082004563644	76.84253867715597	Palakkad - Koduvayur - Meenakshipuram Highway, 678534, Moolathara, Palakkad, Kerala, India	10.635039826393449	76.84936221688986	2020-02-05 13:52:42.07616+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		3	0	0	0	0	0	0	f		0	0	0			5	f	\N	0	0
7	2020-02-07 10:47:12.477937+00	2020-02-07 10:47:12.477937+00	\N	2	1	3	3	0		135, Palakkad - Pollachi Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.664631650755993	76.99786372482777	State Highway 78A, 642005, Thalakkarai, Coimbatore, Tamil Nadu, India	10.674765780824403	76.96036107838154	2020-02-07 10:47:12.092726+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		3	5	0	0	0	0	0	f		0	0	0			5	f	t	0	0
8	2020-02-07 10:49:54.390064+00	2020-02-07 10:49:54.390064+00	\N	2	1	3	3	0		135, Palakkad - Pollachi Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.664631650755993	76.99786372482777	State Highway 78A, 642005, Thalakkarai, Coimbatore, Tamil Nadu, India	10.674765780824403	76.96036107838154	2020-02-07 10:49:54.074871+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		3	5	0	0	0	0	0	f		0	0	0			5	f	t	0	0
11	2020-02-07 11:42:48.216983+00	2020-02-07 11:42:48.216983+00	\N	2	1	3	1	3		135, Palakkad - Pollachi Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.664574979334827	76.99791703373194	101, Palakkad - Pollachi Road, 642001, Bodipalayam village, Coimbatore, Tamil Nadu, India	10.669167307125733	76.98000557720661	2020-02-07 11:42:47.959644+00	2020-02-07 11:43:24.254253+00	2020-02-07 11:43:27.462118+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			3	f	t	0	0
10	2020-02-07 11:39:07.940888+00	2020-02-07 11:39:07.940888+00	\N	2	1	0	1	3		78, Palakkad - Pollachi Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.668961052536083	76.98045048862696	Palakkad - Pollachi Road, 642004, Nallur, Coimbatore, Tamil Nadu, India	10.667527150947588	76.98564123362303	2020-02-07 11:39:07.662006+00	2020-02-07 11:39:51.409554+00	2020-02-07 11:39:55.890261+00	2020-02-07 11:40:04.008192+00	0	f	0	0.14	8.11793055s	1	0	0	2.8	0	0	2.29	f		25.09	0	0			4	f	t	0	0
13	2020-02-07 11:49:39.626055+00	2020-02-07 11:49:39.626055+00	\N	2	1	3	1	3		135, Palakkad - Pollachi Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.66465076088402	76.99787009507418	101, Palakkad - Pollachi Road, 642001, Bodipalayam village, Coimbatore, Tamil Nadu, India	10.668916902174448	76.98008939623833	2020-02-07 11:49:39.255927+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			1	f	t	0	0
14	2020-02-07 11:52:59.106808+00	2020-02-07 11:52:59.106808+00	\N	2	1	3	1	3		135, Palakkad - Pollachi Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.664428028972603	76.99786305427551	78, Palakkad - Pollachi Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.669024971705019	76.98047764599323	2020-02-07 11:52:58.661776+00	2020-02-07 11:53:37.025748+00	2020-02-07 11:53:43.749055+00	2020-02-07 11:53:50.282769+00	0	f	0	0.11	6.533714478s	1	3	0	0.11	0	0	2.02	f		22.13	0	0			4	f	t	0	0
16	2020-02-11 04:10:31.502216+00	2020-02-11 04:10:31.502216+00	\N	6	1	3	1	3		642001, Bodipalayam village, Coimbatore, Tamil Nadu, India	10.665272827844001	76.99522677809	Pollachi Main Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.661027393256225	77.00784388929605	2020-02-11 04:10:31.107781+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			6	f	f	0	0
15	2020-02-07 11:56:49.573572+00	2020-02-07 11:56:49.573572+00	\N	2	1	0	1	3		78, Palakkad - Pollachi Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.669024971705019	76.98047764599323	11, Krishna Anaicut Road, 642005, Bodipalayam village, Coimbatore, Tamil Nadu, India	10.66941968874959	76.97926729917526	2020-02-07 11:56:49.305146+00	2020-02-07 11:57:26.118119+00	2020-02-07 11:57:30.789575+00	2020-02-07 11:57:39.032331+00	0	f	0	0.14	8.242756252s	1	0	0	2.8	0	0	2.29	f		25.09	0	0			4	f	t	0	0
17	2020-02-11 04:48:21.842553+00	2020-02-11 04:48:21.842553+00	\N	6	1	3	1	3		Pollachi Main Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.661027393256225	77.00784388929605	Pollachi Main Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.661027393256225	77.00784388929605	2020-02-11 04:48:21.532737+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			6	f	f	0	0
18	2020-02-11 04:52:17.223856+00	2020-02-11 04:52:17.223856+00	\N	6	1	3	1	3		Pollachi Main Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.661027393256225	77.00784388929605	Pollachi Main Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.661027393256225	77.00784388929605	2020-02-11 04:52:16.990296+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			6	f	f	0	0
1	2020-02-02 18:58:22.576433+00	2020-02-02 18:58:22.576433+00	\N	2	1	0	1	1		Palakkad - Koduvayur - Meenakshipuram Highway, 642103, Meenakshipuram, Coimbatore, Tamil Nadu, India	10.63342683652533	76.86228509992361	Meenkarai Road, 642103, Valathyamaram, Coimbatore, Tamil Nadu, India	10.629672936189516	76.8800949677825	2020-02-02 18:58:22.256423+00	2020-02-02 19:00:18.98773+00	2020-02-02 19:00:22.807986+00	2020-02-02 19:00:28.543098+00	0	f	0	0.1	5.735112201s	1	0	0	0.1	0	0	2.02	f		22.12	0	0			4	f	\N	0	0
19	2020-02-11 06:15:07.020266+00	2020-02-11 06:15:07.020266+00	\N	6	1	3	1	3		Palaghat, 642001, Coimbatore, Tamil Nadu, India	10.664429346913384	76.99688639491796	Perumal Chetty Street, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.661111742397468	77.0074150711298	2020-02-11 06:15:06.572175+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			6	f	f	0	0
20	2020-02-11 06:28:45.175813+00	2020-02-11 06:28:45.175813+00	\N	6	1	3	1	3		642001, Bodipalayam village, Coimbatore, Tamil Nadu, India	10.664991777621816	76.99416864663363	Kettimalanpudur, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.659424755131578	77.00687158852816	2020-02-11 06:28:44.795647+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			6	f	f	0	0
21	2020-02-11 06:32:30.156631+00	2020-02-11 06:32:30.156631+00	\N	6	1	3	1	3		street, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.661646172741236	77.00581312179565	Palaghat Road, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.666228989727252	76.98961962014437	2020-02-11 06:32:29.659781+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			6	f	f	0	0
22	2020-02-11 06:43:08.239371+00	2020-02-11 06:43:08.239371+00	\N	6	1	3	1	3		Palaghat Road, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.666228989727252	76.98961962014437	135, Palakkad - Pollachi Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.66418256240206	76.99797369539738	2020-02-11 06:43:07.898906+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			6	f	f	0	0
23	2020-02-11 07:07:20.408551+00	2020-02-11 07:07:20.408551+00	\N	6	1	3	1	3		14, Imaam Khan Street, Puliampatti, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.657714694045051	77.00795520097017	Pollachi Main Road, 642002, Jeeva Nagar, Coimbatore, Tamil Nadu, India	10.67807071651588	77.00412098318338	2020-02-11 07:07:19.941944+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			6	f	t	0	0
24	2020-02-11 07:08:56.880327+00	2020-02-11 07:08:56.880327+00	\N	6	1	3	1	3		Pollachi Main Road, 642002, Jeeva Nagar, Coimbatore, Tamil Nadu, India	10.67807071651588	77.00412098318338	S.S Kovil Street, Puliampatti, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.657068229143961	77.01024312525988	2020-02-11 07:08:56.605195+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			1	f	t	0	0
26	2020-02-11 07:17:14.296579+00	2020-02-11 07:17:14.296579+00	\N	6	1	3	1	3		Palaghat Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.6637417105339	76.99932754039764	Tiruneelakandar Street, Puliampatti, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.652663862173014	77.00833976268768	2020-02-11 07:17:13.814503+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			6	f	t	0	0
28	2020-02-11 07:31:51.984001+00	2020-02-11 07:31:51.984001+00	\N	6	1	3	1	3		106, State Highway 78A, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.662870548461921	76.999098546803	201, Market Road, Vinayagar Kovil, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.654407230726248	77.00064349919558	2020-02-11 07:31:51.493125+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			6	f	t	0	0
25	2020-02-11 07:11:29.209206+00	2020-02-11 07:11:29.209206+00	\N	6	1	0	1	3		State Highway 78A, 642005, Nallur, Coimbatore, Tamil Nadu, India	10.668521196035055	76.9827052205801	Raja Mill Road, Vinayagar Kovil, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.660227393207643	77.00739562511444	2020-02-11 07:11:28.909404+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	0	0	0	0	0	0	f		0	0	0			6	f	t	0	0
27	2020-02-11 07:25:03.411858+00	2020-02-11 07:25:03.411858+00	\N	6	1	3	1	3		642001, Bodipalayam village, Coimbatore, Tamil Nadu, India	10.665119617614392	76.99446368962526	Pollachi - Udumalpet Road, 642001, Makkinampatti, Coimbatore, Tamil Nadu, India	10.655391433816519	77.02121436595917	2020-02-11 07:25:02.909304+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			6	f	t	0	0
29	2020-02-11 07:36:42.692206+00	2020-02-11 07:36:42.692206+00	\N	6	1	3	1	3		3, State Highway 78A, 642001, Bodipalayam village, Coimbatore, Tamil Nadu, India	10.666918926862525	76.98888435959816	Pollachi - Valparai Road, BK Kovail Street, 642006, Pollachi, Coimbatore, Tamil Nadu, India	10.649880250951052	77.0073664560914	2020-02-11 07:36:42.260687+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			6	f	t	0	0
30	2020-02-11 07:39:35.990662+00	2020-02-11 07:39:35.990662+00	\N	6	1	3	1	3		Vinayagar Kovil, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.664810231560676	76.99566464871168	Perumal Chetty Street, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.661267590553951	77.00739495456219	2020-02-11 07:39:35.532222+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			1	f	t	0	0
31	2020-02-11 09:42:32.711921+00	2020-02-11 09:42:32.711921+00	\N	6	1	0	1	3		Palakkad - Pollachi Road, 642005, Zamin Muthur, Coimbatore, Tamil Nadu, India	10.670682252683948	76.9731729850173	642002, Jeeva Nagar, Coimbatore, Tamil Nadu, India	10.687924687294782	77.00445525348186	2020-02-11 09:42:32.295646+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	0	0	0	0	0	0	f		0	0	0			1	f	t	0	0
32	2020-02-11 10:23:51.112263+00	2020-02-11 10:23:51.112263+00	\N	6	1	3	1	3		106, State Highway 78A, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.663426062839818	77.00053587555885	642001, Nallur, Coimbatore, Tamil Nadu, India	10.66761512258618	76.98803309351206	2020-02-11 10:23:50.643614+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			1	f	t	0	0
33	2020-02-11 10:25:46.316051+00	2020-02-11 10:25:46.316051+00	\N	6	1	3	1	3		South V V Naidu Street, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.662064622478638	77.00281340628862	Pollachi Main Road, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.672439356031912	77.00624663382769	2020-02-11 10:25:45.842197+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			1	f	t	0	0
34	2020-02-11 10:29:32.298293+00	2020-02-11 10:29:32.298293+00	\N	6	1	3	1	3		135, Palakkad - Pollachi Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.661984886431322	77.00390573590994	T . Kottampatty, Mahalakshmi Nagar, 642002, Pollachi, Coimbatore, Tamil Nadu, India	10.673385281643688	77.0199554041028	2020-02-11 10:29:31.899207+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			1	f	t	0	0
35	2020-02-11 10:35:43.766382+00	2020-02-11 10:35:43.766382+00	\N	6	1	3	1	3		135, Palakkad - Pollachi Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.662517338318237	77.00468625873327	642001, Vadugapalayam, Coimbatore, Tamil Nadu, India	10.668731734169974	76.99141468852758	2020-02-11 10:35:43.365755+00	2020-02-11 10:35:54.010786+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			2	f	t	0	0
36	2020-02-11 10:48:22.047946+00	2020-02-11 10:48:22.047946+00	\N	6	1	3	1	3		State Highway 78A, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.663158190374954	76.99047423899174	Bharathi Street, Arumugam Nagar, 642002, Pollachi, Coimbatore, Tamil Nadu, India	10.67065523545213	77.00952798128128	2020-02-11 10:48:21.616924+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			6	f	t	0	0
37	2020-02-11 10:51:49.39039+00	2020-02-11 10:51:49.39039+00	\N	6	1	3	1	3		135, Palakkad - Pollachi Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.662782904938396	77.00383432209492	Ramakrishna Puram Street, 642001, Mahalingapuram, Coimbatore, Tamil Nadu, India	10.66044057299776	77.01233390718699	2020-02-11 10:51:48.887782+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			1	f	t	0	0
38	2020-02-11 10:59:28.183408+00	2020-02-11 10:59:28.183408+00	\N	6	1	1	1	0		Palakkad - Koduvayur - Meenakshipuram Highway, 642103, Meenakshipuram, Coimbatore, Tamil Nadu, India	10.63342683652533	76.86228509992361	Krishna Anaicut Road, 642005, Bodipalayam village, Coimbatore, Tamil Nadu, India	10.664071525780956	76.97805024683475	2020-02-11 10:59:27.799389+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	2	0	0	0	0	0	f		0	0	0			5	f	f	0	0
39	2020-02-11 11:00:40.163507+00	2020-02-11 11:00:40.163507+00	\N	6	1	0	1	3		Unnamed Road, 642005, Coimbatore, Tamil Nadu, India	10.67591234817021	76.98031939566135	Jothi Nagar, 642006, Pollachi, Coimbatore, Tamil Nadu, India	10.643642734329696	77.01196510344744	2020-02-11 11:00:39.695091+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	0	0	0	0	0	0	f		0	0	0			1	f	f	0	0
40	2020-02-16 17:13:18.675852+00	2020-02-16 17:13:18.675852+00	\N	12	3	0	1	0		51, Emilii Plater, Śródmieście, 00-669, Warszawa, Warszawa, mazowieckie, Poland	52.23353077326882	21.002691686153412	Gwiaździsta, Żoliborz, 01-651, Warszawa, Warszawa, mazowieckie, Poland	52.27614176494351	20.987934172153473	2020-02-16 17:13:18.321813+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		5	0	0	0	0	0	0	f		0	0	0			5	f	f	0	0
41	2020-02-16 17:18:14.111383+00	2020-02-16 17:18:14.111383+00	\N	12	3	0	1	0		51, Emilii Plater, Śródmieście, 00-669, Warszawa, Warszawa, mazowieckie, Poland	52.233551717743005	21.002691686153412	51, Emilii Plater, Śródmieście, 00-669, Warszawa, Warszawa, mazowieckie, Poland	52.23356218997638	21.002691686153412	2020-02-16 17:18:13.735328+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		5	0	0	0	0	0	0	f		0	0	0			5	f	f	0	0
42	2020-02-20 08:57:08.793864+00	2020-02-20 08:57:08.793864+00	\N	6	1	3	1	3		24, Palaghat Road, 642001, Bodipalayam village, Coimbatore, Tamil Nadu, India	10.665054709171276	76.99341829866171	Pollachi Main Road, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.673236358684415	77.0059210807085	2020-02-20 08:57:08.430509+00	2020-02-20 08:59:52.115756+00	2020-02-20 08:59:57.555586+00	2020-02-20 09:00:06.206527+00	0	f	0	0.15	8.650941065s	1	3	0	0.15	0	0	2.02	f		22.17	0	0			4	f	t	0	0
43	2020-02-21 08:10:56.604371+00	2020-02-21 08:10:56.604371+00	\N	6	1	0	1	3		Palakkad - Pollachi Road, 642004, Nallur, Coimbatore, Tamil Nadu, India	10.666998661616832	76.98403056710958	135, Palakkad - Pollachi Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.662206631668216	77.005066126585	2020-02-21 08:10:56.070633+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	0	0	0	0	0	0	f		0	0	0			5	f	t	0	0
44	2020-02-21 08:14:39.943591+00	2020-02-21 08:14:39.943591+00	\N	6	1	0	1	3		Palakkad - Pollachi Road, 642004, Nallur, Coimbatore, Tamil Nadu, India	10.666998661616832	76.98403056710958	135, Palakkad - Pollachi Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.662206631668216	77.005066126585	2020-02-21 08:14:39.712781+00	2020-02-21 08:18:34.984711+00	2020-02-21 08:18:42.332622+00	2020-02-21 08:19:44.064665+00	0	f	0	1.03	1m1.732043023s	1	0	0	20.6	0	0	4.07	f		44.67	0	0			4	f	t	0	0
45	2020-02-21 08:37:55.1407+00	2020-02-21 08:37:55.1407+00	\N	6	1	3	1	3		135, Palakkad - Pollachi Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.662206631668216	77.005066126585	96, Palaghat Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.66314830577764	76.99749860912561	2020-02-21 08:37:54.745716+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			5	f	t	0	0
46	2020-02-21 08:41:11.171081+00	2020-02-21 08:41:11.171081+00	\N	6	1	3	1	3		135, Palakkad - Pollachi Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.662206631668216	77.005066126585	96, Palaghat Road, Palaghat, 642001, Pollachi, Coimbatore, Tamil Nadu, India	10.66314830577764	76.99749860912561	2020-02-21 08:41:10.978939+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0001-01-01 00:00:00+00	0	f	0	0		1	3	0	0	0	0	0	f		0	0	0			5	f	t	0	0
\.


--
-- Data for Name: sent_ride_requests; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.sent_ride_requests (id, created_at, updated_at, deleted_at, driver_id, ride_id, is_active) FROM stdin;
1	2020-02-02 18:58:23.441699+00	2020-02-02 18:58:23.441699+00	\N	1	1	t
2	2020-02-03 07:36:15.372767+00	2020-02-03 07:36:15.372767+00	\N	1	2	t
3	2020-02-03 08:36:46.955431+00	2020-02-03 08:36:46.955431+00	\N	1	3	t
4	2020-02-03 08:37:11.811882+00	2020-02-03 08:37:11.811882+00	\N	3	3	f
5	2020-02-03 16:51:03.637984+00	2020-02-03 16:51:03.637984+00	\N	17	4	t
6	2020-02-07 11:37:41.662028+00	2020-02-07 11:37:41.662028+00	\N	3	9	t
7	2020-02-07 11:39:09.213259+00	2020-02-07 11:39:09.213259+00	\N	1	10	t
8	2020-02-07 11:39:34.095518+00	2020-02-07 11:39:34.095518+00	\N	3	10	f
9	2020-02-07 11:42:49.491918+00	2020-02-07 11:42:49.491918+00	\N	1	11	t
10	2020-02-07 11:43:14.374266+00	2020-02-07 11:43:14.374266+00	\N	3	11	f
11	2020-02-07 11:47:31.620286+00	2020-02-07 11:47:31.620286+00	\N	3	12	t
12	2020-02-07 11:49:40.906647+00	2020-02-07 11:49:40.906647+00	\N	1	13	t
13	2020-02-07 11:50:06.083578+00	2020-02-07 11:50:06.083578+00	\N	3	13	f
14	2020-02-07 11:53:00.387321+00	2020-02-07 11:53:00.387321+00	\N	1	14	t
15	2020-02-07 11:53:25.267654+00	2020-02-07 11:53:25.267654+00	\N	3	14	f
16	2020-02-07 11:56:50.85316+00	2020-02-07 11:56:50.85316+00	\N	1	15	t
17	2020-02-07 11:57:15.734994+00	2020-02-07 11:57:15.734994+00	\N	3	15	f
18	2020-02-11 04:10:32.366633+00	2020-02-11 04:10:32.366633+00	\N	1	16	t
19	2020-02-11 04:10:57.228591+00	2020-02-11 04:10:57.228591+00	\N	3	16	f
20	2020-02-11 04:48:22.704032+00	2020-02-11 04:48:22.704032+00	\N	3	17	t
21	2020-02-11 04:52:18.085043+00	2020-02-11 04:52:18.085043+00	\N	1	18	t
22	2020-02-11 04:52:42.943912+00	2020-02-11 04:52:42.943912+00	\N	3	18	f
23	2020-02-11 06:15:07.889494+00	2020-02-11 06:15:07.889494+00	\N	1	19	t
24	2020-02-11 06:15:32.759451+00	2020-02-11 06:15:32.759451+00	\N	3	19	f
25	2020-02-11 06:28:46.049603+00	2020-02-11 06:28:46.049603+00	\N	1	20	t
26	2020-02-11 06:29:10.917893+00	2020-02-11 06:29:10.917893+00	\N	3	20	f
27	2020-02-11 06:32:31.07151+00	2020-02-11 06:32:31.07151+00	\N	3	21	t
28	2020-02-11 06:43:09.129873+00	2020-02-11 06:43:09.129873+00	\N	1	22	t
29	2020-02-11 06:43:34.017026+00	2020-02-11 06:43:34.017026+00	\N	3	22	f
30	2020-02-11 07:07:21.653907+00	2020-02-11 07:07:21.653907+00	\N	3	23	t
31	2020-02-11 07:08:58.121201+00	2020-02-11 07:08:58.121201+00	\N	3	24	t
32	2020-02-11 07:11:30.450897+00	2020-02-11 07:11:30.450897+00	\N	1	25	t
33	2020-02-11 07:11:55.310162+00	2020-02-11 07:11:55.310162+00	\N	3	25	f
34	2020-02-11 07:17:15.53729+00	2020-02-11 07:17:15.53729+00	\N	3	26	t
35	2020-02-11 07:25:04.654615+00	2020-02-11 07:25:04.654615+00	\N	1	27	t
36	2020-02-11 07:25:29.514445+00	2020-02-11 07:25:29.514445+00	\N	3	27	f
37	2020-02-11 07:31:53.224146+00	2020-02-11 07:31:53.224146+00	\N	3	28	t
38	2020-02-11 07:36:43.935959+00	2020-02-11 07:36:43.935959+00	\N	1	29	t
39	2020-02-11 07:37:08.795254+00	2020-02-11 07:37:08.795254+00	\N	3	29	f
40	2020-02-11 07:39:37.233709+00	2020-02-11 07:39:37.233709+00	\N	1	30	t
41	2020-02-11 07:40:02.093546+00	2020-02-11 07:40:02.093546+00	\N	3	30	f
42	2020-02-11 09:42:33.985797+00	2020-02-11 09:42:33.985797+00	\N	1	31	t
43	2020-02-11 09:42:58.866699+00	2020-02-11 09:42:58.866699+00	\N	3	31	f
44	2020-02-11 10:23:52.36109+00	2020-02-11 10:23:52.36109+00	\N	3	32	t
45	2020-02-11 10:25:47.564489+00	2020-02-11 10:25:47.564489+00	\N	3	33	t
46	2020-02-11 10:29:33.545585+00	2020-02-11 10:29:33.545585+00	\N	3	34	t
47	2020-02-11 10:35:45.026066+00	2020-02-11 10:35:45.026066+00	\N	3	35	t
48	2020-02-11 10:48:23.310374+00	2020-02-11 10:48:23.310374+00	\N	1	36	t
49	2020-02-11 10:48:48.182369+00	2020-02-11 10:48:48.182369+00	\N	3	36	f
50	2020-02-11 10:51:50.657749+00	2020-02-11 10:51:50.657749+00	\N	3	37	t
51	2020-02-11 11:00:41.036883+00	2020-02-11 11:00:41.036883+00	\N	1	39	t
52	2020-02-11 11:01:05.899895+00	2020-02-11 11:01:05.899895+00	\N	3	39	f
53	2020-02-20 08:57:10.045659+00	2020-02-20 08:57:10.045659+00	\N	1	42	t
54	2020-02-20 08:57:34.908814+00	2020-02-20 08:57:34.908814+00	\N	3	42	f
55	2020-02-21 08:10:57.990599+00	2020-02-21 08:10:57.990599+00	\N	3	43	t
56	2020-02-21 08:14:41.203995+00	2020-02-21 08:14:41.203995+00	\N	3	44	t
57	2020-02-21 08:37:56.407615+00	2020-02-21 08:37:56.407615+00	\N	1	45	t
58	2020-02-21 08:38:21.285405+00	2020-02-21 08:38:21.285405+00	\N	3	45	f
59	2020-02-21 08:41:12.422322+00	2020-02-21 08:41:12.422322+00	\N	1	46	t
60	2020-02-21 08:41:37.297951+00	2020-02-21 08:41:37.297951+00	\N	3	46	f
\.


--
-- Data for Name: spatial_ref_sys; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.spatial_ref_sys (srid, auth_name, auth_srid, srtext, proj4text) FROM stdin;
\.


--
-- Data for Name: vehicle_categories; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.vehicle_categories (id, created_at, updated_at, deleted_at, name, description, is_active) FROM stdin;
2	2020-02-03 07:52:27.443273+00	2020-02-03 07:52:27.443273+00	\N	Premium		t
3	2020-03-02 16:43:46.309376+00	2020-03-02 16:43:46.309376+00	\N	ch-test-vihicle-type		t
1	2020-02-02 17:01:03.863955+00	2020-02-02 17:01:03.863955+00	\N	Economy		t
\.


--
-- Data for Name: vehicle_types; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.vehicle_types (id, created_at, updated_at, deleted_at, name, image, vehicle_category_id, description, image_active, seat_capacity, is_active) FROM stdin;
1	2020-02-02 17:01:38.72984+00	2020-02-02 17:01:38.72984+00	\N	Taasai Go	public/vehicletype/1580662898631226429_inactive_go.png	1	Affordable and convinient rides	public/vehicletype/1580662898631224629_active_go.png	0	t
2	2020-02-03 07:49:54.426315+00	2020-02-03 07:49:54.426315+00	\N	Taasai Basic	public/vehicletype/1580716194327901309_inactive_go.png	1	Affordable rides	public/vehicletype/1580716194327899268_active_go.png	0	t
3	2020-02-03 07:53:01.923107+00	2020-02-03 07:53:01.923107+00	\N	Taasai Premium	public/vehicletype/1580716381824425916_inactive_go.png	2	Enjoy Luxury rides	public/vehicletype/1580716381824424609_active_go.png	0	t
\.


--
-- Data for Name: zone_fares; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.zone_fares (id, created_at, updated_at, deleted_at, vehicle_type_id, zone_id, base_fare, minimum_fare, waiting_time_limit, waiting_fee, cancellation_time_limit, cancellation_fee, duration_fare, distance_fare, tax, traffic_factor, is_active) FROM stdin;
2	2020-02-03 03:14:54.165308+00	2020-02-03 03:14:54.165308+00	\N	1	1	30	50	5	2	2	10	1	2	10	10	t
1	2020-02-02 20:50:00.10019+00	2020-02-02 20:50:00.10019+00	2020-02-03 03:18:49.627302+00	1	2	2.5	0	0	0	0	0	0.15	1.25	0	0	t
3	2020-02-03 12:23:48.524896+00	2020-02-03 12:23:48.524896+00	\N	1	3	20	25	1	10	1	10	1	2	10	10	t
4	2020-02-03 12:24:14.957689+00	2020-02-03 12:24:14.957689+00	\N	2	3	10	15	1	10	1	10	1	2	10	10	t
5	2020-02-03 12:24:42.153736+00	2020-02-03 12:24:42.153736+00	\N	3	3	30	40	2	10	1	10	2	5	10	15	t
\.


--
-- Data for Name: zones; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.zones (id, created_at, updated_at, deleted_at, name, operator_id, is_active, polygon) FROM stdin;
1	\N	\N	\N	Meenakshipuram Railway station	1	t	01030000000100000009000000EF6FD05E7D4C25408410902FA1375340B00111E2CA49254047020D367538534052EE3EC7474325407F2F8507CD38534036CAFACDC43C25405F44DB3175385340F75B3B51123A25408410902FA137534036CAFACDC43C2540A9DC442DCD36534052EE3EC74743254071AFCC5B75365340B00111E2CA492540C11E1329CD365340EF6FD05E7D4C25408410902FA1375340
2	\N	\N	\N	Heathrow	2	t	01030000000100000009000000AB93331477BE49400B08AD872F13DDBF7BC03C64CABD49409D83674293C4DBBFD3BF249529BC4940FF058200193ADBBF9C3237DF88BA4940917D9065C1C4DBBFAD4ECE50DCB949400B08AD872F13DDBF9C3237DF88BA494058AA0B789961DEBFD3BF249529BC4940EA211ADD41ECDEBF7BC03C64CABD49404CA4349BC761DEBFAB93331477BE49400B08AD872F13DDBF
3	\N	\N	\N	Pollachi Bus Stand	1	t	0103000000010000000A000000C896E5EB325C2540DB148F8B6A4053408928266F8059254087C43D963E4153402B155454FD522540A7AFE76B964153400FF10F5B7A4C25409E060C923E4153401172DEFFC7492540DB148F8B6A4053400FF10F5B7A4C254018231285963F5340BE8575E3DD512540832F4CA60A3F534049D74CBED9562540828DEBDFF53E53408928266F805925403065E080963F5340C896E5EB325C2540DB148F8B6A405340
\.


--
-- Name: admins_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.admins_id_seq', 1, true);


--
-- Name: driver_document_uploads_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.driver_document_uploads_id_seq', 37, true);


--
-- Name: driver_documents_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.driver_documents_id_seq', 6, true);


--
-- Name: drivers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.drivers_id_seq', 54, true);


--
-- Name: fares_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.fares_id_seq', 5, true);


--
-- Name: fleets_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.fleets_id_seq', 30, true);


--
-- Name: operators_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.operators_id_seq', 3, true);


--
-- Name: otps_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.otps_id_seq', 250, true);


--
-- Name: passengers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.passengers_id_seq', 13, true);


--
-- Name: pickup_points_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.pickup_points_id_seq', 7, true);


--
-- Name: ride_event_logs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.ride_event_logs_id_seq', 151, true);


--
-- Name: ride_locations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.ride_locations_id_seq', 1, false);


--
-- Name: ride_messages_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.ride_messages_id_seq', 66, true);


--
-- Name: ride_stops_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.ride_stops_id_seq', 28, true);


--
-- Name: rides_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.rides_id_seq', 46, true);


--
-- Name: sent_ride_requests_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.sent_ride_requests_id_seq', 60, true);


--
-- Name: vehicle_categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.vehicle_categories_id_seq', 3, true);


--
-- Name: vehicle_types_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.vehicle_types_id_seq', 3, true);


--
-- Name: zone_fares_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.zone_fares_id_seq', 5, true);


--
-- Name: zones_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.zones_id_seq', 3, true);


--
-- Name: admins admins_email_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.admins
    ADD CONSTRAINT admins_email_key UNIQUE (email);


--
-- Name: admins admins_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.admins
    ADD CONSTRAINT admins_pkey PRIMARY KEY (id);


--
-- Name: driver_document_uploads driver_document_uploads_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.driver_document_uploads
    ADD CONSTRAINT driver_document_uploads_pkey PRIMARY KEY (id);


--
-- Name: driver_documents driver_documents_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.driver_documents
    ADD CONSTRAINT driver_documents_pkey PRIMARY KEY (id);


--
-- Name: drivers drivers_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.drivers
    ADD CONSTRAINT drivers_pkey PRIMARY KEY (id);


--
-- Name: drivers drivers_vehicle_number_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.drivers
    ADD CONSTRAINT drivers_vehicle_number_key UNIQUE (vehicle_number);


--
-- Name: fares fares_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.fares
    ADD CONSTRAINT fares_pkey PRIMARY KEY (id);


--
-- Name: fleets fleets_email_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.fleets
    ADD CONSTRAINT fleets_email_key UNIQUE (email);


--
-- Name: fleets fleets_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.fleets
    ADD CONSTRAINT fleets_pkey PRIMARY KEY (id);


--
-- Name: operators operators_email_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.operators
    ADD CONSTRAINT operators_email_key UNIQUE (email);


--
-- Name: operators operators_location_name_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.operators
    ADD CONSTRAINT operators_location_name_key UNIQUE (location_name);


--
-- Name: operators operators_name_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.operators
    ADD CONSTRAINT operators_name_key UNIQUE (name);


--
-- Name: operators operators_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.operators
    ADD CONSTRAINT operators_pkey PRIMARY KEY (id);


--
-- Name: otps otps_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.otps
    ADD CONSTRAINT otps_pkey PRIMARY KEY (id);


--
-- Name: passengers passengers_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.passengers
    ADD CONSTRAINT passengers_pkey PRIMARY KEY (id);


--
-- Name: pickup_points pickup_points_name_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.pickup_points
    ADD CONSTRAINT pickup_points_name_key UNIQUE (name);


--
-- Name: pickup_points pickup_points_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.pickup_points
    ADD CONSTRAINT pickup_points_pkey PRIMARY KEY (id);


--
-- Name: ride_event_logs ride_event_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.ride_event_logs
    ADD CONSTRAINT ride_event_logs_pkey PRIMARY KEY (id);


--
-- Name: ride_locations ride_locations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.ride_locations
    ADD CONSTRAINT ride_locations_pkey PRIMARY KEY (id);


--
-- Name: ride_messages ride_messages_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.ride_messages
    ADD CONSTRAINT ride_messages_pkey PRIMARY KEY (id);


--
-- Name: ride_stops ride_stops_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.ride_stops
    ADD CONSTRAINT ride_stops_pkey PRIMARY KEY (id);


--
-- Name: rides rides_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.rides
    ADD CONSTRAINT rides_pkey PRIMARY KEY (id);


--
-- Name: sent_ride_requests sent_ride_requests_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sent_ride_requests
    ADD CONSTRAINT sent_ride_requests_pkey PRIMARY KEY (id);


--
-- Name: vehicle_categories vehicle_categories_name_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.vehicle_categories
    ADD CONSTRAINT vehicle_categories_name_key UNIQUE (name);


--
-- Name: vehicle_categories vehicle_categories_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.vehicle_categories
    ADD CONSTRAINT vehicle_categories_pkey PRIMARY KEY (id);


--
-- Name: vehicle_types vehicle_types_name_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.vehicle_types
    ADD CONSTRAINT vehicle_types_name_key UNIQUE (name);


--
-- Name: vehicle_types vehicle_types_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.vehicle_types
    ADD CONSTRAINT vehicle_types_pkey PRIMARY KEY (id);


--
-- Name: zone_fares zone_fares_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.zone_fares
    ADD CONSTRAINT zone_fares_pkey PRIMARY KEY (id);


--
-- Name: zones zones_name_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.zones
    ADD CONSTRAINT zones_name_key UNIQUE (name);


--
-- Name: zones zones_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.zones
    ADD CONSTRAINT zones_pkey PRIMARY KEY (id);


--
-- Name: idx_admin_email; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_admin_email ON public.admins USING btree (email);


--
-- Name: idx_admins_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_admins_deleted_at ON public.admins USING btree (deleted_at);


--
-- Name: idx_driver_document_uploads_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_driver_document_uploads_deleted_at ON public.driver_document_uploads USING btree (deleted_at);


--
-- Name: idx_driver_documents_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_driver_documents_deleted_at ON public.driver_documents USING btree (deleted_at);


--
-- Name: idx_drivers_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_drivers_deleted_at ON public.drivers USING btree (deleted_at);


--
-- Name: idx_fare; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_fare ON public.zone_fares USING btree (vehicle_type_id, zone_id);


--
-- Name: idx_fares_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_fares_deleted_at ON public.fares USING btree (deleted_at);


--
-- Name: idx_fleets_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_fleets_deleted_at ON public.fleets USING btree (deleted_at);


--
-- Name: idx_operators_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_operators_deleted_at ON public.operators USING btree (deleted_at);


--
-- Name: idx_otps_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_otps_deleted_at ON public.otps USING btree (deleted_at);


--
-- Name: idx_passengers_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_passengers_deleted_at ON public.passengers USING btree (deleted_at);


--
-- Name: idx_pickup_points_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_pickup_points_deleted_at ON public.pickup_points USING btree (deleted_at);


--
-- Name: idx_ride; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_ride ON public.rides USING btree (passenger_id, operator_id, vehicle_type_id, driver_id, fare_id, zone_fare_id);


--
-- Name: idx_ride_event_logs_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_ride_event_logs_deleted_at ON public.ride_event_logs USING btree (deleted_at);


--
-- Name: idx_ride_locations_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_ride_locations_deleted_at ON public.ride_locations USING btree (deleted_at);


--
-- Name: idx_ride_messages_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_ride_messages_deleted_at ON public.ride_messages USING btree (deleted_at);


--
-- Name: idx_ride_stops_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_ride_stops_deleted_at ON public.ride_stops USING btree (deleted_at);


--
-- Name: idx_rides_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_rides_deleted_at ON public.rides USING btree (deleted_at);


--
-- Name: idx_sent_ride_requests_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_sent_ride_requests_deleted_at ON public.sent_ride_requests USING btree (deleted_at);


--
-- Name: idx_vehicle; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_vehicle ON public.drivers USING btree (vehicle_type_id);


--
-- Name: idx_vehicle_categories_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_vehicle_categories_deleted_at ON public.vehicle_categories USING btree (deleted_at);


--
-- Name: idx_vehicle_types_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_vehicle_types_deleted_at ON public.vehicle_types USING btree (deleted_at);


--
-- Name: idx_zone_fares_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_zone_fares_deleted_at ON public.zone_fares USING btree (deleted_at);


--
-- Name: idx_zones_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_zones_deleted_at ON public.zones USING btree (deleted_at);


--
-- Name: ride_locations_latlng_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX ride_locations_latlng_idx ON public.ride_locations USING gist (latlng);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: -
--

REVOKE ALL ON SCHEMA public FROM cloudsqladmin;
REVOKE ALL ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO cloudsqlsuperuser;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

