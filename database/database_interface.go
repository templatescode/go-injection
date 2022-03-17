package database

type databaser interface {
	Create(db string) error
}

type databaseImp struct {
	db string
}

func NewDatabase(db string) databaser {
	return &databaseImp{
		db: db,
	}
}
