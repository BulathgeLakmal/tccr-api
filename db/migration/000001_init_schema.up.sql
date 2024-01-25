CREATE TABLE "user" (
  "user_id" bigserial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "role" varchar NOT NULL,
  "username" varchar NOT NULL
);

CREATE TABLE "userDetails" (
  "user_details_id" bigserial PRIMARY KEY,
  "user_id" bigint,
  "phone" integer NOT NULL,
  "address_line1" varchar NOT NULL,
  "address_line2" varchar NOT NULL
);

CREATE TABLE "userRole" (
  "role_id" bigserial PRIMARY KEY,
  "role" varchar NOT NULL
);

CREATE TABLE "course" (
  "course_id" bigserial PRIMARY KEY,
  "course_name" varchar NOT NULL,
  "course_desc" varchar NOT NULL,
  "category" varchar NOT NULL
);

CREATE TABLE "category" (
  "category_id" bigserial PRIMARY KEY,
  "category_name" varchar NOT NULL,
  "category_desc" varchar NOT NULL
);

CREATE TABLE "courseModule" (
  "module_id" bigserial PRIMARY KEY,
  "course_id" bigint,
  "module_name" varchar NOT NULL
);

CREATE TABLE "lectures" (
  "lecture_id" bigserial PRIMARY KEY,
  "course_module_id" bigint,
  "lecture_desc" varchar NOT NULL,
  "lecture_number" integer,
  "video_URL" varchar NOT NULL,
  "status" varchar NOT NULL
);

CREATE TABLE "assignment" (
  "assignment_id" bigserial PRIMARY KEY,
  "user_id" bigint,
  "course_module_id" bigint,
  "lecture_id" bigint
);

CREATE TABLE "assignment_file" (
  "assignment_file_id" bigserial PRIMARY KEY,
  "assignment_id" bigint,
  "assignment_link" varchar NOT NULL
);

CREATE TABLE "subscribe" (
  "subscribe_id" bigserial PRIMARY KEY,
  "user_id" bigint,
  "course_id" bigint
);

ALTER TABLE "user" ADD FOREIGN KEY ("role_id") REFERENCES "userRole" ("role");

ALTER TABLE "userDetails" ADD FOREIGN KEY ("user_ID") REFERENCES "user" ("user_id");

ALTER TABLE "course" ADD FOREIGN KEY ("category") REFERENCES "category" ("category_id");

ALTER TABLE "courseModule" ADD FOREIGN KEY ("course_id") REFERENCES "course" ("course_id");

ALTER TABLE "lectures" ADD FOREIGN KEY ("course_module_id") REFERENCES "courseModule" ("module_id");

ALTER TABLE "assignment" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");

ALTER TABLE "assignment" ADD FOREIGN KEY ("course_module_id") REFERENCES "courseModule" ("module_id");

ALTER TABLE "assignment" ADD FOREIGN KEY ("lecture_id") REFERENCES "lectures" ("lecture_id");

ALTER TABLE "assignment_file" ADD FOREIGN KEY ("assignment_id") REFERENCES "assignment" ("assignment_id");

ALTER TABLE "subscribe" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");

ALTER TABLE "subscribe" ADD FOREIGN KEY ("course_id") REFERENCES "course" ("course_id");
