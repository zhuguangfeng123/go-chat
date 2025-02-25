package dao

type DBConn interface {
	Rollback()
	Commit()
}

type Transaction interface {
	Action(func(conn DBConn) error) error
}
