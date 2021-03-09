package cli

import (
	"github.com/zhilyaev/helmwave/pkg/helmwave"
	"github.com/zhilyaev/helmwave/pkg/kubedog"
	"helm.sh/helm/v3/pkg/cli"
)

func New() *helmwave.Config {
	return &helmwave.Config{
		Version: helmwave.Version,
		Helm:    cli.New(),
		Logger:  &helmwave.Log{},
		Kubedog: &kubedog.Config{},
	}
}
