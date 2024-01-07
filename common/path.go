package common

type SynthPath interface {
	GetMonsterId() int
	GetSynthPathType() SynthPathType
}

type SynthPathType int

const (
	SynthPathType_Root SynthPathType = iota + 1
	SynthPathType_NominalSynth
	SynthPathType_SpecialSynth
	SynthPathType_SpecialSynthWithRank
	SynthPathType_QuadSynth
)

// SynthPathBase 配合ツリーのノード。各ノードには配合対象となるモンスターが記述される
type SynthPathBase struct {
	MonsterId int
}

func (p *SynthPathBase) GetMonsterId() int {
	return p.MonsterId
}

// SynthPathRoot スカウト限定、配信限定、タマゴ限定など
type SynthPathRoot struct {
	*SynthPathBase
}

func (p *SynthPathRoot) GetSynthPathType() SynthPathType {
	return SynthPathType_Root
}

// SynthPathNominalSynth 通常配合
type SynthPathNominalSynth struct {
	*SynthPathBase
	XFamilyId    Family
	XRank        Rank
	YMonsterId   Family
	YMonsterRank Rank
}

func (p *SynthPathNominalSynth) GetSynthPathType() SynthPathType {
	return SynthPathType_NominalSynth
}

// SynthPathSpecialSynth 特殊配合
type SynthPathSpecialSynth struct {
	*SynthPathBase
	XMonsterId int
	YMonsterId int
}

func (p *SynthPathSpecialSynth) GetSynthPathType() SynthPathType {
	return SynthPathType_SpecialSynth
}

// SynthPathSpecialSynthWithRank 特殊配合(片方が系統・ランクによる配合)
type SynthPathSpecialSynthWithRank struct {
	*SynthPathBase
	XMonsterId int
	YFamily    Family
	YRank      Rank
}

func (p *SynthPathSpecialSynthWithRank) GetSynthPathType() SynthPathType {
	return SynthPathType_SpecialSynthWithRank
}

// SynthPathQuadSynth 4体配合
type SynthPathQuadSynth struct {
	*SynthPathBase
	AMonsterId int
	BMonsterId int
	CMonsterId int
	DMonsterId int
}

func (p *SynthPathQuadSynth) GetSynthPathType() SynthPathType {
	return SynthPathType_QuadSynth
}
