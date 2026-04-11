package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/srv-api/merchant/configs"

	h_permission "github.com/srv-api/merchant/handlers/dashboard/permission"
	r_permission "github.com/srv-api/merchant/repositories/dashboard/permission"
	s_permission "github.com/srv-api/merchant/services/dashboard/permission"

	h_role "github.com/srv-api/merchant/handlers/dashboard/role"
	r_role "github.com/srv-api/merchant/repositories/dashboard/role"
	s_role "github.com/srv-api/merchant/services/dashboard/role"

	h_role_user "github.com/srv-api/merchant/handlers/dashboard/roleuser"
	r_role_user "github.com/srv-api/merchant/repositories/dashboard/roleuser"
	s_role_user "github.com/srv-api/merchant/services/dashboard/roleuser"

	h_role_user_permission "github.com/srv-api/merchant/handlers/dashboard/roleuserpermission"
	r_role_user_permission "github.com/srv-api/merchant/repositories/dashboard/roleuserpermission"
	s_role_user_permission "github.com/srv-api/merchant/services/dashboard/roleuserpermission"

	h_subscribe "github.com/srv-api/merchant/handlers/subscribe"
	r_subscribe "github.com/srv-api/merchant/repositories/subscribe"
	s_subscribe "github.com/srv-api/merchant/services/subscribe"

	h_paymentmethod "github.com/srv-api/merchant/handlers/subscribe/paymentmethod"
	r_paymentmethod "github.com/srv-api/merchant/repositories/subscribe/paymentmethod"
	s_paymentmethod "github.com/srv-api/merchant/services/subscribe/paymentmethod"

	h_transactionmethode "github.com/srv-api/merchant/handlers/transactionmethode/qris"
	r_transactionmethode "github.com/srv-api/merchant/repositories/transactionmethode/qris"
	s_transactionmethode "github.com/srv-api/merchant/services/transactionmethode/qris"

	h_history "github.com/srv-api/merchant/handlers/subscribe/history"
	r_history "github.com/srv-api/merchant/repositories/subscribe/history"
	s_history "github.com/srv-api/merchant/services/subscribe/history"

	h_user "github.com/srv-api/merchant/handlers/user"
	r_user "github.com/srv-api/merchant/repositories/user"
	s_user "github.com/srv-api/merchant/services/user"

	h_usermerchant "github.com/srv-api/merchant/handlers/usermerchant"
	r_usermerchant "github.com/srv-api/merchant/repositories/usermerchant"
	s_usermerchant "github.com/srv-api/merchant/services/usermerchant"

	h_product "github.com/srv-api/merchant/handlers/product"
	r_product "github.com/srv-api/merchant/repositories/product"
	s_product "github.com/srv-api/merchant/services/product"

	h_merchant "github.com/srv-api/merchant/handlers/merchant"
	r_merchant "github.com/srv-api/merchant/repositories/merchant"
	s_merchant "github.com/srv-api/merchant/services/merchant"

	h_contentsetting "github.com/srv-api/merchant/handlers/dashboard/contentsetting"
	r_contentsetting "github.com/srv-api/merchant/repositories/dashboard/contentsetting"
	s_contentsetting "github.com/srv-api/merchant/services/dashboard/contentsetting"

	h_deleteaccount "github.com/srv-api/merchant/handlers/deleteaccount"
	r_deleteaccount "github.com/srv-api/merchant/repositories/deleteaccount"
	s_deleteaccount "github.com/srv-api/merchant/services/deleteaccount"

	h_pin "github.com/srv-api/merchant/handlers/pin"
	r_pin "github.com/srv-api/merchant/repositories/pin"
	s_pin "github.com/srv-api/merchant/services/pin"

	"github.com/srv-api/middlewares/middlewares"
)

