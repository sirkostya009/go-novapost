package novapost

type (
	SettlementRequest struct {
		CityName string
		Limit    int `json:",string"`
		Page     int `json:",string"`
	}

	SettlementAddress struct {
		Present                string
		Warehouses             int `json:",string"`
		AddressDeliveryAllowed bool
		StreetsAvailability    bool
		MainDescription        string
		Area                   string
		Region                 string
		SettlementTypeCode     string
		Ref                    string
		DeliveryCity           string
		ParentRegionTypes      string
		ParentRegionCode       string
		RegionTypes            string
		RegionTypesCode        string
	}

	Settlement struct {
		TotalCount int                 `json:",string"`
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
		Page          int `json:",string,omitempty" xml:",omitempty"`
		Limit         int `json:",string,omitempty" xml:",omitempty"`
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
		TotalCount string                    `json:",string"`
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
	Note            string `json:",omitempty" xml:",omitempty"`
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
// TODO: replace with Ref
func (c *Client) DeleteAddress(rd RefDescription) (*Response[RefDescription], error) {
	return RawRequest[RefDescription](c, AddressModel, "delete", rd)
}

type UpdateAddressRequest struct {
	CounterpartyRef string `json:",omitempty" xml:",omitempty"`
	StreetRef       string `json:",omitempty" xml:",omitempty"`
	BuildingNumber  string `json:",omitempty" xml:",omitempty"`
	Flat            string `json:",omitempty" xml:",omitempty"`
	Note            string `json:",omitempty" xml:",omitempty"`
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
		AreaRef      string  `json:",omitempty" xml:",omitempty"`
		Ref          string  `json:",omitempty" xml:",omitempty"`
		RegionRef    string  `json:",omitempty" xml:",omitempty"`
		Warehouse    IntBool `json:",string,omitempty" xml:",omitempty"` // true or false as a string dont work here FIXME needs a fix
		FindByString string  `json:",omitempty" xml:",omitempty"`
		Page         int     `json:",string,omitempty" xml:",omitempty"`
		Limit        int     `json:",string,omitempty" xml:",omitempty"`
	}

	SettlementGeneral struct {
		Ref                               string
		SettlementType                    string
		Latitude                          float64 `json:",string"`
		Longitude                         float64 `json:",string"`
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
		Delivery1                         IntBool `json:",string"`
		Delivery2                         IntBool `json:",string"`
		Delivery3                         IntBool `json:",string"`
		Delivery4                         IntBool `json:",string"`
		Delivery5                         IntBool `json:",string"`
		Delivery6                         IntBool `json:",string"`
		Delivery7                         IntBool `json:",string"`
		Warehouse                         IntBool `json:",string"`
		SpecialCashCheck                  IntBool `json:",string"`
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
		Ref          string `json:",omitempty" xml:",omitempty"`
		FindByString string `json:",omitempty" xml:",omitempty"`
		Page         int    `json:",string,omitempty" xml:",omitempty"`
		Limit        int    `json:",string,omitempty" xml:",omitempty"`
	}

	City struct {
		Delivery1                  IntBool `json:",string"`
		Delivery2                  IntBool `json:",string"`
		Delivery3                  IntBool `json:",string"`
		Delivery4                  IntBool `json:",string"`
		Delivery5                  IntBool `json:",string"`
		Delivery6                  IntBool `json:",string"`
		Delivery7                  IntBool `json:",string"`
		IsBranch                   IntBool `json:",string"`
		PreventEntryNewStreetsUser IntBool `json:",string"`
		SpecialCashCheck           IntBool `json:",string"`
		CityID                     int     `json:",string"`
		Description                string
		Ref                        string
		Area                       string
		SettlementType             string
		SettlementTypeDescription  string
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
		PostFinance        IntBool `json:",string,omitempty" xml:",omitempty"`
		BicycleParking     IntBool `json:",string,omitempty" xml:",omitempty"`
		POSTerminal        IntBool `json:",string,omitempty" xml:",omitempty"`
		WarehouseId        int     `json:",string,omitempty" xml:",omitempty"`
		FindByString       string  `json:",omitempty" xml:",omitempty"`
		CityName           string  `json:",omitempty" xml:",omitempty"`
		CityRef            string  `json:",omitempty" xml:",omitempty"`
		Page               int     `json:",string,omitempty" xml:",omitempty"`
		Limit              int     `json:",string,omitempty" xml:",omitempty"`
		Language           string  `json:",omitempty" xml:",omitempty"`
		TypeOfWarehouseRef string  `json:",omitempty" xml:",omitempty"`
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
		SiteKey                          string
		WorkInMobileAwis                 IntBool `json:",string"`
		DenyToSelect                     IntBool `json:",string"`
		CanGetMoneyTransfer              IntBool `json:",string"`
		HasMirror                        IntBool `json:",string"`
		HasFittingRoom                   IntBool `json:",string"`
		OnlyReceivingParcel              IntBool `json:",string"`
		PostFinance                      IntBool `json:",string"`
		BicycleParking                   IntBool `json:",string"`
		PaymentAccess                    IntBool `json:",string"`
		POSTerminal                      IntBool `json:",string"`
		InternationalShipping            IntBool `json:",string"`
		SelfServiceWorkplacesCount       IntBool `json:",string"`
		WarehouseForAgent                IntBool `json:",string"`
		GeneratorEnabled                 IntBool `json:",string"`
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
		Longitude                        float64 `json:",string"`
		Latitude                         float64 `json:",string"`
		TotalMaxWeightAllowed            float64 `json:",string"`
		PlaceMaxWeightAllowed            float64 `json:",string"`
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
		MaxDeclaredCost                  float64 `json:",string"`
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
		Page         int `json:",string,omitempty" xml:",omitempty"`
		Limit        int `json:",string,omitempty" xml:",omitempty"`
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
