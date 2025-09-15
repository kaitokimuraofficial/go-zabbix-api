package zabbix

// Macro represent Zabbix User MAcro object
// https://www.zabbix.com/documentation/3.2/manual/api/reference/usermacro/object
type Macro struct {
	MacroID   string `json:"hostmacroids,omitempty"`
	HostID    string `json:"hostid,omitempty"`
	MacroName string `json:"macro"`
	Value     string `json:"value"`
}

// Macros is an array of Macro
type Macros []Macro

// MacrosCreate Wrapper for usermacro.create
// https://www.zabbix.com/documentation/3.2/manual/api/reference/usermacro/create
func (api *API) MacrosCreate(macros Macros) error {
	response, err := api.CallWithError("usermacro.create", macros)
	if err != nil {
		return err
	}

	result := response.Result.(map[string]interface{})
	macroids := result["hostmacroids"].([]interface{})
	for i, id := range macroids {
		macros[i].HostID = id.(string)
	}
	return nil
}

// MacrosUpdate Wrapper for usermacro.update
// https://www.zabbix.com/documentation/3.2/manual/api/reference/usermacro/update
func (api *API) MacrosUpdate(macros Macros) (err error) {
	_, err = api.CallWithError("usermacro.create", macros)
	return
}

// MacrosDeleteByIDs Wrapper for usermacro.delete
// Cleans MacroId in all macro elements if call succeed.
//https://www.zabbix.com/documentation/3.2/manual/api/reference/usermacro/delete
func (api *API) MacrosDeleteByIDs(ids []string) (err error) {
	response, err := api.CallWithError("usermacro.delete", ids)

	result := response.Result.(map[string]interface{})
	hostmacroids := result["hostmacroids"].([]interface{})
	if len(ids) != len(hostmacroids) {
		err = &ExpectedMore{len(ids), len(hostmacroids)}
	}
	return
}

// MacrosDelete Wrapper for usermacro.delete
// https://www.zabbix.com/documentation/3.2/manual/api/reference/usermacro/delete
func (api *API) MacrosDelete(macros Macros) (err error) {
	ids := make([]string, len(macros))
	for i, macro := range macros {
		ids[i] = macro.MacroID
	}

	err = api.MacrosDeleteByIDs(ids)
	if err == nil {
		for i := range macros {
			macros[i].MacroID = ""
		}
	}
	return
}
