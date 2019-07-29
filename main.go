package main

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/url"
	"os"
	"time"

	_userHttpDeliver "github.com/user-management-with-go/user/delivery/http"
	_userRepo "github.com/user-management-with-go/user/repository"
	_userUcase "github.com/user-management-with-go/user/usecase"

	_pageHttpDeliver "github.com/user-management-with-go/page/delivery/http"
	_pageRepo "github.com/user-management-with-go/page/repository"
	_pageUcase "github.com/user-management-with-go/page/usecase"

	_roleHttpDeliver "github.com/user-management-with-go/role/delivery/http"
	_roleRepo "github.com/user-management-with-go/role/repository"
	_roleUcase "github.com/user-management-with-go/role/usecase"

	_actelementHttpDeliver "github.com/user-management-with-go/actionelement/delivery/http"
	_actelementRepo "github.com/user-management-with-go/actionelement/repository"
	_actelementUcase "github.com/user-management-with-go/actionelement/usecase"

	_menuHttpDeliver "github.com/user-management-with-go/menu/delivery/http"
	_menuRepo "github.com/user-management-with-go/menu/repository"
	_menuUcase "github.com/user-management-with-go/menu/usecase"

	_elementHttpDeliver "github.com/user-management-with-go/element/delivery/http"
	_elementRepo "github.com/user-management-with-go/element/repository"
	_elementUcase "github.com/user-management-with-go/element/usecase"

	_privilegeHttpDeliver "github.com/user-management-with-go/privilege/delivery/http"
	_privilegeRepo "github.com/user-management-with-go/privilege/repository"
	_privilegeUcase "github.com/user-management-with-go/privilege/usecase"

	_privilegetypeHttpDeliver "github.com/user-management-with-go/privilegetype/delivery/http"
	_privilegetypeRepo "github.com/user-management-with-go/privilegetype/repository"
	_privilegetypeUcase "github.com/user-management-with-go/privilegetype/usecase"

	_healthHttpDeliver "github.com/user-management-with-go/health/delivery/http"

	"github.com/user-management-with-go/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

}

func main() {

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer dbConn.Close()
	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)
	userRepo := _userRepo.NewMysqlUserRepository(dbConn)
	pageRepo := _pageRepo.NewMysqlPageRepository(dbConn)
	roleRepo := _roleRepo.NewMysqlRoleRepository(dbConn)
	actelementRepo := _actelementRepo.NewMysqlActionelementRepository(dbConn)
	menuRepo := _menuRepo.NewMysqlMenuRepository(dbConn)
	elementRepo := _elementRepo.NewMysqlElementRepository(dbConn)
	privilegeRepo := _privilegeRepo.NewMysqlPrivilegeRepository(dbConn)
	privilegetypeRepo := _privilegetypeRepo.NewMysqlPrivilegeTypeRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	au := _userUcase.NewUserUsecase(userRepo, timeoutContext)
	_userHttpDeliver.NewUserHttpHandler(e, au)

	ab := _pageUcase.NewPageUsecase(pageRepo, timeoutContext)
	_pageHttpDeliver.NewPageHttpHandler(e, ab)

	ac := _roleUcase.NewRoleUsecase(roleRepo, timeoutContext)
	_roleHttpDeliver.NewRoleHttpHandler(e, ac)

	ad := _actelementUcase.NewActionelementUsecase(actelementRepo, timeoutContext)
	_actelementHttpDeliver.NewActionelementHttpHandler(e, ad)

	ae := _menuUcase.NewMenuUsecase(menuRepo, timeoutContext)
	_menuHttpDeliver.NewMenuHttpHandler(e, ae)

	af := _elementUcase.NewElementUsecase(elementRepo, timeoutContext)
	_elementHttpDeliver.NewElementHttpHandler(e, af)

	ag := _privilegeUcase.NewPrivilegeUsecase(privilegeRepo, timeoutContext)
	_privilegeHttpDeliver.NewPrivilegeHttpHandler(e, ag)

	ah := _privilegetypeUcase.NewPrivilegeTypeUsecase(privilegetypeRepo, timeoutContext)
	_privilegetypeHttpDeliver.NewPrivilegeTypeHttpHandler(e, ah)

	_healthHttpDeliver.NewHealthHttpHandler(e, nil)


	e.Start(viper.GetString("server.address"))
}
