package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
	"time"
	"website-api/internal/database"
	"website-api/internal/models"
)

type ProjectHandler struct {
	db    *database.DB
	cache *cache.Cache
}

func NewProjectHandler(db *database.DB) *ProjectHandler {
	return &ProjectHandler{
		db:    db,
		cache: cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (ph *ProjectHandler) GetProjects(ctx echo.Context) error {
	cacheKey := "all_projects"
	if cachedProjects, found := ph.cache.Get(cacheKey); found {
		return ctx.JSON(http.StatusOK, cachedProjects)
	}

	projects, err := ph.db.GetProjects(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	ph.cache.Set(cacheKey, projects, cache.DefaultExpiration)
	return ctx.JSON(http.StatusOK, projects)
}

func (ph *ProjectHandler) GetProjectList(ctx echo.Context) error {
	cacheKey := "project_list"
	if cachedProjects, found := ph.cache.Get(cacheKey); found {
		return ctx.JSON(http.StatusOK, cachedProjects)
	}

	projects, err := ph.db.GetProjectList(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	ph.cache.Set(cacheKey, projects, cache.DefaultExpiration)
	return ctx.JSON(http.StatusOK, projects)
}

func (ph *ProjectHandler) GetProject(ctx echo.Context) error {
	name := ctx.Param("name")
	log.Println("Fetching project", name)
	project, err := ph.db.GetProject(ctx.Request().Context(), name)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, project)
}

func (ph *ProjectHandler) EditProject(c echo.Context) error {
	var project models.Project
	if err := c.Bind(&project); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid project data"})
	}

	err := ph.db.EditProject(c.Request().Context(), &project)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to edit project"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Project updated successfully"})
}

func (ph *ProjectHandler) DeleteProject(c echo.Context) error {
	name := c.Param("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Project name is required"})
	}

	err := ph.db.DeleteProject(c.Request().Context(), name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete project"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Project deleted successfully"})
}

func (ph *ProjectHandler) CreateProject(c echo.Context) error {
	var project models.Project
	if err := c.Bind(&project); err != nil {
		log.Printf("Binding error: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid project data"})
	}

	err := ph.db.CreateProject(c.Request().Context(), &project)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create project"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Project created successfully"})
}
