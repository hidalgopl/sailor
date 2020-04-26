package metrics

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

)

type testCase struct {
	name           string
	of             OriginFinder
	expectedResult RunnerEnv
}


func TestNewOriginFinder(t *testing.T) {
	testCases := []struct{
		name string
		envVars map[string]string
		expectedEnvVars string
	}{
		{
			"happy path",
			map[string]string{
				"FIRST_ENVVAR": "NOTIMPROTANT",
				"SECOND_ENVVAR": "NOTIMPORTANT",
			},
			"FIRST_ENVVAR,SECOND_ENVVAR",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			os.Clearenv()
			for k, v := range tc.envVars {
				os.Setenv(k, v)
			}
			defer os.Clearenv()
			of := NewOriginFinder()
			assert.Equal(t, tc.expectedEnvVars, of.DetectedEnvVars)
		})
	}
}


func TestDetectOrigin(t *testing.T) {
	testCases := []testCase{
		{
			"check sailor action",
			OriginFinder{

				"SAILOR_GH_ACTION,GITHUB_RUN_ID",
				KnownOrigins,
			},
			SailorAction,
		},
		{
			"check bitbucket pipeline",
			OriginFinder{

				"BITBUCKET_BUILD_NUMBER,RANDOM_OTHER_STUFF",
				KnownOrigins,
			},
			BitBucketPipeline,
		},
		{
			"check github action",
			OriginFinder{

				"GITHUB_RUN_ID",
				KnownOrigins,
			},
			GitHubAction,
		},
		{
			"check jenkins slave",
			OriginFinder{

				"JENKINS_URL",
				KnownOrigins,
			},
			JenkinsSlave,
		},
		{
			"check gitlab runner",
			OriginFinder{

				"GITLAB_CI",
				KnownOrigins,
			},
			GitLabRunner,
		},
		{
			"check travis job",
			OriginFinder{

				"TRAVIS_JOB_ID",
				KnownOrigins,
			},
			TravisJob,
		},
		{
			"check none of these found",
			OriginFinder{

				"",
				KnownOrigins,
			},
			LocalOrNotFound,
		},
		{
			"check none of these found",
			OriginFinder{

				"USER,SYSTEM,RANDOM,BOLLOCKS",
				KnownOrigins,
			},
			LocalOrNotFound,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			origin := tc.of.DetectOrigin()
			assert.Equal(t, tc.expectedResult, origin)
		})

	}
}
