package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/persistence"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/message"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/middleware"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/util"
)

var (
	localizationRouter *gin.RouterGroup
)

func initLocalizeService() {
	localizationRouter = http.RegisterServices("/localization")

	localizationRouter.Use(
		append(
			http.EngineMiddlewares,
			middleware.NewQPSLimiter(middlewareConfig),
		)...,
	)
	localizationRouter.GET(
		"/language_pack/:id",
		queryLanguagePackServiceHandler(),
	)
	localizationRouter.GET(
		"/translate",
		translateServiceServiceHandler(),
	)
}

func queryLanguagePackServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if exist, result := util.QueryPath(ctx, ":id"); !exist {
			ctx.AbortWithStatus(400)
		} else if has, record := persistence.LocalizationPersistence.QueryByID(result); !has {
			ctx.AbortWithStatus(404)
		} else {
			response := message.LocalizationQueryResponse{LanguagePack: record.Pack()}
			ctx.JSON(200, response)
		}
	}
}

func translateServiceServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var (
			exist        bool
			languagePack string
			destLanguage int
			request      message.TranslationRequest
		)
		if exist, languagePack = util.QueryPath(ctx, "language_package"); !exist {
			ctx.AbortWithStatus(400)
		} else if exist, destLanguage = util.QueryPathInt(ctx, "target_language"); !exist {
			ctx.AbortWithStatus(400)
		} else if !util.BindJson(ctx, &request) {
			ctx.AbortWithStatus(400)
		} else {
			if has, dictionary := persistence.LocalizationPersistence.QueryByID(languagePack); !has {
				ctx.AbortWithStatus(404)
			} else {
				response := message.TranslationResponse{Translation: map[string]string{}}
				language := enum.Language(destLanguage)
				for _, word := range request.Words {
					if ok, result := dictionary.Translate(word, language); ok {
						response.Translation[word] = result
					}
				}
				ctx.JSON(200, response)
			}
		}
	}
}
