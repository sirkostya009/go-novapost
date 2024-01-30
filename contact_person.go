package novapost

type (
	ContactPersonRequest struct {
		CounterpartyRef string `json:"CounterpartyRef"`
		FirstName       string `json:"FirstName"`
		LastName        string `json:"LastName"`
		MiddleName      string `json:"MiddleName"`
		Phone           string `json:"Phone"`
	}

	ContactPerson struct {
		Ref         string `json:"Ref"`
		Description string `json:"Description"`
		Phones      string `json:"Phones"`
		Email       string `json:"Email"`
		LastName    string `json:"LastName"`
		FirstName   string `json:"FirstName"`
		MiddleName  string `json:"MiddleName"`
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
func (c Client) SaveContactPerson(req ContactPersonRequest) (Response[ContactPerson], error) {
	return request[ContactPerson](c, ContactPersonModel, "save", req)
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
func (c Client) UpdateContactPerson(req ContactPersonRequest) (Response[ContactPerson], error) {
	return request[ContactPerson](c, ContactPersonModel, "update", req)
}

// DeleteContactPerson Видалити Контактну особу Контрагента
//
// Метод «delete», працює в моделі «ContactPerson», цей метод необхідний для видалення контактної особи контрагента
// відправника/отримувача. Видалять дані контактної особи контрагента можуть лише юридичні особи.
func (c Client) DeleteContactPerson(ref Ref) (Response[Ref], error) {
	return request[Ref](c, ContactPersonModel, "delete", ref)
}
