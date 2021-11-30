package assembler

import (
	"jarjarbinks/pkg/domain"
	"jarjarbinks/pkg/domain/assembler"
	"jarjarbinks/pkg/domain/contracts"
)

type cacheAssembler struct{

}

func (c cacheAssembler) ToContract(cache domain.Cache) *contracts.CacheEntryContract {
	panic("implement me")
}

func NewCacheAssembler() assembler.CacheAssembler{
	return &cacheAssembler{}
}
