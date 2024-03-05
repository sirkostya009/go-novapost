package novapost

type (
	RedeliveryCalculate struct {
		CargoType string
		Amount    int
	}

	CargoDetail struct {
		CargoDescription string
		Amount           int
	}

	DocumentPriceRequest struct {
		CitySender          string
		CityRecipient       string
		Weight              float64
		ServiceType         string
		Cost                int
		CargoType           string
		SeatsAmount         int
		RedeliveryCalculate RedeliveryCalculate
		PackCount           int
		PackRef             string
		Amount              int
		CargoDetails        []CargoDetail `xml:"CargoDetails>item"`
		CargoDescription    string
	}

	TimezoneInfo struct {
		TimezoneName string
		TimezoneID   int
	}

	DocumentPrice struct {
		AssessedCost   float64
		Cost           float64
		CostRedelivery float64
		TimezoneInfo   TimezoneInfo
		CostPack       float64
	}
)

// GetDocumentPrice Розрахувати вартість послуг
//
// Метод «getDocumentPrice», працює в моделі «InternetDocument», цей метод дозволяє розрахувати вартості доставки
// вантажу.
//
// Метод дозволять прораховувати не тільки відправлення вантажів, але і шин та дисків, палет, а також документів.
func (c *Client) GetDocumentPrice(req DocumentPriceRequest) (*Response[DocumentPrice], error) {
	return RawRequest[DocumentPrice](c, InternetDocumentModel, "getDocumentPrice", req)
}

type (
	DocumentDeliveryDateRequest struct {
		DateTime      string
		ServiceType   string
		CitySender    string
		CityRecipient string
	}

	DeliveryDate struct {
		Date         string `json:"date" xml:"date"`
		TimezoneType int    `json:"timezone_type" xml:"timezone_type"`
		Timezone     string `json:"timezone" xml:"timezone"`
	}

	DocumentDeliveryDate struct {
		DeliveryDate DeliveryDate
	}
)

// GetDocumentDeliveryDate Прогноз дати доставки вантажу
//
// Метод «getDocumentDeliveryDate», працює в моделі «InternetDocument», цей метод у відповіді виведе орієнтовну дату
// доставки.
//
// Якщо в ЕН будуть такі послуги як: зворотна доставка, доставка щодня та інші, їх також потрібно вказувати у запиті як
// обов'язкові елементи запиту для розрахунку орієнтовної дати доставки.
func (c *Client) GetDocumentDeliveryDate(req DocumentDeliveryDateRequest) (*Response[DocumentDeliveryDate], error) {
	return RawRequest[DocumentDeliveryDate](c, InternetDocumentModel, "getDocumentDeliveryDate", req)
}

type (
	InternetDocument struct {
		Ref                   string
		CostOnSite            int
		EstimatedDeliveryDate string
		IntDocNumber          string
		TypeDocument          string
	}

	Services struct {
		Attorney                bool
		WaybillNewPostWithStamp bool
		UserActions             string
	}

	BackwardDeliveryData struct {
		PayerType string
		CargoType string
		Services  Services
	}

	OptionSeat struct {
		VolumetricVolume float64
		VolumetricWidth  float64
		VolumetricLength float64
		VolumetricHeight float64
		Weight           float64
		PackRef          string
	}

	InternetDocumentRequest struct {
		SenderWarehouseIndex    string
		RecipientWarehouseIndex string
		PayerType               string
		PaymentMethod           string
		DateTime                string
		CargoType               string
		VolumeGeneral           float64
		Weight                  float64
		ServiceType             string
		SeatsAmount             int
		Description             string
		Cost                    float64
		CitySender              string
		Sender                  string
		SenderAddress           string
		ContactSender           string
		SendersPhone            string
		CityRecipient           string
		Recipient               string
		RecipientAddress        string
		ContactRecipient        string
		RecipientsPhone         string
		RedBoxBarcode           string
		NewAddress              int
		BackwardDeliveryData    []BackwardDeliveryData `xml:"BackwardDeliveryData>item"`
		OptionsSeat             []OptionSeat           `xml:"OptionsSeat>item"`
		RecipientAddressNote    string
		RecipientCityName       string
		RecipientArea           string
		RecipientAreaRegions    string
		RecipientAddressName    string
		RecipientHouse          string
		RecipientFlat           string
		RecipientName           string
		RecipientType           string
		SettlementType          string
		OwnershipForm           string
		RecipientContactName    string
		EDRPOU                  string
	}
)

