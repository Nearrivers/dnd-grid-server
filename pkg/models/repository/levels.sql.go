// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: levels.sql

package repository

import (
	"context"
	"database/sql"
)

const deleteLevel = `-- name: DeleteLevel :exec
DELETE FROM "levels" WHERE id = ?
`

func (q *Queries) DeleteLevel(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteLevel, id)
	return err
}

const getLevel = `-- name: GetLevel :one
SELECT id, name, image_path, grid_width, grid_color, grid_spacing FROM "levels" WHERE id = ?
`

func (q *Queries) GetLevel(ctx context.Context, id int64) (Levels, error) {
	row := q.db.QueryRowContext(ctx, getLevel, id)
	var i Levels
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ImagePath,
		&i.GridWidth,
		&i.GridColor,
		&i.GridSpacing,
	)
	return i, err
}

const getLevelWithEntities = `-- name: GetLevelWithEntities :many
SELECT id, name, image_path, grid_width, grid_color, grid_spacing, level_id, entity_id, health_points, x_coord, y_coord, number 
FROM "levels"
INNER JOIN "levels_entities" ON "levels_entities.level_id" = "levels.id"
WHERE id = ?
`

type GetLevelWithEntitiesRow struct {
	ID           int64         `json:"id"`
	Name         string        `json:"name"`
	ImagePath    string        `json:"image_path"`
	GridWidth    int64         `json:"grid_width"`
	GridColor    string        `json:"grid_color"`
	GridSpacing  sql.NullInt64 `json:"grid_spacing"`
	LevelID      int64         `json:"level_id"`
	EntityID     int64         `json:"entity_id"`
	HealthPoints int64         `json:"health_points"`
	XCoord       int64         `json:"x_coord"`
	YCoord       int64         `json:"y_coord"`
	Number       int64         `json:"number"`
}

func (q *Queries) GetLevelWithEntities(ctx context.Context, id int64) ([]GetLevelWithEntitiesRow, error) {
	rows, err := q.db.QueryContext(ctx, getLevelWithEntities, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetLevelWithEntitiesRow
	for rows.Next() {
		var i GetLevelWithEntitiesRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ImagePath,
			&i.GridWidth,
			&i.GridColor,
			&i.GridSpacing,
			&i.LevelID,
			&i.EntityID,
			&i.HealthPoints,
			&i.XCoord,
			&i.YCoord,
			&i.Number,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLevels = `-- name: GetLevels :many
SELECT id, name, image_path, grid_width, grid_color, grid_spacing FROM "levels"
`

func (q *Queries) GetLevels(ctx context.Context) ([]Levels, error) {
	rows, err := q.db.QueryContext(ctx, getLevels)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Levels
	for rows.Next() {
		var i Levels
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ImagePath,
			&i.GridWidth,
			&i.GridColor,
			&i.GridSpacing,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const newLevel = `-- name: NewLevel :exec
INSERT INTO "levels" (name, grid_width, grid_spacing, grid_color, image_path) VALUES (?, ?, ?, ?, ?)
`

type NewLevelParams struct {
	Name        string        `json:"name"`
	GridWidth   int64         `json:"grid_width"`
	GridSpacing sql.NullInt64 `json:"grid_spacing"`
	GridColor   string        `json:"grid_color"`
	ImagePath   string        `json:"image_path"`
}

func (q *Queries) NewLevel(ctx context.Context, arg NewLevelParams) error {
	_, err := q.db.ExecContext(ctx, newLevel,
		arg.Name,
		arg.GridWidth,
		arg.GridSpacing,
		arg.GridColor,
		arg.ImagePath,
	)
	return err
}

const updateLevel = `-- name: UpdateLevel :exec
UPDATE "levels" SET name = ?, grid_width = ?, grid_spacing = ?, grid_color = ? WHERE id = ?
`

type UpdateLevelParams struct {
	Name        string        `json:"name"`
	GridWidth   int64         `json:"grid_width"`
	GridSpacing sql.NullInt64 `json:"grid_spacing"`
	GridColor   string        `json:"grid_color"`
	ID          int64         `json:"id"`
}

func (q *Queries) UpdateLevel(ctx context.Context, arg UpdateLevelParams) error {
	_, err := q.db.ExecContext(ctx, updateLevel,
		arg.Name,
		arg.GridWidth,
		arg.GridSpacing,
		arg.GridColor,
		arg.ID,
	)
	return err
}
