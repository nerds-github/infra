package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devbookhq/orchestration-services/api/internal/api"
	"github.com/gin-gonic/gin"
)

func (a *APIStore) PostEnvsCodeSnippetID(
	c *gin.Context,
	codeSnippetID string,
	params api.PostEnvsCodeSnippetIDParams,
) {
	_, keyErr := a.validateAPIKey(params.ApiKey)
	if keyErr != nil {
		fmt.Printf("error with API key: %+v", keyErr)
		sendAPIStoreError(c, http.StatusUnauthorized, "Error with API token")
		return
	}

	var env api.PostEnvsCodeSnippetIDJSONBody
	if err := c.Bind(&env); err != nil {
		sendAPIStoreError(c, http.StatusBadRequest, fmt.Sprintf("Error when parsing request: %s", err))
		return
	}

	// TODO: Download the base Dockerfile based on a template field in `env`.
	// TODO: Add deps to the Dockerfile.
	err := a.nomadClient.CreateEnv(codeSnippetID, string(env.Template), env.Deps)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		sendAPIStoreError(
			c,
			http.StatusInternalServerError,
			fmt.Sprintf("Failed to delete env for code snippet '%s': %s", codeSnippetID, err),
		)
		return
	}

	c.Status(http.StatusNoContent)
}

func (a *APIStore) DeleteEnvsCodeSnippetID(
	c *gin.Context,
	codeSnippetID string,
	params api.DeleteEnvsCodeSnippetIDParams,
) {
	_, keyErr := a.validateAPIKey(params.ApiKey)
	if keyErr != nil {
		fmt.Printf("error with API key: %+v", keyErr)
		sendAPIStoreError(c, http.StatusUnauthorized, "Error with API token")
		return
	}

	err := a.supabase.DB.
		From("envs").
		Delete().
		Eq("code_snippet_id", codeSnippetID).
		Execute(nil)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		sendAPIStoreError(
			c,
			http.StatusBadRequest,
			fmt.Sprintf("Failed to delete env for code snippet '%s': %s", codeSnippetID, err),
		)
		return
	}

	err = a.nomadClient.DeleteEnv(codeSnippetID)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		fmt.Printf("error: %v\n", err)
		sendAPIStoreError(
			c,
			http.StatusInternalServerError,
			fmt.Sprintf("Failed to delete env for code snippet '%s': %s", codeSnippetID, err),
		)
		return
	}

	c.Status(http.StatusNoContent)
}

func (a *APIStore) PutEnvsCodeSnippetIDState(
	c *gin.Context,
	codeSnippetID string,
	params api.PutEnvsCodeSnippetIDStateParams,
) {
	_, keyErr := a.validateAPIKey(params.ApiKey)
	if keyErr != nil {
		sendAPIStoreError(c, http.StatusUnauthorized, fmt.Sprintf("Error with API token: %s", keyErr))
		return
	}

	var envStateUpdate api.PutEnvsCodeSnippetIDStateJSONBody
	if err := c.Bind(&envStateUpdate); err != nil {
		sendAPIStoreError(
			c,
			http.StatusBadRequest,
			fmt.Sprintf("Error when parsing request: %s", err),
		)
		return
	}

	body := map[string]interface{}{"state": envStateUpdate.State}
	err := a.supabase.DB.
		From("envs").
		Update(body).
		Eq("code_snippet_id", codeSnippetID).
		Execute(nil)

	if err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			fmt.Printf("syntax error at byte offset %d", e.Offset)
		}
		fmt.Printf("error: %v\n", err)
		sendAPIStoreError(
			c,
			http.StatusBadRequest,
			fmt.Sprintf("Failed to update env for code snippet '%s': %s", codeSnippetID, err),
		)
		return
	}

	c.Status(http.StatusNoContent)
}

func (a *APIStore) PatchEnvsCodeSnippetID(
	c *gin.Context,
	codeSnippetID string,
	params api.PatchEnvsCodeSnippetIDParams,
) {
	_, keyErr := a.validateAPIKey(params.ApiKey)
	if keyErr != nil {
		fmt.Printf("error with API key: %+v", keyErr)
		sendAPIStoreError(c, http.StatusUnauthorized, "Error with API token")
		return
	}

	session, err := a.sessionsCache.FindEditSession(codeSnippetID)
	if err != nil {
		fmt.Printf("cannot find active edit session for the code snippet '%s': %v - will use saved rootfs", codeSnippetID, err)
	}

	err = a.nomadClient.UpdateEnv(codeSnippetID, session)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		sendAPIStoreError(
			c,
			http.StatusInternalServerError,
			fmt.Sprintf("Failed to delete env for code snippet '%s': %+v", codeSnippetID, err),
		)
		return
	}

	c.Status(http.StatusNoContent)
}
