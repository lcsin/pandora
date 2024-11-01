package handler

// import (
// 	"errors"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v5"
// 	"github.com/lcsin/pandora/api"
// 	"github.com/lcsin/pandora/api/message"
// 	"github.com/lcsin/pandora/internal/domain"
// 	"github.com/lcsin/pandora/internal/service"
// 	"github.com/spf13/viper"
// )

// // UserHandler user handler
// type UserHandler struct {
// 	userSrv service.IUserService
// }

// // NewUserHandler user handler 构造器
// //
// //	@param userSrv
// //	@return *UserHandler
// func NewUserHandler(userSrv service.IUserService) *UserHandler {
// 	return &UserHandler{userSrv: userSrv}
// }

// // RegisterRoutes  注册用户服务路由
// //
// //	@receiver uh
// //	@param v1
// func (uh *UserHandler) RegisterRoutes(v1 *gin.RouterGroup) {
// 	ug := v1.Group("/users")

// 	ug.POST("/login", uh.Login)
// 	ug.POST("/register", uh.Register)
// 	ug.POST("/info", uh.Info)
// }

// // Info 获取用户信息
// //
// //	@receiver uh
// //	@param c
// func (uh *UserHandler) Info(c *gin.Context) {
// 	uid := c.GetInt64("uid")
// 	user, err := uh.userSrv.GetByID(c, uid)
// 	if err != nil {
// 		if errors.Is(err, service.ErrUserNotFound) {
// 			api.ResponseErrorMessage(c, message.UserNotFound, "用户不存在")
// 			return
// 		}
// 		api.ResponseError(c, message.Failed)
// 		return
// 	}

// 	api.ResponseOK(c, user)
// }

// // Login 用户登录
// //
// //	@receiver uh
// //	@param c
// func (uh *UserHandler) Login(c *gin.Context) {
// 	type LoginReq struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}

// 	var req LoginReq
// 	if err := c.ShouldBind(&req); err != nil {
// 		api.ResponseError(c, message.BadRequest)
// 		return
// 	}
// 	if len(req.Email) == 0 || len(req.Password) == 0 {
// 		api.ResponseErrorMessage(c, message.BadRequest, "用户名或密码不能为空")
// 		return
// 	}

// 	// 登录
// 	user, err := uh.userSrv.Login(c, req.Email, req.Password)
// 	if err != nil {
// 		if errors.Is(err, service.ErrUserNotFound) {
// 			api.ResponseError(c, message.UserNotFound)
// 			return
// 		}
// 		api.ResponseError(c, message.Failed)
// 		return
// 	}

// 	// 设置JWT
// 	tokenInfo, err := uh.setJWTToken(c, user.ID)
// 	if err != nil {
// 		api.ResponseError(c, message.Failed)
// 		return
// 	}

// 	api.ResponseOK(c, tokenInfo)
// }

// // Register 用户注册
// //
// //	@receiver uh
// //	@param c
// func (uh *UserHandler) Register(c *gin.Context) {
// 	type RegiserReq struct {
// 		Email           string `json:"email"`
// 		Password        string `json:"password"`
// 		ConfirmPassword string `json:"confirmPassword"`
// 	}

// 	var req RegiserReq
// 	if err := c.ShouldBind(&req); err != nil {
// 		api.ResponseError(c, message.BadRequest)
// 		return
// 	}

// 	if len(req.Email) == 0 || len(req.Password) == 0 || len(req.ConfirmPassword) == 0 {
// 		api.ResponseError(c, message.BadRequest)
// 		return
// 	}

// 	if req.Password != req.ConfirmPassword {
// 		api.ResponseErrorMessage(c, message.BadRequest, "密码不一致")
// 		return
// 	}

// 	if err := uh.userSrv.Regiser(c, domain.User{
// 		Email:    req.Email,
// 		Password: req.Password,
// 	}); err != nil {
// 		if errors.Is(err, service.ErrUserExisted) {
// 			api.ResponseError(c, message.UserExisted)
// 			return
// 		}
// 		api.ResponseError(c, message.Failed)
// 		return
// 	}

// 	api.ResponseOK(c, nil)
// }

// // 设置JWT Token
// func (uh *UserHandler) setJWTToken(c *gin.Context, uid int64) (*domain.TokenInfo, error) {
// 	accessTokenExpire := time.Now().Add(time.Hour * 24 * 7)
// 	claims := domain.Claims{
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(accessTokenExpire),
// 		},
// 		UID: uid,
// 	}

// 	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(viper.GetString("jwt_secret")))
// 	if err != nil {
// 		api.ResponseError(c, message.Failed)
// 		return nil, err
// 	}
// 	// 请求头添加token
// 	c.Header("Authorization", accessToken)

// 	return &domain.TokenInfo{
// 		AccessToken:       accessToken,
// 		AccessTokenExpire: accessTokenExpire.Unix(),
// 	}, nil
// }
