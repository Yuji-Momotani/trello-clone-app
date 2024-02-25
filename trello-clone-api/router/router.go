package router

import (
	// 各環境に合わせてcontrollerをimport

	"net/http"
	"os"
	"trello-colen-api/controller"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// DefaultCORSConfig = CORSConfig{
		// 	Skipper:      DefaultSkipper,
		// 	AllowOrigins: []string{"*"},
		// 	AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		// }
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowOrigin, echo.HeaderXCSRFToken},
		AllowCredentials: true,
	}))

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		// DefaultCSRFConfig = CSRFConfig{
		// 	Skipper:      DefaultSkipper,
		// 	TokenLength:  32,
		// 	TokenLookup:  "header:" + echo.HeaderXCSRFToken,
		// 	ContextKey:   "csrf",
		// 	CookieName:   "_csrf",
		// 	CookieMaxAge: 86400,
		// }
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieSecure:   true,
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		// CookieSameSite: http.SameSiteDefaultMode, //PostMan確認用。（SameSiteNoneModeだとSecureが自動でtrueになるため）
	}))

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)
	e.GET("/csrf", uc.GetCsrf)

	t := e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		// **************************
		// /hoge/*に対してjwtを使用する例
		// **************************
		// Optional. Default value "header:Authorization".
		TokenLookup: "cookie:token",
		SigningKey:  []byte(os.Getenv("SECRET")),
	}))
	t.GET("", tc.GetTaskCardsAndTasks)
	// task-cardの処理
	t.POST("/card", tc.RegistTaskCard)
	t.PUT("/card/:taskCardId", tc.UpdateTaskCard)
	t.DELETE("/card/:taskCardId", tc.DeleteTaskCard)
	// taskの処理
	t.POST("/task", tc.RegistTask)
	t.PUT("/task/:taskId", tc.UpdateTask)
	t.DELETE("/task/:taskId", tc.DeleteTask)

	e.Use(middleware.Logger())
	return e
}
