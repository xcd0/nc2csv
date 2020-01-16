all:
	make clean
	cd nc2csv/src && make release && make
	git add -A .
	git commit
	git push

tag:
	cd nc2csv/src && make release && make
	cd nc2csv && rm -rf ${ARG} && mkdir ${ARG} && cp -rf build/* ${ARG}
	git tag -a ${ARG} -m ${ARG}
	git push origin ${ARG}

clean:
	find . -name *.html -type f -print0 | xargs -0 rm -rf
	find . -name *.csv -type f -print0 | xargs -0 rm -rf
	find . -name .*.swp -type f -print0 | xargs -0 rm -rf
	find . -name .*.*.swp -type f -print0 | xargs -0 rm -rf

