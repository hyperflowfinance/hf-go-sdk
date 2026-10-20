package codegen

import "fmt"

func introspectionQuery() string {
	ref := typeRefSelection(10)
	return fmt.Sprintf(`
query HFSDKIntrospection {
  __schema {
    queryType { name }
    mutationType { name }
    subscriptionType { name }
    types {
      kind name
      fields(includeDeprecated: true) {
        name
        args { name defaultValue type { %s } }
        type { %s }
      }
      inputFields { name defaultValue type { %s } }
      enumValues(includeDeprecated: true) { name }
      possibleTypes { kind name }
    }
  }
}`, ref, ref, ref)
}

func typeRefSelection(depth int) string {
	if depth == 0 {
		return "kind name"
	}
	return "kind name ofType { " + typeRefSelection(depth-1) + " }"
}
