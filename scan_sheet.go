package novapost

import "time"

type (
	InsertDocumentsRequest struct {
		Ref          string   `json:"Ref"`
		DocumentRefs []string `json:"DocumentRefs"`
		Date         string   `json:"Date"`
	}

	RegistryDocument struct {
		Ref         string      `json:"Ref"`
		Description string      `json:"Description"`
		Number      string      `json:"Number"`
		Date        string      `json:"Date"`
		Errors      []string    `json:"Errors"`
		Warnings    []string    `json:"Warnings"`
		Data        any         `json:"Data"`
		Success     []RefNumber `json:"Success"`
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
func (c Client) InsertDocuments(req InsertDocumentsRequest) (Response[RegistryDocument], error) {
	return request[RegistryDocument](c, ScanSheetModel, "insertDocuments", req)
}

type (
	ScanSheetRequest struct {
		Ref             string `json:"Ref"`
		CounterpartyRef string `json:"CounterpartyRef"`
	}

	ScanSheet struct {
		Ref              string    `json:"Ref"`
		Number           string    `json:"Number"`
		DateTime         time.Time `json:"DateTime"`
		Count            string    `json:"Count"`
		CitySenderRef    string    `json:"CitySenderRef"`
		CitySender       string    `json:"CitySender"`
		SenderAddressRef string    `json:"SenderAddressRef"`
		SenderAddress    string    `json:"SenderAddress"`
		SenderRef        string    `json:"SenderRef"`
		Sender           string    `json:"Sender"`
		Printed          string    `json:"Printed"`
	}
)

// GetScanSheet Завантажити інформацію по одному реєстру
//
// # Робота з реєстрами прийманя-передачі відправлень
//
// Для передачі оформлених відправлень по Реєстру, інтегрується функціонал формування та видалення реєстрів
// приймання-передачі відправлень. При передачі відправлень по Реєстру необхідно на кожному відправленні розміщувати
// маркування та роздрукувати два екземпляри Реєстру.
func (c Client) GetScanSheet(req ScanSheetRequest) (Response[ScanSheet], error) {
	return request[ScanSheet](c, ScanSheetModel, "getScanSheet", req)
}

// GetScanSheetList Завантажити список реєстрів
//
// Для передачі оформлених відправлень по Реєстру, інтегрується функціонал формування та видалення реєстрів
// приймання-передачі відправлень. При передачі відправлень по Реєстру необхідно на кожнома відправленні розміщувати
// маркування та роздрукувати два екземпляри Реєстру.
func (c Client) GetScanSheetList() (Response[ScanSheet], error) {
	return request[ScanSheet](c, ScanSheetModel, "getScanSheetList", nil)
}

type (
	ScanSheetRefs struct {
		ScanSheetRefs []string `json:"ScanSheetRefs"`
	}

	ScanSheetRef struct {
		Ref    string `json:"Ref"`
		Number string `json:"Number"`
		Error  string `json:"Error"`
	}

	RemoveDocumentsRequest struct {
		DocumentRefs []string `json:"DocumentRefs"`
		Ref          string   `json:"Ref"`
	}
)

// DeleteScanSheet Видалити (розформувати) реєстр відправлень
//
// Після видалення реєстра, з інформаційної системи «Нова пошта" видаляється номер реєстру, а експрес-накладні, що були
// включені в нього звільняються, але не видаляються (відбувається розформування рееєстру). Для видалення реєстру
// необхідно сформувати відповідний запит.
func (c Client) DeleteScanSheet(req ScanSheetRefs) (Response[ScanSheetRef], error) {
	return request[ScanSheetRef](c, ScanSheetModel, "delete", req)
}

// RemoveDocuments Видалити експрес-накладні з реєстру
//
// Після видалення реєстру в інформаційній системі «Нова пошта» видаляється номер реєстру, але експрес-накладні, які
// знаходилися в реєстрі, не видаляються (відбувається розформування реєстру). Для видалення експрес-накладних необхідно
// сформувати відповідний запит.
func (c Client) RemoveDocuments(req RemoveDocumentsRequest) (Response[ScanSheetRef], error) {
	return request[ScanSheetRef](c, ScanSheetModel, "removeDocuments", req)
}
