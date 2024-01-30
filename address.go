package novapost

type (
	SettlementRequest struct {
		CityName string
		Limit    int
		Page     int
	}

	SettlementAddress struct {
		Present                string `json:"Present"`
		Warehouses             int    `json:"Warehouses"`
		MainDescription        string `json:"MainDescription"`
		Area                   string `json:"Area"`
		Region                 string `json:"Region"`
		SettlementTypeCode     string `json:"SettlementTypeCode"`
		Ref                    string `json:"Ref"`
		DeliveryCity           string `json:"DeliveryCity"`
		AddressDeliveryAllowed bool   `json:"AddressDeliveryAllowed"`
		StreetsAvailability    bool   `json:"StreetsAvailability"`
		ParentRegionTypes      string `json:"ParentRegionTypes"`
		ParentRegionCode       string `json:"ParentRegionCode"`
		RegionTypes            string `json:"RegionTypes"`
		RegionTypesCode        string `json:"RegionTypesCode"`
	}

	Settlement struct {
		TotalCount string              `json:"TotalCount"`
		Addresses  []SettlementAddress `json:"Addresses"`
	}
)

// SearchSettlements Онлайн пошук в довіднику населених пунктів
//
// Метод «searchSettlements», працює в моделі «Address», цей метод необхідний для ОНЛАЙН ПОШУКУ населених пунктів.
func (c Client) SearchSettlements(req SettlementRequest) (Response[Settlement], error) {
	return request[Settlement](c, AddressModel, "getSettlements", req)
}

type (
	SettlementStreetRequest struct {
		StreetName    string `json:"StreetName"`
		SettlementRef string `json:"SettlementRef"`
		Page          int    `json:"Page"`
	}

	Location struct {
		Latitude  string `json:"lat"`
		Longitude string `json:"lon"`
	}

	SettlementStreetAddress struct {
		SettlementRef               string `json:"SettlementRef"`
		SettlementStreetRef         string `json:"SettlementStreetRef"`
		SettlementStreetDescription string `json:"SettlementStreetDescription"`
		Present                     string `json:"Present"`
		StreetsType                 string `json:"StreetsType"`
		StreetsTypeDescription      string `json:"StreetsTypeDescription"`
		Location                    `json:"Location"`
	}

	SettlementStreet struct {
		TotalCount string                    `json:"TotalCount"`
		Addresses  []SettlementStreetAddress `json:"Addresses"`
	}
)

// SearchSettlementStreets Онлайн пошук вулиць в довіднику населених пунктів
//
// Метод «searchSettlementStreets», працює в моделі «Address», цей метод необхідний для ОНЛАЙН ПОШУКУ вулиць в обраному
// населеному пункті.
func (c Client) SearchSettlementStreets(req SettlementStreetRequest) (Response[SettlementStreet], error) {
	return request[SettlementStreet](c, AddressModel, "getSettlementStreets", req)
}

type SaveAddressRequest struct {
	CounterpartyRef string `json:"CounterpartyRef"`
	StreetRef       string `json:"StreetRef"`
	BuildingNumber  string `json:"BuildingNumber"`
	Flat            string `json:"Flat"`
	Note            string `json:"Note"`
}

// SaveAddress Створити адресу контрагента (відправник / одержувач)
//
// Метод «save», працює в моделі «Address», цей метод зберігає адресу контрагента отримувача/відправника.
func (c Client) SaveAddress(req SaveAddressRequest) (Response[RefDescription], error) {
	return request[RefDescription](c, AddressModel, "save", req)
}

// DeleteAddress Видалити адресу контрагента (відправник / одержувач)
//
// Метод «delete», працює в моделі «Address», цей метод необхідний для видалення адреси контрагента
// відправника/отримувача.
//
// Редагувати дані контрагента можливо лише з моменту його створення та до моменту створення ІД з даним контаргентом.
func (c Client) DeleteAddress(req RefDescription) (Response[RefDescription], error) {
	return request[RefDescription](c, AddressModel, "delete", req)
}

type UpdateAddressRequest struct {
	CounterpartyRef string `json:"CounterpartyRef"`
	StreetRef       string `json:"StreetRef"`
	BuildingNumber  string `json:"BuildingNumber"`
	Flat            string `json:"Flat"`
	Note            string `json:"Note"`
	Ref             string `json:"Ref"`
}

// UpdateAddress Редагувати адресу контрагента (відправник / одержувач)
//
// Метод «update», працює в моделі «Address», цей метод необхідний для оновлення адреси контрагента
// відправника/отримувача.
//
// Редагувати дані контрагента можливо лише з моменту його створення та до моменту створення ІД з даним контаргентом.
func (c Client) UpdateAddress(req UpdateAddressRequest) (Response[RefDescription], error) {
	return request[RefDescription](c, AddressModel, "update", req)
}

