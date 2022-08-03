package categories

type Core struct {
	ID       int
	Category string
	Type     string
}

type Business interface {
	AddCategories(dataCategory Core) (err error)
	GetCategories() (response []Core, err error)
	UpdateCategories(dataCategory Core, ID int) (err error)
	DestroyCategories(ID int) (err error)
}

type Data interface {
	Insert(dataCategory Core) (err error)
	Get() (response []Core, err error)
	Update(dataCategory Core, ID int) (err error)
	Destroy(ID int) (err error)
}
