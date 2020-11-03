CREATE SEQUENCE IF NOT EXISTS questions_id_seq;

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

CREATE TABLE "public"."scores" (
    "username" varchar NOT NULL,
    "score" int2 NOT NULL
);