type (
	GetSettlementsRequest struct {
		AreaRef      string `json:"AreaRef"`
		Ref          string `json:"Ref"`
		RegionRef    string `json:"RegionRef"`
		Page         int    `json:"Page"`
		Warehouse    string `json:"Warehouse"`
		FindByString string `json:"FindByString"`
		Limit        int    `json:"Limit"`
	}

	SettlementGeneral struct {
		Ref                               string `json:"Ref"`
		SettlementType                    string `json:"SettlementType"`
		Latitude                          string `json:"Latitude"`
		Longitude                         string `json:"Longitude"`
		Description                       string `json:"Description"`
		DescriptionTranslit               string `json:"DescriptionTranslit"`
		SettlementTypeDescription         string `json:"SettlementTypeDescription"`
		SettlementTypeDescriptionTranslit string `json:"SettlementTypeDescriptionTranslit"`
		Region                            string `json:"Region"`
		RegionsDescription                string `json:"RegionsDescription"`
		RegionsDescriptionTranslit        string `json:"RegionsDescriptionTranslit"`
		Area                              string `json:"Area"`
		AreaDescription                   string `json:"AreaDescription"`
		AreaDescriptionTranslit           string `json:"AreaDescriptionTranslit"`
		Index1                            string `json:"Index1"`
		Index2                            string `json:"Index2"`
		IndexCOATSU1                      string `json:"IndexCOATSU1"`
		Delivery1                         string `json:"Delivery1"`
		Delivery2                         string `json:"Delivery2"`
		Delivery3                         string `json:"Delivery3"`
		Delivery4                         string `json:"Delivery4"`
		Delivery5                         string `json:"Delivery5"`
		Delivery6                         string `json:"Delivery6"`
		Delivery7                         string `json:"Delivery7"`
		Warehouse                         string `json:"Warehouse"`
		SpecialCashCheck                  int    `json:"SpecialCashCheck"`
	}
)

// GetSettlements Довідник населених пунктів України
//
// Метод «getSettlements», працює в моделі «AddressGeneral», цей метод надає можливість завантажити довідник міст
// України (Українською або Російською мовами), до яких здійснюється доставка вантажів компанією «Нова пошта».
//
// Слід вразовувати, що метод «getSettlements» для кожного населеного пункту повертає область та район. Метод повертає
// не більше 150 записів на сторінку. Для перегляду більш ніж 150 записів, необхідно використовувати параметр "Page"
// разом з параметром "Limit".
//
// Також для методу доступний пошук по рядку, для нього потрібно зазначити параметр FindByString. Важливо! Пошук
// можливий лише Українською мовою.
//
// Параметр "Warehouse" із значенням "1 или 0" дозволить відобразити лише ті населені пункти в яких наявні відділення
// "Нова Пошта".
func (c Client) GetSettlements(req GetSettlementsRequest) (Response[SettlementGeneral], error) {
	return request[SettlementGeneral](c, AddressGeneralModel, "getSettlements", req)
}

type (
	CityRequest struct {
		Ref          string `json:"Ref"`
		Page         int    `json:"Page"`
		FindByString string `json:"FindByString"`
		Limit        int    `json:"Limit"`
	}

	City struct {
		Description                string `json:"Description"`
		Ref                        string `json:"Ref"`
		Delivery1                  string `json:"Delivery1"`
		Delivery2                  string `json:"Delivery2"`
		Delivery3                  string `json:"Delivery3"`
		Delivery4                  string `json:"Delivery4"`
		Delivery5                  string `json:"Delivery5"`
		Delivery6                  string `json:"Delivery6"`
		Delivery7                  string `json:"Delivery7"`
		Area                       string `json:"Area"`
		SettlementType             string `json:"SettlementType"`
		IsBranch                   string `json:"IsBranch"`
		PreventEntryNewStreetsUser string `json:"PreventEntryNewStreetsUser"`
		CityID                     string `json:"CityID"`
		SettlementTypeDescription  string `json:"SettlementTypeDescription"`
		SpecialCashCheck           int    `json:"SpecialCashCheck"`
		AreaDescription            string `json:"AreaDescription"`
		AreaDescriptionRu          string `json:"AreaDescriptionRu"`
	}
)

