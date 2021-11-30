package assembler

import (
	"jarjarbinks/pkg/domain"
	"jarjarbinks/pkg/domain/contracts"
)

type CacheAssembler interface{
	ToContract(cache domain.Cache) *contracts.CacheEntryContract
}
