package handler

import (
	"github.com/gin-gonic/gin"
	oauth2V1 "github.com/versegeek/go-generated/gen/rest/oauth2/v1"
)

// Implement the Client CRUD handler.

func (h *handler) GetOauth2V1Auth(c *gin.Context, params oauth2V1.GetOauth2V1AuthParams) {}

func (h *handler) PostOauth2V1ExtensionsBlacklist(c *gin.Context) {}

func (h *handler) PostOauth2V1Introspect(c *gin.Context, params oauth2V1.PostOauth2V1IntrospectParams) {
}

func (h *handler) PostOauth2V1Revoke(c *gin.Context, params oauth2V1.PostOauth2V1RevokeParams) {}

func (h *handler) PostOauth2V1Token(c *gin.Context, params oauth2V1.PostOauth2V1TokenParams) {}

func (h *handler) PostOauth2V1Userinfo(c *gin.Context, params oauth2V1.PostOauth2V1UserinfoParams) {
}

func (h *handler) GetOauth2V1TenantWellKnownOpenidConfiguration(c *gin.Context, tenant string) {}
