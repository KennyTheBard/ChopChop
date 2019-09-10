package component

type BlueprintRegister map[string]Blueprint

func (register BlueprintRegister) GetBlueprint(bpName string) Blueprint {
	return register[bpName]
}

func (register BlueprintRegister) HasBlueprint(bpName string) bool {
	_, ok := register[bpName]
	return ok
}

func (register BlueprintRegister) AddBlueprint(bp Blueprint) {
	register[bp.GetItem()] = bp
}

func (register BlueprintRegister) GetAllBlueprints() []string {
	bps := make([]string, 0, len(register))
	for bp := range register {
		bps = append(bps, bp)
	}
	return bps
}
