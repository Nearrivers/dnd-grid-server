-- +goose Up
CREATE TABLE "levels_entities" (
	"level_id"	INTEGER NOT NULL,
	"entity_id"	INTEGER NOT NULL,
	"health_points"	INTEGER NOT NULL,
	"x_coord"	INTEGER NOT NULL,
	"y_coord"	INTEGER NOT NULL,
	"number"	INTEGER NOT NULL,
	PRIMARY KEY("level_id","entity_id"),
	FOREIGN KEY("level_id") REFERENCES levels(id),
	FOREIGN KEY("entity_id") REFERENCES entities(id)
);
