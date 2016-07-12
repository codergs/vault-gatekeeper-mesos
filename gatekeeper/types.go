package gatekeeper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type VaultError struct {
	Code   int      `json:"-"`
	Errors []string `json:"errors"`
}

func (e VaultError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, strings.Join(e.Errors, ", "))
}

func buildVaultError(resp *http.Response) error {
	if resp.StatusCode == 200 {
		return nil
	}

	vaultErr := &VaultError{Code: resp.StatusCode}
	if err := json.NewDecoder(resp.Body).Decode(vaultErr); err != nil {
		vaultErr.Errors = []string{"communication error", err.Error()}
	}

	return vaultErr
}

type gkTokenReq struct {
	TaskId string `json:"task_id"`
}

type gkTokenResp struct {
	OK     bool   `json:"ok"`
	Token  string `json:"token"`
	Status string `json:"status"`
	Error  string `json:"error"`
}

type cubbyholeSecret struct {
	Data wrappedResponseData `json:"data"`
}

type wrappedResponseData struct {
	WrappedSecret vaultSecretData `json:"response"`
}

type vaultSecretData struct {
	Token string `json:"token"`
}
