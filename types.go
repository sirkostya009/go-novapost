package novapost

type (
	Response[T any] struct {
		Success      bool     `json:"success" xml:"success"`
		Data         []T      `json:"data" xml:"data"`
		Errors       []string `json:"errors" xml:"errors"`
		Warnings     []string `json:"warnings" xml:"warnings"`
		Info         any      `json:"info" xml:"info"`
		MessageCodes []string `json:"messageCodes" xml:"messageCodes"`
		ErrorCodes   []string `json:"errorCodes" xml:"errorCodes"`
		WarningCodes []string `json:"warningCodes" xml:"warningCodes"`
		InfoCodes    []string `json:"infoCodes" xml:"infoCodes"`
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
		Ref         string `json:"Ref"`
		Description string `json:"Description"`
	}

	RefNumber struct {
		Ref    string `json:"Ref"`
		Number string `json:"Number"`
	}
)
