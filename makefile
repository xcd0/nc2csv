all:
	make clean
	cd nc2csv/src && make release && make
	git add -A .
	git commit
	git push

tag:
	# make tag v=v1.0 のようにすること
	cd nc2csv/src && make release && make
	cd nc2csv && rm -rf ${v} && mkdir ${v} && cp -rf build/* ${v}
	git add -A .
	git commit -m "${v}"
	git push
	git tag -a ${v} -m ${v}
	git push origin ${v}

clean:
	find . -name *.html -type f -print0 | xargs -0 rm -rf
	find . -name *.csv -type f -print0 | xargs -0 rm -rf
	find . -name .*.swp -type f -print0 | xargs -0 rm -rf
	find . -name .*.*.swp -type f -print0 | xargs -0 rm -rf

