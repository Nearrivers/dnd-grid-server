-- name: GetEntities :many
SELECT * FROM "entities";

-- name: GetEntity :one
SELECT * FROM "entities" WHERE id = ?;

-- NewEntity :exec
INSERT INTO "entities" (name, faction) VALUES (?, ?);

-- UpdateEntity :exec
UPDATE "entities" SET name = ?, faction = ? WHERE id = ?;

-- DeleteEntity :exec
DELETE FROM "entities" WHERE id = ?;
