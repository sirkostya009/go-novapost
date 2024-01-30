package novapost

type PossibilityCreateReturn struct {
	NonCash       bool   `json:"NonCash"`
	City          string `json:"City"`
	Counterparty  string `json:"Counterparty"`
	ContactPerson string `json:"ContactPerson"`
	Address       string `json:"Address"`
	Phone         string `json:"Phone"`
	Ref           string `json:"Ref"`
}

// CheckPossibilityCreateReturn Перевірка можливості створення заявки на повернення
//
// Метод «CheckPossibilityCreateReturn», працює в моделі «AdditionalService», цей метод дозволяє перевірити можливість
// створення заявки на повернення. Метод доступний лише клієнтам відправників. У разі успіху повертаються блоки адрес,
// які можна використовуватиме створення заявки.
func (c Client) CheckPossibilityCreateReturn(req Number) (Response[PossibilityCreateReturn], error) {
	return request[PossibilityCreateReturn](c, AdditionalServiceModel, "CheckPossibilityCreateReturn", req)
}

// GetReturnReasons Отримання списку причин повернення
//
// Метод «getReturnReasons», працює в моделі «AdditionalService», цей метод дозволяє отримати список доступних причин
// повернення.
func (c Client) GetReturnReasons() (Response[RefDescription], error) {
	return request[RefDescription](c, AdditionalServiceModel, "getReturnReasons", nil)
}

type (
	ReasonRef struct {
		ReasonRef string `json:"ReasonRef"`
	}

	ReturnReasonSubtype struct {
		Ref         string `json:"Ref"`
		Description string `json:"Description"`
		ReasonRef   string `json:"ReasonRef"`
	}
)

// GetReturnReasonSubtypes Отримання списку підтипів причини повернення
//
// Метод «getReturnReasonsSubtypes», працює в моделі «AdditionalService», цей метод дозволяє отримати список підтипів
// доступних причин повернення.
func (c Client) GetReturnReasonSubtypes(req ReasonRef) (Response[ReturnReasonSubtype], error) {
	return request[ReturnReasonSubtype](c, AdditionalServiceModel, "getReturnReasonSubtypes", req)
}

type NewReturnOrderRequest struct {
	IntDocNumber              string `json:"IntDocNumber"`
	PaymentMethod             string `json:"PaymentMethod"`
	Reason                    string `json:"Reason"`
	SubtypeReason             string `json:"SubtypeReason"`
	Note                      string `json:"Note"`
	OrderType                 string `json:"OrderType"`
	SenderContactName         string `json:"SenderContactName"`
	SenderPhone               string `json:"SenderPhone"`
	Recipient                 string `json:"Recipient"`
	RecipientContactName      string `json:"RecipientContactName"`
	RecipientPhone            string `json:"RecipientPhone"`
	PayerType                 string `json:"PayerType"`
	Customer                  string `json:"Customer"`
	ServiceType               string `json:"ServiceType"`
	RecipientSettlement       string `json:"RecipientSettlement"`
	RecipientSettlementStreet string `json:"RecipientSettlementStreet"`
	RecipientWarehouse        string `json:"RecipientWarehouse"`
	BuildingNumber            string `json:"BuildingNumber"`
	NoteAddressRecipient      string `json:"NoteAddressRecipient"`
	ReturnAddressRef          string `json:"ReturnAddressRef"`
}

// CreateNewReturn Створення заявки на повернення
//
// Використовується для створення заявки на повернення а також зміни даних. Метод доступний лише клієнтам відправників.
func (c Client) CreateNewReturn(req NewReturnOrderRequest) (Response[RefNumber], error) {
	return request[RefNumber](c, AdditionalServiceModel, "save", req)
}

