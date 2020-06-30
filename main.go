package main

import (
	"PNM/controller"
	"PNM/controller/crud"
	"PNM/model"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
)

func main() {
	/*
	 convert database to struct
	 run terminal -> tables-to-go -t mysql -h localhost -d pnm -u root -of database
	*/
	e := echo.New()
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(err)
	}

	e.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{os.Getenv("FRONTENDURL")},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}),
		middleware.Recover(),   //recover server on production if it's stop
		middleware.Logger(),    //logging
		middleware.RequestID(), //add request ID in every route
		middleware.Secure(),    //secure
	)

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		c.Logger().Error(report)

		var (
			code    = report.Code
			message = report.Message
		)

		switch report.Message {
		case "not found":
			code = http.StatusNotFound
			break
		case "sql: no rows in result set":
			code = http.StatusNotFound
			message = "data not found"
			break
		}

		_ = c.JSON(code, model.ErrorResponse{Message: message})
	}

	e.GET("/migrate", controller.Migrate)

	user := e.Group("/auth")
	user.POST("/login", controller.Login)

	api := e.Group("/api")

	// jwt token middleware

	//api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	//	SigningKey:  []byte(os.Getenv("JWTTOKEN")),
	//	TokenLookup: "header:token",
	//}))

	// edit user
	api.POST("/getUser", crud.GetALLUser)
	api.POST("/getUserById", crud.GetUserById)
	api.PUT("/editUser", crud.EditUser)
	api.POST("/getListUser", crud.GetListUser)
	api.POST("/getDivisionDropdown", crud.GetDivisionDropdown)

	// CRUD surat priority route
	api.POST("/getSuratPriority", crud.GetSuratPriorityALL)
	api.POST("/getSuratPriorityById", crud.GetSuratPriorById)
	api.POST("/createSuratPriority", crud.AddSuratPriority)
	api.PUT("/editSuratPriority", crud.EditSuratPriority)
	api.POST("/deleteSuratPriority", crud.DeleteSuratPriority)
	api.POST("/getSuratPriorityDropdown", crud.GetSuratPriorityDropdown)

	// CRUD surat category route
	api.POST("/getSuratCategory", crud.GetAllCategory)
	api.POST("/getSuratCategoryById", crud.GetCategoryById)
	api.POST("/createSuratCategory", crud.AddCategory)
	api.PUT("/editSuratCategory", crud.EditCategory)
	api.POST("/deleteSuratCategory", crud.DeleteCategory)
	api.POST("/getSuratCategoryDropdown", crud.CategoryDropdown)

	// CRUD surat type route
	api.POST("/getSuratType", crud.GetSuratTypeALL)
	api.POST("/getSuratTypeById", crud.GetSuratTypeById)
	api.POST("/createSuratType", crud.AddSuratType)
	api.PUT("/editSuratType", crud.EditSuratType)
	api.POST("/deleteSuratType", crud.DeleteSuratType)
	api.POST("/getSuratTypeDropdown", crud.GetSuratTypeDropdown)

	// CRUD surat status route
	api.POST("/getSuratStatusByCode", crud.GetStatusByCode)
	api.PUT("/editSuratStatus", crud.EditSuratStatus)
	api.POST("/getSuratStatus", crud.GetAllSuratStatus)

	// CRUD memo template
	api.POST("/getMemoTemplate", crud.GetMemoTemplateAll)
	api.POST("/getMemoTemplateById", crud.GetMemoTemplateById)
	api.POST("/createMemoTemplate", crud.CreateMemoTemplate)
	api.PUT("/updateMemoTemplate", crud.UpdateMemoTemplate)
	api.POST("/deleteMemoTemplate", crud.DeleteMemoTemplate)
	api.POST("/getMemoTemplateDropdown", crud.GetMemoTemplateDropdown)

	// CRUD surat masuk
	api.POST("/addSuratMasuk", crud.CreateSuratMasuk)
	api.PUT("/editSuratMasuk", crud.EditSuratMasuk)
	api.POST("/viewByStatus", crud.ViewByStatus)
	api.POST("/editStatus", crud.EditStatus)
	api.POST("/viewById", crud.ViewById)
	api.POST("/printPdf", crud.PrintPDF)

	// List assigned user
	api.POST("/addAssignedUser", crud.AddRecipient)

	// middleware for SSO custom
	//
	// user.Use(middleware2.PNMMiddleware)

	// getting jwt token example
	//
	//user.POST("/protectedRoute", func(context echo.Context) error {
	//	user := context.Get("user").(*jwt.Token)
	//	claims := user.Claims.(jwt.MapClaims)
	//	name := claims["name"].(string)
	//	return context.JSON(200, model.Response{Message: "hello "+name})
	//})

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
