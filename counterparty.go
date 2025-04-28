package novapost

type (
	CounterpartyRequest struct {
		FirstName            string
		MiddleName           string
		LastName             string
		Phone                string
		Email                string `json:",omitempty" xml:",omitempty"`
		CounterpartyType     string
		EDRPOU               string
		CounterpartyProperty string
		CityRef              string
	}

	Counterparty struct {
		Ref                      string
		Description              string
		FirstName                string
		MiddleName               string
		LastName                 string
		Counterparty             string
		OwnershipForm            string
		OwnershipFormDescription string
		EDRPOU                   string
		CounterpartyType         string
		ContactPerson            Response[ContactPerson]
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
func (c *Client) SaveCounterparty(req CounterpartyRequest) (*Response[Counterparty], error) {
	return RawRequest[Counterparty](c, CounterpartyModel, "save", req)
}

// UpdateCounterparty Оновити дані Контрагента
//
// Метод «update», працює в моделі «Counterparty», цей метод використовується для оновлення/внесення змін в дані
// контрагента отримувача.
func (c *Client) UpdateCounterparty(req CounterpartyRequest) (*Response[Counterparty], error) {
	return RawRequest[Counterparty](c, CounterpartyModel, "update", req)
}

// DeleteCounterparty Видалити Контрагента одержувача
//
// Метод «delete», працює в моделі «Counterparty», цей метод використовується для видалення контрагента отримувача. Для
// видалення контрагента відправника необхідно звертатися до Вашого менеджера
func (c *Client) DeleteCounterparty(ref Ref) (*Response[Ref], error) {
	return RawRequest[Ref](c, CounterpartyModel, "delete", ref)
}

type CounterpartyAddressRequest struct {
	Ref                  string
	CounterpartyProperty string
}

// GetCounterpartyAddresses Завантажити список адрес Контрагентів
//
// Метод «getCounterpartyAddresses», працює в моделі «Counterparty», цей метод завантажує список контрагентів
// відправників/отримувачів.
//
// Необхідно зберігати копію довідника на стороні клієнта та підтримувати його в актуальному стані.
//
// Рекомендовано проводити оновлення довідників раз на місяць.
func (c *Client) GetCounterpartyAddresses(req CounterpartyAddressRequest) (*Response[RefDescription], error) {
	return RawRequest[RefDescription](c, CounterpartyModel, "getCounterpartyAddresses", req)
}

type (
	DebtorParams struct {
		AgreementId   int `json:",string"`
		BlockedStatus int `json:",string"`
		PaymTermId    string
		PastDueDebts  float64 `json:",string"`
	}

	Vbf struct {
		Vbf int `xml:"vbf" json:"vbf"`
	}

	CounterpartyOptions struct {
		DebtorParams                      []DebtorParams `xml:"debtorParams>item"`
		PrintMarkingAllowedTypes          Vbf
		AddressDocumentDelivery           bool
		AfterpaymentType                  bool
		BackwardDeliverySubtypesDocuments bool
		BlockInternationalSenderLKK       bool
		CalculationByFactualWeight        bool
		CanAfterpaymentOnGoodsCost        bool
		CanCreditDocuments                bool
		CanEWTransporter                  bool
		CanForwardingService              bool
		CanNonCashPayment                 bool
		CanPayTheThirdPerson              bool
		CanSameDayDelivery                bool
		CanSameDayDeliveryStandart        bool
		CanSentFromPostomat               bool
		CanSignedDocuments                bool
		CarCall                           bool
		CreditDocuments                   bool
		CustomerReturn                    bool
		DayCustomerReturn                 bool
		DeliveryByHand                    bool
		DescentFromFloor                  bool
		FillingWarranty                   bool
		HaveMoneyWallets                  bool
		HideDeliveryCost                  bool
		InternationalDelivery             bool
		InternationalDeliveryServiceType  bool
		LoyaltyProgram                    bool
		MainCounterparty                  bool
		PickupService                     bool
		SecurePayment                     bool
		Services                          bool
		ShowDeliveryByHand                bool
		SignedDocuments                   bool
		TransferPricingConditions         bool
	}
)

// GetCounterpartyOptions Завантажити параметри Контрагента
//
// Метод «getCounterpartyOptions», працює в моделі «Counterparty», цей метод використовується для отримання параметрів
// контрагента відправника в розрізі можливостей замовлення додаткових послуг, з розділу: «Формування запиту на створення
// ЕН із додатковими послугами» Замовити посслуги можливо через особистого менеджера.
func (c *Client) GetCounterpartyOptions(ref Ref) (*Response[CounterpartyOptions], error) {
	return RawRequest[CounterpartyOptions](c, CounterpartyModel, "getCounterpartyOptions", ref)
}

type CounterpartyContactPersonRequest struct {
	Ref  string
	Page int `json:",string,omitempty" xml:"omitempty"`
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
func (c *Client) GetCounterpartyContactPersons(req CounterpartyContactPersonRequest) (*Response[ContactPerson], error) {
	return RawRequest[ContactPerson](c, CounterpartyModel, "getCounterpartyContactPersons", req)
}

type CounterpartiesRequest struct {
	CounterpartyProperty string
	Page                 int`json:",string,omitempty" xml:"omitempty"`
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
func (c *Client) GetCounterparties(req CounterpartiesRequest) (*Response[Counterparty], error) {
	return RawRequest[Counterparty](c, CounterpartyModel, "getCounterparties", req)
}