// SaveInternetDocument Створити експрес-накладну
//
// Базова: https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/a965630e-8512-11ec-8ced-005056b2dbe1
//
// Поштомат: https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/0227072e-8f38-11ec-8ced-005056b2dbe1
//
// Різні типи вантажу: https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/3fdc36f1-93ad-11ec-8ced-005056b2dbe1
//
// Додаткові послуги: https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/751067b8-9337-11ec-8ced-005056b2dbe1
//
// До відділення (рядком): https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/a19ced79-8f32-11ec-8ced-005056b2dbe1
//
// Із зворотньою доставкою: https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/c7c8aaf5-8f40-11ec-8ced-005056b2dbe1
//
// Із поштомату: https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/df942798-68d3-11ee-a60f-48df37b921db
//
// До адреси (рядком): https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/f74a0918-8f18-11ec-8ced-005056b2dbe1
func (c *Client) SaveInternetDocument(req InternetDocumentRequest) (*Response[InternetDocument], error) {
	return RawRequest[InternetDocument](c, InternetDocumentModel, "save", req)
}

// UpdateInternetDocument Редагувати експрес-накладну
//
// Метод «update», працює в моделі «InternetDocument», цей метод дозволяє редагувати експрес-накладну
// (інтернет-документу). У цьому методі обов'язкові до заповнення всі параметри, після оновлення замінюються лише ті, що
// були змінені.
func (c *Client) UpdateInternetDocument(req InternetDocumentRequest) (*Response[InternetDocument], error) {
	return RawRequest[InternetDocument](c, InternetDocumentModel, "update", req)
}

type (
	DocumentListRequest struct {
		DateTimeFrom    string
		DateTimeTo      string
		Page            int
		GetFullList     int
		DateTime        string
		RedeliveryMoney int
	}

	Document struct {
		Ref                     string
		DateTime                string
		PreferredDeliveryDate   string
		RecipientDateTime       string
		EWDateCreated           string
		Weight                  float64
		SeatsAmount             int
		IntDocNumber            string
		Cost                    float64
		CitySender              string
		CityRecipient           string
		SenderAddress           string
		RecipientAddress        string
		CostOnSite              float64
		PayerType               string
		PaymentMethod           string
		AfterPaymentOnGoodsCost int `xml:"AfterpaymentOnGoodsCost" json:"AfterpaymentOnGoodsCost"`
		PackingNumber           int
		RejectionReason         string
		StateId                 int
		StateName               string
	}
)

// GetDocumentList Отримати список ЕН
//
// Метод «getDocumentList», працює в моделі «InternetDocument», цей метод у відповіді виведе всі номери ЕН, які створено
// в даному в особистому кабінеті та їх ідентифікатори (Ref) в АРI.
//
// У запиті у властивостях методу можна вказати дату, за яку необхідно отримати дані по ЕН у параметрі "DateTime".
//
// Якщо не вказувати дату, система автоматично видасть Вам інформацію по всіх ЕН фактичного дня. У запиті у властивостях
// методу також можна вказувати проміжок дат, за які необхідно отримати дані по ЕН "DateTimeFrom" і "DateTimeTo".
//
// Додатково впроваджена можливість вивантажувати сторінковий список, використовуючи параметр "GetFullList". Якщо 0, то
// працює сторінкове завантаження, якщо 1 - весь список (але не більше 500 документів).
//
// Додано можливість відображення всіх ЕН зі зворотною доставкою (Грошовий переказ або Каса форпост) "RedeliveryMoney"
// зі значенням "1".
func (c *Client) GetDocumentList(req DocumentListRequest) (*Response[Document], error) {
	return RawRequest[Document](c, InternetDocumentModel, "getDocumentList", req)
}

type DocumentRefs struct {
	DocumentRefs []string
}

// DeleteInternetDocument Видалити експрес-накладну
//
// Метод «delete», працює в моделі «InternetDocument», цей метод необхідний для видалення експрес-накладної
// (інтернет-документа)
func (c *Client) DeleteInternetDocument(dr DocumentRefs) (*Response[Ref], error) {
	return RawRequest[Ref](c, InternetDocumentModel, "delete", dr)
}

type (
	ReportRequest struct {
		DocumentRefs []string
		Type         string
		DateTime     string
	}

	Report struct {
		Ref                     string
		DateTime                string
		PreferredDeliveryDate   string
		Weight                  float64
		SeatsAmount             int
		IntDocNumber            string
		Cost                    float64
		CitySender              string
		CityRecipient           string
		State                   string
		SenderAddress           string
		RecipientAddress        string
		CostOnSite              float64
		PayerType               string
		PaymentMethod           string
		AfterPaymentOnGoodsCost float64 `xml:"AfterpaymentOnGoodsCost" json:"AfterpaymentOnGoodsCost"`
		PackingNumber           int
	}
)

// GenerateReport Формування запиту для отримання повного звіту за накладними
//
// Метод «generateReport», працює в моделі «InternetDocument», цей метод формує запит для отримання повного звіту про
// накладні
func (c *Client) GenerateReport(req ReportRequest) (*Response[Report], error) {
	return RawRequest[Report](c, InternetDocumentModel, "generateReport", req)
}
