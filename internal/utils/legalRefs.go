package utils

import (
	"regexp"
	"strings"
)

var legalRef = regexp.MustCompile(`(ст\.|п\.|Приложение|ПП|Федерального закона|№.*|РТН|ГрК|ТР ТС).*`)

func CleanLegalRefs(text string) string {
	return strings.TrimSpace(legalRef.ReplaceAllString(text, ""))
}
