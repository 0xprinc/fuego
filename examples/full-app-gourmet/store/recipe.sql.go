// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: recipe.sql

package store

import (
	"context"
	"database/sql"
)

const createRecipe = `-- name: CreateRecipe :one
INSERT INTO recipe (
  id,
  name,
  description,
  instructions,
  prep_time,
  cook_time,
  category,
  image_url,
  published,
  servings,
  when_to_eat
) 
VALUES (?,?,?,?,?,?,?,?,?,?,?) RETURNING id, created_at, name, description, instructions, category, published, created_by, calories, cost, prep_time, cook_time, servings, image_url, disclaimer, when_to_eat
`

type CreateRecipeParams struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Instructions string `json:"instructions"`
	PrepTime     int64  `json:"prep_time"`
	CookTime     int64  `json:"cook_time"`
	Category     string `json:"category"`
	ImageUrl     string `json:"image_url"`
	Published    bool   `json:"published"`
	Servings     int64  `json:"servings"`
	WhenToEat    string `json:"when_to_eat"`
}

func (q *Queries) CreateRecipe(ctx context.Context, arg CreateRecipeParams) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, createRecipe,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Instructions,
		arg.PrepTime,
		arg.CookTime,
		arg.Category,
		arg.ImageUrl,
		arg.Published,
		arg.Servings,
		arg.WhenToEat,
	)
	var i Recipe
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Description,
		&i.Instructions,
		&i.Category,
		&i.Published,
		&i.CreatedBy,
		&i.Calories,
		&i.Cost,
		&i.PrepTime,
		&i.CookTime,
		&i.Servings,
		&i.ImageUrl,
		&i.Disclaimer,
		&i.WhenToEat,
	)
	return i, err
}

const deleteRecipe = `-- name: DeleteRecipe :exec
DELETE FROM recipe WHERE id = ?
`

func (q *Queries) DeleteRecipe(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteRecipe, id)
	return err
}

const getRandomRecipes = `-- name: GetRandomRecipes :many
SELECT id, created_at, name, description, instructions, category, published, created_by, calories, cost, prep_time, cook_time, servings, image_url, disclaimer, when_to_eat FROM recipe ORDER BY RANDOM() DESC LIMIT 10
`

func (q *Queries) GetRandomRecipes(ctx context.Context) ([]Recipe, error) {
	rows, err := q.db.QueryContext(ctx, getRandomRecipes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Recipe
	for rows.Next() {
		var i Recipe
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Name,
			&i.Description,
			&i.Instructions,
			&i.Category,
			&i.Published,
			&i.CreatedBy,
			&i.Calories,
			&i.Cost,
			&i.PrepTime,
			&i.CookTime,
			&i.Servings,
			&i.ImageUrl,
			&i.Disclaimer,
			&i.WhenToEat,
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

const getRecipe = `-- name: GetRecipe :one
SELECT id, created_at, name, description, instructions, category, published, created_by, calories, cost, prep_time, cook_time, servings, image_url, disclaimer, when_to_eat FROM recipe WHERE id = ?
`

func (q *Queries) GetRecipe(ctx context.Context, id string) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, getRecipe, id)
	var i Recipe
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Description,
		&i.Instructions,
		&i.Category,
		&i.Published,
		&i.CreatedBy,
		&i.Calories,
		&i.Cost,
		&i.PrepTime,
		&i.CookTime,
		&i.Servings,
		&i.ImageUrl,
		&i.Disclaimer,
		&i.WhenToEat,
	)
	return i, err
}

const getRecipes = `-- name: GetRecipes :many
SELECT id, created_at, name, description, instructions, category, published, created_by, calories, cost, prep_time, cook_time, servings, image_url, disclaimer, when_to_eat FROM recipe
`

