// package api

// import (
// 	"database/sql"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	db "github.com/it21152832/Learning-Backend/db/sqlc"
// 	"github.com/lib/pq"
// )

// type createUserRequest struct {
// 	FirstName string `json:"first_name" binding:"required"`
// 	LastName  string `json:"last_name" binding:"required"`
// 	Email     string `json:"email" binding:"required"`
// 	HashedPassword  string `json:"hashed_password" binding:"required"`
// 	Role      string `json:"role" binding:"required"`
// }

// func (server *Server) createUser(ctx *gin.Context) {
// 	var req createUserRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	arg := db.CreateUserParams{
// 		FirstName: req.FirstName,
// 		LastName:req.LastName,
// 		Email:req.Email,
// 		HashedPassword:req.HashedPassword,
// 		Role:req.Role,

// 	}

// 	user, err := server.store.CreateUser(ctx, arg)
// 	if err != nil {
// 		if pqErr, ok := err.(*pq.Error); ok{
// 			switch pqErr.Code.Name(){
// 			case "foreign_key_violation","unique_violation":
// 				ctx.JSON(http.StatusForbidden, errorResponse(err))
// 				return
// 			}
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, user)
// }

// type getUserRequest struct {
// 	UserID int64 `uri:"user_id" binding:"required,min=1"`
// }

// func (server *Server) getUser(ctx *gin.Context) {
// 	var req getUserRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	user, err := server.store.GetUser(ctx, req.UserID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, user)
// }

// package api

// import (
// 	"database/sql"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	db "github.com/it21152832/Learning-Backend/db/sqlc"
// 	"github.com/it21152832/Learning-Backend/util"
// 	"github.com/lib/pq"
// )

// type createUserRequest struct {
// 	FirstName string `json:"first_name" binding:"required"`
// 	LastName  string `json:"last_name" binding:"required"`
// 	Email     string `json:"email" binding:"required,email"`
// 	HashedPassword  string `json:"hashed_password" binding:"required"`
// 	Role      string `json:"role" binding:"required"`
// 	Username      string `json:"username" binding:"required"`

// }

// type userResponse struct {
// 	Username          string    `json:"username" `
// 	FirstName          string    `json:"first_name" `
// 	Email             string    `json:"email" `
// 	PasswordChangedAt time.Time `json:"password_changed_at"`
// 	CreatedAt         time.Time `json:"created_at"`
// }

// // type createAccountRequest struct {
// // 	Owner    string `json:"owner" binding:"required"`
// // 	Currency string `json:"currency" binding:"required,oneof=USD EUR"`

// // }

// // type accountResponse struct {
// // 	Owner    string `json:"owner" binding:"required"`
// // 	Currency string `json:"currency" binding:"required,oneof=USD EUR"`

// // }

// func newUserResponse(user db.User) userResponse{
// 	return userResponse{
// 		Username: user.Username,
// 		Email:             user.Email,
// 	}
// }

// func (server *Server) createUser(ctx *gin.Context) {
// 	var req createUserRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	arg := db.CreateUserParams{
// 		Username:       req.Username,
// 		Email:          req.Email,
// 	}

// 	User, err := server.store.CreateUser(ctx, arg)
// 	if err != nil {
// 		if pqErr, ok := err.(*pq.Error); ok {
// 			switch pqErr.Code.Name() {
// 			case "foreign_key_violation", "unique_violation":
// 				ctx.JSON(http.StatusForbidden, errorResponse(err))
// 				return
// 			}
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	rsp := newUserResponse(User)
// 	ctx.JSON(http.StatusOK, rsp)
// }

// // func newAccountResponse(account db.Account) accountResponse{
// // 	return accountResponse{

// // 		Owner:          account.Owner,
// // 		Currency: account.Currency,
// // 	}
// // }

// // func (server *Server) createAccount(ctx *gin.Context) {
// // 	var req createAccountRequest
// // 	if err := ctx.ShouldBindJSON(&req); err != nil {
// // 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// // 		return
// // 	}

// // 	arg := db.CreateAccountParams{
// // 		Owner:       req.Owner,
// // 		Currency: req.Currency,
// // 	}

// // 	Account, err := server.store.CreateAccount(ctx, arg)
// // 	if err != nil {
// // 		if pqErr, ok := err.(*pq.Error); ok {
// // 			switch pqErr.Code.Name() {
// // 			case "foreign_key_violation", "unique_violation":
// // 				ctx.JSON(http.StatusForbidden, errorResponse(err))
// // 				return
// // 			}
// // 		}
// // 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// // 		return
// // 	}
// // 	rsp := newAccountResponse(Account)
// // 	ctx.JSON(http.StatusOK, rsp)
// // }

