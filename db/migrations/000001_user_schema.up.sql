CREATE TABLE "users" (
    "username" varchar PRIMARY KEY,
    "full_name" varchar NOT NULL,
    "phone_number" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);