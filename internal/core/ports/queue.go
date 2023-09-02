package ports

import "github.com/thalesb16/koifood/internal/core/domain"

type OrderQueue interface {
	Write(domain.Order) error
}
