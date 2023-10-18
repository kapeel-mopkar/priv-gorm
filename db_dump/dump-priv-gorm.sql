-- Table structure for table "privileges"

DROP TABLE IF EXISTS "privileges";

CREATE TABLE "privileges" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(100)
);

-- Dumping data for table "privileges"

INSERT INTO "privileges" ("id", "name") VALUES (1,'Admin'),(2,'ReadUser'),(3,'User'),(4,'WriteUser');

-- Table structure for table "scoped_privileges"

DROP TABLE IF EXISTS "scoped_privileges";

CREATE TABLE "scoped_privileges" (
  "privilege_name" varchar(100) PRIMARY KEY
);

-- Dumping data for table "scoped_privileges"

INSERT INTO "scoped_privileges" ("privilege_name") VALUES ('Admin'),('ReadUser'),('User'),('WriteUser');

-- Table structure for table "user_scoped_privileges"

DROP TABLE IF EXISTS "user_scoped_privileges";

CREATE TABLE "user_scoped_privileges" (
  "user_id" serial NOT NULL,
  "privilege_name" varchar(100) NOT NULL,
  PRIMARY KEY ("user_id","privilege_name")
);

-- Dumping data for table "user_scoped_privileges"

INSERT INTO "user_scoped_privileges" ("user_id", "privilege_name") VALUES (1,'ReadUser'),(1,'WriteUser'),(2,'ReadUser'),(2,'User'),(3,'Admin'),(3,'WriteUser');

-- Table structure for table "users"

DROP TABLE IF EXISTS "users";

CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "email" varchar(100) NOT NULL,
  "name" varchar(50) NOT NULL,
  "last_login" timestamp NOT NULL DEFAULT current_timestamp,
  "created_at" timestamp NOT NULL DEFAULT current_timestamp,
  "updated_at" timestamp NOT NULL DEFAULT current_timestamp
);

-- Dumping data for table "users"

INSERT INTO "users" ("id", "email", "name", "last_login", "created_at", "updated_at") VALUES
(1,'kapeel@viu.com','Kapeel Mopkar','2023-10-18 06:12:03','2023-10-18 06:12:03','2023-10-18 06:12:03'),
(2,'ganesh@viu.com','Ganesh Patil','2023-10-18 06:12:03','2023-10-18 06:12:03','2023-10-18 06:12:03'),
(3,'ashish@viu.com','Ashish Mishra','2023-10-18 06:12:03','2023-10-18 06:12:03','2023-10-18 06:12:03');

-- Table structure for table "users_privileges"

DROP TABLE IF EXISTS "users_privileges";

CREATE TABLE "users_privileges" (
  "user_id" serial NOT NULL,
  "privilege_id" serial NOT NULL
);

-- Dumping data for table "users_privileges"

INSERT INTO "users_privileges" ("user_id", "privilege_id") VALUES (1,2),(1,4),(2,2),(2,3),(3,1),(3,4);
