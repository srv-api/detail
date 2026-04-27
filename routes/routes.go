package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/srv-api/detail/configs"

	h_permission "github.com/srv-api/detail/handlers/dashboard/permission"
	r_permission "github.com/srv-api/detail/repositories/dashboard/permission"
	s_permission "github.com/srv-api/detail/services/dashboard/permission"

	h_role "github.com/srv-api/detail/handlers/dashboard/role"
	r_role "github.com/srv-api/detail/repositories/dashboard/role"
	s_role "github.com/srv-api/detail/services/dashboard/role"

	h_role_user "github.com/srv-api/detail/handlers/dashboard/roleuser"
	r_role_user "github.com/srv-api/detail/repositories/dashboard/roleuser"
	s_role_user "github.com/srv-api/detail/services/dashboard/roleuser"

	h_role_user_permission "github.com/srv-api/detail/handlers/dashboard/roleuserpermission"
	r_role_user_permission "github.com/srv-api/detail/repositories/dashboard/roleuserpermission"
	s_role_user_permission "github.com/srv-api/detail/services/dashboard/roleuserpermission"

	h_user "github.com/srv-api/detail/handlers/user"
	r_user "github.com/srv-api/detail/repositories/user"
	s_user "github.com/srv-api/detail/services/user"

	h_userdetail "github.com/srv-api/detail/handlers/userdetail"
	r_userdetail "github.com/srv-api/detail/repositories/userdetail"
	s_userdetail "github.com/srv-api/detail/services/userdetail"

	h_contentsetting "github.com/srv-api/detail/handlers/dashboard/contentsetting"
	r_contentsetting "github.com/srv-api/detail/repositories/dashboard/contentsetting"
	s_contentsetting "github.com/srv-api/detail/services/dashboard/contentsetting"

	h_deleteaccount "github.com/srv-api/detail/handlers/deleteaccount"
	r_deleteaccount "github.com/srv-api/detail/repositories/deleteaccount"
	s_deleteaccount "github.com/srv-api/detail/services/deleteaccount"

	h_pin "github.com/srv-api/detail/handlers/pin"
	r_pin "github.com/srv-api/detail/repositories/pin"
	s_pin "github.com/srv-api/detail/services/pin"

	handler "github.com/srv-api/detail/handlers/like"
	repository "github.com/srv-api/detail/repositories/like"
	service "github.com/srv-api/detail/services/like"

	h_match "github.com/srv-api/detail/handlers/match"
	r_match "github.com/srv-api/detail/repositories/match"
	s_match "github.com/srv-api/detail/services/match"

	"github.com/srv-api/middlewares/middlewares"
)

var (
	DB = configs.InitDB()

	pp = configs.InitApp()

	JWT = middlewares.NewJWTService()

	userdetailR = r_userdetail.NewUserDetailRepository(DB)
	userdetailS = s_userdetail.NewUserDetailService(userdetailR, JWT)
	userdetailH = h_userdetail.NewUserDetailHandler(userdetailS)

	contentsettingR = r_contentsetting.NewContentSettingRepository(DB)
	contentsettingS = s_contentsetting.NewContentSettingService(contentsettingR, JWT)
	contentsettingH = h_contentsetting.NewContentSettingHandler(contentsettingS)

	permissionR = r_permission.NewPermissionRepository(DB)
	permissionS = s_permission.NewPermissionService(permissionR, JWT)
	permissionH = h_permission.NewPermissionHandler(permissionS)

	roleR = r_role.NewRoleRepository(DB)
	roleS = s_role.NewRoleService(roleR, JWT)
	roleH = h_role.NewRoleHandler(roleS)

	roleuserR = r_role_user.NewRoleUserRepository(DB)
	roleuserS = s_role_user.NewRoleUserService(roleuserR, JWT)
	roleuserH = h_role_user.NewRoleUserHandler(roleuserS)

	roleuserpermissionR = r_role_user_permission.NewRoleUserPermissionRepository(DB)
	roleuserpermissionS = s_role_user_permission.NewRoleUserPermissionService(roleuserpermissionR, JWT)
	roleuserpermissionH = h_role_user_permission.NewRoleUserPermissionHandler(roleuserpermissionS)

	deleteaccountR = r_deleteaccount.NewDeleteAccountRepository(DB)
	deleteaccountS = s_deleteaccount.NewDeleteAccountService(deleteaccountR, JWT)
	deleteaccountH = h_deleteaccount.NewRequestDeleteHandler(deleteaccountS)

	userR = r_user.NewUserRepository(DB)
	userS = s_user.NewUserService(userR, JWT)
	userH = h_user.NewUserHandler(userS)

	pinR = r_pin.NewPinRepository(DB)
	pinS = s_pin.NewPinService(pinR, JWT)
	pinH = h_pin.NewPinHandler(pinS)

	matchHandler = h_match.NewMatchHandler(matchService)
	matchRepo    = r_match.NewMatchRepository(DB)
	matchService = s_match.NewMatchService(matchRepo)

	likeHandler = handler.NewLikeHandler(likeService)
	likeRepo    = repository.NewLikeRepository(DB)
	likeService = service.NewLikeService(likeRepo, matchService)
)

