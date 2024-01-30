package novapost

type (
	CounterpartyRequest struct {
		FirstName            string `json:"FirstName"`
		MiddleName           string `json:"MiddleName"`
		LastName             string `json:"LastName"`
		Phone                string `json:"Phone"`
		Email                string `json:"Email"`
		CounterpartyType     string `json:"CounterpartyType"`
		EDRPOU               string `json:"EDRPOU"`
		CounterpartyProperty string `json:"CounterpartyProperty"`
		CityRef              string `json:"CityRef"`
	}

	Counterparty struct {
		Ref                      string                  `json:"Ref"`
		Description              string                  `json:"Description"`
		FirstName                string                  `json:"FirstName"`
		MiddleName               string                  `json:"MiddleName"`
		LastName                 string                  `json:"LastName"`
		Counterparty             string                  `json:"Counterparty"`
		OwnershipForm            string                  `json:"OwnershipForm"`
		OwnershipFormDescription string                  `json:"OwnershipFormDescription"`
		EDRPOU                   string                  `json:"EDRPOU"`
		CounterpartyType         string                  `json:"CounterpartyType"`
		ContactPerson            Response[ContactPerson] `json:"ContactPerson"`
	}
)

// SaveCounterparty Створити Контрагента
//
// Приватна особа: https://developers.novaposhta.ua/view/model/a28f4b04-8512-11ec-8ced-005056b2dbe1/method/0ae5dd75-8a5f-11ec-8ced-005056b2dbe1
//
// Третя особа: https://developers.novaposhta.ua/view/model/a28f4b04-8512-11ec-8ced-005056b2dbe1/method/b0fdf818-8a8e-11ec-8ced-005056b2dbe1
//
// Організація: https://developers.novaposhta.ua/view/model/a28f4b04-8512-11ec-8ced-005056b2dbe1/method/bc3c44c7-8a8a-11ec-8ced-005056b2dbe1
//
// Дані по отримувачу вносяться виключно Українською мовою.
//
// Рекомендовано проводити оновлення довідника один раз на місяць.
func (c Client) SaveCounterparty(req CounterpartyRequest) (Response[Counterparty], error) {
	return request[Counterparty](c, CounterpartyModel, "save", req)
}

// UpdateCounterparty Оновити дані Контрагента
//
// Метод «update», працює в моделі «Counterparty», цей метод використовується для оновлення/внесення змін в дані
// контрагента отримувача.
func (c Client) UpdateCounterparty(req CounterpartyRequest) (Response[Counterparty], error) {
	return request[Counterparty](c, CounterpartyModel, "update", req)
}

// DeleteCounterparty Видалити Контрагента одержувача
//
// Метод «delete», працює в моделі «Counterparty», цей метод використовується для видалення контрагента отримувача. Для
// видалення контрагента відправника необхідно звертатися до Вашого менеджера
func (c Client) DeleteCounterparty(ref Ref) (Response[Ref], error) {
	return request[Ref](c, CounterpartyModel, "delete", ref)
}

type CounterpartyAddressRequest struct {
	Ref                  string `json:"Ref"`
	CounterpartyProperty string `json:"CounterpartyProperty"`
}

// GetCounterpartyAddresses Завантажити список адрес Контрагентів
//
// Метод «getCounterpartyAddresses», працює в моделі «Counterparty», цей метод завантажує список контрагентів
// відправників/отримувачів.
//
// Необхідно зберігати копію довідника на стороні клієнта та підтримувати його в актуальному стані.
//
// Рекомендовано проводити оновлення довідників раз на місяць.
func (c Client) GetCounterpartyAddresses(req CounterpartyAddressRequest) (Response[RefDescription], error) {
	return request[RefDescription](c, CounterpartyModel, "getCounterpartyAddresses", req)
}

