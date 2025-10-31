/*
 Navicat Premium Data Transfer

 Source Server         : Postgres Bootcamp
 Source Server Type    : PostgreSQL
 Source Server Version : 180000 (180000)
 Source Host           : localhost:5432
 Source Catalog        : bioskop_db
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 180000 (180000)
 File Encoding         : 65001

*/


-- ----------------------------
-- Table structure for bioskop
-- ----------------------------
DROP TABLE IF EXISTS "public"."bioskop";
CREATE TABLE "public"."bioskop" (
  "id" int4 NOT NULL DEFAULT nextval('bioskop_id_seq'::regclass),
  "nama" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "lokasi" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "rating" numeric(3,1) DEFAULT 0.0,
  "created_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP
)
;

-- ----------------------------
-- Primary Key structure for table bioskop
-- ----------------------------
ALTER TABLE "public"."bioskop" ADD CONSTRAINT "bioskop_pkey" PRIMARY KEY ("id");


-- CREATE SEQUENCE bioskop_id_seq;

-- CREATE TABLE bioskop (
--     id INTEGER NOT NULL DEFAULT nextval('bioskop_id_seq'),
--     nama VARCHAR(255) NOT NULL,
--     lokasi VARCHAR(255),
--     rating FLOAT,
--     PRIMARY KEY (id)
-- );

-- ALTER SEQUENCE bioskop_id_seq OWNED BY bioskop.id;