package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/ljg-cqu/core/esignbox_swagin/common"
	"github.com/ljg-cqu/core/esignbox_swagin/file_"
	"github.com/ljg-cqu/core/esignbox_swagin/template"
	"github.com/ljg-cqu/core/esignbox_swagin/token"
	"github.com/ljg-cqu/core/middleware"
	"github.com/ljg-cqu/core/postgres"
	"github.com/long2ice/swagin"
	"io/ioutil"
)

const (
	EsignSandBoxHost = "https://smlopenapi.esign.cn"
)

func main() {
	client := resty.New().SetDebug(true).SetBaseURL(EsignSandBoxHost)
	common.Client = client
	common.PgxPool = postgres.PgxPool(postgres.TestDBAliConnStr)

	//// apply DB migration before Client can work as expected
	//_, err := common.PgxPool.Exec(context.Background(), models.Schema)
	//if err != nil {
	//	log.Printf("failed to to do DB migration for gue queue on top of PostgreSQL:%+v", err)
	//	os.Exit(1)
	//}

	// Use customize Gin engine
	r := gin.New()

	// Registering func(c *gin.Context) is accepted,
	// but the OpenAPI generator will ignore the operation and it won't appear in the specification.
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/upload", func(c *gin.Context) {
		// Single file
		fileH, _ := c.FormFile("file")

		file, _ := fileH.Open()
		defer file.Close()
		bytes, _ := ioutil.ReadAll(file)

		fmt.Println(string(bytes))

		c.JSON(200, "")

		fmt.Println("file name:", fileH.Filename)

		//c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	app := swagin.NewFromEngine(r, NewSwagger())
	//subApp := swagin.NewFromEngine(r, NewSwagger())
	//
	//subApp.GET("/noModel", noModel)
	//app.Mount("/sub", subApp)
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}), middleware.BasicAuth("Esign", map[string]string{
		"admin": "admin",
	}))

	_token := app.Group("/esign", swagin.Tags("e签宝OAuth2.0鉴权接口"))
	_token.POST("/token", token.GetTokenRequestH())

	contractTemplate := app.Group("/contractTemplate", swagin.Tags(" PDF合同模板接口（集成e签宝）"))
	{
		contractTemplate.POST("/upload", template.UploadPDFTemplFileRequestH())
		//contractTemplate.POST("/uploadUrl", template.GetTemplUploadUrlRequestH())
		//contractTemplate.POST("/fillControls//:templateId/:ids", template.DeleteFillControlRequestH())
		//contractTemplate.POST("/fillControls/:templateId", template.AddFillControlRequestH())
		contractTemplate.POST("/fill", template.FillTemplateContentRequestH())

		//contractTemplate.POST("/uploadStatus/:templateId", template.GetTemplUploadStatusRequestH())
		//contractTemplate.GET("/officialTemplateInfo", template.GetTemplInfoRequestH())
		contractTemplate.POST("/details/:templateId", template.GetTemplDetailsRequestH())
		contractTemplate.POST("/details/list", template.GetTemplListRequestH())
	}

	contractFile := app.Group("/contractFile", swagin.Tags("PDF合同文件接口（集成e签宝）"))
	contractFile.POST("/details/:fileId", file_.GetPdfFileDetailsRequestH())
	contractFile.POST("/details/list", file_.GetPdfFileDetailsListRequestH())
	contractFile.POST("/merge/:fileIds", file_.MergeThenUploadFilesRequestH())

	port := ":8085"

	fmt.Printf("Now you can visit http://127.0.0.1%v/docs, http://127.0.0.1%v/redoc or http://127.0.0.1%v/rapidoc to see the api docs. Have fun!\n\n", port, port, port)
	if err := app.Run(port); err != nil {
		panic(err)
	}
}