// GetCities Довідник міст компанії
//
// Отримання довідника міст компанії «Нова Пошта» українською та російською мовами. Метод «getCities» працює в моделі
// «Address», цей метод завантажує довідник населених пунктів України.
//
// Варто враховувати, що довідник вивантажується тільки з населеними пунктами, де є відділення "Нова Пошта" і можна
// оформити доставку на відділення, а також доставку за адресою.
//
// Якщо до цього запиту додати параметр «FindByString» (пошук за рядками) та у його властивостях прописати назву
// населеного пункту (Бровари), який потрібно знайти, то отримаємо запит за допомогою якого в довіднику знаходиться
// населений пункт.
func (c Client) GetCities(req CityRequest) (Response[City], error) {
	return request[City](c, AddressModel, "getCities", req)
}

type Area struct {
	Ref         string `json:"Ref"`
	AreasCenter string `json:"AreasCenter"`
	Description string `json:"Description"`
}

// GetAreas Довідник географічних областей України
//
// Метод «getAreas», працює в моделі «Address», цей метод необхідий для завантаження довідника географічних областей
// України, компанії «Новая почша».
func (c Client) GetAreas() (Response[Area], error) {
	return request[Area](c, AddressModel, "getAreas", nil)
}

type (
	WarehouseRequest struct {
		PostFinance        bool   `json:"PostFinance"`
		BicycleParking     bool   `json:"BicycleParking"`
		POSTerminal        bool   `json:"POSTerminal"`
		FindByString       string `json:"FindByString"`
		CityName           string `json:"CityName"`
		CityRef            string `json:"CityRef"`
		Page               int    `json:"Page"`
		Limit              int    `json:"Limit"`
		Language           string `json:"Language"`
		TypeOfWarehouseRef string `json:"TypeOfWarehouseRef"`
		WarehouseId        string `json:"WarehouseId"`
	}

	SendingLimitationsOnDimensions struct {
		Width  int `json:"Width"`
		Height int `json:"Height"`
		Length int `json:"Length"`
	}

	ReceivingLimitationsOnDimensions struct {
		Width  int `json:"Width"`
		Height int `json:"Height"`
		Length int `json:"Length"`
	}

	Reception struct {
		Monday    string `json:"Monday"`
		Tuesday   string `json:"Tuesday"`
		Wednesday string `json:"Wednesday"`
		Thursday  string `json:"Thursday"`
		Friday    string `json:"Friday"`
		Saturday  string `json:"Saturday"`
		Sunday    string `json:"Sunday"`
	}

	Delivery struct {
		Monday    string `json:"Monday"`
		Tuesday   string `json:"Tuesday"`
		Wednesday string `json:"Wednesday"`
		Thursday  string `json:"Thursday"`
		Friday    string `json:"Friday"`
		Saturday  string `json:"Saturday"`
		Sunday    string `json:"Sunday"`
	}

	Schedule struct {
		Monday    string `json:"Monday"`
		Tuesday   string `json:"Tuesday"`
		Wednesday string `json:"Wednesday"`
		Thursday  string `json:"Thursday"`
		Friday    string `json:"Friday"`
		Saturday  string `json:"Saturday"`
		Sunday    string `json:"Sunday"`
	}

	Warehouse struct {
		SiteKey                          string `json:"SiteKey"`
		Description                      string `json:"Description"`
		ShortAddress                     string `json:"ShortAddress"`
		Phone                            string `json:"Phone"`
		TypeOfWarehouse                  string `json:"TypeOfWarehouse"`
		Ref                              string `json:"Ref"`
		Number                           string `json:"Number"`
		CityRef                          string `json:"CityRef"`
		CityDescription                  string `json:"CityDescription"`
		SettlementRef                    string `json:"SettlementRef"`
		SettlementDescription            string `json:"SettlementDescription"`
		SettlementAreaDescription        string `json:"SettlementAreaDescription"`
		SettlementRegionsDescription     string `json:"SettlementRegionsDescription"`
		SettlementTypeDescription        string `json:"SettlementTypeDescription"`
		Longitude                        string `json:"Longitude"`
		Latitude                         string `json:"Latitude"`
		PostFinance                      string `json:"PostFinance"`
		BicycleParking                   string `json:"BicycleParking"`
		PaymentAccess                    string `json:"PaymentAccess"`
		POSTerminal                      string `json:"POSTerminal"`
		InternationalShipping            string `json:"InternationalShipping"`
		SelfServiceWorkplacesCount       string `json:"SelfServiceWorkplacesCount"`
		TotalMaxWeightAllowed            string `json:"TotalMaxWeightAllowed"`
		PlaceMaxWeightAllowed            string `json:"PlaceMaxWeightAllowed"`
		SendingLimitationsOnDimensions   `json:"SendingLimitationsOnDimensions"`
		ReceivingLimitationsOnDimensions `json:"ReceivingLimitationsOnDimensions"`
		Reception                        `json:"Reception"`
		Delivery                         `json:"Delivery"`
		Schedule                         `json:"Schedule"`
		DistrictCode                     string `json:"DistrictCode"`
		WarehouseStatus                  string `json:"WarehouseStatus"`
		WarehouseStatusDate              string `json:"WarehouseStatusDate"`
		WarehouseIllusha                 string `json:"WarehouseIllusha"`
		CategoryOfWarehouse              string `json:"CategoryOfWarehouse"`
		Direct                           string `json:"Direct"`
		RegionCity                       string `json:"RegionCity"`
		WarehouseForAgent                string `json:"WarehouseForAgent"`
		GeneratorEnabled                 string `json:"GeneratorEnabled"`
		MaxDeclaredCost                  string `json:"MaxDeclaredCost"`
		WorkInMobileAwis                 string `json:"WorkInMobileAwis"`
		DenyToSelect                     string `json:"DenyToSelect"`
		CanGetMoneyTransfer              string `json:"CanGetMoneyTransfer"`
		HasMirror                        string `json:"HasMirror"`
		HasFittingRoom                   string `json:"HasFittingRoom"`
		OnlyReceivingParcel              string `json:"OnlyReceivingParcel"`
		PostMachineType                  string `json:"PostMachineType"`
		PostalCodeUA                     string `json:"PostalCodeUA"`
		WarehouseIndex                   string `json:"WarehouseIndex"`
		BeaconCode                       string `json:"BeaconCode"`
		PostomatFor                      string `json:"PostomatFor"`
	}
)

