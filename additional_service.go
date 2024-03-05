package novapost

type PossibilityCreateReturn struct {
	NonCash       bool
	City          string
	Counterparty  string
	ContactPerson string
	Address       string
	Phone         string
	Ref           string
}

// CheckPossibilityCreateReturn Перевірка можливості створення заявки на повернення
//
// Метод «CheckPossibilityCreateReturn», працює в моделі «AdditionalService», цей метод дозволяє перевірити можливість
// створення заявки на повернення. Метод доступний лише клієнтам відправників. У разі успіху повертаються блоки адрес,
// які можна використовуватиме створення заявки.
func (c *Client) CheckPossibilityCreateReturn(n Number) (*Response[PossibilityCreateReturn], error) {
	return RawRequest[PossibilityCreateReturn](c, AdditionalServiceModel, "CheckPossibilityCreateReturn", n)
}

// GetReturnReasons Отримання списку причин повернення
//
// Метод «getReturnReasons», працює в моделі «AdditionalService», цей метод дозволяє отримати список доступних причин
// повернення.
func (c *Client) GetReturnReasons() (*Response[RefDescription], error) {
	return RawRequest[RefDescription](c, AdditionalServiceModel, "getReturnReasons", nil)
}

type (
	ReasonRef struct {
		ReasonRef string
	}

	ReturnReasonSubtype struct {
		Ref         string
		Description string
		ReasonRef   string
	}
)

// GetReturnReasonSubtypes Отримання списку підтипів причини повернення
//
// Метод «getReturnReasonsSubtypes», працює в моделі «AdditionalService», цей метод дозволяє отримати список підтипів
// доступних причин повернення.
func (c *Client) GetReturnReasonSubtypes(ref ReasonRef) (*Response[ReturnReasonSubtype], error) {
	return RawRequest[ReturnReasonSubtype](c, AdditionalServiceModel, "getReturnReasonSubtypes", ref)
}

type NewReturnOrderRequest struct {
	IntDocNumber              string
	PaymentMethod             string
	Reason                    string
	SubtypeReason             string
	Note                      string
	OrderType                 string
	SenderContactName         string
	SenderPhone               string
	Recipient                 string
	RecipientContactName      string
	RecipientPhone            string
	PayerType                 string
	Customer                  string
	ServiceType               string
	RecipientSettlement       string
	RecipientSettlementStreet string
	RecipientWarehouse        string
	BuildingNumber            string
	NoteAddressRecipient      string
	ReturnAddressRef          string
}

// CreateNewReturn Створення заявки на повернення
//
// Використовується для створення заявки на повернення а також зміни даних. Метод доступний лише клієнтам відправників.
func (c *Client) CreateNewReturn(req NewReturnOrderRequest) (*Response[RefNumber], error) {
	return RawRequest[RefNumber](c, AdditionalServiceModel, "save", req)
}

type (
	ReturnOrderRequest struct {
		Number    string
		Ref       string
		BeginDate string
		EndDate   string
		Page      int
		Limit     int
	}

	ReturnOrder struct {
		OrderRef               string
		OrderNumber            string
		OrderStatus            string
		DocumentNumber         string
		CounterpartyRecipient  string
		ContactPersonRecipient string
		AddressRecipient       string
		DeliveryCost           float64
		EstimatedDeliveryDate  string
		ExpressWaybillNumber   string
		ExpressWaybillStatus   string
	}
)

// GetReturnOrdersList Отримання списку заявок на повернення
//
// Метод «getReturnOrdersList», працює в моделі «AdditionalService», цей метод дозволяє отримати список заявок на
// повернення.
func (c *Client) GetReturnOrdersList(req ReturnOrderRequest) (*Response[ReturnOrder], error) {
	return RawRequest[ReturnOrder](c, AdditionalServiceModel, "getReturnOrdersList", req)
}

// DeleteReturnOrder Видалення заявки
//
// Метод «delete», працює в моделі «AdditionalService», цей метод дозволяє видалити:
//
// - заявку на повернення;
//
// - заявку щодо зміни даних (можна видалити заявку лише зі статусом «Прийнято»);
//
// - заявку переадресації.
func (c *Client) DeleteReturnOrder(ref Ref) (*Response[ReturnOrder], error) {
	return RawRequest[ReturnOrder](c, AdditionalServiceModel, "delete", ref)
}

