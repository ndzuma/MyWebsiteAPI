package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"website-api/internal/models"
)

type DB struct {
	pool *pgxpool.Pool
}

// NewDB connects to the database
func NewDB(dbUrl string) (*DB, error) {
	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}
	return &DB{pool: pool}, nil
}

// Close closes the database connection
func (db *DB) Close() {
	db.pool.Close()
}

func (db *DB) GetProjects(ctx context.Context) ([]models.Project, error) {
	query := `select id, created_at, name, "type", github_link, project_url, "year", dev_time, technologies, main_image, secondary_image, overview, detailed_description, in_progress, bg_color, text_color from projects`
	rows, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query projects: %w", err)
	}
	defer rows.Close()

	results, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Project])
	if err != nil {
		return nil, fmt.Errorf("failed to collect rows: %w", err)
	}

	return results, nil
}

func (db *DB) GetProjectList(ctx context.Context) ([]models.ProjectList, error) {
	query := `select id, name, main_image, in_progress, bg_color, text_color from projects`
	rows, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query project list: %w", err)
	}
	defer rows.Close()
	results, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ProjectList])
	if err != nil {
		return nil, fmt.Errorf("failed to collect rows: %w", err)
	}
	return results, nil
}

func (db *DB) GetProject(ctx context.Context, name string) (*models.Project, error) {
	query := `select  name, "type", github_link, project_url, "year", dev_time, technologies, main_image, secondary_image, overview, detailed_description, in_progress from projects where name = $1`
	row := db.pool.QueryRow(ctx, query, name)

	var project models.Project
	err := row.Scan(&project.Name, &project.Type, &project.GithubLink, &project.ProjectURL, &project.Year, &project.DevTime, &project.Technologies, &project.MainImage, &project.SecondaryImage, &project.Overview, &project.DetailedDescription, &project.InProgress)
	if err != nil {
		return nil, fmt.Errorf("failed to scan project: %w", err)
	}
	return &project, nil
}

func (db *DB) EditProject(ctx context.Context, project *models.Project) error {
	query := `update projects set "type" = $1, github_link = $2, project_url = $3, "year" = $4, dev_time = $5, technologies = $6, main_image = $7, secondary_image = $8, overview = $9, detailed_description = $10, in_progress = $11, bg_color = $12, text_color =$13 where name = $14`
	_, err := db.pool.Exec(ctx, query, project.Type, project.GithubLink, project.ProjectURL, project.Year, project.DevTime, project.Technologies, project.MainImage, project.SecondaryImage, project.Overview, project.DetailedDescription, project.InProgress, project.BGColor, project.TextColor, project.Name)
	if err != nil {
		return fmt.Errorf("failed to update project: %w", err)
	}
	return nil
}

func (db *DB) DeleteProject(ctx context.Context, name string) error {
	query := `delete from projects where name = $1`
	_, err := db.pool.Exec(ctx, query, name)
	if err != nil {
		return fmt.Errorf("failed to delete project: %w", err)
	}
	return nil
}

func (db *DB) CreateProject(ctx context.Context, project *models.Project) error {
	query := `insert into projects (name, "type", github_link, project_url, "year", dev_time, technologies, main_image, secondary_image, overview, detailed_description, in_progress, bg_color, text_color) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`
	_, err := db.pool.Exec(ctx, query, project.Name, project.Type, project.GithubLink, project.ProjectURL, project.Year, project.DevTime, project.Technologies, project.MainImage, project.SecondaryImage, project.Overview, project.DetailedDescription, project.InProgress, project.BGColor, project.TextColor)
	if err != nil {
		return fmt.Errorf("failed to create project: %w", err)
	}
	return nil
}
