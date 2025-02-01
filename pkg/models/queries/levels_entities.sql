-- name: NewEntityLevel :exec
INSERT INTO "levels_entities" (level_id, entity_id, health_points, x_coord, y_coord, number) VALUES (?, ?, ?, ?, ?, ?);
