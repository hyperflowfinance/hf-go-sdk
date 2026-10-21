package codegen

import (
	"fmt"
	"strconv"
	"strings"
)

func renderOperations(directory string, schema *Schema, types map[string]Type) error {
	var source strings.Builder
	source.WriteString(sourceHeader())
	if err := appendOperation(&source, "NewQuery", "query", schema.QueryType, types); err != nil {
		return err
	}
	if schema.MutationType != nil {
		if err := appendOperation(&source, "NewMutation", "mutation", schema.MutationType, types); err != nil {
			return err
		}
	}
	if schema.SubscriptionType != nil {
		if err := appendOperation(&source, "NewSubscription", "subscription", schema.SubscriptionType, types); err != nil {
			return err
		}
	}
	return writeSource(directory, "zz_generated_operations.go", source.String())
}

func appendOperation(source *strings.Builder, function, operation string, root *RootType, types map[string]Type) error {
	typeDef, ok := types[root.Name]
	if !ok {
		return fmt.Errorf("operation root %q is missing from schema", root.Name)
	}
	if typeDef.Kind != "OBJECT" {
		return fmt.Errorf("operation root %q is not an object", root.Name)
	}
	fields := goName(root.Name) + "Fields"
	source.WriteString("// " + function + " creates a " + operation + " document from selected fields.\n")
	source.WriteString("func " + function + "(selectFields func(*" + fields + ")) Document {\n")
	source.WriteString("\troot := &selectionNode{}\n")
	source.WriteString("\tif selectFields != nil { selectFields(&" + fields + "{node: root}) }\n")
	source.WriteString("\treturn Document{operation: " + strconv.Quote(operation) + ", root: root}\n}\n\n")
	return nil
}
