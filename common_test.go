package novapost_test

import (
	"testing"

	. "github.com/sirkostya009/go-novapost"
)

func TestClient_GetTimeIntervals(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	intervals, err := c.GetTimeIntervals(TimeIntervalRequest{
		DateTime:         "01.01.2026",
		RecipientCityRef: KyivCityRef,
	})
	if err != nil {
		t.Error(err)
	}

	if !intervals.Success {
		t.Log("No time intervals found", intervals)
	}
}

func TestClient_GetCargoTypes(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	cargoTypes, err := c.GetCargoTypes()
	if err != nil {
		t.Error(err)
	}

	if !cargoTypes.Success {
		t.Log("No cargo types found", cargoTypes)
	}
}

func TestClient_GetBackwardDeliveryCargoTypes(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	cargoTypes, err := c.GetBackwardDeliveryCargoTypes()
	if err != nil {
		t.Error(err)
	}

	if !cargoTypes.Success {
		t.Log("No backward delivery cargo types found", cargoTypes)
	}
}

func TestClient_GetPalletsList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	pallets, err := c.GetPalletsList()
	if err != nil {
		t.Error(err)
	}

	if !pallets.Success {
		t.Log("No pallets found", pallets)
	}
}

func TestClient_GetTypesOfPayersForRedelivery(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	payers, err := c.GetTypesOfPayersForRedelivery()
	if err != nil {
		t.Error(err)
	}

	if !payers.Success {
		t.Log("No payers found", payers)
	}
}

func TestClient_GetPackList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	packs, err := c.GetPackList(PackRequest{
		Length:           10,
		Width:            2,
		Height:           15,
		VolumetricWeight: 8.54,
	})
	if err != nil {
		t.Error(err)
	}

	if !packs.Success {
		t.Log("No packs found", packs)
	}
}

func TestClient_GetTiresWheelsList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	tiresWheels, err := c.GetTiresWheelsList()
	if err != nil {
		t.Error(err)
	}

	if !tiresWheels.Success {
		t.Log("No tires and wheels found", tiresWheels)
	}
}

func TestClient_GetCargoDescriptionList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	cargoDescriptions, err := c.GetCargoDescriptionList(CargoDescriptionRequest{
		FindByString: "Вантаж",
		Page:         1,
	})
	if err != nil {
		t.Error(err)
	}

	if !cargoDescriptions.Success {
		t.Log("No cargo descriptions found", cargoDescriptions)
	}
}

func TestClient_GetMessageCodeText(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	message, err := c.GetMessageCodeText()
	if err != nil {
		t.Error(err)
	}

	if !message.Success {
		t.Log("No message found", message)
	}
}

func TestClient_GetServiceTypes(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	serviceTypes, err := c.GetServiceTypes()
	if err != nil {
		t.Error(err)
	}

	if !serviceTypes.Success {
		t.Log("No service types found", serviceTypes)
	}
}

func TestClient_GetOwnershipFormsList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	c := newTestClient()
	ownershipForms, err := c.GetOwnershipFormsList()
	if err != nil {
		t.Error(err)
	}

	if !ownershipForms.Success {
		t.Log("No ownership forms found", ownershipForms)
	}
}
