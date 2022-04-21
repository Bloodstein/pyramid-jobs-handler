package repository

type Repository interface {
	PopJob()
	SaveResult()
	StoreJob()
}
