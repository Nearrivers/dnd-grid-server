-- +goose Up
CREATE TABLE "entities" (
	"id"	INTEGER PRIMARY KEY,
	"name"	TEXT NOT NULL,
	"faction" TEXT NOT NULL
);
