package store

// UserStore структура яка використовує collection
type UserStore struct {
	collection *Collection
} // посилаємося на базову колекцію з файлу collection

// NewUserStore - ств конструктор у якому використовуємо вже готову функцію NewCollection з файлу collection

func NewUserStore() *UserStore {
	return &UserStore{
		collection: NewCollection(),
	}
}

// реалізовуємо методи з collection через UserStore

func (us *UserStore) Create(doc Document) {
	us.collection.Create(doc)
}

func (us *UserStore) Get(key string) (*Document, bool) {
	return us.collection.Get(key)
}

func (us *UserStore) Delete(key string) bool {
	return us.collection.Delete(key)
}

func (us *UserStore) List() []Document {
	return us.collection.List()
}
