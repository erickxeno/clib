.PHONY : test

test:
	./test/run_local_test.sh

setup:
	./test/docker_script/setup.sh

clean:
	./test/docker_script/clean.sh