type (
	PossibilityChangeEW struct {
		CanChangeSender                     bool
		CanChangeRecipient                  bool
		CanChangePayerTypeOrPaymentMethod   bool
		CanChangeBackwardDeliveryDocuments  bool
		CanChangeBackwardDeliveryMoney      bool
		CanChangeCash2Card                  bool
		CanChangeBackwardDeliveryOther      bool
		CanChangeAfterpaymentType           bool
		CanChangeLiftingOnFloor             bool
		CanChangeLiftingOnFloorWithElevator bool
		CanChangeFillingWarranty            bool
		SenderCounterparty                  string
		ContactPersonSender                 string
		SenderPhone                         string
		RecipientCounterparty               string
		ContactPersonRecipient              string
		RecipientPhone                      string
		PayerType                           string
		PaymentMethod                       string
	}

	IntDocNumber struct {
		IntDocNumber string
	}
)

// CheckPossibilityChangeEW Перевірка можливості створення заявки зі зміни даних
//
// Метод «CheckPossibilityChangeEW», працює в моделі «AdditionalService», цей метод дозволяє перевірити можливість
// створення заявки щодо зміни даних. У разі успіху повертаються необхідні параметри створення заявки.
func (c *Client) CheckPossibilityChangeEW(req IntDocNumber) (*Response[PossibilityChangeEW], error) {
	return RawRequest[PossibilityChangeEW](c, AdditionalServiceModel, "CheckPossibilityChangeEW", req)
}

type (
	ChangeEWRequest struct {
		Number    string
		Ref       string
		BeginDate string
		EndDate   string
		Page      int
		Limit     int
	}

	ChangeEW struct {
		OrderRef                            string
		OrderNumber                         string
		OrderStatus                         string
		DocumentNumber                      string
		DateTime                            string
		BeforeChangeSenderCounterparty      string
		AfterChangeChangeSenderCounterparty string
		Cost                                float64
		BeforeChangeSenderPhone             string
		AfterChangeSenderPhone              string
	}
)

// GetChangeEWOrdersList Отримання списку заявок
//
// Метод «getChangeEWOrdersList», працює в моделі «AdditionalService», цей метод дозволяє отримати список заявок щодо
// зміни даних.
func (c *Client) GetChangeEWOrdersList(req ChangeEWRequest) (*Response[ChangeEW], error) {
	return RawRequest[ChangeEW](c, AdditionalServiceModel, "getChangeEWOrdersList", req)
}

type (
	Number struct {
		Number string
	}

	PossibilityRedirecting struct {
		Ref                              string
		Number                           string
		PayerType                        string
		PaymentMethod                    string
		WarehouseRef                     string
		WarehouseDescription             string
		AddressDescription               string
		StreetDescription                string
		BuildingNumber                   string
		CityRecipient                    string
		CityRecipientDescription         string
		SettlementRecipient              string
		SettlementRecipientDescription   string
		SettlementType                   string
		CounterpartyRecipientRef         string
		CounterpartyRecipientDescription string
		RecipientName                    string
		PhoneSender                      string
		PhoneRecipient                   string
		DocumentWeight                   float64
	}
)

// CheckPossibilityForRedirecting Перевірка можливості створення заявки на переадресацію відправлення
//
// Метод «checkPossibilityForRedirecting», працює в моделі «AdditionalServiceGeneral», цей метод дозволяє перевірити
// можливість створення заявки на переадресацію відправлення.
//
// Метод доступний лише клієнтам відправників.
//
// У разі успіху повертаються необхідні параметри створення заявки.
func (c *Client) CheckPossibilityForRedirecting(req Number) (*Response[PossibilityRedirecting], error) {
	return RawRequest[PossibilityRedirecting](c, AdditionalServiceModel, "CheckPossibilityForRedirecting", req)
}

type (
	Redirection struct {
		OrderRef              string
		OrderNumber           string
		DateTime              string
		DocumentNumber        string
		Note                  string
		CityRecipient         string
		RecipientAddress      string
		CounterpartyRecipient string
		RecipientName         string
		PhoneRecipient        string
		PayerType             string
		DeliveryCost          float64
		EstimatedDeliveryDate string
		ExpressWaybillNumber  string
		ExpressWaybillStatus  string
	}

	RedirectionRequest struct {
		Number    string
		Ref       string
		BeginDate string
		EndDate   string
		Page      int
		Limit     int
	}
)

// GetRedirectionOrdersList Отримання списку заявок
//
// Метод «getRedirectionOrdersList», працює в моделі «AdditionalService», цей метод дозволяє отримати список заявок
// переадресації відправлень.
func (c *Client) GetRedirectionOrdersList(req RedirectionRequest) (*Response[Redirection], error) {
	return RawRequest[Redirection](c, AdditionalServiceModel, "getRedirectionOrdersList", req)
}
