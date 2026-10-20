package hyperflow

func NewQuery(selectFields func(*QueryFields)) Document {
	root := &selectionNode{}
	if selectFields != nil {
		selectFields(&QueryFields{node: root})
	}
	return Document{operation: "query", root: root}
}

func NewMutation(selectFields func(*MutationFields)) Document {
	root := &selectionNode{}
	if selectFields != nil {
		selectFields(&MutationFields{node: root})
	}
	return Document{operation: "mutation", root: root}
}

func NewSubscription(selectFields func(*SubscriptionFields)) Document {
	root := &selectionNode{}
	if selectFields != nil {
		selectFields(&SubscriptionFields{node: root})
	}
	return Document{operation: "subscription", root: root}
}
