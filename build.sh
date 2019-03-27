mkdir -p output/conf
cp -rf conf ./output/
export GO15VENDOREXPERIMENT="1"
go build -o ./output/run
