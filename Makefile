info_color := \033[0;32m
no_color := \033[0m

linting := .linting
golangci_lint_ver := 2.8.0

.DEFAULT_GOAL := lint

${hxgo_linting}:
	mkdir -p ${hxgo_linting}

install-golangci-lint: ${hxgo_linting}
	@echo -e "$(info_color)==> $@ $(no_color)"
	@if [ -f "${hxgo_linting}/golangci-lint" ] && "${hxgo_linting}/golangci-lint" --version >/dev/null 2>&1; then \
		echo "golangci-lint is already installed"; \
	else \
		echo "golangci-lint not found or not executable, installing v${golangci_lint_ver}..."; \
		rm -f "${hxgo_linting}/golangci-lint"; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b ${hxgo_linting} v${golangci_lint_ver} || { echo "Failed to install golangci-lint"; exit 1; }; \
		echo "Installed golangci-lint, version: $$("${hxgo_linting}/golangci-lint" --version 2>&1)" || { echo "Failed to verify golangci-lint version"; exit 1; }; \
	fi

deps: install-golangci-lint
	@echo -e "$(info_color)==> $@ $(no_color)"

hxgo_linting := .linting

lint: deps
	@${hxgo_linting}/golangci-lint config path
	@${hxgo_linting}/golangci-lint config verify
	@${hxgo_linting}/golangci-lint run --config ${PWD}/.golangci.yaml

