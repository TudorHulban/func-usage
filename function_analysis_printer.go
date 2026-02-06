package funcusage

type Printer struct {
	columns []string
}

func NewPrinter() *Printer {
	return &Printer{
		columns: make([]string, 0, 7),
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

func (p *Printer) WithTotal() *Printer {
	p.columns = append(p.columns, _LabelTotal)

	return p
}
