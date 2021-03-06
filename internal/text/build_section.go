package text

import (
	"strings"

	"github.com/aevea/quoad"
	"github.com/aevea/release-notary/internal"
)

func (r *ReleaseNotes) buildSection(category string, commits []quoad.Commit) string {
	builder := strings.Builder{}

	builder.WriteString(r.buildHeading(category))
	builder.WriteString(r.buildCommitLog(commits, internal.PredefinedSections[category].OpenByDefault))

	return builder.String()
}
