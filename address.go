package novapost

type (
	SettlementRequest struct {
		CityName string
		Limit    int
		Page     int
	}

	SettlementAddress struct {
		Present                string
		Warehouses             int
		MainDescription        string
		Area                   string
		Region                 string
		SettlementTypeCode     string
		Ref                    string
		DeliveryCity           string
		AddressDeliveryAllowed bool
		StreetsAvailability    bool
		ParentRegionTypes      string
		ParentRegionCode       string
		RegionTypes            string
		RegionTypesCode        string
	}

	Settlement struct {
		TotalCount int
		Addresses  []SettlementAddress `xml:"Addresses>item"`
	}
)

// SearchSettlements Онлайн пошук в довіднику населених пунктів
//
// Метод «searchSettlements», працює в моделі «Address», цей метод необхідний для ОНЛАЙН ПОШУКУ населених пунктів.
func (c *Client) SearchSettlements(req SettlementRequest) (*Response[Settlement], error) {
	return RawRequest[Settlement](c, AddressModel, "searchSettlements", req)
}

type (
	SettlementStreetRequest struct {
		StreetName    string
		SettlementRef string
		Page          int
	}

	Location struct {
		Latitude  float64 `xml:"lat" json:"lat"`
		Longitude float64 `xml:"lon" json:"lon"`
	}

	SettlementStreetAddress struct {
		SettlementRef               string
		SettlementStreetRef         string
		SettlementStreetDescription string
		Present                     string
		StreetsType                 string
		StreetsTypeDescription      string
		Location                    Location
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
func (c *Client) SearchSettlementStreets(req SettlementStreetRequest) (*Response[SettlementStreet], error) {
	return RawRequest[SettlementStreet](c, AddressModel, "searchSettlementStreets", req)
}

type SaveAddressRequest struct {
	CounterpartyRef string
	StreetRef       string
	BuildingNumber  string
	Flat            string
	Note            string
}

// SaveAddress Створити адресу контрагента (відправник / одержувач)
//
// Метод «save», працює в моделі «Address», цей метод зберігає адресу контрагента отримувача/відправника.
func (c *Client) SaveAddress(req SaveAddressRequest) (*Response[RefDescription], error) {
	return RawRequest[RefDescription](c, AddressModel, "save", req)
}

// DeleteAddress Видалити адресу контрагента (відправник / одержувач)
//
// Метод «delete», працює в моделі «Address», цей метод необхідний для видалення адреси контрагента
// відправника/отримувача.
//
// Редагувати дані контрагента можливо лише з моменту його створення та до моменту створення ІД з даним контаргентом.
func (c *Client) DeleteAddress(rd RefDescription) (*Response[RefDescription], error) {
	return RawRequest[RefDescription](c, AddressModel, "delete", rd)
}

type UpdateAddressRequest struct {
	CounterpartyRef string
	StreetRef       string
	BuildingNumber  string
	Flat            string
	Note            string
	Ref             string
}

// UpdateAddress Редагувати адресу контрагента (відправник / одержувач)
//
// Метод «update», працює в моделі «Address», цей метод необхідний для оновлення адреси контрагента
// відправника/отримувача.
//
// Редагувати дані контрагента можливо лише з моменту його створення та до моменту створення ІД з даним контаргентом.
func (c *Client) UpdateAddress(req UpdateAddressRequest) (*Response[RefDescription], error) {
	return RawRequest[RefDescription](c, AddressModel, "update", req)
}

type (
	GetSettlementsRequest struct {
		AreaRef      string
		Ref          string
		RegionRef    string
		Page         int
		Warehouse    bool
		FindByString string
		Limit        int
	}

	SettlementGeneral struct {
		Ref                               string
		SettlementType                    string
		Latitude                          float64
		Longitude                         float64
		Description                       string
		DescriptionTranslit               string
		SettlementTypeDescription         string
		SettlementTypeDescriptionTranslit string
		Region                            string
		RegionsDescription                string
		RegionsDescriptionTranslit        string
		Area                              string
		AreaDescription                   string
		AreaDescriptionTranslit           string
		Index1                            string
		Index2                            string
		IndexCOATSU1                      string
		Delivery1                         bool
		Delivery2                         bool
		Delivery3                         bool
		Delivery4                         bool
		Delivery5                         bool
		Delivery6                         bool
		Delivery7                         bool
		Warehouse                         bool
		SpecialCashCheck                  bool
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
func (c *Client) GetSettlements(req GetSettlementsRequest) (*Response[SettlementGeneral], error) {
	return RawRequest[SettlementGeneral](c, AddressGeneralModel, "getSettlements", req)
}

type (
	CityRequest struct {
		Ref          string
		Page         int
		FindByString string
		Limit        int
	}

	City struct {
		Description                string
		Ref                        string
		Delivery1                  bool
		Delivery2                  bool
		Delivery3                  bool
		Delivery4                  bool
		Delivery5                  bool
		Delivery6                  bool
		Delivery7                  bool
		Area                       string
		SettlementType             string
		IsBranch                   bool
		PreventEntryNewStreetsUser bool
		CityID                     int
		SettlementTypeDescription  string
		SpecialCashCheck           bool
		AreaDescription            string
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
func (c *Client) GetCities(req CityRequest) (*Response[City], error) {
	return RawRequest[City](c, AddressModel, "getCities", req)
}

type Area struct {
	Ref         string
	AreasCenter string
	Description string
}

// GetAreas Довідник географічних областей України
//
// Метод «getAreas», працює в моделі «Address», цей метод необхідий для завантаження довідника географічних областей
// України, компанії «Новая почша».
func (c *Client) GetAreas() (*Response[Area], error) {
	return RawRequest[Area](c, AddressModel, "getAreas", nil)
}

type (
	WarehouseRequest struct {
		PostFinance        bool
		BicycleParking     bool
		POSTerminal        bool
		FindByString       string
		CityName           string
		CityRef            string
		Page               int
		Limit              int
		Language           string
		TypeOfWarehouseRef string
		WarehouseId        int
	}

	SendingLimitationsOnDimensions struct {
		Width  float64
		Height float64
		Length float64
	}

	ReceivingLimitationsOnDimensions struct {
		Width  float64
		Height float64
		Length float64
	}

	Reception struct {
		Monday    string
		Tuesday   string
		Wednesday string
		Thursday  string
		Friday    string
		Saturday  string
		Sunday    string
	}

	Delivery struct {
		Monday    string
		Tuesday   string
		Wednesday string
		Thursday  string
		Friday    string
		Saturday  string
		Sunday    string
	}

	Schedule struct {
		Monday    string
		Tuesday   string
		Wednesday string
		Thursday  string
		Friday    string
		Saturday  string
		Sunday    string
	}

	Warehouse struct {
		SiteKey                          int
		Description                      string
		ShortAddress                     string
		Phone                            string
		TypeOfWarehouse                  string
		Ref                              string
		Number                           string
		CityRef                          string
		CityDescription                  string
		SettlementRef                    string
		SettlementDescription            string
		SettlementAreaDescription        string
		SettlementRegionsDescription     string
		SettlementTypeDescription        string
		Longitude                        float64
		Latitude                         float64
		PostFinance                      bool
		BicycleParking                   bool
		PaymentAccess                    bool
		POSTerminal                      bool
		InternationalShipping            bool
		SelfServiceWorkplacesCount       bool
		TotalMaxWeightAllowed            float64
		PlaceMaxWeightAllowed            float64
		SendingLimitationsOnDimensions   SendingLimitationsOnDimensions
		ReceivingLimitationsOnDimensions ReceivingLimitationsOnDimensions
		Reception                        Reception
		Delivery                         Delivery
		Schedule                         Schedule
		DistrictCode                     string
		WarehouseStatus                  string
		WarehouseStatusDate              string
		WarehouseIllusha                 string
		CategoryOfWarehouse              string
		Direct                           string
		RegionCity                       string
		WarehouseForAgent                bool
		GeneratorEnabled                 bool
		MaxDeclaredCost                  float64
		WorkInMobileAwis                 bool
		DenyToSelect                     bool
		CanGetMoneyTransfer              bool
		HasMirror                        bool
		HasFittingRoom                   bool
		OnlyReceivingParcel              bool
		PostMachineType                  string
		PostalCodeUA                     string
		WarehouseIndex                   string
		BeaconCode                       string
		PostomatFor                      string
	}
)

// GetWarehouses Довідник відділень та поштоматів
//
// Метод «getWarehouses», працює в моделі «Address», цей метод завантажує довідник відділень «Нова пошта» в розрізі
// населених пунктів України.
func (c *Client) GetWarehouses(req WarehouseRequest) (*Response[Warehouse], error) {
	return RawRequest[Warehouse](c, AddressModel, "getWarehouses", req)
}

// GetWarehouseTypes Довідник типів відділень
//
// Метод «getWarehouseTypes», працює в моделі «Address», цей метод завантажує довідник типів відділень «Нова пошта».
func (c *Client) GetWarehouseTypes() (*Response[RefDescription], error) {
	return RawRequest[RefDescription](c, AddressModel, "getWarehouseTypes", nil)
}

type (
	StreetRequest struct {
		CityRef      string
		FindByString string
		Page         int
		Limit        int
	}

	Street struct {
		Ref            string
		Description    string
		StreetsTypeRef string
		StreetsType    string
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
func (c *Client) GetStreet(req StreetRequest) (*Response[Street], error) {
	return RawRequest[Street](c, AddressModel, "getStreet", req)
}

type (
	AreaRef struct {
		AreaRef string
	}

	AreaRegion struct {
		Ref         string
		AreasCenter string
		Description string
		RegionType  string
	}
)

// GetSettlementCountryRegion Довідник районів областей населених пунктів
//
// Метод дозволяє отримати довідник районів облестей населених пунктів, тих самих які можна побачити викорисовуючи метод
// «getSettlements»
func (c *Client) GetSettlementCountryRegion(ref AreaRef) (*Response[AreaRegion], error) {
	return RawRequest[AreaRegion](c, AddressModel, "getSettlementCountryRegion", ref)
}

// GetSettlementAreas Довідник областей населених пунктів
//
// Метод дозволяє отримати довідник облестей населених пунктів, тих самих які можна побачити викорисовуючи метод
// «getSettlements»
func (c *Client) GetSettlementAreas(ref Ref) (*Response[AreaRegion], error) {
	return RawRequest[AreaRegion](c, AddressModel, "getSettlementAreas", ref)
}
