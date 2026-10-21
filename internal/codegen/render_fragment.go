package codegen

import (
	"fmt"
	"strconv"
	"strings"
)

func renderFragment(receiver, method, typeName string, types map[string]Type) (string, error) {
	target, ok := types[typeName]
	if !ok || target.Kind != "OBJECT" {
		return "", fmt.Errorf("possible type %q is missing or not an object", typeName)
	}
	childType := goName(typeName) + "Fields"
	var source strings.Builder
	source.WriteString("// " + method + " adds an inline fragment for GraphQL " + typeName + ".\n")
	source.WriteString("func (s *" + receiver + ") " + method + "(selectFields func(*" + childType + ")) *" + receiver + " {\n")
	source.WriteString("\tchild := s.node.inlineFragment(" + strconv.Quote(typeName) + ")\n")
	source.WriteString("\tif selectFields != nil { selectFields(&" + childType + "{node: child}) }\n")
	source.WriteString("\treturn s\n}\n")
	return source.String(), nil
}
