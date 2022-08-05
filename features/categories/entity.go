package categories

type Core struct {
	ID       int
	Category string
	Type     string
}

type Business interface {
	AddCategories(dataCategory Core) (row int, err error)
	GetCategories() (response []Core, err error)
	UpdateCategories(dataCategory Core, ID int) (row int, err error)
	DestroyCategories(ID int) (row int, err error)
}

type Data interface {
	Insert(dataCategory Core) (row int, err error)
	Get() (response []Core, err error)
	Update(dataCategory Core, ID int) (row int, err error)
	Destroy(ID int) (row int, err error)
}
