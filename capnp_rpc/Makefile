ROLES := server client
TYPES := unary serverStreaming clientStreaming clientStreamingV2 bidirectionalStreaming
RUN_COMMANDS := $(foreach role,$(ROLES),$(foreach type,$(TYPES),$(type)_$(role)))

compile:
	capnp compile -I ./capnp/go-capnp/std -ogo ./capnp/schema/*.capnp
	mv ./capnp/schema/*.capnp.go ./capnp/

$(RUN_COMMANDS):
	go run ./$(subst _,/,$@)