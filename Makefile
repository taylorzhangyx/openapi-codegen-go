.PHONY: api
api:
	bash autogen.sh

.PHONY: run
run:
	./taylorpetstore

.PHONY: buildrun
buildrun:
	make api
	make run
