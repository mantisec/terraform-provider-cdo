package util

import (
	"github.com/mantisec/terraform-provider-cdo/internal/util/sliceutil"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"strings"
)

func DiagSummary(d diag.Diagnostics) string {
	summaries := sliceutil.Map(d, func(t diag.Diagnostic) string {
		return t.Summary() + "\n" + t.Detail()
	})
	return strings.Join(summaries, "\n\n")
}
