package date

type DateRepository interface {
}

type mysqlDateRepository struct {
}

func NewMysqlDateRepository() DateRepository {
	return &mysqlDateRepository{}
}