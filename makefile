all:
	make clean
	git add -A . && git commit -m "update `date +"%Y.%m.%d.%H.%M.%S"`"
	git push

clean:
	find . -name *.html -type f | xargs rm -rf
	find . -name *.csv -type f | xargs rm -rf
	find . -name .*.swp -type f | xargs rm -rf
	find . -name .*.*.swp -type f | xargs rm -rf

