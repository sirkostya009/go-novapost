package novapost

type (
	StatusDocumentRequest struct {
		DocumentNumber string `json:"DocumentNumber"`
		Phone          string `json:"Phone"`
	}

	Documents struct {
		Documents []StatusDocumentRequest `json:"Documents"`
	}

	StatusDocument struct {
		PossibilityCreateReturn              bool   `json:"PossibilityCreateReturn"`
		PossibilityCreateRefusal             bool   `json:"PossibilityCreateRefusal"`
		PossibilityChangeEW                  bool   `json:"PossibilityChangeEW"`
		PossibilityCreateRedirecting         bool   `json:"PossibilityCreateRedirecting"`
		Number                               string `json:"Number"`
		Redelivery                           string `json:"Redelivery"`
		RedeliverySum                        string `json:"RedeliverySum"`
		RedeliveryNum                        string `json:"RedeliveryNum"`
		RedeliveryPayer                      string `json:"RedeliveryPayer"`
		OwnerDocumentType                    string `json:"OwnerDocumentType"`
		LastCreatedOnTheBasisDocumentType    string `json:"LastCreatedOnTheBasisDocumentType"`
		LastCreatedOnTheBasisPayerType       string `json:"LastCreatedOnTheBasisPayerType"`
		LastCreatedOnTheBasisDateTime        string `json:"LastCreatedOnTheBasisDateTime"`
		LastTransactionStatusGM              string `json:"LastTransactionStatusGM"`
		LastTransactionDateTimeGM            string `json:"LastTransactionDateTimeGM"`
		LastAmountTransferGM                 string `json:"LastAmountTransferGM"`
		DateCreated                          string `json:"DateCreated"`
		DocumentWeight                       string `json:"DocumentWeight"`
		FactualWeight                        string `json:"FactualWeight"`
		VolumeWeight                         string `json:"VolumeWeight"`
		CheckWeight                          string `json:"CheckWeight"`
		CheckWeightMethod                    string `json:"CheckWeightMethod"`
		DocumentCost                         string `json:"DocumentCost"`
		CalculatedWeight                     string `json:"CalculatedWeight"`
		SumBeforeCheckWeight                 string `json:"SumBeforeCheckWeight"`
		PayerType                            string `json:"PayerType"`
		RecipientFullName                    string `json:"RecipientFullName"`
		RecipientDateTime                    string `json:"RecipientDateTime"`
		ScheduledDeliveryDate                string `json:"ScheduledDeliveryDate"`
		PaymentMethod                        string `json:"PaymentMethod"`
		CargoDescriptionString               string `json:"CargoDescriptionString"`
		CargoType                            string `json:"CargoType"`
		CitySender                           string `json:"CitySender"`
		CityRecipient                        string `json:"CityRecipient"`
		WarehouseRecipient                   string `json:"WarehouseRecipient"`
		CounterpartyType                     string `json:"CounterpartyType"`
		AfterpaymentOnGoodsCost              string `json:"AfterpaymentOnGoodsCost"`
		ServiceType                          string `json:"ServiceType"`
		UndeliveryReasonsSubtypeDescription  string `json:"UndeliveryReasonsSubtypeDescription"`
		WarehouseRecipientNumber             string `json:"WarehouseRecipientNumber"`
		LastCreatedOnTheBasisNumber          string `json:"LastCreatedOnTheBasisNumber"`
		PhoneRecipient                       string `json:"PhoneRecipient"`
		RecipientFullNameEW                  string `json:"RecipientFullNameEW"`
		WarehouseRecipientInternetAddressRef string `json:"WarehouseRecipientInternetAddressRef"`
		MarketplacePartnerToken              string `json:"MarketplacePartnerToken"`
		ClientBarcode                        string `json:"ClientBarcode"`
		RecipientAddress                     string `json:"RecipientAddress"`
		CounterpartyRecipientDescription     string `json:"CounterpartyRecipientDescription"`
		CounterpartySenderType               string `json:"CounterpartySenderType"`
		DateScan                             string `json:"DateScan"`
		PaymentStatus                        string `json:"PaymentStatus"`
		PaymentStatusDate                    string `json:"PaymentStatusDate"`
		AmountToPay                          string `json:"AmountToPay"`
		AmountPaid                           string `json:"AmountPaid"`
		Status                               string `json:"Status"`
		StatusCode                           string `json:"StatusCode"`
		RefEW                                string `json:"RefEW"`
		BackwardDeliverySubTypesActions      string `json:"BackwardDeliverySubTypesActions"`
		BackwardDeliverySubTypesServices     string `json:"BackwardDeliverySubTypesServices"`
		UndeliveryReasons                    string `json:"UndeliveryReasons"`
		DatePayedKeeping                     string `json:"DatePayedKeeping"`
		InternationalDeliveryType            string `json:"InternationalDeliveryType"`
		SeatsAmount                          string `json:"SeatsAmount"`
		CardMaskedNumber                     string `json:"CardMaskedNumber"`
		ExpressWaybillPaymentStatus          string `json:"ExpressWaybillPaymentStatus"`
		ExpressWaybillAmountToPay            string `json:"ExpressWaybillAmountToPay"`
		PhoneSender                          string `json:"PhoneSender"`
		TrackingUpdateDate                   string `json:"TrackingUpdateDate"`
		WarehouseSender                      string `json:"WarehouseSender"`
		DateReturnCargo                      string `json:"DateReturnCargo"`
		DateMoving                           string `json:"DateMoving"`
		DateFirstDayStorage                  string `json:"DateFirstDayStorage"`
		RefCityRecipient                     string `json:"RefCityRecipient"`
		RefCitySender                        string `json:"RefCitySender"`
		RefSettlementRecipient               string `json:"RefSettlementRecipient"`
		RefSettlementSender                  string `json:"RefSettlementSender"`
		SenderAddress                        string `json:"SenderAddress"`
		SenderFullNameEW                     string `json:"SenderFullNameEW"`
		AnnouncedPrice                       string `json:"AnnouncedPrice"`
		AdditionalInformationEW              string `json:"AdditionalInformationEW"`
		ActualDeliveryDate                   string `json:"ActualDeliveryDate"`
		PostomatV3CellReservationNumber      string `json:"PostomatV3CellReservationNumber"`
		OwnerDocumentNumber                  string `json:"OwnerDocumentNumber"`
		LastAmountReceivedCommissionGM       string `json:"LastAmountReceivedCommissionGM"`
		DeliveryTimeframe                    string `json:"DeliveryTimeframe"`
		CreatedOnTheBasis                    string `json:"CreatedOnTheBasis"`
		UndeliveryReasonsDate                string `json:"UndeliveryReasonsDate"`
		RecipientWarehouseTypeRef            string `json:"RecipientWarehouseTypeRef"`
		WarehouseRecipientRef                string `json:"WarehouseRecipientRef"`
		CategoryOfWarehouse                  string `json:"CategoryOfWarehouse"`
		WarehouseRecipientAddress            string `json:"WarehouseRecipientAddress"`
		WarehouseSenderInternetAddressRef    string `json:"WarehouseSenderInternetAddressRef"`
		WarehouseSenderAddress               string `json:"WarehouseSenderAddress"`
		AviaDelivery                         string `json:"AviaDelivery"`
		BarcodeRedBox                        string `json:"BarcodeRedBox"`
		CargoReturnRefusal                   string `json:"CargoReturnRefusal"`
		DaysStorageCargo                     string `json:"DaysStorageCargo"`
		Packaging                            any    `json:"Packaging"`
		PartialReturnGoods                   any    `json:"PartialReturnGoods"`
		SecurePayment                        string `json:"SecurePayment"`
		PossibilityChangeCash2Card           bool   `json:"PossibilityChangeCash2Card"`
		PossibilityChangeDeliveryIntervals   bool   `json:"PossibilityChangeDeliveryIntervals"`
		PossibilityTermExtensio              bool   `json:"PossibilityTermExtensio"`
		StorageAmount                        string `json:"StorageAmount"`
		StoragePrice                         string `json:"StoragePrice"`
		FreeShipping                         string `json:"FreeShipping"`
		LoyaltyCardRecipient                 string `json:"LoyaltyCardRecipient"`
	}
)