// type loginUserRequest struct {
// 	Username string `json:"username" binding:"required,alphanum"`
// 	Password string `json:"password" binding:"required,min=6"`
// }

// type loginUserResponse struct {
// 	AccessToken string       `json:"access_token"`
// 	User        userResponse `json:"user"`
// 	// Account		accountResponse `json:"account"`
// }

// func (server *Server) loginUser(ctx *gin.Context) {
// 	var req loginUserRequest

// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	user, err := server.store.GetUser(ctx, req.Username)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	err = util.CheckPassword(req.Password, user.HashedPassword)
// 	if err != nil {
// 		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
// 		return
// 	}

// 	accessToken, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	rsp := loginUserResponse{
// 		AccessToken: accessToken,
// 		User:        newUserResponse(user),
// 		// Account: newAccountResponse(account),
// 	}
// 	ctx.JSON(http.StatusOK, rsp)
// }

// package api

// import (
// 	"database/sql"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	db "github.com/it21152832/Learning-Backend/db/sqlc"
// 	"github.com/it21152832/Learning-Backend/util"
// 	"github.com/lib/pq"
// )

// type createUserRequest struct {
// 	FirstName       string `json:"first_name" binding:"required"`
// 	LastName        string `json:"last_name" binding:"required"`
// 	Email           string `json:"email" binding:"required,email"`
// 	HashedPassword  string `json:"hashed_password" binding:"required"`
// 	Role            string `json:"role" binding:"required"`
// 	Username        string `json:"username" binding:"required"`
// }

// type userResponse struct {
// 	FirstName       string `json:"first_name"`
// 	LastName        string `json:"last_name"`
// 	Email           string `json:"email"`
// 	HashedPassword  string `json:"hashed_password"`
// 	Role            string `json:"role"`
// 	Username        string `json:"username"`
// }

// func newUserResponse(user db.User) userResponse {
// 	return userResponse{
// 		Username:          user.Username,
// 		Email:             user.Email,
// 	}
// }

// func (server *Server) createUser(ctx *gin.Context) {
// 	var req createUserRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	arg := db.CreateUserParams{
// 		Username: req.Username,
// 		Email:    req.Email,
// 		// Add other fields as needed
// 	}

// 	user, err := server.store.CreateUser(ctx, arg)
// 	if err != nil {
// 		if pqErr, ok := err.(*pq.Error); ok {
// 			switch pqErr.Code.Name() {
// 			case "foreign_key_violation", "unique_violation":
// 				ctx.JSON(http.StatusForbidden, errorResponse(err))
// 				return
// 			}
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	rsp := newUserResponse(user)
// 	ctx.JSON(http.StatusOK, rsp)
// }

// type loginUserRequest struct {
// 	Username string `json:"username" binding:"required,alphanum"`
// 	Password string `json:"password" binding:"required,min=6"`
// }

// type loginUserResponse struct {
// 	AccessToken string       `json:"access_token"`
// 	User        userResponse `json:"user"`
// 	// Add other fields like Account if needed
// }

// func (server *Server) loginUser(ctx *gin.Context) {
// 	var req loginUserRequest

// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	user, err := server.store.GetUser(ctx, req.Username)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	err = util.CheckPassword(req.Password, user.HashedPassword)
// 	if err != nil {
// 		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
// 		return
// 	}

// 	accessToken, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

//		rsp := loginUserResponse{
//			AccessToken: accessToken,
//			User:        newUserResponse(user),
//		}
//		ctx.JSON(http.StatusOK, rsp)
//	}
package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/it21152832/Learning-Backend/db/sqlc"
	"github.com/it21152832/Learning-Backend/util"
	"github.com/lib/pq"
)

type createUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	Role      string `json:"role" binding:"required"`
	Username  string `json:"username" binding:"required"`
}

type userResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Username  string `json:"username"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{

		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.Role,
		Username:  user.Username,
	}
}


// create user
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.CreateUserParams{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		HashedPassword: hashedPassword,
		Role:           req.Role,
		Username:       req.Username,
	}

	User, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newUserResponse(User)
	ctx.JSON(http.StatusOK, rsp)
}


type loginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string         `json:"access_token"`
	User        userResponse   `json:"user"`
}

// login user
func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(user.Email, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)
}
