package macho

import (
	"net/http"
	"strings"

	"github.com/blacktop/go-macho"
	"github.com/blacktop/ipsw/api/types"
	"github.com/gin-gonic/gin"
)

// Info is the struct for the macho info route parameters
type Info struct {
	Path string `form:"path" json:"path" binding:"required"`
	Arch string `form:"arch" json:"arch"`
}

// swagger:response
type machoInfoResponse struct {
	Path string      `json:"path"`
	Arch string      `json:"arch"`
	Info *macho.File `json:"info"`
}

func machoInfo(c *gin.Context) {
	var m *macho.File
	var params Info

	if err := c.BindQuery(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, types.GenericError{Error: err.Error()})
		return
	}

	fat, err := macho.OpenFat(params.Path)
	if err != nil {
		if err == macho.ErrNotFat { // not a fat binary
			m, err = macho.Open(params.Path)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, types.GenericError{Error: err.Error()})
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.GenericError{Error: err.Error()})
			return
		}
	} else { // fat binary
		if params.Arch == "" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.GenericError{Error: "'arch' query parameter is required for universal binaries"})
			return
		}
		for _, farch := range fat.Arches {
			if strings.EqualFold(farch.SubCPU.String(farch.CPU), params.Arch) {
				m = farch.File
			}
		}
	}
	c.IndentedJSON(http.StatusOK, machoInfoResponse{Path: params.Path, Arch: params.Arch, Info: m})
}
