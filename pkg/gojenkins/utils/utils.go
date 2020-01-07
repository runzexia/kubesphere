package utils

import (
	"github.com/asaskevich/govalidator"
	"kubesphere.io/kubesphere/pkg/simple/client/devops/jenkins"
	"net/http"
	"strconv"
)

func GetJenkinsStatusCode(jenkinsErr error) int {
	if code, err := strconv.Atoi(jenkinsErr.Error()); err == nil {
		message := http.StatusText(code)
		if !govalidator.IsNull(message) {
			return code
		}
	}
	if jErr, ok := jenkinsErr.(*jenkins.ErrorResponse); ok {
		return jErr.Response.StatusCode
	}
	return http.StatusInternalServerError
}
