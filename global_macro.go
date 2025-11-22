package zabbix

// GlobalMacro represent Zabbix global User Macro object
// https://www.zabbix.com/documentation/7.0/manual/api/reference/usermacro/object
type GlobalMacro struct {
	MacroID     string `json:"globalmacroid,omitempty"`
	MacroName   string `json:"macro"`
	Value       string `json:"value"`
	Description string `json:"description,omitempty"`
}

// GlobalMacros is an array of GlobalMacro
type GlobalMacros []GlobalMacro

// GlobalMacrosGet Wrapper for usermacro.get
// https://www.zabbix.com/documentation/7.0/manual/api/reference/usermacro/get
func (api *API) GlobalMacrosGet(params Params) (res GlobalMacros, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	params["globalmacro"] = true
	err = api.CallWithErrorParse("usermacro.get", params, &res)
	return
}

// GlobalMacroGetByID Get macro by macro ID if there is exactly 1 matching macro
func (api *API) GlobalMacroGetByID(id string) (res *GlobalMacro, err error) {
	macros, err := api.GlobalMacrosGet(Params{"globalmacroids": id})
	if err != nil {
		return
	}

	if len(macros) == 1 {
		res = &macros[0]
	} else {
		e := ExpectedOneResult(len(macros))
		err = &e
	}
	return
}

// GlobalMacrosCreate Wrapper for usermacro.createglobal
// https://www.zabbix.com/documentation/7.0/manual/api/reference/usermacro/createglobal
func (api *API) GlobalMacrosCreate(globalmacros GlobalMacros) error {
	response, err := api.CallWithError("usermacro.createglobal", globalmacros)
	if err != nil {
		return err
	}

	result := response.Result.(map[string]interface{})
	globalmacroids := result["globalmacroids"].([]interface{})
	for i, id := range globalmacroids {
		globalmacros[i].MacroID = id.(string)
	}
	return nil
}

// GlobalMacrosUpdate Wrapper for usermacro.updateglobal
// https://www.zabbix.com/documentation/7.0/manual/api/reference/usermacro/updateglobal
func (api *API) GlobalMacrosUpdate(globalmacros GlobalMacros) (err error) {
	_, err = api.CallWithError("usermacro.updateglobal", globalmacros)
	return
}

// GlobalMacrosDeleteByIDs Wrapper for usermacro.deleteglobal
// Cleans MacroId in all macro elements if call succeed.
// https://www.zabbix.com/documentation/7.0/manual/api/reference/usermacro/deleteglobal
func (api *API) GlobalMacrosDeleteByIDs(ids []string) (err error) {
	response, err := api.CallWithError("usermacro.deleteglobal", ids)

	result := response.Result.(map[string]interface{})
	globalmacroids := result["globalmacroids"].([]interface{})
	if len(ids) != len(globalmacroids) {
		err = &ExpectedMore{len(ids), len(globalmacroids)}
	}
	return
}

// GlobalMacrosDelete Wrapper for usermacro.deleteglobal
// https://www.zabbix.com/documentation/7.0/manual/api/reference/usermacro/deleteglobal
func (api *API) GlobalMacrosDelete(globalmacros GlobalMacros) (err error) {
	ids := make([]string, len(globalmacros))
	for i, globalmacro := range globalmacros {
		ids[i] = globalmacro.MacroID
	}

	err = api.GlobalMacrosDeleteByIDs(ids)
	if err == nil {
		for i := range globalmacros {
			globalmacros[i].MacroID = ""
		}
	}
	return
}
