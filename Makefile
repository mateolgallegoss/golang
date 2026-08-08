.DEFAULT_GOAL := help

# El Path de un directorio de Go al cuál se busca compilar el módulo.
DIR ?= hello-world

.PHONY: help fmt vet build clear

# El comando de ayuda explicando el funcionamiento
help:
	@echo "Targets:"
	@echo "  make build DIR=<dir>   Build the Go module/program in the given sub-directory."
	@echo "  make clean DIR=<dir>   Remove all compiler temp/trash files, keeping the executable."
	@echo ""
	@echo "Examples:"
	@echo "  make build            Build the default module (hello-world)."
	@echo "  make build DIR=data-types"
	@echo "  make build DIR=functions && make clean DIR=functions"

# Validar que el directorio elejido contenga un módulo de go
$(DIR)/main.go:
	@test -f go.mod && echo "Using module: $(DIR)" || (echo "No go.mod in $(DIR)" && exit 1)

fmt:
	go fmt ./...

build: $(DIR)/main.go
	cd $(DIR) && go build -o $(notdir $(DIR)).exe

# Elimina cualquier artefacto o residuo que haya quedado de la compilación, exceptuando el ejecutable
clean:
	@cd $(DIR) && go clean -cache -testcache
	@cd $(DIR) && del /q *.test *.out *.cover *.tmp *.exe~ 2>nul
	@echo Cleaned $(DIR): kept $(DIR)\$(notdir $(DIR)).exe
