package models

import "github.com/jackc/pgtype"

type (
	Object struct {
		ID   pgtype.UUID
		Data pgtype.JSONB
	}
	Objects []Object
)

// type Order struct {
// 	OrderUID    pgtype.Varchar
// 	TrackNumber pgtype.Varchar
// 	Entry       pgtype.Varchar

// 	Delivery struct {
// 		Name    pgtype.Varchar
// 		Phone   pgtype.Varchar
// 		Zip     pgtype.Varchar
// 		City    pgtype.Varchar
// 		Address pgtype.Varchar
// 		Region  pgtype.Varchar
// 		Email   pgtype.Varchar
// 	}

// 	Payment struct {
// 		Transaction  pgtype.Varchar
// 		RequestID    pgtype.Varchar
// 		Currency     pgtype.Varchar
// 		Provider     pgtype.Varchar
// 		Amount       pgtype.Int4
// 		PaymentDT    pgtype.Int4
// 		Bank         pgtype.Varchar
// 		DeliveryCost pgtype.Int4
// 		GoodsTotal   pgtype.Int4
// 		CustomFee    pgtype.Int4
// 	}

// 	Items []struct {
// 		ChrtID      pgtype.Int4
// 		TrackNumber pgtype.Varchar
// 		Price       pgtype.Int4
// 		Rid         pgtype.Varchar
// 		Name        pgtype.Varchar
// 		Sale        pgtype.Int4
// 		Size        pgtype.Varchar
// 		TotalPrice  pgtype.Int4
// 		NmID        pgtype.Int4
// 		Brand       pgtype.Varchar
// 		Status      pgtype.Int4
// 	}

// 	Locale            pgtype.Varchar
// 	InternalSignature pgtype.Varchar
// 	CustomerID        pgtype.Varchar
// 	DeliveryService   pgtype.Varchar
// 	Shardkey          pgtype.Varchar
// 	SmID              pgtype.Int4
// 	DateCreated       pgtype.Timestamptz
// 	OofShard          pgtype.Int4
// }
