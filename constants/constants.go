package temp11

const (
	PlatformName                         = "DataVerifier"
	PeersKey                             = "peers"
	PeersListKey                         = "peersList"
	MaxRetryForPeerListRequest           = 5
	TotalLeadingDuration                 = 150 // 150 seconds
	EffectiveLeadingDuration             = 120 // 120 seconds
	LeadingPowerReleaseDuration          = 20  // 30 seconds
	BreakTime                            = 10  // 5 seconds
	WaitingTimeForPeersInEachRound       = 5   // 5 seconds
	WaitingTimeForRemovalOfUnJoinedPeers = 5   // 5 seconds
	MaxWaitTimeForPeerToJoinNetwork      = 150 // 150 seconds
	RegistrationRequestMessage           = "Can You Add Me In Network"
	EthereumNodeSocketURL                = "wss://polygon-mumbai.g.alchemy.com/v2/barNHxwKcvdxJuDoKlbor5qx6mhT2C_O"
	EthereumNodeRPCURL                   = "https://polygon-mumbai.g.alchemy.com/v2/VUq8Nxd1_VYC2aA7R0ySc1kqL05cw0CF"
	NodeManagerContract                  = "0x7F152299960132F4F5f0E78a53c0345cBCe8b257"
	RequestMangerContract                = "0x7Ae4f1C3371e1c882c5e653454EEeab5f22dAabC"
	NodeAddedTopicHash                   = "0xe7aeb8c7da993ec76eb1440800ed15083dbf35b984ff73ec80e8d0cca20d86a6"
	NodeLeftTopicHash                    = "0xb03075d77f987a47c111a292ed269586f40b157471734a6fa420ccdedd400a94"
	RequestReceivedEventTopicHash        = "0x85dcc3449f98075b485886e476487433e044f2db02a65b5bf5d29e92392b5c2c"
	MinETHBalanceForRequestFulfillment   = 100000000000000000
	ChainID                              = 80001
	MaxGasForTransaction                 = 5000000
)
