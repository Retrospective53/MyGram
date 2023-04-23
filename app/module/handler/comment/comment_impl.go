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


// @Tags Comment
// @Summary finds all comment records
// @Schemes http
// @Description fetch all comment records
// @Param Authorization header string true "Bearer Token"
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=[]string}
// @Success 401 {object} response.SuccessResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /comments/all [get]
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


// @Tags Comment
// @Summary Find a comment by ID
// @Schemes http
// @Description Fetch a comment with the given id
// @Accept json
// @Param id path string true "Comment ID"
// @Param Authorization header string true "Bearer Token"
// @Produce json
// @Success 200 {object} response.SuccessResponse{}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 401 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /comments/{id} [get]
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


// @Tags Comment
// @Summary Create a new comment
// @Schemes http
// @Description Creates a new comment with the provided data
// @Accept json
// @Param body body commentcreatemodel.CommentCreate true "Create Comment Request Body"
// @Param Authorization header string true "Bearer Token"
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=object}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /comments [post]
func (c *CommentHandlerImpl) CreateCommentHdl(ctx *gin.Context) {
	// get user_id from context first
	accessClaimI, ok := ctx.Get(middleware.AccessClaim.String())
	if !ok {
		err := errors.New("error get claim from context")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
				Message: response.SomethingWentWrong,
				Error:   err.Error(),
			})
			return
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


// @Tags Comment
// @Summary Update an existing photo by id
// @Schemes http
// @Description Updates an existing photo with the provided data
// @Accept json
// @Param id path string true "Photo ID"
// @Param Authorization header string true "Bearer Token"
// @Param request body models.Comment true "Create Comment Request Body"
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=object}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 401 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /comments/{id} [put]
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
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
				Message: response.SomethingWentWrong,
				Error:   err.Error(),
			})
			return
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

	comment, err := c.commentService.UpdateCommentSvc(ctx, updateComment, commentId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InternalServer,
			Error:   "something went wrong",
		})
		return
	}

	if comment.Message == "" {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InvalidParam,
			Error:   "comment not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success",
		Data:    comment,
	})
}

// @Tags Comment
// @Summary Delete a comment by ID
// @Schemes http
// @Description Deletes a comment with the given id
// @Accept json
// @Param id path string true "Comment ID"
// @Param Authorization header string true "Bearer Token"
// @Produce json
// @Success 200 {object} response.SuccessResponse{}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 401 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /comments/{id} [delete]
func (c *CommentHandlerImpl) DeleteCommentByIdHdl(ctx *gin.Context) {
	commentId := ctx.Param("id")

	// get user_id from context first
	accessClaimI, ok := ctx.Get(middleware.AccessClaim.String())
	if !ok {
		err := errors.New("error get claim from context")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
				Message: response.SomethingWentWrong,
				Error:   err.Error(),
			})
			return
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
