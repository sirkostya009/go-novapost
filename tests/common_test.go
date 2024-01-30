package tests

import (
	"github.com/sirkostya009/go-novapost"
	"testing"
)

func TestGetTimeIntervals(t *testing.T) {
	intervals, err := c.GetTimeIntervals(novapost.TimeIntervalRequest{
		DateTime:         "01.01.2026",
		RecipientCityRef: KyivCityRef,
	})
	if err != nil {
		t.Error(err)
	}

	if !intervals.Success {
		t.Error("No time intervals found", intervals)
	}
}

func TestGetCargoTypes(t *testing.T) {
	cargoTypes, err := c.GetCargoTypes()
	if err != nil {
		t.Error(err)
	}

	if !cargoTypes.Success {
		t.Error("No cargo types found", cargoTypes)
	}
}

func TestGetBackwardDeliveryCargoTypes(t *testing.T) {
	cargoTypes, err := c.GetBackwardDeliveryCargoTypes()
	if err != nil {
		t.Error(err)
	}

	if !cargoTypes.Success {
		t.Error("No backward delivery cargo types found", cargoTypes)
	}
}

func TestGetPalletsList(t *testing.T) {
	pallets, err := c.GetPalletsList()
	if err != nil {
		t.Error(err)
	}

	if !pallets.Success {
		t.Error("No pallets found", pallets)
	}
}

func TestGetTypesOfPayersForRedelivery(t *testing.T) {
	payers, err := c.GetTypesOfPayersForRedelivery()
	if err != nil {
		t.Error(err)
	}

	if !payers.Success {
		t.Error("No payers found", payers)
	}
}

func TestGetPackList(t *testing.T) {
	packs, err := c.GetPackList(novapost.PackRequest{
		Lengthstring:           "10",
		Widthstring:            "2",
		Heightstring:           "15",
		VolumetricWeightstring: "8.54",
	})
	if err != nil {
		t.Error(err)
	}

	if !packs.Success {
		t.Error("No packs found", packs)
	}
}

func TestGetTiresWheelsList(t *testing.T) {
	tiresWheels, err := c.GetTiresWheelsList()
	if err != nil {
		t.Error(err)
	}

	if !tiresWheels.Success {
		t.Error("No tires and wheels found", tiresWheels)
	}
}

func TestGetCargoDescriptionList(t *testing.T) {
	cargoDescriptions, err := c.GetCargoDescriptionList(novapost.CargoDescriptionRequest{
		FindByString: "Вантаж",
		Page:         1,
	})
	if err != nil {
		t.Error(err)
	}

	if !cargoDescriptions.Success {
		t.Error("No cargo descriptions found", cargoDescriptions)
	}
}

func TestGetMessageCodeText(t *testing.T) {
	message, err := c.GetMessageCodeText()
	if err != nil {
		t.Error(err)
	}

	if !message.Success {
		t.Error("No message found", message)
	}
}

func TestGetServiceTypes(t *testing.T) {
	serviceTypes, err := c.GetServiceTypes()
	if err != nil {
		t.Error(err)
	}

	if !serviceTypes.Success {
		t.Error("No service types found", serviceTypes)
	}
}

func TestGetOwnershipFormsList(t *testing.T) {
	ownershipForms, err := c.GetOwnershipFormsList()
	if err != nil {
		t.Error(err)
	}

	if !ownershipForms.Success {
		t.Error("No ownership forms found", ownershipForms)
	}
}
