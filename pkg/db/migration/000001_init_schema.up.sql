CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "owner" bigserial NOT NULL,
  "title" varchar NOT NULL,
  "content" varchar NOT NULL,
  "upvotes" bigint NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "upvotes" (
  "userId" bigserial NOT NULL,
  "postId" bigserial NOT NULL,
  "value" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "posts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id");

ALTER TABLE "upvotes" ADD FOREIGN KEY ("userId") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "upvotes" ADD FOREIGN KEY ("postId") REFERENCES "posts" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

CREATE UNIQUE INDEX ON "upvotes" ("userId", "postId");
