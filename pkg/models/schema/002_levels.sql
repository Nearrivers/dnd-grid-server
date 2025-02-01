-- +goose Up
CREATE TABLE "levels" (
	"id"	INTEGER NOT NULL PRIMARY KEY,
	"name"	TEXT NOT NULL,
	"image_path"	TEXT NOT NULL,
	"grid_width"	INTEGER NOT NULL,
	"grid_color"	TEXT NOT NULL,
	"grid_spacing"	INTEGER
);
