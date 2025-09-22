package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type QRAsset struct {
	ProductID                        string   `json:"productId"`
	ProductNameEn                    string   `json:"product_name_en"`
	ProductNameBn                    string   `json:"product_name_bn"`
	SpeciesEn                        string   `json:"species_en"`
	SpeciesBn                        string   `json:"species_bn"`
	ProductImage                     string   `json:"product_image"`
	ProcessingTypeEn                 string   `json:"processing_type_en"`
	ProcessingTypeBn                 string   `json:"processing_type_bn"`
	DateOfHarvesting                 string   `json:"date_of_harvesting"`
	DateOfPackaging                  string   `json:"date_of_packaging"`
	ExpiredDate                      string   `json:"expired_date"`
	MRP                              float64  `json:"mrp"`
	HasBlastFreezer                  bool     `json:"has_blast_freezer"`
	HasIQF                           bool     `json:"has_iqf"`
	HasVacuumPackage                 bool     `json:"has_vacuum_package"`
	HasFoodGradePackageLDPE4         bool     `json:"has_food_grade_package_ldpe_4"`
	StorageEn                        string   `json:"storage_en"`
	StorageBn                        string   `json:"storage_bn"`
	WashingEn                        []string `json:"washing_en"`
	WashingBn                        []string `json:"washing_bn"`
	HasFreezerVanTransportation      bool     `json:"has_freezer_van_transportation"`
	CookingTemperatureEn             string   `json:"cooking_temperature_en"`
	CookingTemperatureBn             string   `json:"cooking_temperature_bn"`
	BatchNumber                      string   `json:"batch_number"`
	LotNumber                        string   `json:"lot_number"`
	NetWeight                        float64  `json:"net_weight"`
	CertificationEn                  []string `json:"certification_en"`
	CertificationBn                  []string `json:"certification_bn"`
	SourceOfFishEn                   string   `json:"source_of_fish_en"`
	SourceOfFishBn                   string   `json:"source_of_fish_bn"`
	ProductionLatitude               float64  `json:"production_latitude"`
	ProductionLongitude              float64  `json:"production_longitude"`
	ProducerOrganizationEn           string   `json:"producer_organization_en"`
	ProducerOrganizationBn           string   `json:"producer_organization_bn"`
	FishCollectionCenterLatitude     float64  `json:"fish_collection_center_latitude"`
	FishCollectionCenterLongitude    float64  `json:"fish_collection_center_longitude"`
	CollectorOrganizationEn          string   `json:"collector_organization_en"`
	CollectorOrganizationBn          string   `json:"collector_organization_bn"`
	FishProcessingUnitLatitude       float64  `json:"fish_processing_unit_latitude"`
	FishProcessingUnitLongitude      float64  `json:"fish_processing_unit_longitude"`
	ProcessorOrganizationEn          string   `json:"processor_organization_en"`
	ProcessorOrganizationBn          string   `json:"processor_organization_bn"`
	DocType                          string   `json:"docType"`
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assets := []QRAsset{
		{
			ProductID: "1",
			ProductNameEn: "Frozen Hilsa Fish",
			ProductNameBn: "Frozen Hilsa Fish",
			SpeciesEn: "Hilsa",
			SpeciesBn: "Hilsa",
			ProductImage: "https://fish.com/hilsa.jpg",
			ProcessingTypeEn: "Frozen",
			ProcessingTypeBn: "Frozen",
			DateOfHarvesting: "2025-09-01",
			DateOfPackaging: "2025-09-03",
			ExpiredDate: "2026-03-01",
			MRP: 1200.5,
			HasBlastFreezer: true,
			HasIQF: false,
			HasVacuumPackage: true,
			HasFoodGradePackageLDPE4: true,
			StorageEn: "Cold Storage Dhaka",
			StorageBn: "Cold Storage Dhaka",
			WashingEn: []string{"Filtered water", "Arsenic"},
			WashingBn: []string{"Filtered water", "Arsenic"},
			HasFreezerVanTransportation: true,
			CookingTemperatureEn: "N/A",
			CookingTemperatureBn: "N/A",
			BatchNumber: "BATCH-001",
			LotNumber: "LOT-001",
			NetWeight: 2.5,
			CertificationEn: []string{"ISO22000", "HACCP"},
			CertificationBn: []string{"ISO22000", "HACCP"},
			SourceOfFishEn: "Padma River",
			SourceOfFishBn: "Padma River",
			ProductionLatitude: 23.8103,
			ProductionLongitude: 90.4125,
			ProducerOrganizationEn: "Padma Fisheries Ltd",
			ProducerOrganizationBn: "Padma Fisheries Ltd",
			FishCollectionCenterLatitude: 23.90,
			FishCollectionCenterLongitude: 90.44,
			CollectorOrganizationEn: "Dhaka Fish Collectors",
			CollectorOrganizationBn: "Dhaka Fish Collectors",
			FishProcessingUnitLatitude: 23.75,
			FishProcessingUnitLongitude: 90.39,
			ProcessorOrganizationEn: "Bangladesh Fish Processing Ltd",
			ProcessorOrganizationBn: "Bangladesh Fish Processing Ltd",
			DocType: "asset",
		},
	}

	for _, asset := range assets {
		key := fmt.Sprintf("QR:%s", asset.ProductID)
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}
		if err := ctx.GetStub().PutState(key, assetJSON); err != nil {
			return fmt.Errorf("failed to put to world state: %v", err)
		}
	}
	return nil
}

func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, productID string) (bool, error) {
	key := fmt.Sprintf("QR:%s", productID)
	assetJSON, err := ctx.GetStub().GetState(key)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}
	return assetJSON != nil, nil
}

func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, assetJSON string) error {
	var asset QRAsset
	if err := json.Unmarshal([]byte(assetJSON), &asset); err != nil {
		return fmt.Errorf("failed to unmarshal asset: %v", err)
	}

	exists, err := s.AssetExists(ctx, asset.ProductID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", asset.ProductID)
	}

	asset.DocType = "asset"
	key := fmt.Sprintf("QR:%s", asset.ProductID)
	assetBytes, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	if err := ctx.GetStub().PutState(key, assetBytes); err != nil {
		return err
	}

	event := map[string]string{"productId": asset.ProductID}
	eventBytes, _ := json.Marshal(event)
	if err := ctx.GetStub().SetEvent("QRCreated", eventBytes); err != nil {
		return fmt.Errorf("failed to set event: %v", err)
	}

	return nil
}


func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, productID string) (*QRAsset, error) {
	key := fmt.Sprintf("QR:%s", productID)
	assetJSON, err := ctx.GetStub().GetState(key)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", productID)
	}

	var asset QRAsset
	if err := json.Unmarshal(assetJSON, &asset); err != nil {
		return nil, err
	}
	return &asset, nil
}