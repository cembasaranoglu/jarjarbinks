package errors

import "errors"

const (
	cacheEntryKeyCouldNotBeNilErrorCode   = "ERR_1001"
	cacheEntryValueCouldNotBeNilErrorCode = "ERR_1002"
	cacheEntryDoesNotExistsErrorCode = "ERR_1003"
)

var (
	cacheEntryKeyCouldNotBeNilError            = errors.New("cache entry key could not be nil")
	cacheEntryValueCouldNotBeNilError            = errors.New("cache entry value could not be nil")
	cacheEntryDoesNotExistsError = errors.New("cache entry does not exist")
)


func ThrowCacheEntryKeyCouldNotBeNilError() *DomainError {
	return &DomainError{
		ErrorBase: NewError(cacheEntryKeyCouldNotBeNilErrorCode, cacheEntryKeyCouldNotBeNilError),
	}
}

func ThrowCacheEntryValueCouldNotBeNilError() *DomainError {
	return &DomainError{
		ErrorBase: NewError(cacheEntryValueCouldNotBeNilErrorCode, cacheEntryValueCouldNotBeNilError),
	}
}

func ThrowCacheEntryDoesNotExistsError() *DoesNotExistsDomainError {
	return &DoesNotExistsDomainError{
		ErrorBase: NewError(cacheEntryDoesNotExistsErrorCode, cacheEntryDoesNotExistsError),
	}
}