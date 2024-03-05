package novapost

type (
	Response[T any] struct {
		Success      bool     `json:"success" xml:"success"`
		Data         []T      `json:"data" xml:"data>item"`
		Errors       []string `json:"errors" xml:"errors>item"`
		Warnings     []string `json:"warnings" xml:"warnings>item"`
		Info         any      `json:"info" xml:"info"`
		MessageCodes []string `json:"messageCodes" xml:"messageCodes>item"`
		ErrorCodes   []string `json:"errorCodes" xml:"errorCodes>item"`
		WarningCodes []string `json:"warningCodes" xml:"warningCodes>item"`
		InfoCodes    []string `json:"infoCodes" xml:"infoCodes>item"`
	}

	Request struct {
		ApiKey           string `json:"apiKey" xml:"apiKey"`
		ModelName        string `json:"modelName" xml:"modelName"`
		CalledMethod     string `json:"calledMethod" xml:"calledMethod"`
		MethodProperties any    `json:"methodProperties" xml:"methodProperties"`
	}

	Ref struct {
		Ref string
	}

	RefDescription struct {
		Ref         string
		Description string
	}

	RefNumber struct {
		Ref    string
		Number string
	}
)