func New() *echo.Echo {

	e := echo.New()
	// e.POST("/menu/order", orderH.Order)
	e.PUT("/user/update", userdetailH.LongLat)
	// cron.StartDailyReset(DB)

	merchant := e.Group("/merchant", middlewares.AuthorizeJWT(JWT))
	{
		merchant.PUT("/update", userdetailH.Update)
		merchant.GET("/get", userdetailH.Get)
		merchant.GET("/explore", userdetailH.Explore)
	}
	contentsetting := e.Group("/merchant", middlewares.AuthorizeJWT(JWT))
	{
		contentsetting.PUT("/contentsetting/update", contentsettingH.Update)
		contentsetting.GET("/contentsetting/get", contentsettingH.Get)
	}
	web := e.Group("/merchant")
	{
		web.GET("/web/get/content", contentsettingH.Get)
		web.PUT("/web/update/content", contentsettingH.Update)
	}

	pin := e.Group("/merchant", middlewares.AuthorizeJWT(JWT))
	{
		pin.POST("/pin/create", pinH.Create)
		pin.GET("/pin/pagination", pinH.Get)
		pin.GET("/pin/:id", pinH.GetById)
		pin.PUT("/pin/update/:id", pinH.Update)
		pin.DELETE("/pin/:id", pinH.Delete)
		pin.DELETE("/pin/bulk-delete", pinH.BulkDelete)
		pin.POST("/verify-pin", pinH.VerifyPIN)
		pin.GET("/pin/status", pinH.GetPinStatus)

	}
	permission := e.Group("/merchant", middlewares.AuthorizeJWT(JWT))
	{
		permission.POST("/permission/create", permissionH.Create)
		permission.GET("/permission", permissionH.Get)
		permission.GET("/permission/pagination", permissionH.Pagination)
		permission.PUT("/permission/update/:id", permissionH.Update)
		permission.DELETE("/permission/:id", permissionH.Delete)
		permission.DELETE("/permission/bulk-delete", permissionH.BulkDelete)
	}
	role := e.Group("/merchant", middlewares.AuthorizeJWT(JWT))
	{
		role.POST("/role/create", roleH.Create)
		permission.GET("/role", permissionH.Get)
		role.GET("/role/pagination", roleH.Pagination)
		role.GET("/role_user", roleH.RoleUser)
		role.PUT("/role/update/:id", roleH.Update)
		role.DELETE("/role/:id", roleH.Delete)
		role.DELETE("/role/bulk-delete", roleH.BulkDelete)
	}
	roleuser := e.Group("/merchant", middlewares.AuthorizeJWT(JWT))
	{
		roleuser.POST("/roleuser/create", roleuserH.Create)
		roleuser.GET("/roleuser", roleuserH.Get)
		roleuser.GET("/roleuser/pagination", roleuserH.Pagination)
		roleuser.PUT("/roleuser/update/:id", roleuserH.Update)
		roleuser.DELETE("/roleuser/:id", roleuserH.Delete)
		roleuser.DELETE("/roleuser/bulk-delete", roleuserH.BulkDelete)
	}
	roleuserpermission := e.Group("/merchant", middlewares.AuthorizeJWT(JWT))
	{
		roleuserpermission.POST("/roleuserpermission/create", roleuserpermissionH.Create)
		roleuserpermission.GET("/roleuserpermission", roleuserpermissionH.Get)
		roleuserpermission.GET("/roleuserpermission/pagination", roleuserpermissionH.Pagination)
		roleuserpermission.PUT("/roleuserpermission/update/:id", roleuserpermissionH.Update)
		roleuserpermission.DELETE("/roleuserpermission/:id", roleuserpermissionH.Delete)
		roleuserpermission.DELETE("/roleuserpermission/bulk-delete", roleuserpermissionH.BulkDelete)
	}

	user := e.Group("/merchant", middlewares.AuthorizeJWT(JWT))
	{
		user.POST("/user/create", userH.Create)
		user.GET("/user/pagination", userH.Get)
		user.PUT("/user/update/:id", userH.Update)
		user.DELETE("/user/:id", userH.Delete)
		user.DELETE("/user/bulk-delete", userH.BulkDelete)
	}

	deleteAccount := e.Group("api/account", middlewares.AuthorizeJWT(JWT))
	{
		deleteAccount.POST("/request-delete", deleteaccountH.Create)
		// deleteAccount.GET("/unit/pagination", unitH.Get)
		// deleteAccount.PUT("/unit/:id", unitH.Update)
		// deleteAccount.DELETE("/unit/:id", unitH.Delete)
		// deleteAccount.DELETE("/unit/bulk-delete", unitH.BulkDelete)
	}

	like := e.Group("api/account", middlewares.AuthorizeJWT(JWT))
	{
		like.POST("/like", likeHandler.LikeUser)
		like.GET("/like/me", likeHandler.Me)
	}
	match := e.Group("api/account", middlewares.AuthorizeJWT(JWT))
	{
		match.GET("/matches", matchHandler.GetMatches)
	}

	return e
}
