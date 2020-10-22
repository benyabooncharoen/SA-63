package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/benyabooncharoen/app/controllers"
	_ "github.com/benyabooncharoen/app/docs"
	"github.com/benyabooncharoen/app/ent"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Genders struct {
	Gender []Gender
}

type Gender struct {
	genderName string
}

type Rightoftreatments struct {
	Rightoftreatment []Rightoftreatment
}

type Rightoftreatment struct {
	rightoftreatmentName string
}

type Systemmembers struct {
	Systemmember []Systemmember
}

type Systemmember struct {
	systemmemberName 	string
	Password			string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewGenderController(v1, client)
	controllers.NewRightoftreatmentController(v1, client)
	controllers.NewSystemmemberController(v1, client)
	controllers.NewPatientController(v1, client)

	// Set Genders Data
	genders := Genders{
		Gender: []Gender{
			Gender{"Male"},
			Gender{"Famale"},
		},
	}
	for _, g := range genders.Gender {
		client.Gender.
			Create().
			SetGenderName(g.genderName).
			Save(context.Background())
	}

	// Set Rightoftreatments Data
	rightoftreatments := Rightoftreatments{
		Rightoftreatment: []Rightoftreatment{
			Rightoftreatment{"สิทธิสวัสดิการการรักษาพยาบาลของข้าราชการ"},
			Rightoftreatment{"สิทธิประกันสังคม"},
			Rightoftreatment{"สิทธิหลักประกับสุขภาพ 30 บาท"},
		},
	}
	for _, r := range rightoftreatments.Rightoftreatment {
		client.Rightoftreatment.
			Create().
			SetRightoftreatmentName(r.rightoftreatmentName).
			Save(context.Background())
	}

	// Set Systemmembers Data
	systemmembers := Systemmembers{
		Systemmember: []Systemmember{
			Systemmember{"Ji soo","12345687"},
			Systemmember{"Li saa","23649523"},
			Systemmember{"Min Nie","68499986"},

		},
	}
	for _, sm := range systemmembers.Systemmember {
		client.Systemmember.
			Create().
			SetSystemmemberName(sm.systemmemberName).
			SetPassword(sm.Password).
			Save(context.Background())
	}
	
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}