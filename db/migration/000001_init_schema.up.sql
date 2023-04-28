DO $$
BEGIN
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role') THEN
        CREATE TYPE "role" AS ENUM (
          'admin',
          'user',
          'owner'
        );
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'enum_gender') THEN
        CREATE TYPE "enum_gender" AS ENUM (
          'male',
          'female',
          'other'
        );
    END IF;
    --more types here...
END$$;

CREATE TABLE "users" (
    "id" UUID NOT NULL DEFAULT (uuid_generate_v4()) ,
    "email" VARCHAR NOT NULL ,
    "verified" BOOLEAN NOT NULL,
    "password" VARCHAR NOT NULL,
    "role" role DEFAULT 'user',
    "created_at" TIMESTAMP NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMP NOT NULL,

    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "users_email_key" ON "users"("email");

CREATE TABLE "user_detail" (
  "id" bigserial PRIMARY KEY,
  "user_id" UUID NOT NULL,
  "name" varchar,
  "gender" enum_gender,
  "age" int,
  "address" varchar,
  "phone" int
);

CREATE TABLE "payment" (
  "id" bigserial PRIMARY KEY,
  "payment_method" varchar NOT NULL,
  "payment_status" varchar NOT NULL,
  "started_at" timestamp NOT NULL DEFAULT (now()),
  "end_at" timestamp NOT NULL,
  "transaction_id" bigint NOT NULL
);

CREATE TABLE "pets" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "age" int NOT NULL,
  "gender" enum_gender NOT NULL,
  "images" varchar,
  "movies" varchar,
  "contraception" bool,
  "condition" varchar,
  "owner_id" bigint NOT NULL,
  "transaction_id" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "transaction" (
  "id" bigserial PRIMARY KEY,
  "transaction_status" varchar,
  "payment_type" varchar,
  "pet_id" varchar
);

CREATE TABLE "transaction_detail" (
  "id" bigserial PRIMARY KEY,
  "name" varchar,
  "price" int,
  "transaction_id" bigint NOT NULL
);

CREATE TABLE "pet_tag" (
  "id" bigserial PRIMARY KEY,
  "pet_id" bigint NOT NULL,
  "tag_id" bigint NOT NULL
);

CREATE TABLE "fav_food" (
  "id" bigserial PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "pet_fav_food" (
  "id" bigserial PRIMARY KEY,
  "pet_id" bigint NOT NULL,
  "food_id" bigint NOT NULL
);

CREATE TABLE "tags" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "immunization" (
  "id" bigserial PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "pet_immunization" (
  "id" bigserial PRIMARY KEY,
  "pet_id" bigint NOT NULL,
  "immunization_id" bigint NOT NULL
);

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "user_detail" ("name");

CREATE INDEX ON "user_detail" ("address");

CREATE INDEX ON "user_detail" ("phone");

CREATE INDEX ON "pets" ("name");

CREATE INDEX ON "tags" ("name");

ALTER TABLE "user_detail" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "payment" ADD FOREIGN KEY ("transaction_id") REFERENCES "transaction" ("id");

ALTER TABLE "pets" ADD FOREIGN KEY ("owner_id") REFERENCES "user_detail" ("id");

ALTER TABLE "pets" ADD FOREIGN KEY ("transaction_id") REFERENCES "transaction" ("id");

ALTER TABLE "transaction_detail" ADD FOREIGN KEY ("transaction_id") REFERENCES "transaction" ("id");

ALTER TABLE "pet_tag" ADD FOREIGN KEY ("pet_id") REFERENCES "pets" ("id");

ALTER TABLE "pet_tag" ADD FOREIGN KEY ("tag_id") REFERENCES "tags" ("id");

ALTER TABLE "pet_fav_food" ADD FOREIGN KEY ("pet_id") REFERENCES "pets" ("id");

ALTER TABLE "pet_fav_food" ADD FOREIGN KEY ("food_id") REFERENCES "fav_food" ("id");

ALTER TABLE "pet_immunization" ADD FOREIGN KEY ("pet_id") REFERENCES "pets" ("id");

ALTER TABLE "pet_immunization" ADD FOREIGN KEY ("immunization_id") REFERENCES "immunization" ("id");
