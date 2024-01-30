package novapost

type (
	TimeIntervalRequest struct {
		RecipientCityRef string `json:"RecipientCityRef"`
		DateTime         string `json:"DateTime"`
	}

	TimeInterval struct {
		Number string `json:"Number"`
		Start  string `json:"Start"`
		End    string `json:"End"`
	}
)

// GetTimeIntervals Види часових інтервалів
//
// Метод «getTimeIntervals», працює в моделі «Common», цей метод необхідний для отримання списку часових інтервалів (для
// замовлення послуги «Часові інтервали»). Для оновлення даних, довідник потрібно завантажувати один раз на місяць.
func (c Client) GetTimeIntervals(req TimeIntervalRequest) (Response[TimeInterval], error) {
	return request[TimeInterval](c, CommonModel, "getTimeIntervals", req)
}

// GetCargoTypes Види вантажу
//
// Метод «getCargoTypes», працює в моделі «Common», цей метод необхідний для завантаження списку типів вантажу
// українською мовою. Для оновлення даних, довідник потрібно завантажувати один раз на місяць.
func (c Client) GetCargoTypes() (Response[RefDescription], error) {
	return request[RefDescription](c, CommonModel, "getCargoTypes", nil)
}

// GetBackwardDeliveryCargoTypes Види зворотної доставки вантажу
//
// Метод «getBackwardDeliveryCargoTypes», працює в моделі «Common», цей метод необхідний для відображення списку видів
// зворотної доставки вантажу українською та російською мовами. Для оновлення даних, довідник потрібно завантажувати
// один раз на місяць.
func (c Client) GetBackwardDeliveryCargoTypes() (Response[RefDescription], error) {
	return request[RefDescription](c, CommonModel, "getBackwardDeliveryCargoTypes", nil)
}

type Pallet struct {
	Ref         string `json:"Ref"`
	Description string `json:"Description"`
	Weight      string `json:"Weight"`
}

// GetPalletsList Види палет
//
// Метод «getPalletsList», працює в моделі «Common», цей метод необхідний для отримання списку видів палет. Для
// оновлення даних, довідник потрібно завантажувати один раз на місяць.
func (c Client) GetPalletsList() (Response[Pallet], error) {
	return request[Pallet](c, CommonModel, "getPalletsList", nil)
}

// GetTypesOfPayersForRedelivery Види платників зворотної доставки
//
// Метод «getTypesOfPayersForRedelivery», працює в моделі «Common», цей метод необхідний для завантаження списку видів
// платників послуги зворотної доставки українською та англійською мовами: Sender, Recipient.
//
// Для оновлення даних, довідник потрібно завантажувати один раз на місяць.
func (c Client) GetTypesOfPayersForRedelivery() (Response[RefDescription], error) {
	return request[RefDescription](c, CommonModel, "getTypesOfPayersForRedelivery", nil)
}

type (
	PackRequest struct {
		Lengthstring           string `json:"Lengthstring"`
		Widthstring            string `json:"Widthstring"`
		Heightstring           string `json:"Heightstring"`
		VolumetricWeightstring string `json:"VolumetricWeightstring"`
		TypeOfPackingstring    string `json:"TypeOfPackingstring"`
	}

	Pack struct {
		Ref               string `json:"Ref"`
		Description       string `json:"Description"`
		Length            string `json:"Length"`
		Width             string `json:"Width"`
		Height            string `json:"Height"`
		VolumetricWeight  string `json:"VolumetricWeight"`
		TypeOfPacking     string `json:"TypeOfPacking"`
		PackagingForPlace string `json:"PackagingForPlace"`
	}
)

// GetPackList Види упаковки
//
// Метод «getPackList», працює в моделі «Common», цей метод необхідний для завантаження видів упаковки вантажу
// українською або російською мовами. Для оновлення даних, довідник потрібно завантажувати один раз на місяць.
func (c Client) GetPackList(req PackRequest) (Response[Pack], error) {
	return request[Pack](c, CommonModel, "getPackList", req)
}

type TireWheel struct {
	Ref             string `json:"Ref"`
	Description     string `json:"Description"`
	Weight          string `json:"Weight"`
	DescriptionType string `json:"DescriptionType"`
}

// GetTiresWheelsList Види шин і дисків
//
// Метод «getTiresWheelsList», працює в моделі «Common», цей метод дозволяє завантажити список шин і дисків
// (використовується, якщо вид вантажу шини-диски) українською та російською мовами.
//
// Необхідно зберігати копію довідників на стороні клієнта та підтримувати її в актуальному стані.
//
// Рекомендується проводити оновлення довідників раз на місяць.
func (c Client) GetTiresWheelsList() (Response[TireWheel], error) {
	return request[TireWheel](c, CommonModel, "getTiresWheelsList", nil)
}

type CargoDescriptionRequest struct {
	FindByString string `json:"FindByString"`
	Page         int    `json:"Page"`
}

// GetCargoDescriptionList Описи вантажу
//
// Метод «getCargoDescriptionList», працює в моделі «Common», цей метод дозволяє віддати опис вантажу українською та
// російською мовами. "FindByString": "абажур", - також доступний пошук по рядках, не обов'язковий параметр.
//
// Необхідно зберігати копію довідників на стороні клієнта та підтримувати її в актуальному стані.
//
// Рекомендується проводити оновлення довідників раз на місяць.
func (c Client) GetCargoDescriptionList(req CargoDescriptionRequest) (Response[RefDescription], error) {
	return request[RefDescription](c, CommonModel, "getCargoDescriptionList", req)
}

type MessageCodeText struct {
	MessageCode        string `json:"MessageCode"`
	MessageText        string `json:"MessageText"`
	MessageDescription string `json:"MessageDescriptionUA"`
}

// GetMessageCodeText Перелік помилок
//
// Метод «getMessageCodeText», працює в моделі «CommonGeneral», цей метод необхідний для завантаження довідника з описом
// переліку помилок.
//
// Метод постійно поповнюється новим описом трьома мовами.
//
// Для оновлення даних, довідник потрібно завантажувати один раз на місяць.
func (c Client) GetMessageCodeText() (Response[MessageCodeText], error) {
	return request[MessageCodeText](c, CommonGeneralModel, "getMessageCodeText", nil)
}

// GetServiceTypes Технології доставки
//
// Метод «getServiceTypes», працює в моделі «Common», цей метод дозволяє завантажити список типів технологій доставки:
// «склад-склад», «двері-двері», «склад-двері», «двері-склад» українською та російською мовами.
//
// Необхідно зберігати копію довідників на стороні клієнта та підтримувати її в актуальному стані.
//
// Рекомендується проводити оновлення довідників раз на місяць.
func (c Client) GetServiceTypes() (Response[RefDescription], error) {
	return request[RefDescription](c, CommonModel, "getServiceTypes", nil)
}

type OwnershipForm struct {
	Ref         string `json:"Ref"`
	Description string `json:"Description"`
	FullName    string `json:"FullName"`
}

// GetOwnershipFormsList Форми власності
//
// Метод «getOwnershipFormsList», працює в моделі «Common», цей метод дозволяє завантажити список форм власності
// українською мовою.
//
// Необхідно зберігати копію довідників на стороні клієнта та підтримувати її в актуальному стані.
//
// Рекомендується проводити оновлення довідників раз на місяць.
func (c Client) GetOwnershipFormsList() (Response[OwnershipForm], error) {
	return request[OwnershipForm](c, CommonModel, "getOwnershipFormsList", nil)
}
