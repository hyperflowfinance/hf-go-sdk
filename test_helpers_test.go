package hyperflow

func testQueryDocument() Document {
	root := &selectionNode{}
	metadata := root.object("metadata")
	metadata.field("chainId")
	return Document{operation: "query", root: root}
}
