package novapost

type (
	RedeliveryCalculate struct {
		CargoType string `json:"CargoType"`
		Amount    string `json:"Amount"`
	}

	CargoDetail struct {
		CargoDescription string `json:"CargoDescription"`
		Amount           string `json:"Amount"`
	}

	DocumentPriceRequest struct {
		CitySender          string `json:"CitySender"`
		CityRecipient       string `json:"CityRecipient"`
		Weight              string `json:"Weight"`
		ServiceType         string `json:"ServiceType"`
		Cost                string `json:"Cost"`
		CargoType           string `json:"CargoType"`
		SeatsAmount         string `json:"SeatsAmount"`
		RedeliveryCalculate `json:"RedeliveryCalculate"`
		PackCount           string        `json:"PackCount"`
		PackRef             string        `json:"PackRef"`
		Amount              string        `json:"Amount"`
		CargoDetails        []CargoDetail `json:"CargoDetails"`
		CargoDescription    string        `json:"CargoDescription"`
	}

	TimezoneInfo struct {
		TimezoneName string `json:"TzoneName"`
		TimezoneID   string `json:"TzoneID"`
	}

	DocumentPrice struct {
		AssessedCost   string `json:"AssessedCost"`
		Cost           string `json:"Cost"`
		CostRedelivery string `json:"CostRedelivery"`
		TimezoneInfo   `json:"TZoneInfo"`
		CostPack       string `json:"CostPack"`
	}
)

// GetDocumentPrice Розрахувати вартість послуг
//
// Метод «getDocumentPrice», працює в моделі «InternetDocument», цей метод дозволяє розрахувати вартості доставки
// вантажу.
//
// Метод дозволять прораховувати не тільки відправлення вантажів, але і шин та дисків, палет, а також документів.
func (c Client) GetDocumentPrice(req DocumentPriceRequest) (Response[DocumentPrice], error) {
	return request[DocumentPrice](c, InternetDocumentModel, "getDocumentPrice", req)
}

type (
	DocumentDeliveryDateRequest struct {
		DateTime      string `json:"DateTime"`
		ServiceType   string `json:"ServiceType"`
		CitySender    string `json:"CitySender"`
		CityRecipient string `json:"CityRecipient"`
	}

	DeliveryDate struct {
		Date         string `json:"date"`
		TimezoneType string `json:"timezone_type"`
		Timezone     string `json:"timezone"`
	}

	DocumentDeliveryDate struct {
		DeliveryDate `json:"DeliveryDate"`
	}
)

// GetDocumentDeliveryDate Прогноз дати доставки вантажу
//
// Метод «getDocumentDeliveryDate», працює в моделі «InternetDocument», цей метод у відповіді виведе орієнтовну дату
// доставки.
//
// Якщо в ЕН будуть такі послуги як: зворотна доставка, доставка щодня та інші, їх також потрібно вказувати у запиті як
// обов'язкові елементи запиту для розрахунку орієнтовної дати доставки.
func (c Client) GetDocumentDeliveryDate(req DocumentDeliveryDateRequest) (Response[DocumentDeliveryDate], error) {
	return request[DocumentDeliveryDate](c, InternetDocumentModel, "getDocumentDeliveryDate", req)
}

