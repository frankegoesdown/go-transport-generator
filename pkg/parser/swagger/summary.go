package swagger

import (
	"errors"
	"strings"

	"github.com/wildberries-ru/go-transport-generator/pkg/api"
)

var (
	errSwaggerSummaryDidNotSet = errors.New("swagger summary did not set")
)

type summary struct {
	prefix string
	suffix string
	next   Parser
}

// Parse ...
func (t *summary) Parse(info *api.SwaggerInfo, firstTag string, tags ...string) (err error) {
	if strings.HasPrefix(firstTag, t.prefix) && strings.HasSuffix(firstTag, t.suffix) {
		if len(tags) == 1 {
			info.Summary = &tags[0]
			return
		}
		return errSwaggerSummaryDidNotSet
	}
	return t.next.Parse(info, firstTag, tags...)
}

// NewSummary ...
func NewSummary(prefix string, suffix string, next Parser) Parser {
	return &summary{
		prefix: prefix,
		suffix: suffix,
		next:   next,
	}
}
