# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

CORES=2

include ../../SharedMakefile.mk

MAX_RESOURCE_COUNT=500000000

# define standard colors
ifneq (,$(findstring xterm,${TERM}))
	RED          := $(shell tput -Txterm setaf 1)
	YELLOW       := $(shell tput -Txterm setaf 3)
	RESET := $(shell tput -Txterm sgr0)
else
	RED          := ""
	YELLOW       := ""
	RESET        := ""
endif

# Overriding this target just to generate the project.properties file
polymorph_dafny :
	cd $(CODEGEN_CLI_ROOT); \
	$(GRADLEW) run --args="\
	--dafny-version $(DAFNY_VERSION) \
	--library-root $(LIBRARY_ROOT) \
	--properties-file $(PROJECT_ROOT)/$(STD_LIBRARY)/project.properties \
	--model $(PROJECT_ROOT)/$(STD_LIBRARY)/Model \
	--namespace aws.polymorph \
	--dependent-model $(PROJECT_ROOT)/$(STD_LIBRARY)/../Model \
	";

polymorph_dotnet :
	echo "Skipping polymorph_dotnet for StandardLibrary"

polymorph_java :
	echo "Skipping polymorph_java for StandardLibrary"

# Using this target for the side-effect of maintaining patches,
# Since Dafny Rust code generation is incomplete and also needs to be patched.
# That only works if you run transpile_rust first.
polymorph_rust :
	cd $(CODEGEN_CLI_ROOT); \
	$(GRADLEW) run --args="\
	--dafny-version $(DAFNY_VERSION) \
	--library-root $(LIBRARY_ROOT) \
	--patch-files-dir $(if $(DIR_STRUCTURE_V2),$(LIBRARY_ROOT)/codegen-patches/$(SERVICE),$(LIBRARY_ROOT)/codegen-patches) \
	--output-rust $(LIBRARY_ROOT)/runtimes/rust \
	--model $(PROJECT_ROOT)/$(STD_LIBRARY)/Model \
	--namespace aws.polymorph \
	--dependent-model $(PROJECT_ROOT)/$(STD_LIBRARY)/../Model \
	$(POLYMORPH_OPTIONS) \
	";

_polymorph_dependencies: 
	echo "Skipping _polymorph_dependencies for StandardLibrary"

# Override SharedMakefile's transpile_net to not transpile
# StandardLibrary as a dependency
transpile_net: | transpile_implementation_net transpile_test_net

# Override SharedMakefile's transpile_java to not transpile
# StandardLibrary as a dependency
transpile_java: | transpile_implementation_java transpile_test_java

# Override SharedMakefile's transpile_java to not transpile
# StandardLibrary as a dependency
transpile_rust: | transpile_implementation_rust

RUST_TEST_INDEX=test_for_rust

# Override the targets to relocate translated Rust code directly into the main crate,
# instead of the dafny_impl sub-crate.
# We also don't want to delete the src directory since we have an extern in there too.
_mv_implementation_rust:
	mv implementation_from_dafny-rust/src/implementation_from_dafny.rs runtimes/rust/src/implementation_from_dafny.rs
	rustfmt runtimes/rust/src/implementation_from_dafny.rs
	rm -rf implementation_from_dafny-rust
_mv_test_rust:
	mv tests_from_dafny-rust/src/tests_from_dafny.rs runtimes/rust/tests/tests_from_dafny.rs
	rustfmt runtimes/rust/tests/tests_from_dafny.rs
	rm -rf tests_from_dafny-rust

transpile_dependencies:
	echo "Skipping transpile_dependencies for StandardLibrary"

# Override SharedMakefile's transpile_implementation to not reference
# StandardLibrary as a Library
transpile_implementation:
	find ./dafny/**/$(SRC_INDEX_TRANSPILE)/ ./$(SRC_INDEX_TRANSPILE)/ -name 'Index.dfy' | sed -e 's/^/include "/' -e 's/$$/"/' | dafny \
		-stdin \
		-noVerify \
		-vcsCores:$(CORES) \
		-compileTarget:$(TARGET) \
		-spillTargetCode:3 \
		-compile:0 \
		-optimizeErasableDatatypeWrapper:0 \
		-compileSuffix:1 \
		-unicodeChar:0 \
		-functionSyntax:3 \
		-useRuntimeLib \
		$(DAFNY_OPTIONS) \
		-out $(OUT)

# Override SharedMakefile's build_java to not install
# StandardLibrary as a dependency
build_java: transpile_java
	gradle -p runtimes/java build

# Override SharedMakefile's mvn_local_deploy to
# issue warning
mvn_local_deploy:
	@echo "${RED}Warning!!${YELLOW} Installing TestModel's STD to Maven Local replaces ESDK's STD!\n$(RESET)" >&2
	gradle -p runtimes/java publishToMavenLocal

dafny_benerate: DAFNY_RUST_OUTPUT=\
	runtimes/rust/implementation_from_dafny
dafny_benerate: DAFNY_RUST_OUTPUT_FOLDER=\
	$(DAFNY_RUST_OUTPUT)-rust/src
dafny_benerate: DAFNY_RUST_OUTPUT_RS=\
	$(DAFNY_RUST_OUTPUT_FOLDER)/implementation_from_dafny.rs
dafny_benerate: DAFNY_RUST_OUTPUT_RS_DTR=\
	$(DAFNY_RUST_OUTPUT_FOLDER)/implementation_from_dafny-rs.dtr
dafny_benerate: DAFNY_RUST_IMPL=\
	runtimes/rust/src/implementation_from_dafny.rs
dafny_benerate:
	$(DAFNY) translate rs src/Index.dfy  --emit-uncompilable-code --no-verify   --optimize-erasable-datatype-wrapper false   --include-runtime   --compile-suffix true   --unicode-char false   --function-syntax 3   --allow-warnings   --output $(DAFNY_RUST_OUTPUT)
	rustfmt $(DAFNY_RUST_OUTPUT_RS)
	# echo "Diffs remaining till code generation works:"
	# diff -u $(DAFNY_RUST_OUTPUT_RS) $(DAFNY_RUST_IMPL) | wc -l
	mv $(DAFNY_RUST_OUTPUT_RS) $(DAFNY_RUST_IMPL)
	rm $(DAFNY_RUST_OUTPUT_RS_DTR);
	rm -r $(DAFNY_RUST_OUTPUT_FOLDER);