type (
	InternetDocument struct {
		Ref                   string `json:"Ref"`
		CostOnSite            string `json:"CostOnSite"`
		EstimatedDeliveryDate string `json:"EstimatedDeliveryDate"`
		IntDocNumber          string `json:"IntDocNumber"`
		TypeDocument          string `json:"TypeDocument"`
	}

	Services struct {
		Attorney                string `json:"Attorney"`
		WaybillNewPostWithStamp string `json:"WaybillNewPostWithStamp"`
		UserActions             string `json:"UserActions"`
	}

	BackwardDeliveryData struct {
		PayerType string `json:"PayerType"`
		CargoType string `json:"CargoType"`
		Services  `json:"Services"`
	}

	OptionSeat struct {
		VolumetricVolume string `json:"volumetricVolume"`
		VolumetricWidth  string `json:"volumetricWidth"`
		VolumetricLength string `json:"volumetricLength"`
		VolumetricHeight string `json:"volumetricHeight"`
		Weight           string `json:"weight"`
		PackRef          string `json:"packRef"`
	}

	InternetDocumentRequest struct {
		SenderWarehouseIndex    string                 `json:"SenderWarehouseIndex"`
		RecipientWarehouseIndex string                 `json:"RecipientWarehouseIndex"`
		PayerType               string                 `json:"PayerType"`
		PaymentMethod           string                 `json:"PaymentMethod"`
		DateTime                string                 `json:"DateTime"`
		CargoType               string                 `json:"CargoType"`
		VolumeGeneral           string                 `json:"VolumeGeneral"`
		Weight                  string                 `json:"Weight"`
		ServiceType             string                 `json:"ServiceType"`
		SeatsAmount             string                 `json:"SeatsAmount"`
		Description             string                 `json:"Description"`
		Cost                    string                 `json:"Cost"`
		CitySender              string                 `json:"CitySender"`
		Sender                  string                 `json:"Sender"`
		SenderAddress           string                 `json:"SenderAddress"`
		ContactSender           string                 `json:"ContactSender"`
		SendersPhone            string                 `json:"SendersPhone"`
		CityRecipient           string                 `json:"CityRecipient"`
		Recipient               string                 `json:"Recipient"`
		RecipientAddress        string                 `json:"RecipientAddress"`
		ContactRecipient        string                 `json:"ContactRecipient"`
		RecipientsPhone         string                 `json:"RecipientsPhone"`
		RedBoxBarcode           string                 `json:"RedBoxBarcode"`
		NewAddress              string                 `json:"NewAddress"`
		BackwardDeliveryData    []BackwardDeliveryData `json:"BackwardDeliveryData"`
		OptionsSeat             []OptionSeat           `json:"OptionsSeat"`
		RecipientAddressNote    string                 `json:"RecipientAddressNote"`
		RecipientCityName       string                 `json:"RecipientCityName"`
		RecipientArea           string                 `json:"RecipientArea"`
		RecipientAreaRegions    string                 `json:"RecipientAreaRegions"`
		RecipientAddressName    string                 `json:"RecipientAddressName"`
		RecipientHouse          string                 `json:"RecipientHouse"`
		RecipientFlat           string                 `json:"RecipientFlat"`
		RecipientName           string                 `json:"RecipientName"`
		RecipientType           string                 `json:"RecipientType"`
		SettlementType          string                 `json:"SettlementType"`
		OwnershipForm           string                 `json:"OwnershipForm"`
		RecipientContactName    string                 `json:"RecipientContactName"`
		EDRPOU                  string                 `json:"EDRPOU"`
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
func (c Client) SaveInternetDocument(req InternetDocumentRequest) (Response[InternetDocument], error) {
	return request[InternetDocument](c, InternetDocumentModel, "save", req)
}

// UpdateInternetDocument Редагувати експрес-накладну
//
// Метод «update», працює в моделі «InternetDocument», цей метод дозволяє редагувати експрес-накладну
// (інтернет-документу). У цьому методі обов'язкові до заповнення всі параметри, після оновлення замінюються лише ті, що
// були змінені.
func (c Client) UpdateInternetDocument(req InternetDocumentRequest) (Response[InternetDocument], error) {
	return request[InternetDocument](c, InternetDocumentModel, "update", req)
}

type (
	DocumentListRequest struct {
		DateTimeFrom string `json:"DateTimeFrom"`
		DateTimeTo   string `json:"DateTimeTo"`
		Page         int    `json:"Page"`
		GetFullList  string `json:"GetFullList"`
		DateTime     string `json:"DateTime"`
	}

	Document struct {
		Ref                     string `json:"Ref"`
		DateTime                string `json:"DateTime"`
		PreferredDeliveryDate   string `json:"PreferredDeliveryDate"`
		RecipientDateTime       string `json:"RecipientDateTime"`
		EWDateCreated           string `json:"EWDateCreated"`
		Weight                  string `json:"Weight"`
		SeatsAmount             string `json:"SeatsAmount"`
		IntDocNumber            string `json:"IntDocNumber"`
		Cost                    string `json:"Cost"`
		CitySender              string `json:"CitySender"`
		CityRecipient           string `json:"CityRecipient"`
		SenderAddress           string `json:"SenderAddress"`
		RecipientAddress        string `json:"RecipientAddress"`
		CostOnSite              string `json:"CostOnSite"`
		PayerType               string `json:"PayerType"`
		PaymentMethod           string `json:"PaymentMethod"`
		AfterpaymentOnGoodsCost string `json:"AfterpaymentOnGoodsCost"`
		PackingNumber           string `json:"PackingNumber"`
		RejectionReason         string `json:"RejectionReason"`
		StateId                 string `json:"StateId"`
		StateName               string `json:"StateName"`
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
func (c Client) GetDocumentList(req DocumentListRequest) (Response[Document], error) {
	return request[Document](c, InternetDocumentModel, "getDocumentList", req)
}

type DocumentRefs struct {
	DocumentRefs []string `json:"DocumentRefs"`
}

// DeleteInternetDocument Видалити експрес-накладну
//
// Метод «delete», працює в моделі «InternetDocument», цей метод необхідний для видалення експрес-накладної
// (інтернет-документа)
func (c Client) DeleteInternetDocument(documentRefs DocumentRefs) (Response[Ref], error) {
	return request[Ref](c, InternetDocumentModel, "delete", documentRefs)
}

type (
	ReportRequest struct {
		DocumentRefs []string `json:"DocumentRefs"`
		Type         string   `json:"Type"`
		DateTime     string   `json:"DateTime"`
	}

	Report struct {
		Ref                     string `json:"Ref"`
		DateTime                string `json:"DateTime"`
		PreferredDeliveryDate   string `json:"PreferredDeliveryDate"`
		Weight                  string `json:"Weight"`
		SeatsAmount             string `json:"SeatsAmount"`
		IntDocNumber            string `json:"IntDocNumber"`
		Cost                    string `json:"Cost"`
		CitySender              string `json:"CitySender"`
		CityRecipient           string `json:"CityRecipient"`
		State                   string `json:"State"`
		SenderAddress           string `json:"SenderAddress"`
		RecipientAddress        string `json:"RecipientAddress"`
		CostOnSite              string `json:"CostOnSite"`
		PayerType               string `json:"PayerType"`
		PaymentMethod           string `json:"PaymentMethod"`
		AfterpaymentOnGoodsCost string `json:"AfterpaymentOnGoodsCost"`
		PackingNumber           string `json:"PackingNumber"`
	}
)

// GenerateReport Формування запиту для отримання повного звіту за накладними
//
// Метод «generateReport», працює в моделі «InternetDocument», цей метод формує запит для отримання повного звіту про
// накладні
func (c Client) GenerateReport(req ReportRequest) (Response[Report], error) {
	return request[Report](c, InternetDocumentModel, "generateReport", req)
}
