package sqlclosecheck

import (
	"github.com/ryanrolds/sqlclosecheck/pkg/analyzer"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.SqlCloseCheckSettings) *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(analyzer.NewAnalyzer(settings.ExtraPackages)).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