var (
	DB = configs.InitDB()

	pp = configs.InitApp()

	JWT = middlewares.NewJWTService()

	merchantR = r_merchant.NewMerchantRepository(DB)
	merchantS = s_merchant.NewMerchantService(merchantR, JWT)
	merchantH = h_merchant.NewMerchantHandler(merchantS)

	contentsettingR = r_contentsetting.NewContentSettingRepository(DB)
	contentsettingS = s_contentsetting.NewContentSettingService(contentsettingR, JWT)
	contentsettingH = h_contentsetting.NewContentSettingHandler(contentsettingS)

	subscribeR = r_subscribe.NewSubscribeRepository(DB, pp)
	subscribeS = s_subscribe.NewSubscribeService(subscribeR, JWT)
	subscribeH = h_subscribe.NewSubscribeHandler(subscribeS)

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

	productR = r_product.NewProductRepository(DB)
	productS = s_product.NewProductService(productR, JWT)
	productH = h_product.NewProductHandler(productS)

	paymentmethodR = r_paymentmethod.NewPaymentRepository(DB)
	paymentmethodS = s_paymentmethod.NewPaymentMethodService(paymentmethodR, JWT)
	paymentmethodH = h_paymentmethod.NewPaymentHandler(paymentmethodS)

	transactionmethodeR = r_transactionmethode.NewQrisRepository(DB)
	transactionmethodeS = s_transactionmethode.NewQrisService(transactionmethodeR, JWT)
	transactionmethodeH = h_transactionmethode.NewQrisHandler(transactionmethodeS)

	historyR = r_history.NewHistoryRepository(DB)
	historyS = s_history.NewHistoryService(historyR, JWT)
	historyH = h_history.NewHistoryHandler(historyS)

	deleteaccountR = r_deleteaccount.NewDeleteAccountRepository(DB)
	deleteaccountS = s_deleteaccount.NewDeleteAccountService(deleteaccountR, JWT)
	deleteaccountH = h_deleteaccount.NewRequestDeleteHandler(deleteaccountS)
	userR          = r_user.NewUserRepository(DB)
	userS          = s_user.NewUserService(userR, JWT)
	userH          = h_user.NewUserHandler(userS)

	usermerchantR = r_usermerchant.NewUserMerchantRepository(DB)
	usermerchantS = s_usermerchant.NewUserMerchantService(usermerchantR, JWT)
	usermerchantH = h_usermerchant.NewUserMerchantHandler(usermerchantS)

	pinR = r_pin.NewPinRepository(DB)
	pinS = s_pin.NewPinService(pinR, JWT)
	pinH = h_pin.NewPinHandler(pinS)
)

