package routes

import (
	"sitax/controller"
	"sitax/controller/c_file"
	"sitax/controller/c_group"
	"sitax/controller/c_kantor"
	"sitax/controller/c_kewenangan"
	"sitax/controller/c_menu"
	"sitax/controller/c_pajak"
	"sitax/controller/c_pajak_detail"
	"sitax/controller/c_panduan_pajak"
	"sitax/controller/c_referensi"
	"sitax/controller/c_user"
	"sitax/handler"
	"sitax/middleware"
	"sitax/repository/r_file"
	"sitax/repository/r_group"
	"sitax/repository/r_kantor"
	"sitax/repository/r_kewenangan"
	"sitax/repository/r_menu"
	"sitax/repository/r_pajak"
	"sitax/repository/r_pajak_detail"
	"sitax/repository/r_panduan_pajak"
	"sitax/repository/r_referensi"
	"sitax/repository/r_user"

	"sitax/service/s_file"
	"sitax/service/s_group"
	"sitax/service/s_kantor"
	"sitax/service/s_kewenangan"
	"sitax/service/s_menu"
	"sitax/service/s_pajak"
	"sitax/service/s_pajak_detail"
	"sitax/service/s_panduan_pajak"
	"sitax/service/s_referensi"
	"sitax/service/s_user"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Secure JSON prefix
	r.SecureJsonPrefix(")]}',\n")

	// Setup auth middleware
	authMiddleware, err := middleware.AuthMiddleware()
	if err != nil {
		panic(err)
	}

	//User
	registerUserRepo := r_user.NewUserRepository()
	registerUser := c_user.NewUserController(registerUserRepo)

	getUserRepo := r_user.NewGetUserRepository()
	userGetService := s_user.NewGetUserService(getUserRepo)

	updateUserRepo := r_user.NewUpdateUserRepository()
	userUpdateService := s_user.NewUpdateUserService(updateUserRepo)

	userDeleteRepo := r_user.NewDeleteUserRepository()
	userDeleteService := s_user.NewDeleteUserService(userDeleteRepo)

	//Group
	getGroupRepo := r_group.NewGetGroupRepository()
	groupGetService := s_group.NewGetGroupService(getGroupRepo)

	addGroupRepo := r_group.NewAddGroupRepository()
	groupAddService := s_group.NewAddGroupService(addGroupRepo)

	updateGroupRepo := r_group.NewUpdateGroupRepository()
	groupUpdateService := s_group.NewUpdateGroupService(updateGroupRepo)

	groupDeleteRepo := r_group.NewDeleteGroupRepository()
	groupDeleteService := s_group.NewDeleteGroupService(groupDeleteRepo)

	//Menu
	getMenuRepo := r_menu.NewGetMenuRepository()
	menuGetService := s_menu.NewGetMenuService(getMenuRepo)

	addMenuRepo := r_menu.NewAddMenuRepository()
	menuAddService := s_menu.NewAddMenuService(addMenuRepo)

	updateMenuRepo := r_menu.NewUpdateMenuRepository()
	menuUpdateService := s_menu.NewUpdateMenuService(updateMenuRepo)

	menuDeleteRepo := r_menu.NewDeleteMenuRepository()
	menuDeleteService := s_menu.NewDeleteMenuService(menuDeleteRepo)

	//Kewenangan
	getKewenanganRepo := r_kewenangan.NewGetKewenanganRepository()
	kewenanganGetService := s_kewenangan.NewGetKewenanganService(getKewenanganRepo)

	addKewenanganRepo := r_kewenangan.NewAddKewenanganRepository()
	kewenanganAddService := s_kewenangan.NewAddKewenanganService(addKewenanganRepo)

	updateKewenanganRepo := r_kewenangan.NewUpdateKewenanganRepository()
	kewenanganUpdateService := s_kewenangan.NewUpdateKewenanganService(updateKewenanganRepo)

	kewenanganDeleteRepo := r_kewenangan.NewDeleteKewenanganRepository()
	kewenanganDeleteService := s_kewenangan.NewDeleteKewenanganService(kewenanganDeleteRepo)

	//Kantor
	getKantorRepo := r_kantor.NewGetKantorRepository()
	kantorGetService := s_kantor.NewGetKantorService(getKantorRepo)

	addKantorRepo := r_kantor.NewAddKantorRepository()
	kantorAddService := s_kantor.NewAddKantorService(addKantorRepo)

	updateKantorRepo := r_kantor.NewUpdateKantorRepository()
	kantorUpdateService := s_kantor.NewUpdateKantorService(updateKantorRepo)

	kantorDeleteRepo := r_kantor.NewDeleteKantorRepository()
	kantorDeleteService := s_kantor.NewDeleteKantorService(kantorDeleteRepo)

	//Pajak
	getPajakRepo := r_pajak.NewGetPajakRepository()
	pajakGetService := s_pajak.NewGetPajakService(getPajakRepo)

	addPajakRepo := r_pajak.NewAddPajakRepository()
	pajakAddService := s_pajak.NewAddPajakService(addPajakRepo)

	updatePajakRepo := r_pajak.NewUpdatePajakRepository()
	pajakUpdateService := s_pajak.NewUpdatePajakService(updatePajakRepo)

	pajakDeleteRepo := r_pajak.NewDeletePajakRepository()
	pajakDeleteService := s_pajak.NewDeletePajakService(pajakDeleteRepo)

	//Pajak Detail
	getPajakDetailRepo := r_pajak_detail.NewGetPajakDetailRepository()
	pajak_detailGetService := s_pajak_detail.NewGetPajakDetailService(getPajakDetailRepo)

	addPajakDetailRepo := r_pajak_detail.NewAddPajakDetailRepository()
	pajak_detailAddService := s_pajak_detail.NewAddPajakDetailService(addPajakDetailRepo)

	updatePajakDetailRepo := r_pajak_detail.NewUpdatePajakDetailRepository()
	pajak_detailUpdateService := s_pajak_detail.NewUpdatePajakDetailService(updatePajakDetailRepo)

	pajak_detailDeleteRepo := r_pajak_detail.NewDeletePajakDetailRepository()
	pajak_detailDeleteService := s_pajak_detail.NewDeletePajakDetailService(pajak_detailDeleteRepo)

	//Panduan Pajak
	getPanduanPajakRepo := r_panduan_pajak.NewGetPanduanPajakRepository()
	panduan_pajakGetService := s_panduan_pajak.NewGetPanduanPajakService(getPanduanPajakRepo)

	addPanduanPajakRepo := r_panduan_pajak.NewAddPanduanPajakRepository()
	panduan_pajakAddService := s_panduan_pajak.NewAddPanduanPajakService(addPanduanPajakRepo)

	updatePanduanPajakRepo := r_panduan_pajak.NewUpdatePanduanPajakRepository()
	panduan_pajakUpdateService := s_panduan_pajak.NewUpdatePanduanPajakService(updatePanduanPajakRepo)

	panduan_pajakDeleteRepo := r_panduan_pajak.NewDeletePanduanPajakRepository()
	panduan_pajakDeleteService := s_panduan_pajak.NewDeletePanduanPajakService(panduan_pajakDeleteRepo)

	//File
	getFileRepo := r_file.NewGetFileRepository()
	fileGetService := s_file.NewGetFileService(getFileRepo)

	addFileRepo := r_file.NewAddFileRepository()
	fileAddService := s_file.NewAddFileService(addFileRepo)

	updateFileRepo := r_file.NewUpdateFileRepository()
	fileUpdateService := s_file.NewUpdateFileService(updateFileRepo)

	fileDeleteRepo := r_file.NewDeleteFileRepository()
	fileDeleteService := s_file.NewDeleteFileService(fileDeleteRepo)

	//Referensi
	getReferensiRepo := r_referensi.NewGetReferensiRepository()
	referensiGetService := s_referensi.NewGetReferensiService(getReferensiRepo)

	// Create controller instance
	userGetController := c_user.NewGetUserController(userGetService)
	userUpdateController := c_user.NewUpdateUserController(userUpdateService)
	userDeleteController := c_user.NewUserDeleteController(userDeleteService)

	groupGetController := c_group.NewGetGroupController(groupGetService)
	groupAddController := c_group.NewGroupAddController(groupAddService)
	groupUpdateController := c_group.NewUpdateGroupController(groupUpdateService)
	groupDeleteController := c_group.NewGroupDeleteController(groupDeleteService)

	menuGetController := c_menu.NewGetMenuController(menuGetService)
	menuAddController := c_menu.NewMenuAddController(menuAddService)
	menuUpdateController := c_menu.NewUpdateMenuController(menuUpdateService)
	menuDeleteController := c_menu.NewMenuDeleteController(menuDeleteService)

	kewenanganGetController := c_kewenangan.NewGetKewenanganController(kewenanganGetService)
	kewenanganAddController := c_kewenangan.NewKewenanganAddController(kewenanganAddService)
	kewenanganUpdateController := c_kewenangan.NewUpdateKewenanganController(kewenanganUpdateService)
	kewenanganDeleteController := c_kewenangan.NewKewenanganDeleteController(kewenanganDeleteService)

	kantorGetController := c_kantor.NewGetKantorController(kantorGetService)
	kantorAddController := c_kantor.NewKantorAddController(kantorAddService)
	kantorUpdateController := c_kantor.NewUpdateKantorController(kantorUpdateService)
	kantorDeleteController := c_kantor.NewKantorDeleteController(kantorDeleteService)

	pajakGetController := c_pajak.NewGetPajakController(pajakGetService)
	pajakAddController := c_pajak.NewPajakAddController(pajakAddService)
	pajakUpdateController := c_pajak.NewUpdatePajakController(pajakUpdateService)
	pajakDeleteController := c_pajak.NewPajakDeleteController(pajakDeleteService)

	pajak_detailGetController := c_pajak_detail.NewGetPajakDetailController(pajak_detailGetService)
	pajak_detailAddController := c_pajak_detail.NewPajakDetailAddController(pajak_detailAddService)
	pajak_detailUpdateController := c_pajak_detail.NewUpdatePajakDetailController(pajak_detailUpdateService)
	pajak_detailDeleteController := c_pajak_detail.NewPajakDetailDeleteController(pajak_detailDeleteService)

	panduan_pajakGetController := c_panduan_pajak.NewGetPanduanPajakController(panduan_pajakGetService)
	panduan_pajakAddController := c_panduan_pajak.NewPanduanPajakAddController(panduan_pajakAddService)
	panduan_pajakUpdateController := c_panduan_pajak.NewUpdatePanduanPajakController(panduan_pajakUpdateService)
	panduan_pajakDeleteController := c_panduan_pajak.NewPanduanPajakDeleteController(panduan_pajakDeleteService)

	fileGetController := c_file.NewGetFileController(fileGetService)
	fileAddController := c_file.NewFileAddController(fileAddService)
	fileUpdateController := c_file.NewUpdateFileController(fileUpdateService)
	fileDeleteController := c_file.NewFileDeleteController(fileDeleteService)

	referensiGetController := c_referensi.NewGetReferensiController(referensiGetService)

	// Apply to public routes
	r.GET("/", controller.Helloworld)
	r.POST("/login", handler.LoginHandler)
	r.POST("/register", registerUser.RegisterUser)

	r.GET("/group", groupGetController.GetAllGroup)
	r.POST("/group/", groupAddController.AddGroup)
	r.PUT("group/:group_id", groupUpdateController.UpdateGroup)
	r.DELETE("group/:group_id", groupDeleteController.DeleteGroup)

	r.GET("/menu", menuGetController.GetAllMenu)
	r.POST("/menu/", menuAddController.AddMenu)
	r.PUT("menu/:menu_id", menuUpdateController.UpdateMenu)
	r.DELETE("menu/:menu_id", menuDeleteController.DeleteMenu)

	r.GET("/kewenangan", kewenanganGetController.GetAllKewenangan)
	r.POST("/kewenangan/", kewenanganAddController.AddKewenangan)
	r.PUT("/kewenangan/:group_id/:menu_id", kewenanganUpdateController.UpdateKewenangan)
	r.DELETE("/kewenangan/:group_id/:menu_id", kewenanganDeleteController.DeleteKewenangan)

	r.GET("/kantor", kantorGetController.GetAllKantor)
	r.POST("/kantor/", kantorAddController.AddKantor)
	r.PUT("kantor/:kd_kantor", kantorUpdateController.UpdateKantor)
	r.DELETE("kantor/:kd_kantor", kantorDeleteController.DeleteKantor)

	r.GET("/referensi", referensiGetController.GetAllReferensi)
	r.GET("/pubpajak", pajakGetController.GetAllPajak)
	r.GET("/pubpajak_detail/:pajak_id", pajakGetController.GetPajakByID)
	// r.GET("/users", userGetController.GetAllUser)
	// Apply auth middleware to routes
	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/users", userGetController.GetAllUser)
		auth.GET("/users/:username", userGetController.GetUserByID)
		auth.PUT("/users/:username", userUpdateController.UpdateUser)
		auth.DELETE("/users/:username", userDeleteController.DeleteUser)

		auth.GET("/pajak", pajakGetController.GetAllPajak)
		auth.POST("/pajak/", pajakAddController.AddPajak)
		auth.PUT("pajak/:pajak_id", pajakUpdateController.UpdatePajak)
		auth.DELETE("pajak/:pajak_id", pajakDeleteController.DeletePajak)

		auth.GET("/pajak-detail", pajak_detailGetController.GetAllPajakDetail)
		auth.POST("/pajak-detail/", pajak_detailAddController.AddPajakDetail)
		auth.PUT("pajak-detail/:pajak_detail_id", pajak_detailUpdateController.UpdatePajakDetail)
		auth.DELETE("pajak-detail/:pajak_detail_id", pajak_detailDeleteController.DeletePajakDetail)

		auth.GET("/panduan_pajak", panduan_pajakGetController.GetAllPanduanPajak)
		auth.POST("/panduan_pajak/", panduan_pajakAddController.AddPanduanPajak)
		auth.PUT("panduan_pajak/:panduan_pajak_id", panduan_pajakUpdateController.UpdatePanduanPajak)
		auth.DELETE("panduan_pajak/:panduan_pajak_id", panduan_pajakDeleteController.DeletePanduanPajak)

		auth.GET("/file", fileGetController.GetAllFile)
		auth.POST("/file/", fileAddController.AddFile)
		auth.PUT("file/:file_id", fileUpdateController.UpdateFile)
		auth.DELETE("file/:file_id", fileDeleteController.DeleteFile)
	}
	return r
}