type (
	DebtorParams struct {
		AgreementId   string `json:"AgreementId"`
		PaymTermId    string `json:"PaymTermId"`
		PastDueDebts  string `json:"PastDueDebts"`
		BlockedStatus string `json:"BlockedStatus"`
	}

	Vbf struct {
		Vbf int `json:"vbf"`
	}

	CounterpartyOptions struct {
		AddressDocumentDelivery           bool           `json:"AddressDocumentDelivery"`
		AfterpaymentType                  bool           `json:"AfterpaymentType"`
		BackwardDeliverySubtypesDocuments bool           `json:"BackwardDeliverySubtypesDocuments"`
		BlockInternationalSenderLKK       bool           `json:"BlockInternationalSenderLKK"`
		CalculationByFactualWeight        bool           `json:"CalculationByFactualWeight"`
		CanAfterpaymentOnGoodsCost        bool           `json:"CanAfterpaymentOnGoodsCost"`
		CanCreditDocuments                bool           `json:"CanCreditDocuments"`
		CanEWTransporter                  bool           `json:"CanEWTransporter"`
		CanForwardingService              bool           `json:"CanForwardingService"`
		CanNonCashPayment                 bool           `json:"CanNonCashPayment"`
		CanPayTheThirdPerson              bool           `json:"CanPayTheThirdPerson"`
		CanSameDayDelivery                bool           `json:"CanSameDayDelivery"`
		CanSameDayDeliveryStandart        bool           `json:"CanSameDayDeliveryStandart"`
		CanSentFromPostomat               bool           `json:"CanSentFromPostomat"`
		CanSignedDocuments                bool           `json:"CanSignedDocuments"`
		CarCall                           bool           `json:"CarCall"`
		CreditDocuments                   bool           `json:"CreditDocuments"`
		CustomerReturn                    bool           `json:"CustomerReturn"`
		DayCustomerReturn                 bool           `json:"DayCustomerReturn"`
		DebtorParams                      []DebtorParams `json:"DebtorParams"`
		DeliveryByHand                    bool           `json:"DeliveryByHand"`
		DescentFromFloor                  bool           `json:"DescentFromFloor"`
		FillingWarranty                   bool           `json:"FillingWarranty"`
		HaveMoneyWallets                  bool           `json:"HaveMoneyWallets"`
		HideDeliveryCost                  bool           `json:"HideDeliveryCost"`
		InternationalDelivery             bool           `json:"InternationalDelivery"`
		InternationalDeliveryServiceType  bool           `json:"InternationalDeliveryServiceType"`
		LoyaltyProgram                    bool           `json:"LoyaltyProgram"`
		MainCounterparty                  bool           `json:"MainCounterparty"`
		PickupService                     bool           `json:"PickupService"`
		PrintMarkingAllowedTypes          Vbf            `json:"PrintMarkingAllowedTypes"`
		SecurePayment                     bool           `json:"SecurePayment"`
		Services                          bool           `json:"Services"`
		ShowDeliveryByHand                bool           `json:"ShowDeliveryByHand"`
		SignedDocuments                   bool           `json:"SignedDocuments"`
		TransferPricingConditions         bool           `json:"TransferPricingConditions"`
	}
)

// GetCounterpartyOptions Завантажити параметри Контрагента
//
// Метод «getCounterpartyOptions», працює в моделі «Counterparty», цей метод використовується для отримання параметрів
// контрагента відправника в розрізі можливостей замовлення додаткових послуг, з розділу: «Формування запиту на створення
// ЕН із додатковими послугами» Замовити посслуги можливо через особистого менеджера.
func (c Client) GetCounterpartyOptions(ref Ref) (Response[CounterpartyOptions], error) {
	return request[CounterpartyOptions](c, CounterpartyModel, "getCounterpartyOptions", ref)
}

type CounterpartyContactPersonRequest struct {
	Ref  string `json:"Ref"`
	Page int    `json:"Page"`
}

// GetCounterpartyContactPersons Завантажити список контактних осіб Контрагента
//
// Метод «getCounterpartyContactPersons», працює в моделі «Counterparty», цей метод завантажує список контактних осіб
// контрагентів.
//
// Якщо в методі більше ніж 100 контактних осіб контрагентів, необхідно використовувати параметр «Page»
//
// Копію довідника необхідно зберігати на стороні клієнта та підтримувати його в актуальному стані.
//
// Рекомендовано проводити оновлення довідників раз на день.
func (c Client) GetCounterpartyContactPersons(req CounterpartyContactPersonRequest) (Response[ContactPerson], error) {
	return request[ContactPerson](c, CounterpartyModel, "getCounterpartyContactPersons", req)
}

type CounterpartiesRequest struct {
	CounterpartyProperty string `json:"CounterpartyProperty"`
	Page                 int    `json:"Page"`
}

// GetCounterparties Завантажити список контрагентів відправників / одержувачів / третя особа
// Метод «getCounterparties», працює в моделі «Counterparty», цей метод завантажує список контрагентів відправників,
// отримувачі та третіх осіб.
//
// Якщо в методі більше ніж 100 контрагентів відправників або отримувачів, необхідно використовувати параметр «Page».
//
// Якщо в запит додати параметр «FindByString» (пошук по рядкам) та у властивостях вказати назву контрагента (Талісман),
// який необхідно знайти, то отримаємо запит за допомогою якого у довіднику знайдеться необхідний контрагент.
//
// Копію довідника необхідно зберігати на сторні клієнта та підтримувати його в актуальному стані.
//
// Рекомендовано проводити оновлення довідників раз на день.
func (c Client) GetCounterparties(req CounterpartiesRequest) (Response[Counterparty], error) {
	return request[Counterparty](c, CounterpartyModel, "getCounterparties", req)
}
