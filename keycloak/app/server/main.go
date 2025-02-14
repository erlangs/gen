package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	//_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
	//_ "github.com/jinzhu/gorm/dialects/mssql"

	"github.com/gin-gonic/gin"
	//"github.com/jinzhu/gorm"
	"github.com/droundy/goopt"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware

	"gorm.io/driver/mysql"
	"keycloak/rest/api/api"
	"keycloak/rest/api/dao"
	//"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	// BuildDate date string of when build was performed filled in by -X compile flag
	BuildDate string

	// LatestCommit date string of when build was performed filled in by -X compile flag
	LatestCommit string

	// BuildNumber date string of when build was performed filled in by -X compile flag
	BuildNumber string

	// BuiltOnIP date string of when build was performed filled in by -X compile flag
	BuiltOnIP string

	// BuiltOnOs date string of when build was performed filled in by -X compile flag
	BuiltOnOs string

	// RuntimeVer date string of when build was performed filled in by -X compile flag
	RuntimeVer string

	// OsSignal signal used to shutdown
	OsSignal chan os.Signal
)

// GinServer launch gin server
func GinServer() (err error) {
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	api.ConfigGinRouter(router)
	router.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting server, the error is '%v'", err)
	}

	return
}

// @title Sample CRUD api for keycloak db
// @version 1.0
// @description Sample CRUD api for keycloak db
// @termsOfService

// @contact.name Me
// @contact.url http://me.com/terms.html
// @contact.email me@me.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	OsSignal = make(chan os.Signal, 1)

	// Define version information
	goopt.Version = fmt.Sprintf(
		`Application build information
  Build date      : %s
  Build number    : %s
  Git commit      : %s
  Runtime version : %s
  Built on OS     : %s
`, BuildDate, BuildNumber, LatestCommit, RuntimeVer, BuiltOnOs)
	goopt.Parse(nil)

	db, err := gorm.Open(mysql.Open("root:rootroot@tcp(localhost:3306)/keycloak?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	//db, err := gorm.Open(postgres.Open("user=postgres password=postgres dbname=keycloak sslmode=disable"), &gorm.Config{})
	//db, err := gorm.Open("postgres", "user=postgres password=postgres dbname=keycloak sslmode=disable")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	db.Logger.LogMode(logger.Silent)
	//db.LogMode(true)
	dao.DB = db

	//db.AutoMigrate(
	//	&model.AdminEventEntity{},
	//	&model.AssociatedPolicy{},
	//	&model.AuthenticationExecution{},
	//	&model.AuthenticationFlow{},
	//	&model.AuthenticatorConfig{},
	//	&model.AuthenticatorConfigEntry{},
	//	&model.BrokerLink{},
	//	&model.Client{},
	//	&model.ClientAttributes{},
	//	&model.ClientAuthFlowBindings{},
	//	&model.ClientInitialAccess{},
	//	&model.ClientNodeRegistrations{},
	//	&model.ClientScope{},
	//	&model.ClientScopeAttributes{},
	//	&model.ClientScopeClient{},
	//	&model.ClientScopeRoleMapping{},
	//	&model.ClientSession{},
	//	&model.ClientSessionAuthStatus{},
	//	&model.ClientSessionNote{},
	//	&model.ClientSessionProtMapper{},
	//	&model.ClientSessionRole{},
	//	&model.ClientUserSessionNote{},
	//	&model.Component{},
	//	&model.ComponentConfig{},
	//	&model.CompositeRole{},
	//	&model.Credential{},
	//	&model.Databasechangelog{},
	//	&model.Databasechangeloglock{},
	//	&model.DefaultClientScope{},
	//	&model.EventEntity{},
	//	&model.FedUserAttribute{},
	//	&model.FedUserConsent{},
	//	&model.FedUserConsentClScope{},
	//	&model.FedUserCredential{},
	//	&model.FedUserGroupMembership{},
	//	&model.FedUserRequiredAction{},
	//	&model.FedUserRoleMapping{},
	//	&model.FederatedIdentity{},
	//	&model.FederatedUser{},
	//	&model.GroupAttribute{},
	//	&model.GroupRoleMapping{},
	//	&model.IdentityProvider{},
	//	&model.IdentityProviderConfig{},
	//	&model.IdentityProviderMapper{},
	//	&model.IdpMapperConfig{},
	//	&model.KeycloakGroup{},
	//	&model.KeycloakRole{},
	//	&model.MigrationModel{},
	//	&model.OfflineClientSession{},
	//	&model.OfflineUserSession{},
	//	&model.PolicyConfig{},
	//	&model.ProtocolMapper{},
	//	&model.ProtocolMapperConfig{},
	//	&model.Realm{},
	//	&model.RealmAttribute{},
	//	&model.RealmDefaultGroups{},
	//	&model.RealmEnabledEventTypes{},
	//	&model.RealmEventsListeners{},
	//	&model.RealmLocalizations{},
	//	&model.RealmRequiredCredential{},
	//	&model.RealmSMTPConfig{},
	//	&model.RealmSupportedLocales{},
	//	&model.RedirectUris{},
	//	&model.RequiredActionConfig{},
	//	&model.RequiredActionProvider{},
	//	&model.ResourceAttribute{},
	//	&model.ResourcePolicy{},
	//	&model.ResourceScope{},
	//	&model.ResourceServer{},
	//	&model.ResourceServerPermTicket{},
	//	&model.ResourceServerPolicy{},
	//	&model.ResourceServerResource{},
	//	&model.ResourceServerScope{},
	//	&model.ResourceUris{},
	//	&model.RoleAttribute{},
	//	&model.ScopeMapping{},
	//	&model.ScopePolicy{},
	//	&model.UserAttribute{},
	//	&model.UserConsent{},
	//	&model.UserConsentClientScope{},
	//	&model.UserEntity{},
	//	&model.UserFederationConfig{},
	//	&model.UserFederationMapper{},
	//	&model.UserFederationMapperConfig{},
	//	&model.UserFederationProvider{},
	//	&model.UserGroupMembership{},
	//	&model.UserRequiredAction{},
	//	&model.UserRoleMapping{},
	//	&model.UserSession{},
	//	&model.UserSessionNote{},
	//	&model.UsernameLoginFailure{},
	//	&model.WebOrigins{},
	//)

	dao.Logger = func(ctx context.Context, sql string) {
		fmt.Printf("SQL: %s\n", sql)
	}

	go GinServer()
	LoopForever()
}

// LoopForever on signal processing
func LoopForever() {
	fmt.Printf("Entering infinite loop\n")

	signal.Notify(OsSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
	_ = <-OsSignal

	fmt.Printf("Exiting infinite loop received OsSignal\n")

}
