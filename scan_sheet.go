package novapost

type (
	InsertDocumentsRequest struct {
		Ref          string
		DocumentRefs []string `xml:"DocumentRefs>item"`
		Date         string
	}

	RegistryDocument struct {
		Ref         string
		Description string
		Number      string
		Date        string
		Errors      []string `xml:"Errors>item"`
		Warnings    []string `xml:"Warnings>item"`
		Data        []any
		Success     []RefNumber `xml:"Success>item"`
	}
)

// InsertDocuments Додати експрес-накладні
//
// Для передачі оформлених відправлень реєстром необхідно сформувати відповідний запит. У відповідь на запит формування
// реєстру повертається номер реєстру та номери відправлень зі статусом "додано" або "не додано" до реєстру.
//
// # Є такі обмеження по роботі з реєстрами:
//
// 1) До реєстру можна додати інтернет-документ (ІД), якщо дані Відправника (місто, контрагент, адреса) ідентичні для
// всіх відправлень, що додаються.
//
// 2) ІД можна додати лише до одного реєстру, тобто один і той же документ не можна одночасно додати до кількох реєстрів.
//
// 3) До реєстру не можна додати ІД, якщо по документу вже отримано друковану форму і дата відправлення (друку) менша,
// ніж вчорашня дата від дати формування реєстру.
//
// 4) ІД можна додати до реєстру, лише до моменту створення експрес накладної на основі ІД (або проведення сканувань
// цього відправлення у відділенні/підрозділі Нової Пошти).
//
// 5) До реєстру не можна додати ІД, що помічено на видалення.
//
// 6) Після отримання друкованої форми реєстру, додання документів до нього блокується.
//
// # Робота з реєстрами приймання-передачі відправлень
//
// Для передачі оформлених відправлень по Реєстру, інтегрується функціонал формування та видалення реєстрів
// приймання-передачі відправлень. При передачі відправлень по Реєстру необхідно на кожному відправленні розміщувати
// маркування і роздрукувати два екземпляри Реєстру.
func (c *Client) InsertDocuments(req InsertDocumentsRequest) (*Response[RegistryDocument], error) {
	return RawRequest[RegistryDocument](c, ScanSheetModel, "insertDocuments", req)
}

type (
	ScanSheetRequest struct {
		Ref             string
		CounterpartyRef string
	}

	ScanSheet struct {
		Ref              string
		Number           string
		DateTime         string
		CitySenderRef    string
		CitySender       string
		SenderAddressRef string
		SenderAddress    string
		SenderRef        string
		Sender           string
		Count            int     `json:",string"`
		Printed          IntBool `json:",string"`
	}
)

// GetScanSheet Завантажити інформацію по одному реєстру
//
// # Робота з реєстрами прийманя-передачі відправлень
//
// Для передачі оформлених відправлень по Реєстру, інтегрується функціонал формування та видалення реєстрів
// приймання-передачі відправлень. При передачі відправлень по Реєстру необхідно на кожному відправленні розміщувати
// маркування та роздрукувати два екземпляри Реєстру.
func (c *Client) GetScanSheet(req ScanSheetRequest) (*Response[ScanSheet], error) {
	return RawRequest[ScanSheet](c, ScanSheetModel, "getScanSheet", req)
}

// GetScanSheetList Завантажити список реєстрів
//
// Для передачі оформлених відправлень по Реєстру, інтегрується функціонал формування та видалення реєстрів
// приймання-передачі відправлень. При передачі відправлень по Реєстру необхідно на кожнома відправленні розміщувати
// маркування та роздрукувати два екземпляри Реєстру.
func (c *Client) GetScanSheetList() (*Response[ScanSheet], error) {
	return RawRequest[ScanSheet](c, ScanSheetModel, "getScanSheetList", nil)
}

type (
	ScanSheetRefs struct {
		ScanSheetRefs []string `xml:"ScanSheetRefs>item"`
	}

	ScanSheetRef struct {
		Ref    string
		Number string
		Error  string
	}
)

// DeleteScanSheet Видалити (розформувати) реєстр відправлень
//
// Після видалення реєстра, з інформаційної системи «Нова пошта" видаляється номер реєстру, а експрес-накладні, що були
// включені в нього звільняються, але не видаляються (відбувається розформування рееєстру). Для видалення реєстру
// необхідно сформувати відповідний запит.
func (c *Client) DeleteScanSheet(ssr ScanSheetRefs) (*Response[ScanSheetRef], error) {
	return RawRequest[ScanSheetRef](c, ScanSheetModel, "delete", ssr)
}

type RemoveDocumentsRequest struct {
	DocumentRefs []string `xml:"DocumentRefs>item"`
	Ref          string   `json:",omitempty" xml:",omitempty"`
}

// RemoveDocuments Видалити експрес-накладні з реєстру
//
// Після видалення реєстру в інформаційній системі «Нова пошта» видаляється номер реєстру, але експрес-накладні, які
// знаходилися в реєстрі, не видаляються (відбувається розформування реєстру). Для видалення експрес-накладних необхідно
// сформувати відповідний запит.
func (c *Client) RemoveDocuments(req RemoveDocumentsRequest) (*Response[ScanSheetRef], error) {
	return RawRequest[ScanSheetRef](c, ScanSheetModel, "removeDocuments", req)
}
