package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// FindAllConfigSources returns a resource configuration from the database.
func FindAllConfigSources(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	kind := c.Request.URL.Query().Get("kind")
	config := c.Request.URL.Query().Get("config")

	rows, err := dbGetConfigSource(kind, id, config)
	if err != nil {
		logrus.Errorf("searching for config source: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"configs": rows,
	})
}

// FindAllConfigSourcesWithJSONBody returns a resource configuration from the database.
func FindAllConfigSourcesWithJSONBody(c *gin.Context) {
	type configResource struct {
		ID     string `json:"id"`
		Kind   string `json:"kind"`
		Config string `json:"config"`
	}

	queryConfig := configResource{}

	if err := c.ShouldBindJSON(&queryConfig); err != nil {
		logrus.Errorf("failed to read json body: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	rows, err := dbGetConfigSource(queryConfig.Kind, queryConfig.ID, queryConfig.Config)
	if err != nil {
		logrus.Errorf("searching for config source: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"configs": rows,
	})

}

// DeleteConfigSource deletes a resource configuration from the database.
func DeleteConfigSource(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")

	err := dbDeleteConfigResource("source", id)
	if err != nil {
		logrus.Errorf("deleting config source: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// FindAllConfigConditions returns condition configurations from the database.
func FindAllConfigConditions(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	kind := c.Request.URL.Query().Get("kind")
	config := c.Request.URL.Query().Get("config")

	rows, err := dbGetConfigCondition(kind, id, config)
	if err != nil {
		logrus.Errorf("searching for config condition: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"configs": rows,
	})
}

// FindAllConfigConditionsWithJSONBody returns condition configurations from the database.
func FindAllConfigConditionsWithJSONBody(c *gin.Context) {
	type configResource struct {
		ID     string `json:"id"`
		Kind   string `json:"kind"`
		Config string `json:"config"`
	}

	queryConfig := configResource{}

	if err := c.ShouldBindJSON(&queryConfig); err != nil {
		logrus.Errorf("failed to read json body: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	rows, err := dbGetConfigCondition(queryConfig.Kind, queryConfig.ID, queryConfig.Config)
	if err != nil {
		logrus.Errorf("searching for config condition: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"configs": rows,
	})
}

// DeleteConfigCondition deletes a resource configuration from the database.
func DeleteConfigCondition(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")

	err := dbDeleteConfigResource("condition", id)
	if err != nil {
		logrus.Errorf("deleting config source: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// FindAllConfigTargets returns target configurations from the database.
func FindAllConfigTargets(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	kind := c.Request.URL.Query().Get("kind")
	config := c.Request.URL.Query().Get("config")

	rows, err := dbGetConfigTarget(kind, id, config)
	if err != nil {
		logrus.Errorf("searching for config target: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"configs": rows,
	})
}

// FindAllConfigTargetsWithJSONBody returns target configurations from the database.
func FindAllConfigTargetsWithJSONBody(c *gin.Context) {
	type configResource struct {
		ID     string `json:"id"`
		Kind   string `json:"kind"`
		Config string `json:"config"`
	}

	queryConfig := configResource{}

	if err := c.ShouldBindJSON(&queryConfig); err != nil {
		logrus.Errorf("failed to read json body: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	rows, err := dbGetConfigTarget(queryConfig.Kind, queryConfig.ID, queryConfig.Config)
	if err != nil {
		logrus.Errorf("searching for config target: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"configs": rows,
	})
}

// DeleteConfigTarget deletes a resource configuration from the database.
func DeleteTargetTarget(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")

	err := dbDeleteConfigResource("target", id)
	if err != nil {
		logrus.Errorf("deleting config source: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
