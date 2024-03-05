package novapost

type (
	ContactPersonRequest struct {
		CounterpartyRef string
		FirstName       string
		LastName        string
		MiddleName      string
		Phone           string
	}

	ContactPerson struct {
		Ref         string
		Description string
		Phones      string
		Email       string
		LastName    string
		FirstName   string
		MiddleName  string
	}
)

// SaveContactPerson Створити контактну особу Контрагента
//
// Метод «save», працює в моделі «ContactPerson», цей метод використовується для створення контактної особі
// відправника/отримувача.
//
// Редагувати дані контактної особи можуть тільки юридичні особи.
//
// Приватні особи можуть редагувати лише телефон контактної особи контрагента.
//
// Редагувати дані контрагента можливо тільки до моменту створення ІД з даним контрагентом.
func (c *Client) SaveContactPerson(req ContactPersonRequest) (*Response[ContactPerson], error) {
	return RawRequest[ContactPerson](c, ContactPersonModel, "save", req)
}

// UpdateContactPerson Оновити дані контактної особи Контрагента
//
// Метод «update», працює в моделі «ContactPerson», цей метод використовується для оновлення контактної особі
// відправника/отримувача.
//
// Редагувати дані контактної особи можуть тільки юридичні особи.
//
// Приватні особи можуть редагувати лише телефон контактної особи контрагента.
//
// Редагувати дані контрагента можливо тільки до моменту створення ІД з даним контрагентом.
func (c *Client) UpdateContactPerson(req ContactPersonRequest) (*Response[ContactPerson], error) {
	return RawRequest[ContactPerson](c, ContactPersonModel, "update", req)
}

// DeleteContactPerson Видалити Контактну особу Контрагента
//
// Метод «delete», працює в моделі «ContactPerson», цей метод необхідний для видалення контактної особи контрагента
// відправника/отримувача. Видалять дані контактної особи контрагента можуть лише юридичні особи.
func (c *Client) DeleteContactPerson(ref Ref) (*Response[Ref], error) {
	return RawRequest[Ref](c, ContactPersonModel, "delete", ref)
}
