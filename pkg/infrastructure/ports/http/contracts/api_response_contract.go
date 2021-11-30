package contracts

import "jarjarbinks/pkg/domain/contracts"

type ApiResponseContract struct {
	Result   interface{}                  `json:"result"`
	Messages *[]contracts.MessageContract `json:"messages"`
}