func (q *Queries) GetRecipes(ctx context.Context) ([]Recipe, error) {
	rows, err := q.db.QueryContext(ctx, getRecipes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Recipe
	for rows.Next() {
		var i Recipe
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Name,
			&i.Description,
			&i.Instructions,
			&i.Category,
			&i.Published,
			&i.CreatedBy,
			&i.Calories,
			&i.Cost,
			&i.PrepTime,
			&i.CookTime,
			&i.Servings,
			&i.ImageUrl,
			&i.Disclaimer,
			&i.WhenToEat,
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

const searchRecipes = `-- name: SearchRecipes :many
SELECT id, created_at, name, description, instructions, category, published, created_by, calories, cost, prep_time, cook_time, servings, image_url, disclaimer, when_to_eat FROM recipe WHERE
  (name LIKE '%' || ?1 || '%')
  AND published = ?2
  AND calories <= ?3
  AND prep_time + cook_time <= ?4
ORDER BY name ASC
LIMIT ?6
OFFSET ?5
`

type SearchRecipesParams struct {
	Search      sql.NullString `json:"search"`
	Published   bool           `json:"published"`
	MaxCalories int64          `json:"max_calories"`
	MaxTime     int64          `json:"max_time"`
	Offset      int64          `json:"offset"`
	Limit       int64          `json:"limit"`
}

// Search anything that contains the given string
func (q *Queries) SearchRecipes(ctx context.Context, arg SearchRecipesParams) ([]Recipe, error) {
	rows, err := q.db.QueryContext(ctx, searchRecipes,
		arg.Search,
		arg.Published,
		arg.MaxCalories,
		arg.MaxTime,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Recipe
	for rows.Next() {
		var i Recipe
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Name,
			&i.Description,
			&i.Instructions,
			&i.Category,
			&i.Published,
			&i.CreatedBy,
			&i.Calories,
			&i.Cost,
			&i.PrepTime,
			&i.CookTime,
			&i.Servings,
			&i.ImageUrl,
			&i.Disclaimer,
			&i.WhenToEat,
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

const updateRecipe = `-- name: UpdateRecipe :one
UPDATE recipe SET 
  name=COALESCE(?1, name),
  description=COALESCE(?2, description),
  instructions=COALESCE(?3, instructions),
  category=COALESCE(?4, category),
  when_to_eat=COALESCE(?5, when_to_eat),
  image_url=COALESCE(?6, image_url),
  cook_time=COALESCE(?7, cook_time),
  prep_time=COALESCE(?8, prep_time),
  servings=COALESCE(?9, servings),
  published=COALESCE(?10, published)
WHERE id = ?11
RETURNING id, created_at, name, description, instructions, category, published, created_by, calories, cost, prep_time, cook_time, servings, image_url, disclaimer, when_to_eat
`

type UpdateRecipeParams struct {
	Name         string         `json:"name"`
	Description  sql.NullString `json:"description"`
	Instructions sql.NullString `json:"instructions"`
	Category     string         `json:"category"`
	WhenToEat    string         `json:"when_to_eat"`
	ImageUrl     string         `json:"image_url"`
	CookTime     int64          `json:"cook_time"`
	PrepTime     int64          `json:"prep_time"`
	Servings     int64          `json:"servings"`
	Published    bool           `json:"published"`
	ID           string         `json:"id"`
}

func (q *Queries) UpdateRecipe(ctx context.Context, arg UpdateRecipeParams) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, updateRecipe,
		arg.Name,
		arg.Description,
		arg.Instructions,
		arg.Category,
		arg.WhenToEat,
		arg.ImageUrl,
		arg.CookTime,
		arg.PrepTime,
		arg.Servings,
		arg.Published,
		arg.ID,
	)
	var i Recipe
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Description,
		&i.Instructions,
		&i.Category,
		&i.Published,
		&i.CreatedBy,
		&i.Calories,
		&i.Cost,
		&i.PrepTime,
		&i.CookTime,
		&i.Servings,
		&i.ImageUrl,
		&i.Disclaimer,
		&i.WhenToEat,
	)
	return i, err
}