// GetWarehouses Довідник відділень та поштоматів
//
// Метод «getWarehouses», працює в моделі «Address», цей метод завантажує довідник відділень «Нова пошта» в розрізі
// населених пунктів України.
func (c Client) GetWarehouses(req WarehouseRequest) (Response[Warehouse], error) {
	return request[Warehouse](c, AddressModel, "getWarehouses", req)
}

// GetWarehouseTypes Довідник типів відділень
//
// Метод «getWarehouseTypes», працює в моделі «Address», цей метод завантажує довідник типів відділень «Нова пошта».
func (c Client) GetWarehouseTypes() (Response[RefDescription], error) {
	return request[RefDescription](c, AddressModel, "getWarehouseTypes", nil)
}

type (
	StreetRequest struct {
		CityRef      string `json:"CityRef"`
		FindByString string `json:"FindByString"`
		Page         int    `json:"Page"`
		Limit        int    `json:"Limit"`
	}

	Street struct {
		Ref            string `json:"Ref"`
		Description    string `json:"Description"`
		StreetsTypeRef string `json:"StreetsTypeRef"`
		StreetsType    string `json:"StreetsType"`
	}
)

// GetStreet Довідник вулиць компанії
//
// Метод «getStreet» працює в моделі «Address», цей метод завантажує довідник вулиць в рамках населених пунктів України
// куди здійснєю доставку компанія «Нова Пошта».
//
// Довідник використовується під час створення відправлень з типом доставки від/до адреси клієнта. Якщо в цей запит
// додати параметр FindByString (пошук по рядках) і в його властивостях прописати назву вулиці (Броварський), який
// потрібно знайти, то отримаємо запит за допомогою якого в довіднику знаходиться проспект або вулиця.
func (c Client) GetStreet(req StreetRequest) (Response[Street], error) {
	return request[Street](c, AddressModel, "getStreet", req)
}

type (
	AreaRef struct {
		AreaRef string `json:"AreaRef"`
	}

	AreaRegion struct {
		Ref         string `json:"Ref"`
		AreasCenter string `json:"AreasCenter"`
		Description string `json:"Description"`
		RegionType  string `json:"RegionType"`
	}
)

// GetSettlementCountryRegion Довідник районів областей населених пунктів
//
// Метод дозволяє отримати довідник районів облестей населених пунктів, тих самих які можна побачити викорисовуючи метод
// «getSettlements»
func (c Client) GetSettlementCountryRegion(ref AreaRef) (Response[AreaRegion], error) {
	return request[AreaRegion](c, AddressModel, "getSettlementCountryRegion", ref)
}

// GetSettlementAreas Довідник областей населених пунктів
//
// Метод дозволяє отримати довідник облестей населених пунктів, тих самих які можна побачити викорисовуючи метод
// «getSettlements»
func (c Client) GetSettlementAreas(ref Ref) (Response[AreaRegion], error) {
	return request[AreaRegion](c, AddressModel, "getSettlementAreas", ref)
}
