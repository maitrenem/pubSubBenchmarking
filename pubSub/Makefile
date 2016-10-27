COMPS=nanomsg\
	redis

all: exe

exe: $(COMPS)
	@$(foreach f,$^, make -C $(f) exe;)

clean: $(COMPS)
	@$(foreach f,$^, make -C $(f) clean;)
