CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "username" varchar NOT NULL,
                         "email" varchar NOT NULL,
                         "birthdate" date NOT NULL,
                         "remember_hash" varchar NOT NULL,
                         "password_hash" varchar NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT 'now()',
                         "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "channels" (
                            "id" bigserial PRIMARY KEY,
                            "name" varchar NOT NULL,
                            "description" varchar NOT NULL,
                            "video" varchar[],
                            "user_id" bigint NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT 'now()',
                            "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "comments" (
                            "id" bigserial PRIMARY KEY,
                            "content" varchar NOT NULL,
                            "user_id" bigint NOT NULL,
                            "video_id" bigint NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT 'now()',
                            "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "followers" (
                             "id" bigserial PRIMARY KEY,
                             "leader_id" bigint NOT NULL,
                             "follower_id" bigint NOT NULL,
                             "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "videos" (
                          "id" bigserial PRIMARY KEY,
                          "title" varchar NOT NULL,
                          "url" varchar NOT NULL,
                          "user_id" bigint NOT NULL,
                          "channel_id" bigint NOT NULL,
                          "created_at" timestamptz NOT NULL DEFAULT 'now()',
                          "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

ALTER TABLE "channels" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("video_id") REFERENCES "videos" ("id");

ALTER TABLE "followers" ADD FOREIGN KEY ("leader_id") REFERENCES "users" ("id");

ALTER TABLE "followers" ADD FOREIGN KEY ("follower_id") REFERENCES "users" ("id");

ALTER TABLE "videos" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "videos" ADD FOREIGN KEY ("channel_id") REFERENCES "channels" ("id");
