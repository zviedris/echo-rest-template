package contexts

import (
	coreDomains "github.com/zviedris/echo-rest-template/modules/core/domains"
	projectDomains "github.com/zviedris/echo-rest-template/modules/projects/domains"
	"github.com/zviedris/echo-rest-template/pkg/constants"
	"github.com/labstack/echo/v4"
)

func GetUserFromContext(c echo.Context) (usr *coreDomains.UserWithRoles, err error) {
	uCtx := c.Get(constants.ContextKeyUser)
	usr, ok := uCtx.(*coreDomains.UserWithRoles)
	if !ok || usr == nil {
		err = constants.ErrUnauthorized
	}
	return
}

func GetOrgFromContext(c echo.Context) *coreDomains.Org {
	org := c.Get(constants.ContextKeyOrg)
	if org != nil {
		return org.(*coreDomains.Org)
	}
	return nil
}

func GetProjectFromContext(c echo.Context) *projectDomains.Project {
	proj := c.Get(constants.ContextKeyProject)
	if proj != nil {
		return proj.(*projectDomains.Project)
	}
	return nil
}
