package funcusage

type Printer struct {
	columns []string
}

func NewPrinter() *Printer {
	return &Printer{
		columns: make([]string, 0, 9),
	}
}

func (p *Printer) WithName() *Printer {
	p.columns = append(p.columns, _LabelName)

	return p
}

func (p *Printer) WithMethodOf() *Printer {
	p.columns = append(p.columns, _LabelMethodOf)

	return p
}

func (p *Printer) WithPosition() *Printer {
	p.columns = append(p.columns, _LabelPosition)

	return p
}

func (p *Printer) WithTotal() *Printer {
	p.columns = append(p.columns, _LabelTotal)

	return p
}

func (p *Printer) WithTypesParams() *Printer {
	p.columns = append(p.columns, _LabelTypesParams)

	return p
}

func (p *Printer) WithTypesResults() *Printer {
	p.columns = append(p.columns, _LabelTypesResults)

	return p
}
