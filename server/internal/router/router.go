// Copyright © 2023 OpenIM open source community. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package router

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// swagger docs.
	_ "github.com/OpenIMSDK/OpenKF/server/docs"

	"github.com/gin-gonic/gin"

	"github.com/OpenIMSDK/OpenKF/server/internal/api"
	"github.com/OpenIMSDK/OpenKF/server/internal/config"
	"github.com/OpenIMSDK/OpenKF/server/internal/middleware"
	urltrie "github.com/OpenIMSDK/OpenKF/server/internal/middleware/hooks/url_trie"
)

// InitRouter init router.
func InitRouter() *gin.Engine {
	if config.GetString("app.debug") == "true" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(urltrie.RunHook(), middleware.EnableCROS())

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	{
		// User register api
		register := apiv1.Group("/register")
		{
			// Register with email
			register.POST("/email/code", api.SendCode)
			register.POST("/admin", api.AdminRegister)
			register.POST("/staff", api.StaffRegister)

			// register.POST("/github", api.GithubRegister)
		}

		// User login api
		login := apiv1.Group("/login")
		{
			login.POST("/account", api.AccountLogin)
			// user.POST("/email", api.GithubRegister)
			// user.POST("/github", api.GithubRegister)
			// user.POST("/wechat", api.GithubRegister)
		}

		// admin := apiv1.Group("/admin")
		// user := apiv1.Group("/user")
		community := apiv1.Group("/community")
		{
			community.POST("/create", api.CreateCommunity)
		}

		// OpenIM callback api
		command := apiv1.Group("/openim/callback")
		{
			command.POST("/", api.OpenIMCallback)
			command.POST("/callbackBeforeSendSingleMsgCommand", api.BeforeSendSingleMsg)
			command.POST("/callbackAfterSendSingleMsgCommand", api.AfterSendSingleMsg)
			command.POST("/callbackMsgModifyCommand", api.MsgModify)
			command.POST("/callbackUserOnlineCommand", api.UserOnline)
			command.POST("/callbackUserOfflineCommand", api.UserOffline)
			command.POST("/callbackOfflinePushCommand", api.OfflinePush)
			command.POST("/callbackOnlinePushCommand", api.OnlinePush)
		}
	}

	return r
}
