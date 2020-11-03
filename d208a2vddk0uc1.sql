-- -------------------------------------------------------------
-- TablePlus 3.7.1(332)
--
-- https://tableplus.com/
--
-- Database: d208a2vddk0uc1
-- Generation Time: 2020-11-03 22:06:46.9120
-- -------------------------------------------------------------


-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS questions_id_seq;

-- Table Definition
CREATE TABLE "public"."questions" (
    "id" int4 NOT NULL DEFAULT nextval('questions_id_seq'::regclass),
    "question" text NOT NULL,
    "first_answer" varchar NOT NULL,
    "second_answer" varchar NOT NULL,
    "third_answer" varchar NOT NULL,
    "fourth_answer" varchar NOT NULL,
    "answer" int2 NOT NULL,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."scores" (
    "username" varchar NOT NULL,
    "score" int2 NOT NULL
);

