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
    PRIMARY KEY ("id")
);

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS scores_id_seq;

-- Table Definition
CREATE TABLE "public"."scores" (
    "id" int4 NOT NULL DEFAULT nextval('scores_id_seq'::regclass),
    "username" varchar NOT NULL,
    "score" int2 NOT NULL,
    PRIMARY KEY ("id")
);

