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
	SourceOfAgroCommoditiesEn        []string `json:"source_of_agro_commodities_en"`
	SourceOfAgroCommoditiesBn        []string `json:"source_of_agro_commodities_bn"`
	BatchNumber                      string   `json:"batch_number"`
	LotNumber                        string   `json:"lot_number"`
	NetWeight                        float64  `json:"net_weight"`
	CertificationEn                  []string `json:"certification_en"`
	CertificationBn                  []string `json:"certification_bn"`
	ProductionLatitude               float64  `json:"production_latitude"`
	ProductionLongitude              float64  `json:"production_longitude"`
	ProducerOrganizationEn           string   `json:"producer_organization_en"`
	ProducerOrganizationBn           string   `json:"producer_organization_bn"`
	CropCollectionCenterLatitude     float64  `json:"crop_collection_center_latitude"`
	CropCollectionCenterLongitude    float64  `json:"crop_collection_center_longitude"`
	CollectorOrganizationEn          string   `json:"collector_organization_en"`
	CollectorOrganizationBn          string   `json:"collector_organization_bn"`
	CropProcessingUnitLatitude       float64  `json:"crop_processing_unit_latitude"`
	CropProcessingUnitLongitude      float64  `json:"crop_processing_unit_longitude"`
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
			SourceOfAgroCommoditiesEn: []string{"Filtered water", "Arsenic"},
			SourceOfAgroCommoditiesBn: []string{"Filtered water", "Arsenic"},
			BatchNumber: "BATCH-001",
			LotNumber: "LOT-001",
			NetWeight: 2.5,
			CertificationEn: []string{"ISO22000", "HACCP"},
			CertificationBn: []string{"ISO22000", "HACCP"},
			ProductionLatitude: 23.8103,
			ProductionLongitude: 90.4125,
			ProducerOrganizationEn: "Padma Fisheries Ltd",
			ProducerOrganizationBn: "Padma Fisheries Ltd",
			CropCollectionCenterLatitude: 23.90,
			CropCollectionCenterLongitude: 90.44,
			CollectorOrganizationEn: "Dhaka Fish Collectors",
			CollectorOrganizationBn: "Dhaka Fish Collectors",
			CropProcessingUnitLatitude: 23.75,
			CropProcessingUnitLongitude: 90.39,
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