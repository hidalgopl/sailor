package metrics

import (
	"os"
	"strings"
)

type OriginFinder struct {
	DetectedEnvVars string
	KnownOrigins map[RunnerEnv]string

}

func NewOriginFinder() *OriginFinder {
	var detectedEnvs []string
	for _, key := range os.Environ() {
		variable := strings.Split(key,"=")
		detectedEnvs = append(detectedEnvs, variable[0])
	}
	return &OriginFinder{
		strings.Join(detectedEnvs, ","),
		KnownOrigins,
	}
}

func (of *OriginFinder) DetectOrigin() (RunnerEnv) {
	for runner, vars := range of.KnownOrigins {
		if strings.Contains(of.DetectedEnvVars, vars) {
			if runner == GitHubAction {
				isSailorAction := strings.Contains(of.DetectedEnvVars, of.KnownOrigins[SailorAction])
				if isSailorAction {
					return SailorAction
				}
			}
			return runner
		}

	}
	return LocalOrNotFound

}
