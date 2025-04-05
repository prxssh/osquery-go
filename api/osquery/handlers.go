package osquery

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prxssh/osquery-go/internal/repo"
	"github.com/prxssh/osquery-go/models"
)

type osqueryInfo struct {
	Version    string `json:"version"   binding:"required"`
	Platform   string `json:"platform"  binding:"required"`
	Extensions string `json:"extension" binding:"required"`
}

type osDetails struct {
	Name     string `json:"name"     binding:"required"`
	Arch     string `json:"arch"     binding:"required"`
	Version  string `json:"version"  binding:"required"`
	Platform string `json:"platform" binding:"required"`
}

type app struct {
	Name         string `json:"name"           binding:"required"`
	Path         string `json:"path"           binding:"required"`
	Version      string `json:"version"        binding:"required"`
	Compiler     string `json:"compiler"       binding:"required"`
	Category     string `json:"category"       binding:"required"`
	PackageType  string `json:"package_type"   binding:"required"`
	LastOpenedAt string `json:"last_opened_at" binding:"required"`
}

type latestDataResp struct {
	OsDetails   *osDetails   `json:"os_details"   binding:"required"`
	OsqueryInfo *osqueryInfo `json:"osquery_info" binding:"required"`
	AppDetails  []*app       `json:"applications" binding:"required"`
}

const (
	errInternalServer    = "ERROR_INTERNAL_SERVER"
	errMsgInternalServer = "Something went wrong!"
	errInvalidParams     = "ERROR_INVALID_PARAMS"
)

func (osq *OsqueryAPIService) latestData(c *gin.Context) {
	ctx := context.Background()

	page, err := parseQueryParamToInt(c, "page", "1")
	if err != nil {
		sendErrorResponse(
			c,
			http.StatusInternalServerError,
			errInvalidParams,
			"malformed param page type",
		)
		return
	}

	limit, err := parseQueryParamToInt(c, "limit", "10")
	if err != nil {
		sendErrorResponse(
			c,
			http.StatusInternalServerError,
			errInvalidParams,
			"malformed param limit type",
		)
		return
	}

	osqueryDetails, err := osq.repo.OsqueryInfo.Get(ctx)
	if err != nil {
		sendErrorResponse(
			c,
			http.StatusInternalServerError,
			errInternalServer,
			errMsgInternalServer,
		)
		return
	}

	osDetails, err := osq.repo.OsVersion.Get(ctx)
	if err != nil {
		sendErrorResponse(
			c,
			http.StatusInternalServerError,
			errInternalServer,
			errMsgInternalServer,
		)
		return
	}

	apps, err := getPaginatedAppDetails(osq.repo, page, limit)
	if err != nil {
		sendErrorResponse(
			c,
			http.StatusInternalServerError,
			errInternalServer,
			errMsgInternalServer,
		)
		return
	}

	totalApps, err := osq.repo.Apps.Count(ctx)
	if err != nil {
		sendErrorResponse(
			c,
			http.StatusInternalServerError,
			errInternalServer,
			errMsgInternalServer,
		)
		return
	}
	fmt.Println("totalApps:", totalApps)
	totalPages := int32(math.Ceil(float64(totalApps) / float64(limit)))

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data": &latestDataResp{
			AppDetails:  apps,
			OsDetails:   parseOsDetails(osDetails),
			OsqueryInfo: parseOsqueryInfo(osqueryDetails),
		},
		"pagination": gin.H{
			"current_page": page,
			"total_pages":  totalPages,
			"limit":        limit,
			"total_items":  totalApps,
		},
	})
}

func parseOsqueryInfo(data models.OsqueryInfo) *osqueryInfo {
	return &osqueryInfo{
		Version:    data.Version.String,
		Platform:   data.BuildPlatform.String,
		Extensions: data.Extensions.String,
	}
}

func parseOsDetails(data models.OsVersion) *osDetails {
	return &osDetails{
		Name:     data.Name,
		Version:  data.Version,
		Arch:     data.Arch.String,
		Platform: data.Platform.String,
	}
}

func getPaginatedAppDetails(
	repo *repo.Repo,
	page, limit int32,
) ([]*app, error) {
	offset := (page - 1) * limit

	apps, err := repo.Apps.List(context.Background(), limit, offset)
	if err != nil {
		return nil, err
	}

	var appDetails []*app
	for _, app := range apps {
		appDetails = append(appDetails, parseApp(app))
	}

	return appDetails, nil
}

func parseApp(data models.App) *app {
	lastOpenedAtTruncSec := int64(math.Floor(data.LastOpenedTime.Float64))

	return &app{
		Name:         data.Name.String,
		Path:         data.Path.String,
		Category:     data.Category.String,
		Compiler:     data.Compiler.String,
		Version:      data.BundleVersion.String,
		LastOpenedAt: time.Unix(lastOpenedAtTruncSec, 0).String(),
	}
}

func sendErrorResponse(
	c *gin.Context,
	statusCode int,
	errCode string,
	message string,
) {
	c.AbortWithStatusJSON(statusCode, gin.H{
		"status": errCode,
		"code":   message,
	})
}

func sendSuccessResponse(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{"status": "OK", "data": data})
}

func parseQueryParamToInt(
	c *gin.Context,
	key string,
	defaultValue string,
) (int32, error) {
	param := c.DefaultQuery(key, defaultValue)
	parsedInt, err := strconv.ParseInt(param, 10, 32)
	if err != nil {
		return -1, err
	}

	return int32(parsedInt), nil
}
