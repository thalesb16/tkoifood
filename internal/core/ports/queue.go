package ports

import "github.com/thalesb16/tkoifood/internal/core/domain"

type OrderQueue interface {
	Write(domain.Order) error
}
