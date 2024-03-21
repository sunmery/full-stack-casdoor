package main

import (
	"fmt"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strings"
)

func LoadConf() error {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./app.yaml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Errorf("fatal error config file: %w", err)
			return err
		} else {
			// Config file was found but another error was produced
			fmt.Errorf("找到配置文件，但生成另一个错误: %w", err)
			return err
		}
	}

	return nil
}

func initAuthConfig() {
	Endpoint := viper.GetString("server.endpoint")
	ClientID := viper.GetString("client.client_id")
	ClientSecret := viper.GetString("client.client_secret")
	cate := viper.GetString("client.certificate")
	Organization := viper.GetString("client.organization")
	Application := viper.GetString("client.application")
	casdoorsdk.InitConfig(
		Endpoint,
		ClientID,
		ClientSecret,
		cate,
		Organization,
		Application,
	)
}

func main() {
	const PORT = "0.0.0.0:8080"
	//	 读取文件
	err := LoadConf()
	if err != nil {
		panic(err)
	}

	// 根据yaml进行初始化
	initAuthConfig()

	r := gin.Default()
	r.Use(Cors())
	r.POST("/api/signin", signinHandler)
	r.POST("/api/userinfo", userinfoHandler)

	fmt.Println("Server listening at: ", PORT)
	err1 := r.Run(PORT)
	if err1 != nil {
		panic(err1)
	}
}

func signinHandler(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")

	fmt.Println("code:", code)
	fmt.Println("state:", state)

	token, err := casdoorsdk.GetOAuthToken(code, state)
	if err != nil {
		fmt.Println("GetOAuthToken() error", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   token.AccessToken,
	})
}

func userinfoHandler(c *gin.Context) {
	// 检索JWT
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		//http.Error(w, "authHeader is empty", http.StatusUnauthorized)
		fmt.Println("authHeader is empty")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "authHeader is empty",
		})
		return
	}

	// 判断JWT正确性
	token := strings.Split(authHeader, "Bearer ")
	if len(token) != 2 {
		//http.Error(w, "token is not valid Bearer token", http.StatusUnauthorized)
		fmt.Println("token is not valid Bearer token")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "token is not valid Bearer token",
		})
		return
	}

	claims, err := casdoorsdk.ParseJwtToken(token[1])
	if err != nil {
		//http.Error(w, "ParseJwtToken() error", http.StatusUnauthorized)
		fmt.Println("ParseJwtToken() error")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "ParseJwtToken() error",
		})
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   claims.User,
	})
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,tokenString")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
			//
			//c.Header("Content-Type", "application/json")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")
			c.Header("Access-Control-Expose-Headers", "Authorization, Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.AbortWithStatus(http.StatusNoContent)
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()
		c.Next()
	}
}
