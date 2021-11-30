package application

import (
	"context"
	"jarjarbinks/pkg/domain/assembler"
	"jarjarbinks/pkg/domain/errors"
	"jarjarbinks/pkg/domain/queries"
	"jarjarbinks/pkg/domain/repository"
	lg "jarjarbinks/pkg/infrastructure/logging/interfaces"
	"jarjarbinks/pkg/infrastructure/mediator"
)

type FindCacheEntryByKeyQueryHandler struct {
	logger     lg.Logger
	repository repository.CacheStore
	assembler  assembler.CacheAssembler
}


func NewFindCacheEntryByKeyQueryHandler(
	logger                        lg.Logger,
	repository repository.CacheStore,
	assembler  assembler.CacheAssembler) FindCacheEntryByKeyQueryHandler {
	return FindCacheEntryByKeyQueryHandler{
		logger:     logger,
		repository: repository,
		assembler:  assembler,
	}
}

func (p FindCacheEntryByKeyQueryHandler) Handle(ctx context.Context, message mediator.Message) (interface{}, error) {
	query := message.(*queries.FindCacheEntryByKeyQuery)
	if len(query.EntryKey) == 0 {
		return nil, errors.ThrowCacheEntryKeyCouldNotBeNilError()
	}
	if cacheEntry, exists := p.repository.Peek(query.EntryKey); !exists{
		return nil, errors.ThrowCacheEntryDoesNotExistsError()
	}else{
		return cacheEntry, nil
	}
}

