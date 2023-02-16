package constant

type EnvironmentType string

var (
	DEV    = EnvironmentType("dev")
	DOCKER = EnvironmentType("docker")
)
