package metrics

type RunnerEnv string

const (

	GitHubAction = RunnerEnv("github_action")
	BitBucketPipeline = RunnerEnv("bitbucket_pipeline")
	JenkinsSlave = RunnerEnv("jenkins_slave")
	SailorAction = RunnerEnv("sailor_action")
	GitLabRunner = RunnerEnv("gitlab_runner")
	TravisJob = RunnerEnv("travis_job")
	LocalOrNotFound = RunnerEnv("local_or_other")
	BitBucketEnvVars = "BITBUCKET_BUILD_NUMBER"
	GitHubEnvVars = "GITHUB_RUN_ID"
	JenkinsEnvVars = "JENKINS_URL"
	SailorActionEnvVars = "SAILOR_GH_ACTION"
	GitLabEnvVars = "GITLAB_CI"
	TravisEnvVars = "TRAVIS_JOB_ID"

)

var KnownOrigins = map[RunnerEnv]string{
	SailorAction: SailorActionEnvVars,
	GitLabRunner: GitLabEnvVars,
	JenkinsSlave: JenkinsEnvVars,
	TravisJob: TravisEnvVars,
	GitHubAction: GitHubEnvVars,
	BitBucketPipeline: BitBucketEnvVars,

}

