package novapost

type (
	RedeliveryCalculate struct {
		CargoType string
		Amount    int `json:",string"`
	}

	CargoDetail struct {
		CargoDescription string
		Amount           int `json:",string"`
	}

	DocumentPriceRequest struct {
		CitySender          string
		CityRecipient       string
		Weight              float64 `json:",string"`
		ServiceType         string
		CargoType           string
		PackRef             string        `json:",omitempty" xml:",omitempty"`
		CargoDetails        []CargoDetail `json:",omitempty" xml:"CargoDetails>item,omitempty"`
		CargoDescription    string        `json:",omitempty" xml:",omitempty"`
		RedeliveryCalculate RedeliveryCalculate
		Amount              int `json:",string,omitempty" xml:",omitempty"`
		PackCount           int `json:",string,omitempty" xml:",omitempty"`
		Cost                int `json:",string"`
		SeatsAmount         int `json:",string"`
	}

	TimezoneInfo struct {
		TimezoneName string
		TimezoneID   int
	}

	DocumentPrice struct {
		AssessedCost   float64
		Cost           float64
		CostRedelivery float64
		CostPack       float64
		TimezoneInfo   TimezoneInfo
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
		DateTime      string `json:",omitempty" xml:",omitempty"`
		ServiceType   string
		CitySender    string
		CityRecipient string
	}

	DeliveryDate struct {
		Date         string `json:"date" xml:"date"`
		Timezone     string `json:"timezone" xml:"timezone"`
		TimezoneType int    `json:"timezone_type" xml:"timezone_type"`
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
		EstimatedDeliveryDate string
		IntDocNumber          string
		TypeDocument          string
		CostOnSite            int `json:",string"`
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
		VolumetricVolume float64 `json:"volumetricVolume,string" xml:"volumetricVolume"`
		VolumetricWidth  float64 `json:"volumetricWidth,string" xml:"volumetricWidth"`
		VolumetricLength float64 `json:"volumetricLength,string" xml:"volumetricLength"`
		VolumetricHeight float64 `json:"volumetricHeight,string" xml:"volumetricHeight"`
		Weight           float64 `json:"weight,string" xml:"weight"`
		PackRef          string  `json:"packRef" xml:"packRef"`
	}

	InternetDocumentRequest struct {
		SenderWarehouseIndex    string                 `json:",omitempty" xml:",omitempty"`
		RecipientWarehouseIndex string                 `json:",omitempty" xml:",omitempty"`
		PayerType               string                 `json:",omitempty" xml:",omitempty"`
		PaymentMethod           string                 `json:",omitempty" xml:",omitempty"`
		DateTime                string                 `json:",omitempty" xml:",omitempty"`
		CargoType               string                 `json:",omitempty" xml:",omitempty"`
		VolumeGeneral           float64                `json:",string,omitempty" xml:",omitempty"`
		Weight                  float64                `json:",string,omitempty" xml:",omitempty"`
		ServiceType             string                 `json:",omitempty" xml:",omitempty"`
		Description             string                 `json:",omitempty" xml:",omitempty"`
		Cost                    float64                `json:",string,omitempty" xml:",omitempty"`
		CitySender              string                 `json:",omitempty" xml:",omitempty"`
		Sender                  string                 `json:",omitempty" xml:",omitempty"`
		SenderAddress           string                 `json:",omitempty" xml:",omitempty"`
		ContactSender           string                 `json:",omitempty" xml:",omitempty"`
		SendersPhone            string                 `json:",omitempty" xml:",omitempty"`
		CityRecipient           string                 `json:",omitempty" xml:",omitempty"`
		Recipient               string                 `json:",omitempty" xml:",omitempty"`
		RecipientAddress        string                 `json:",omitempty" xml:",omitempty"`
		ContactRecipient        string                 `json:",omitempty" xml:",omitempty"`
		RecipientsPhone         string                 `json:",omitempty" xml:",omitempty"`
		RedBoxBarcode           string                 `json:",omitempty" xml:",omitempty"`
		SeatsAmount             int                    `json:",string,omitempty" xml:",omitempty"`
		NewAddress              IntBool                `json:",string,omitempty" xml:",omitempty"`
		BackwardDeliveryData    []BackwardDeliveryData `json:",omitempty" xml:"BackwardDeliveryData>item,omitempty"`
		OptionsSeat             []OptionSeat           `json:",omitempty" xml:"OptionsSeat>item,omitempty"`
		RecipientAddressNote    string                 `json:",omitempty" xml:",omitempty"`
		RecipientCityName       string                 `json:",omitempty" xml:",omitempty"`
		RecipientArea           string                 `json:",omitempty" xml:",omitempty"`
		RecipientAreaRegions    string                 `json:",omitempty" xml:",omitempty"`
		RecipientAddressName    string                 `json:",omitempty" xml:",omitempty"`
		RecipientHouse          string                 `json:",omitempty" xml:",omitempty"`
		RecipientFlat           string                 `json:",omitempty" xml:",omitempty"`
		RecipientName           string                 `json:",omitempty" xml:",omitempty"`
		RecipientType           string                 `json:",omitempty" xml:",omitempty"`
		SettlementType          string                 `json:",omitempty" xml:",omitempty"`
		OwnershipForm           string                 `json:",omitempty" xml:",omitempty"`
		RecipientContactName    string                 `json:",omitempty" xml:",omitempty"`
		EDRPOU                  string                 `json:",omitempty" xml:",omitempty"`
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
		DateTime        string `json:",omitempty" xml:",omitempty"`
		RedeliveryMoney int    `json:",string,omitempty"`
		Page            int    `json:",string,omitempty"`
		GetFullList     int    `json:",string"`
	}

	Document struct {
		Ref                     string
		DateTime                string
		PreferredDeliveryDate   string
		RecipientDateTime       string
		EWDateCreated           string
		Weight                  float64 `json:",string"`
		IntDocNumber            string
		Cost                    float64 `json:",string"`
		CitySender              string
		CityRecipient           string
		SenderAddress           string
		RecipientAddress        string
		CostOnSite              float64 `json:",string"`
		PayerType               string
		PaymentMethod           string
		AfterPaymentOnGoodsCost int `xml:"AfterpaymentOnGoodsCost" json:"AfterpaymentOnGoodsCost,string"`
		PackingNumber           int `json:",string"`
		RejectionReason         string
		StateName               string
		StateId                 int `json:",string"`
		SeatsAmount             int `json:",string"`
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
		Weight                  float64 `json:",string"`
		IntDocNumber            string
		Cost                    float64 `json:",string"`
		CitySender              string
		CityRecipient           string
		State                   string
		SenderAddress           string
		RecipientAddress        string
		CostOnSite              float64 `json:",string"`
		PayerType               string
		PaymentMethod           string
		AfterPaymentOnGoodsCost float64 `xml:"AfterpaymentOnGoodsCost" json:"AfterpaymentOnGoodsCost,string"`
		PackingNumber           int     `json:",string"`
		SeatsAmount             int     `json:",string"`
	}
)

// GenerateReport Формування запиту для отримання повного звіту за накладними
//
// Метод «generateReport», працює в моделі «InternetDocument», цей метод формує запит для отримання повного звіту про
// накладні
func (c *Client) GenerateReport(req ReportRequest) (*Response[Report], error) {
	return RawRequest[Report](c, InternetDocumentModel, "generateReport", req)
}
