package db

// ExistenceChecker is an interface used to check the existence of an entity.
type ExistenceChecker[T any] interface {
	// CheckExistence checks the existence of a record by the id column.
	// For composite primary keys, a struct can be used as a parameter.
	CheckExistence(id T) (bool, error)
}

// DataWriter is an interface used to write data to the database.
type DataWriter[T any, K any] interface {
	// Insert inserts a new record into the database and returns the primary key value.
	// For composite primary keys, a struct can be used as a result.
	Insert(entity T) (K, error)

	// Update updates the given entity by the id value.
	// For composite primary keys, a struct can be used as a parameter.
	Update(id K, entity T) (bool, error)

	// Delete deletes the database entity record by the id value.
	// For composite primary keys, a struct can be used as a parameter.
	Delete(id K) (bool, error)
}

// DataReader is an interface used to read data from the database.
type DataReader[T any, K any] interface {
	// GetById gets the entity record by the id value.
	// For composite primary keys, a struct can be used as a parameter.
	GetById(id K) (T, error)

	// Retrieve retrieves all the records of the entity from the database.
	Retrieve() ([]T, error)
}

// DataFetcher is an interface used to fetch entity records by given limit and offset values.
type DataFetcher[T any] interface {
	// Fetch fetches the entity records by the given limit and offset.
	Fetch(limit, offset int) ([]T, error)
}

// Repository is an interface used for general database operations.
type Repository[T any, K any] interface {
	DataReader[T, K]
	DataFetcher[T]
	ExistenceChecker[K]
	DataWriter[T, K]
}
