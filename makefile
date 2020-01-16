all:
	make clean
	git add -A . && git commit -m "update `date +"%Y.%m.%d.%H.%M.%S"`"
	git push

clean:
	find . -name *.html -type f -print0 | xargs -0 rm -rf
	find . -name *.csv -type f -print0 | xargs -0 rm -rf
	find . -name .*.swp -type f -print0 | xargs -0 rm -rf
	find . -name .*.*.swp -type f -print0 | xargs -0 rm -rf