type (
	ReturnOrderRequest struct {
		Number    string `json:"Number"`
		Ref       string `json:"Ref"`
		BeginDate string `json:"BeginDate"`
		EndDate   string `json:"EndDate"`
		Page      int    `json:"Page"`
		Limit     int    `json:"Limit"`
	}

	ReturnOrder struct {
		OrderRef               string `json:"OrderRef"`
		OrderNumber            string `json:"OrderNumber"`
		OrderStatus            string `json:"OrderStatus"`
		DocumentNumber         string `json:"DocumentNumber"`
		CounterpartyRecipient  string `json:"CounterpartyRecipient"`
		ContactPersonRecipient string `json:"ContactPersonRecipient"`
		AddressRecipient       string `json:"AddressRecipient"`
		DeliveryCost           string `json:"DeliveryCost"`
		EstimatedDeliveryDate  string `json:"EstimatedDeliveryDate"`
		ExpressWaybillNumber   string `json:"ExpressWaybillNumber"`
		ExpressWaybillStatus   string `json:"ExpressWaybillStatus"`
	}
)

// GetReturnOrdersList Отримання списку заявок на повернення
//
// Метод «getReturnOrdersList», працює в моделі «AdditionalService», цей метод дозволяє отримати список заявок на
// повернення.
func (c Client) GetReturnOrdersList(req ReturnOrderRequest) (Response[ReturnOrder], error) {
	return request[ReturnOrder](c, AdditionalServiceModel, "getReturnOrdersList", req)
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
func (c Client) DeleteReturnOrder(req Ref) (Response[ReturnOrder], error) {
	return request[ReturnOrder](c, AdditionalServiceModel, "delete", req)
}

type (
	PossibilityChangeEW struct {
		CanChangeSender                     bool   `json:"CanChangeSender"`
		CanChangeRecipient                  bool   `json:"CanChangeRecipient"`
		CanChangePayerTypeOrPaymentMethod   bool   `json:"CanChangePayerTypeOrPaymentMethod"`
		CanChangeBackwardDeliveryDocuments  bool   `json:"CanChangeBackwardDeliveryDocuments"`
		CanChangeBackwardDeliveryMoney      bool   `json:"CanChangeBackwardDeliveryMoney"`
		CanChangeCash2Card                  bool   `json:"CanChangeCash2Card"`
		CanChangeBackwardDeliveryOther      bool   `json:"CanChangeBackwardDeliveryOther"`
		CanChangeAfterpaymentType           bool   `json:"CanChangeAfterpaymentType"`
		CanChangeLiftingOnFloor             bool   `json:"CanChangeLiftingOnFloor"`
		CanChangeLiftingOnFloorWithElevator bool   `json:"CanChangeLiftingOnFloorWithElevator"`
		CanChangeFillingWarranty            bool   `json:"CanChangeFillingWarranty"`
		SenderCounterparty                  string `json:"SenderCounterparty"`
		ContactPersonSender                 string `json:"ContactPersonSender"`
		SenderPhone                         string `json:"SenderPhone"`
		RecipientCounterparty               string `json:"RecipientCounterparty"`
		ContactPersonRecipient              string `json:"ContactPersonRecipient"`
		RecipientPhone                      string `json:"RecipientPhone"`
		PayerType                           string `json:"PayerType"`
		PaymentMethod                       string `json:"PaymentMethod"`
	}

	IntDocNumber struct {
		IntDocNumber string `json:"IntDocNumber"`
	}
)

// CheckPossibilityChangeEW Перевірка можливості створення заявки зі зміни даних
//
// Метод «CheckPossibilityChangeEW», працює в моделі «AdditionalService», цей метод дозволяє перевірити можливість
// створення заявки щодо зміни даних. У разі успіху повертаються необхідні параметри створення заявки.
func (c Client) CheckPossibilityChangeEW(req IntDocNumber) (Response[PossibilityChangeEW], error) {
	return request[PossibilityChangeEW](c, AdditionalServiceModel, "CheckPossibilityChangeEW", req)
}

type (
	ChangeEWRequest struct {
		Number    string `json:"Number"`
		Ref       string `json:"Ref"`
		BeginDate string `json:"BeginDate"`
		EndDate   string `json:"EndDate"`
		Page      int    `json:"Page"`
		Limit     int    `json:"Limit"`
	}

	ChangeEW struct {
		OrderRef                            string `json:"OrderRef"`
		OrderNumber                         string `json:"OrderNumber"`
		OrderStatus                         string `json:"OrderStatus"`
		DocumentNumber                      string `json:"DocumentNumber"`
		DateTime                            string `json:"DateTime"`
		BeforeChangeSenderCounterparty      string `json:"BeforeChangeSenderCounterparty"`
		AfterChangeChangeSenderCounterparty string `json:"AfterChangeChangeSenderCounterparty"`
		Cost                                string `json:"Cost"`
		BeforeChangeSenderPhone             string `json:"BeforeChangeSenderPhone"`
		AfterChangeSenderPhone              string `json:"AfterChangeSenderPhone"`
	}
)

// GetChangeEWOrdersList Отримання списку заявок
//
// Метод «getChangeEWOrdersList», працює в моделі «AdditionalService», цей метод дозволяє отримати список заявок щодо
// зміни даних.
func (c Client) GetChangeEWOrdersList(req ChangeEWRequest) (Response[ChangeEW], error) {
	return request[ChangeEW](c, AdditionalServiceModel, "getChangeEWOrdersList", req)
}

type (
	Number struct {
		Number string `json:"Number"`
	}

	PossibilityRedirecting struct {
		Ref                              string `json:"Ref"`
		Number                           string `json:"Number"`
		PayerType                        string `json:"PayerType"`
		PaymentMethod                    string `json:"PaymentMethod"`
		WarehouseRef                     string `json:"WarehouseRef"`
		WarehouseDescription             string `json:"WarehouseDescription"`
		AddressDescription               string `json:"AddressDescription"`
		StreetDescription                string `json:"StreetDescription"`
		BuildingNumber                   string `json:"BuildingNumber"`
		CityRecipient                    string `json:"CityRecipient"`
		CityRecipientDescription         string `json:"CityRecipientDescription"`
		SettlementRecipient              string `json:"SettlementRecipient"`
		SettlementRecipientDescription   string `json:"SettlementRecipientDescription"`
		SettlementType                   string `json:"SettlementType"`
		CounterpartyRecipientRef         string `json:"CounterpartyRecipientRef"`
		CounterpartyRecipientDescription string `json:"CounterpartyRecipientDescription"`
		RecipientName                    string `json:"RecipientName"`
		PhoneSender                      string `json:"PhoneSender"`
		PhoneRecipient                   string `json:"PhoneRecipient"`
		DocumentWeight                   string `json:"DocumentWeight"`
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
func (c Client) CheckPossibilityForRedirecting(req Number) (Response[PossibilityRedirecting], error) {
	return request[PossibilityRedirecting](c, AdditionalServiceModel, "CheckPossibilityForRedirecting", req)
}

type (
	Redirection struct {
		OrderRef              string `json:"OrderRef"`
		OrderNumber           string `json:"OrderNumber"`
		DateTime              string `json:"DateTime"`
		DocumentNumber        string `json:"DocumentNumber"`
		Note                  string `json:"Note"`
		CityRecipient         string `json:"CityRecipient"`
		RecipientAddress      string `json:"RecipientAddress"`
		CounterpartyRecipient string `json:"CounterpartyRecipient"`
		RecipientName         string `json:"RecipientName"`
		PhoneRecipient        string `json:"PhoneRecipient"`
		PayerType             string `json:"PayerType"`
		DeliveryCost          string `json:"DeliveryCost"`
		EstimatedDeliveryDate string `json:"EstimatedDeliveryDate"`
		ExpressWaybillNumber  string `json:"ExpressWaybillNumber"`
		ExpressWaybillStatus  string `json:"ExpressWaybillStatus"`
	}

	RedirectionRequest struct {
		Number    string `json:"Number"`
		Ref       string `json:"Ref"`
		BeginDate string `json:"BeginDate"`
		EndDate   string `json:"EndDate"`
		Page      int    `json:"Page"`
		Limit     int    `json:"Limit"`
	}
)

// GetRedirectionOrdersList Отримання списку заявок
//
// Метод «getRedirectionOrdersList», працює в моделі «AdditionalService», цей метод дозволяє отримати список заявок
// переадресації відправлень.
func (c Client) GetRedirectionOrdersList(req RedirectionRequest) (Response[Redirection], error) {
	return request[Redirection](c, AdditionalServiceModel, "getRedirectionOrdersList", req)
}
