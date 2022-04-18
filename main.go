package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/aliepba/go-crypto/app/handlers"
	"github.com/aliepba/go-crypto/app/methods"
	"github.com/aliepba/go-crypto/app/services"
	"github.com/aliepba/go-crypto/auth"
	"github.com/aliepba/go-crypto/helpers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func authMiddleware(authService auth.Service, userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helpers.APIResponse("Unthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helpers.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helpers.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)

		if err != nil {
			response := helpers.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)

	}
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/crypto?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	//repository
	userMethod := methods.NewMethodUser(db)
	coinMethod := methods.NewMethodCoin(db)
	metadataMethod := methods.NewMethodMetadata(db)
	categoryMethod := methods.NewMethodCategory(db)
	airdropMethod := methods.NewMethodAirdrop(db)

	//service
	userService := services.NewServiceUser(userMethod)
	coinService := services.NewServiceCoin(coinMethod)
	metadataService := services.NewServiceMetadata(metadataMethod)
	categoryService := services.NewServiceCategory(categoryMethod)
	airdropService := services.NewServiceAirdrop(airdropMethod)
	authService := auth.NewService()

	//handler
	userHandler := handlers.NewUserHandler(userService, authService)
	coinHandler := handlers.NewCoinHandler(coinService)
	metadataHandler := handlers.NewMetadataHandler(metadataService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	airdropHandler := handlers.NewAirdropHandler(airdropService)

	router := gin.Default()
	api := router.Group("/epiay/v1")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/check-email", userHandler.CheckEmailAvailability)

	api.POST("/coin", authMiddleware(authService, userService), coinHandler.SaveCoin)
	api.GET("/list-coins", authMiddleware(authService, userService), coinHandler.GetCoins)
	api.GET("/crypto/detail/:symbol", authMiddleware(authService, userService), coinHandler.GetCoin)

	api.POST("/metadata", authMiddleware(authService, userService), metadataHandler.SaveMetadata)
	api.GET("/cypto/info", authMiddleware(authService, userService), metadataHandler.GetMetadata)

	api.GET("/categories", authMiddleware(authService, userService), categoryHandler.GetCategories)

	api.POST("/airdrop", authMiddleware(authService, userService), airdropHandler.SaveAirdrop)
	api.GET("/airdrops", authMiddleware(authService, userService), airdropHandler.GetAirdrops)
	api.GET("/airdrop/:id", authMiddleware(authService, userService), airdropHandler.GetAirdrop)
	api.PUT("/airdrop/:id", authMiddleware(authService, userService), airdropHandler.UpdateAirdrop)

	router.Run()
}
