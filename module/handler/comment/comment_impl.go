package comment

import (
	"errors"
	"log"
	"net/http"

	"github.com/Retrospective53/myGram/module/models"
	commentcreatemodel "github.com/Retrospective53/myGram/module/models/comment"
	"github.com/Retrospective53/myGram/module/models/token"
	commentservice "github.com/Retrospective53/myGram/module/service/comment"
	"github.com/Retrospective53/myGram/pkg/json"
	"github.com/Retrospective53/myGram/pkg/middleware"
	"github.com/Retrospective53/myGram/pkg/response"
	"github.com/gin-gonic/gin"
)

type CommentHandlerImpl struct {
	commentService commentservice.CommentService
}

func NewCommentHandlerImpl(commentService commentservice.CommentService) CommentHandler {
	return &CommentHandlerImpl{
		commentService: commentService,
	}
}

func (c *CommentHandlerImpl) FindAllCommentsHdl(ctx *gin.Context) {
	comments, err := c.commentService.FindAllCommentsSvc(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InternalServer,
			Error:   "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success",
		Data:    comments,
	})
}

func (c *CommentHandlerImpl) FindCommentByIdHdl(ctx *gin.Context) {
	commentId := ctx.Param("id")

	comment, err := c.commentService.FindCommentByIdSvc(ctx, commentId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InternalServer,
			Error:   "something went wrong",
		})
		return
	}

	if comment.Message == "" || comment.ID.String() == "" {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InvalidParam,
			Error:   "photo not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success",
		Data:    comment,
	})
}

func (c *CommentHandlerImpl) CreateCommentHdl(ctx *gin.Context) {
	// get user_id from context first
	accessClaimI, ok := ctx.Get(middleware.AccessClaim.String())
	if !ok {
		err := errors.New("error get claim from context")
		if err != nil {
			panic(err)
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: response.InvalidPayload,
			Error:   "invalid user id",
		})
		return
	}

	var accessClaim token.AccessClaim
	if err := json.ObjectMapper(accessClaimI, &accessClaim); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: response.InvalidPayload,
			Error:   "invalid payload",
		})
		return
	}

		// binding payload
		var createComment commentcreatemodel.CommentCreate
		createComment.UserID = accessClaim.UserID
		log.Printf("%s data type is: %T", accessClaim.UserID, accessClaim.UserID)
		if err := ctx.BindJSON(&createComment); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				response.ErrorResponse{
					Message: response.InvalidBody,
					Error:   "error binding payload",
				},
			)
			return
		}

	comment, err := c.commentService.CreateCommentSvc(ctx, createComment, accessClaim.UserID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InternalServer,
			Error:   "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success",
		Data:    comment,
	})
}

func (c *CommentHandlerImpl) UpdateCommentHdl(ctx *gin.Context) {
	commentId := ctx.Param("id")
	
	// binding payload
	var updateComment models.Comment
	if err := ctx.BindJSON(&updateComment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			response.ErrorResponse{
				Message: response.InvalidBody,
				Error:   "error binding payload",
			},
		)
		return
	}

	// get user_id from context first
	accessClaimI, ok := ctx.Get(middleware.AccessClaim.String())
	if !ok {
		err := errors.New("error get claim from context")
		if err != nil {
			panic(err)
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: response.InvalidPayload,
			Error:   "invalid user id",
		})
		return
	}

	var accessClaim token.AccessClaim
	if err := json.ObjectMapper(accessClaimI, &accessClaim); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: response.InvalidPayload,
			Error:   "invalid payload",
		})
		return
	}


	// authorization only admin
	if accessClaim.Role != "ROLE_ADMIN" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
			Message: response.Unauthorized,
			Error: "Unauthorized",
		})
		return
	}

	// //get basicauth username and password
	// // Retrieve the value and check if the key exists
	// basicClaimRaw, ok := ctx.Get("userBasic")

	// // Check if the key exists
	// if !ok {
	// 	// Handle the case when the key does not exist
	// 	err := errors.New("error getting basic auth claim from context")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
	// 		Message: response.InvalidPayload,
	// 		Error:   "invalid username and password",
	// 	})
	// 	return
	// }

	// // Convert the value to the expected type
	// basicClaim, ok := basicClaimRaw.(map[string]string)

	// // Check if the conversion was successful
	// if !ok {
	// 	// Handle the case when the value is not of the expected type
	// 	err := errors.New("error converting basic auth claim to map[string]string")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
	// 		Message: response.InvalidPayload,
	// 		Error:   "invalid username and password",
	// 	})
	// 	return
	// }

	// // Now you can access the values in the basicClaim map
	// if basicClaim["username"] != "admin" || basicClaim["password"] != "admin" {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
	// 		Message: response.InvalidPayload,
	// 		Error:   "invalid username and password",
	// 	})
	// 	return
	// }

	comment, err := c.commentService.UpdateCommentSvc(ctx, updateComment, commentId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InternalServer,
			Error:   "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success",
		Data:    comment,
	})
}

func (c *CommentHandlerImpl) DeleteCommentByIdHdl(ctx *gin.Context) {
	commentId := ctx.Param("id")

	// get user_id from context first
	accessClaimI, ok := ctx.Get(middleware.AccessClaim.String())
	if !ok {
		err := errors.New("error get claim from context")
		if err != nil {
			panic(err)
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: response.InvalidPayload,
			Error:   "invalid user id",
		})
		return
	}

	var accessClaim token.AccessClaim
	if err := json.ObjectMapper(accessClaimI, &accessClaim); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: response.InvalidPayload,
			Error:   "invalid payload",
		})
		return
	}

	// authorization only admin
	if accessClaim.Role != "ROLE_ADMIN" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
			Message: response.Unauthorized,
			Error: "Unauthorized",
		})
		return
	}

	_, err := c.commentService.DeleteCommentByIdSvc(ctx, commentId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InternalServer,
			Error:   "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success",
		Data:    "comment deleted",
	})
}
