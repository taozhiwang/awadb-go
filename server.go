package main

import (
	"context"
	"log"
	"net"

	pb ""

	"google.golang.org/grpc"
)

const (
	PORT = ":50001"
)

type AwaDBServerClient struct {
	pb.AwaDBServerServer
}

func (s *AwaDBServerClient) Create(ctx context.Context, in *pb.DBMeta) (*pb.ResponseStatus, error) {
	dbName := in.GetDbName()
	desc := in.GetDesc()
	tablesMeta := in.GetTablesMeta()

	log.Println("db_name: %s", dbName)
	log.Printf("Description: %s\n", desc)

	for _, tableMeta := range tablesMeta {
		tableName := tableMeta.GetName()
		tableDesc := tableMeta.GetDesc()
		fieldsMeta := tableMeta.GetFieldsMeta()

		log.Printf("  Table Name: %s\n", tableName)
		log.Printf("  Table Description: %s\n", tableDesc)

		// Iterate through fieldsMeta and print field data
		for _, fieldMeta := range fieldsMeta {
			fieldName := fieldMeta.GetName()
			fieldType := fieldMeta.GetType().String()
			isIndex := fieldMeta.GetIsIndex()
			isStore := fieldMeta.GetIsStore()
			reindex := fieldMeta.GetReindex()
			vecMeta := fieldMeta.GetVecMeta()

			log.Printf("    Field Name: %s\n", fieldName)
			log.Printf("	Field Type: %s\n", fieldType)
			log.Printf("	Is Index: %v\n", isIndex)
			log.Printf("	Is Store: %v\n", isStore)
			log.Printf("	Reindex: %v\n", reindex)

			if vecMeta != nil {
				dataType := vecMeta.GetDataType().String()
				dimension := vecMeta.GetDimension()
				storeType := vecMeta.GetStoreType()
				storeParam := vecMeta.GetStoreParam()
				hasSource := vecMeta.GetHasSource()

				log.Printf("		Vector Meta - Data Type: %s\n", dataType)
				log.Printf("		Vector Meta - Dimension: %d\n", dimension)
				log.Printf("		Vector Meta - Store Type: %s\n", storeType)
				log.Printf("		Vector Meta - Store Param: %s\n", storeParam)
				log.Printf("		Vector Meta - Has Source: %v\n", hasSource)
			}
		}
	}

	response := &pb.ResponseStatus{
		Code:       pb.ResponseCode_OK,
		OutputInfo: "Operation successful",
	}

	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAwaDBServerServer(s, &AwaDBServerClient{})

	log.Println("rpc server open")

	// Start server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