// GetStatusDocuments Трекінг
//
// Оновлений метод «getStatusDocuments» працює в моделі «TrackingDocument», цей метод дозволяє переглядати більш
// розширену інформацію щодо статусу відправлення.
//
// При введеному номері телефону можна отримати наступні відомості: дані відправника або одержувача, номер телефону.
//
// Метод дозволяє переглядати одночасно до 100 відправлень.
//
// Доступність: Не вимагає використання API-ключа.
//
// # Актуальні статуси трекінгу
//
// 1
// Відправник самостійно створив цю накладну, але ще не надав до відправки
//
// 2
// Видалено
//
// 3
// Номер не знайдено
//
// 4
// Відправлення у місті ХХXХ. (Статус для межобластных отправлений)
//
// 41
// Відправлення у місті ХХXХ. (Статус для услуг локал стандарт и локал экспресс - доставка в пределах города)
//
// 5
// Відправлення прямує до міста YYYY
//
// 6
// Відправлення у місті YYYY, орієнтовна доставка до ВІДДІЛЕННЯ-XXX dd-mm. Очікуйте додаткове повідомлення про прибуття
//
// 7
// Прибув на відділення
//
// 8
// Прибув на відділення (завантажено в Поштомат)
//
// 9
// Відправлення отримано
//
// 10
// Відправлення отримано %DateReceived%. Протягом доби ви одержите SMS-повідомлення про надходження грошового переказу та зможете отримати його в касі відділення «Нова пошта»
//
// 11
// Відправлення отримано %DateReceived%. Грошовий переказ видано одержувачу.
//
// 12
// Нова Пошта комплектує ваше відправлення
//
// 101
// На шляху до одержувача
//
// 102
// Відмова від отримання (Відправником створено замовлення на повернення)
//
// 103
// Відмова одержувача (отримувач відмовився від відправлення)
//
// 104
// Змінено адресу
//
// 105
// Припинено зберігання
//
// 106
// Одержано і створено ЄН зворотньої доставки
//
// 111
// Невдала спроба доставки через відсутність Одержувача на адресі або зв'язку з ним
//
// 112
// Дата доставки перенесена Одержувачем
func (c Client) GetStatusDocuments(documents Documents) (Response[StatusDocument], error) {
	return request[StatusDocument](c, TrackingDocumentModel, "getStatusDocuments", documents)
}
