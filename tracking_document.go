package novapost

type (
	StatusDocumentRequest struct {
		DocumentNumber string
		Phone          string
	}

	Documents struct {
		Documents []StatusDocumentRequest `xml:"Documents>item"`
	}

	StatusDocument struct {
		Number                               string
		RedeliverySum                        float64
		RedeliveryNum                        string
		RedeliveryPayer                      string
		OwnerDocumentType                    string
		LastCreatedOnTheBasisDocumentType    string
		LastCreatedOnTheBasisPayerType       string
		LastCreatedOnTheBasisDateTime        string
		LastTransactionStatusGM              string
		LastTransactionDateTimeGM            string
		LastAmountTransferGM                 string
		DateCreated                          string
		DocumentWeight                       float64
		FactualWeight                        float64
		VolumeWeight                         float64
		CheckWeight                          string
		CheckWeightMethod                    string
		DocumentCost                         float64
		CalculatedWeight                     float64
		SumBeforeCheckWeight                 string
		PayerType                            string
		RecipientFullName                    string
		RecipientDateTime                    string
		ScheduledDeliveryDate                string
		PaymentMethod                        string
		CargoDescriptionString               string
		CargoType                            string
		CitySender                           string
		CityRecipient                        string
		WarehouseRecipient                   string
		CounterpartyType                     string
		AfterPaymentOnGoodsCost              string `xml:"AfterpaymentOnGoodsCost" json:"AfterpaymentOnGoodsCost"`
		ServiceType                          string
		UndeliveryReasonsSubtypeDescription  string
		WarehouseRecipientNumber             string
		LastCreatedOnTheBasisNumber          string
		PhoneRecipient                       string
		RecipientFullNameEW                  string
		WarehouseRecipientInternetAddressRef string
		MarketplacePartnerToken              string
		ClientBarcode                        string
		RecipientAddress                     string
		CounterpartyRecipientDescription     string
		CounterpartySenderType               string
		DateScan                             string
		PaymentStatus                        string
		PaymentStatusDate                    string
		AmountToPay                          float64
		AmountPaid                           float64
		Status                               string
		RefEW                                string
		BackwardDeliverySubTypesActions      string
		BackwardDeliverySubTypesServices     string
		UndeliveryReasons                    string
		DatePayedKeeping                     string
		InternationalDeliveryType            string
		CardMaskedNumber                     string
		ExpressWaybillPaymentStatus          string
		ExpressWaybillAmountToPay            float64
		PhoneSender                          string
		TrackingUpdateDate                   string
		WarehouseSender                      string
		DateReturnCargo                      string
		DateMoving                           string
		DateFirstDayStorage                  string
		RefCityRecipient                     string
		RefCitySender                        string
		RefSettlementRecipient               string
		RefSettlementSender                  string
		SenderAddress                        string
		SenderFullNameEW                     string
		AnnouncedPrice                       float64
		AdditionalInformationEW              string
		ActualDeliveryDate                   string
		PostomatV3CellReservationNumber      string
		OwnerDocumentNumber                  string
		LastAmountReceivedCommissionGM       float64
		DeliveryTimeframe                    string
		CreatedOnTheBasis                    string
		UndeliveryReasonsDate                string
		RecipientWarehouseTypeRef            string
		WarehouseRecipientRef                string
		CategoryOfWarehouse                  string
		WarehouseRecipientAddress            string
		WarehouseSenderInternetAddressRef    string
		WarehouseSenderAddress               string
		AviaDelivery                         string
		BarcodeRedBox                        string
		LoyaltyCardRecipient                 string
		StatusCode                           int
		SeatsAmount                          int
		DaysStorageCargo                     string
		Packaging                            []any `xml:"Packaging>item"`
		PartialReturnGoods                   []any `xml:"PartialReturnGoods>item"`
		StoragePrice                         float64
		StorageAmount                        int
		FreeShipping                         bool
		CargoReturnRefusal                   bool
		PossibilityCreateReturn              bool
		PossibilityCreateRefusal             bool
		PossibilityChangeEW                  bool
		PossibilityCreateRedirecting         bool
		Redelivery                           bool
		SecurePayment                        bool
		PossibilityChangeCash2Card           bool
		PossibilityChangeDeliveryIntervals   bool
		PossibilityTermExtensio              bool
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
func (c *Client) GetStatusDocuments(documents Documents) (*Response[StatusDocument], error) {
	return RawRequest[StatusDocument](c, TrackingDocumentModel, "getStatusDocuments", documents)
}
