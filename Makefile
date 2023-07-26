.PHONY: protos

protos:
	protoc -I protos/currency  --go-grpc_out=protos/ --go_out=protos/ protos/currency/currency.proto