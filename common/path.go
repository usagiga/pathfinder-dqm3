package common

// PathBase 配合ツリーのノード。各ノードには配合対象となるモンスターが記述される
type PathBase struct {
	MonsterId int
}

// PathRoot スカウト限定、配信限定、タマゴ限定など
type PathRoot struct {
	*PathBase
}

// PathSynth 通常配合
type PathSynth struct {
	*PathBase
	XFamilyId    Family
	XRank        Rank
	YMonsterId   Family
	YMonsterRank Rank
}

// PathSpecialSynth 特殊配合
type PathSpecialSynth struct {
	*PathBase
	XMonsterId int
	YMonsterId int
}

// PathSpecialSynthWithRank 特殊配合(片方が系統・ランクによる配合)
type PathSpecialSynthWithRank struct {
	*PathBase
	XMonsterId int
	YFamily    Family
	YRank      Rank
}

// PathQuadSynth 4体配合
type PathQuadSynth struct {
	*PathBase
	AMonsterId int
	BMonsterId int
	CMonsterId int
	DMonsterId int
}
