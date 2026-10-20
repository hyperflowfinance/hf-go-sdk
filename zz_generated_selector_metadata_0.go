package hyperflow

func (s *MetadataFields) Typename() *MetadataFields {
	s.node.field("__typename")
	return s
}

func (s *MetadataFields) ClientVersion(argChain Optional[string], argChainID Optional[int32]) *MetadataFields {
	s.node.field("clientVersion", argument{name: "chain", gqlType: "String", value: argChain, optional: true}, argument{name: "chainId", gqlType: "Int", value: argChainID, optional: true})
	return s
}

func (s *MetadataFields) ChainID(argChain Optional[string], argChainID Optional[int32]) *MetadataFields {
	s.node.field("chainId", argument{name: "chain", gqlType: "String", value: argChain, optional: true}, argument{name: "chainId", gqlType: "Int", value: argChainID, optional: true})
	return s
}

func (s *MetadataFields) NetVersion(argChain Optional[string], argChainID Optional[int32]) *MetadataFields {
	s.node.field("netVersion", argument{name: "chain", gqlType: "String", value: argChain, optional: true}, argument{name: "chainId", gqlType: "Int", value: argChainID, optional: true})
	return s
}

func (s *MetadataFields) NetListening(argChain Optional[string], argChainID Optional[int32]) *MetadataFields {
	s.node.field("netListening", argument{name: "chain", gqlType: "String", value: argChain, optional: true}, argument{name: "chainId", gqlType: "Int", value: argChainID, optional: true})
	return s
}

func (s *MetadataFields) NetPeerCount(argChain Optional[string], argChainID Optional[int32]) *MetadataFields {
	s.node.field("netPeerCount", argument{name: "chain", gqlType: "String", value: argChain, optional: true}, argument{name: "chainId", gqlType: "Int", value: argChainID, optional: true})
	return s
}

func (s *MetadataFields) EthSyncing(argChain Optional[string], argChainID Optional[int32]) *MetadataFields {
	s.node.field("ethSyncing", argument{name: "chain", gqlType: "String", value: argChain, optional: true}, argument{name: "chainId", gqlType: "Int", value: argChainID, optional: true})
	return s
}

func (s *MetadataFields) EthMining(argChain Optional[string], argChainID Optional[int32]) *MetadataFields {
	s.node.field("ethMining", argument{name: "chain", gqlType: "String", value: argChain, optional: true}, argument{name: "chainId", gqlType: "Int", value: argChainID, optional: true})
	return s
}

func (s *MetadataFields) EthHashrate(argChain Optional[string], argChainID Optional[int32]) *MetadataFields {
	s.node.field("ethHashrate", argument{name: "chain", gqlType: "String", value: argChain, optional: true}, argument{name: "chainId", gqlType: "Int", value: argChainID, optional: true})
	return s
}

func (s *MetadataFields) EthProtocolVersion(argChain Optional[string], argChainID Optional[int32]) *MetadataFields {
	s.node.field("ethProtocolVersion", argument{name: "chain", gqlType: "String", value: argChain, optional: true}, argument{name: "chainId", gqlType: "Int", value: argChainID, optional: true})
	return s
}

func (s *MetadataFields) EthBlockNumber(argChainID Optional[int32]) *MetadataFields {
	s.node.field("ethBlockNumber", argument{name: "chainId", gqlType: "Int", value: argChainID, optional: true})
	return s
}

func (s *MetadataFields) EthGasPrice(argChainID Optional[int32]) *MetadataFields {
	s.node.field("ethGasPrice", argument{name: "chainId", gqlType: "Int", value: argChainID, optional: true})
	return s
}
