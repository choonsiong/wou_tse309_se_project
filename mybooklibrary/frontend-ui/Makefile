## dist: Build dist
dist:
	@echo "Building dist..."
	@npm run build
	@tar zcvf dist.tar.gz dist/
	@echo "Dist built!"


## clean: Clean up dist
clean:
	@echo "Cleaning..."
	@rm -rf dist
	@rm dist.tar.gz
	@echo "Cleaned!"