package conf

import "github.com/erda-project/erda/pkg/envconf"

// Conf action 入参
type Conf struct {
	ActionBranch          string `env:"ACTION_BRANCH" required:"true"`
	ActionApplicationName string `env:"ACTION_APPLICATION_NAME" required:"true"`
	ActionPipelineYmlName string `env:"ACTION_PIPELINE_YML_NAME" required:"true"`

	// env
	MetaFile             string `env:"METAFILE"`
	WorkDir              string `env:"WORKDIR" default:"."`
	DiceVersion          string `env:"DICE_VERSION"`
	PipelineId           uint64 `env:"PIPELINE_ID"`
	OrgId                uint64 `env:"DICE_ORG_ID"`
	TaskId               uint64 `env:"PIPELINE_TASK_ID"`
	ProjectId            uint64 `env:"DICE_PROJECT_ID"`
	PipelineID           uint64 `env:"PIPELINE_ID"`
	SponsorId            string `env:"DICE_USER_ID"`
	CommitId             string `env:"GITTAR_COMMIT"`
	GittarUsername       string `env:"GITTAR_USERNAME"`
	GittarPassword       string `env:"GITTAR_PASSWORD"`
	ProjectName          string `env:"DICE_PROJECT_NAME"`
	BranchName           string `env:"GITTAR_BRANCH"`
	DiceOpenapiToken     string `env:"DICE_OPENAPI_TOKEN" required:"true"`
	DiceOpenapiAddr      string `env:"DICE_OPENAPI_ADDR" require:"true"`
	DiceOpenapiPublicUrl string `env:"DICE_OPENAPI_PUBLIC_URL" require:"true"`
}

var (
	cfg Conf
)

func Load() error {
	return envconf.Load(&cfg)
}

func ProjectId() uint64 {
	return cfg.ProjectId
}

func PipelineID() uint64 {
	return cfg.PipelineID
}

func DiceOpenapiToken() string {
	return cfg.DiceOpenapiToken
}

func MetaFile() string {
	return cfg.MetaFile
}

func DiceOpenapiAddr() string {
	return cfg.DiceOpenapiAddr
}

func DiceOpenapiPublicUrl() string {
	return cfg.DiceOpenapiPublicUrl
}

func ActionBranch() string {
	return cfg.ActionBranch
}

func WorkDir() string {
	return cfg.WorkDir
}

func ActionApplicationName() string {
	return cfg.ActionApplicationName
}

func ActionPipelineYmlName() string {
	return cfg.ActionPipelineYmlName
}