func New() *echo.Echo {

	e := echo.New()
	// e.POST("/menu/order", orderH.Order)

	sub := e.Group("/merchant", middlewares.AuthorizeJWT(JWT))
	{
		sub.GET("/subscribe/transaction/:order_id/status", subscribeH.CheckTransactionStatus)
		sub.POST("/subscribe/midtrans/callback", subscribeH.MidtransCallback)
		sub.POST("/subscribe/charge-bni", subscribeH.ChargeBni)
		sub.POST("/subscribe/charge-permata", subscribeH.ChargePermata)
		sub.POST("/subscribe/charge-mandiri", subscribeH.ChargeMandiri)
		sub.POST("/subscribe/charge-bri", subscribeH.ChargeBri)
		sub.POST("/subscribe/charge-cimb", subscribeH.ChargeCimb)
		sub.POST("/subscribe/charge-qris", subscribeH.ChargeQris)
		sub.POST("/subscribe/charge-gopay", subscribeH.ChargeGopay)
		sub.POST("/subscribe/charge-shopeepay", subscribeH.ChargeShopeePay)
		sub.POST("/subscribe/charge-gpay", subscribeH.ChargeGpay)
		sub.GET("/subscribe/tokenize", subscribeH.TokenizeCardHandler)
		sub.POST("/subscribe/charge-card", subscribeH.CardPayment)
		sub.POST("/subscribe/cancel/:order_id", subscribeH.CancelPay)
		sub.POST("/subscribe/paypal", subscribeH.PayPal)
		sub.GET("/subscribe/paypal/capture/:order_id", subscribeH.CapturePaypalOrder)
	}

	packages := e.Group("/merchant", middlewares.AuthorizeJWT(JWT))
	{
		packages.POST("/packages/create", subscribeH.Create)
	}
	merchant := e.Group("/merchant", middlewares.AuthorizeJWT(JWT))
	{
		merchant.PUT("/update", merchantH.Update)
		merchant.GET("/get", merchantH.Get)
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

	methode := e.Group("/merchant", middlewares.AuthorizeJWT(JWT))
	{
		methode.POST("/methode-pay/qris", transactionmethodeH.Create)
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

	paymentmethod := e.Group("/merchant", middlewares.AuthorizeJWT(JWT))
	{
		paymentmethod.POST("/payment-method/create", paymentmethodH.Create)
		paymentmethod.GET("/payment-method", paymentmethodH.Get)
		paymentmethod.PUT("/payment-method/update/:id", paymentmethodH.Update)
		paymentmethod.DELETE("/payment-method/:id", paymentmethodH.Delete)
		paymentmethod.DELETE("/payment-method/bulk-delete", paymentmethodH.BulkDelete)

	}
	e.GET("/merchant/uploads/:file_name", productH.GetPicture)
	e.GET("/merchant/payment-method/uploads/:file_name", paymentmethodH.GetPicture)

	history := e.Group("/merchant", middlewares.AuthorizeJWT(JWT))
	{
		history.GET("/history/pagination", historyH.Get)
		history.GET("/history/:id", historyH.GetById)
		history.PUT("/history/expire/:order_id", historyH.CheckExpire)
	}

	user := e.Group("/merchant", middlewares.AuthorizeJWT(JWT))
	{
		user.POST("/user/create", userH.Create)
		user.GET("/user/pagination", userH.Get)
		user.PUT("/user/update/:id", userH.Update)
		user.DELETE("/user/:id", userH.Delete)
		user.DELETE("/user/bulk-delete", userH.BulkDelete)
	}

	usermerchant := e.Group("/merchant", middlewares.AuthorizeJWT(JWT))
	{
		usermerchant.POST("/user_merchant/create", usermerchantH.Create)
		usermerchant.GET("/user_merchant/pagination", usermerchantH.Get)
		usermerchant.GET("/user_merchant/:id", usermerchantH.GetById)
		usermerchant.PUT("/user_merchant/update/:id", usermerchantH.Update)
		usermerchant.DELETE("/user_merchant/:id", usermerchantH.Delete)
		usermerchant.DELETE("/user_merchant/bulk-delete", usermerchantH.BulkDelete)
	}

	deleteAccount := e.Group("api/account", middlewares.AuthorizeJWT(JWT))
	{
		deleteAccount.POST("/request-delete", deleteaccountH.Create)
		// deleteAccount.GET("/unit/pagination", unitH.Get)
		// deleteAccount.PUT("/unit/:id", unitH.Update)
		// deleteAccount.DELETE("/unit/:id", unitH.Delete)
		// deleteAccount.DELETE("/unit/bulk-delete", unitH.BulkDelete)
	}

	product := e.Group("/merchant", middlewares.AuthorizeJWT(JWT))
	{
		product.POST("/product/create", productH.Create)
		product.GET("/product/:id", productH.GetById)
		product.DELETE("/product/:id", productH.Delete)
		product.DELETE("/product/bulk-delete", productH.BulkDelete)
		product.PUT("/product/bulk-edit", productH.BulkEdit)
		product.PUT("/product/update/:id", productH.Update)
		product.GET("/product/pagination", productH.Get)
		product.PUT("/product/upload/:id", productH.UploadImage)
	}

	return e
}
