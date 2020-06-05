srv=$1

protos="
    agv
    area
    cameramonitor
    config
    devices
    ledscreen
    log
    lpr
    material
    opc
    order
    point
    rack
    racklot
    racktype
    role
    route
    script
    segment
    station
    system
    task
    user
"

if [ ! -n "${srv}" ]; then
    for p in ${protos}; do
        echo protoc --proto_path=. --micro_out=. --go_out=. proto/${p}/${p}.proto
        protoc --proto_path=. --micro_out=. --go_out=. proto/${p}/${p}.proto
    done
else
    echo protoc --proto_path=. --micro_out=. --go_out=. proto/${srv}/${srv}.proto
    protoc --proto_path=. --micro_out=. --go_out=. proto/${srv}/${srv}.proto 
fi

