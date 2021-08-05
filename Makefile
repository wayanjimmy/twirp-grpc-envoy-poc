include .env
export

wire:
	cd app && wire

generate:
	/usr/bin/buf-Linux-x86_64 generate --path ./protos

gen_doc:
	./doc_gen.sh