package db_lib

import (
	"os"
	"strings"

	"github.com/ansible-semaphore/semaphore/pkg/task_logger"
)

func isSensitiveVar(v string) bool {
	sensitives := []string{
		"SEMAPHORE_ACCESS_KEY_ENCRYPTION",
		"SEMAPHORE_ADMIN_PASSWORD",
		"SEMAPHORE_DB_USER",
		"SEMAPHORE_DB_NAME",
		"SEMAPHORE_DB_HOST",
		"SEMAPHORE_DB_PASS",
		"SEMAPHORE_LDAP_PASSWORD",
		"SEMAPHORE_RUNNER_TOKEN",
		"SEMAPHORE_RUNNER_ID",
	}

	for _, s := range sensitives {
		if strings.HasPrefix(v, s+"=") {
			return true
		}
	}

	return false
}

func removeSensitiveEnvs(envs []string) (res []string) {

	for _, e := range envs {
		if !isSensitiveVar(e) {
			res = append(res, e)
		}
	}

	return res
}

type LocalApp interface {
	SetLogger(logger task_logger.Logger) task_logger.Logger
	InstallRequirements(environmentVars *[]string) error
	Run(args []string, environmentVars *[]string, inputs map[string]string, cb func(*os.Process)) error
}
