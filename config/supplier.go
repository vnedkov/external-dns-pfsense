package config

// Supplier is a configuration supplier.
//
//go:generate mockery --name Supplier
type Supplier interface {
	Get(key string) (string, bool)
}
