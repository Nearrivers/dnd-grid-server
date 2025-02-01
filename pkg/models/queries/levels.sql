-- name: GetLevels :many
SELECT * FROM "levels";

-- name: GetLevel :one
SELECT * FROM "levels" WHERE id = ?;

-- name: GetLevelWithEntities :many
SELECT * 
FROM "levels"
INNER JOIN "levels_entities" ON "levels_entities.level_id" = "levels.id"
WHERE id = ?;

-- name: NewLevel :exec
INSERT INTO "levels" (name, grid_width, grid_spacing, grid_color, image_path) VALUES (?, ?, ?, ?, ?);

-- name: UpdateLevel :exec
UPDATE "levels" SET name = ?, grid_width = ?, grid_spacing = ?, grid_color = ? WHERE id = ?;

-- name: DeleteLevel :exec
DELETE FROM "levels" WHERE id = ?;
