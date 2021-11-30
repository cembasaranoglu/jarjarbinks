package application

import (
	"context"
	"jarjarbinks/pkg/domain/assembler"
	"jarjarbinks/pkg/domain/commands"
	"jarjarbinks/pkg/domain/errors"
	"jarjarbinks/pkg/domain/repository"
	lg "jarjarbinks/pkg/infrastructure/logging/interfaces"
	"jarjarbinks/pkg/infrastructure/mediator"
	"time"
)

type CreateCacheEntryCommandHandler struct {
	logger     lg.Logger
	repository repository.CacheStore
	assembler  assembler.CacheAssembler
}


func NewCreateCacheEntryCommandHandler(
	logger                        lg.Logger,
	repository repository.CacheStore,
	assembler  assembler.CacheAssembler) CreateCacheEntryCommandHandler {
	return CreateCacheEntryCommandHandler{
		logger:     logger,
		repository: repository,
		assembler:  assembler,
	}
}

func (p CreateCacheEntryCommandHandler) Handle(ctx context.Context, message mediator.Message) (interface{}, error) {
	command := message.(*commands.CreateCacheEntryCommand)
	if len(command.EntryKey) == 0 {
		return nil, errors.ThrowCacheEntryKeyCouldNotBeNilError()
	}
	if command.EntryValue == nil {
		return nil, errors.ThrowCacheEntryValueCouldNotBeNilError()
	}
	p.repository.StoreWithTTL(command.EntryKey, command.EntryValue, time.Second * command.ExpireAt)
	return nil, nil
}

