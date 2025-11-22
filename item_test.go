package zabbix_test

import (
	"testing"

	zapi "github.com/claranet/go-zabbix-api"
)

func testCreateItem(host *zapi.Host, t *testing.T) *zapi.Item {
	items := zapi.Items{{
		HostID: host.HostID,
		Key:    "key.lala.laa",
		Name:   "name for key",
		Type:   zapi.ZabbixTrapper,
		Delay:  "0",
	}}
	err := testGetAPI(t).ItemsCreate(items)
	if err != nil {
		t.Fatal(err)
	}
	return &items[0]
}

func testDeleteItem(item *zapi.Item, t *testing.T) {
	err := testGetAPI(t).ItemsDelete(zapi.Items{*item})
	if err != nil {
		t.Fatal(err)
	}
}

func TestItems(t *testing.T) {
	api := testGetAPI(t)

	group := testCreateHostGroup(t)
	defer testDeleteHostGroup(group, t)

	host := testCreateHost(group, t)
	defer testDeleteHost(host, t)

	item := testCreateItem(host, t)

	_, err := api.ItemGetByID(item.ItemID)
	if err != nil {
		t.Fatal(err)
	}

	item.Name = "another name"
	item.HostID = ""
	err = api.ItemsUpdate(zapi.Items{*item})
	if err != nil {
		t.Error(err)
	}

	testDeleteItem(item, t)
}
