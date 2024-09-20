package face

type PortfolioType string

const (
	PortfolioTypeEmpty     PortfolioType = ""
	PortfolioTypeUndefined PortfolioType = "UNDEFINED"
	PortfolioTypeDefault   PortfolioType = "DEFAULT"
	PortfolioTypeConsumer  PortfolioType = "CONSUMER"
	PortfolioTypeINTX      PortfolioType = "INTX"
)

type Portfolio struct {
	Name    string        `json:"name"`
	Uuid    string        `json:"uuid"`
	Type    PortfolioType `json:"type"`
	Deleted bool          `json:"deleted"`
}

type ListPortfolios struct {
	Portfolios []*Portfolio `json:"portfolios"`
}

func (l *ListPortfolios) GetPortfolios() []*Portfolio {
	if l == nil {
		return nil
	}

	return l.Portfolios
}
