#!/bin/bash
#TODO Enable
#COMPILE_TIMES=30
#STARTUP_TIMES=30
#LOAD_TIMES=30

COMPILE_TIMES=1
STARTUP_TIMES=1
LOAD_TIMES=1

function check(){
    prepareDocker "$1" "$3"
    compileTime "$1" "$2"
    startup     "$1"
    load        "$1" "$3"
    cleanDocker "$1"
}

function prepareDocker () {
    # Delete everything and set up postgres new
    docker-compose stop
    docker-compose rm -f
    docker-compose up --build -d postgres
    if [ "$2" = "" ]
    then
        echo "No setup is needed"
    else
        echo "Setup DB"
        docker-compose run $1 $2
        if [ $? -ne 0 ]
        then
            fail "Could not prepare DB $1"
        fi
    fi
}

function compileTime(){
    for (( i=0; i<COMPILE_TIMES; i++))
    do
        #make a clean first as we want to measure a full rebuild
        clean "$1" "$2"

        #Build the application and store the time needed to results
        startNS=$(date +"%s%N")
        compile "$1" "$2"
        buildImage "$1"
        endNS=$(date +"%s%N")
        compiletime=$(echo "scale=2;($endNS-$startNS)/1000000000" | bc)
        echo "$1, Compile time, $compiletime" >> results.csv
    done
}

function clean() {
    pushd ../$1
    if [ "$2" = "mvn" ]
    then
        ./mvnw clean
    elif [ "$2" = "gradle" ]
    then
        ./gradlew clean
    elif [ "$2" = "npm" ]
    then
        npm run clean
    elif [ "$2" = "go" ]
    then
        go clean
    elif [ "$2" = "none" ]
    then
        echo "No build needed"
    else
        popd
        fail "Do not know how to build $2"
    fi
    if [ $? -ne 0 ]
    then
        popd
        fail "Could not clean folder $1"
    fi
    popd
}

function compile(){
    pushd ../$1
    if [ "$2" = "mvn" ]
    then
        ./mvnw package
    elif [ "$2" = "gradle" ]
    then
        ./gradlew assemble
    elif [ "$2" = "npm" ]
    then
        npm install
        npm run build
    elif [ "$2" = "go" ]
    then
        export CGO_ENABLED=0
        go build
    elif [ "$2" = "none" ]
    then
        echo "No build needed"
    else
        popd
        fail "Do not know how to build $2"
    fi
    if [ $? -ne 0 ]
    then
        popd
        fail "Could not build folder $1"
    fi
    popd
}
 
function buildImage(){
    docker-compose build $1
    if [ $? -ne 0 ]
    then
        fail "Could not build image $1"
    fi
}

function startup(){
    for (( start=0; start<STARTUP_TIMES; start++))
    do
        #Recreate the container to always have a startup from null
        disposeContainer "$1"

        #Start the container and measure how long it takes untill we get a valid result
        startNS=$(date +"%s%N")
        startContainer "$1"
        endNS=$(date +"%s%N")
        startuptime=$(echo "scale=2;($endNS-$startNS)/1000000000" | bc)
        echo "$1, Startup time, $startuptime" >> results.csv

        #Measure memory
        memory=$(docker stats --format "{{.MemUsage}}" --no-stream "compare_$1_1" | awk 'match($0,/[0-9\.]+/) {print substr($0, RSTART, RLENGTH)}')
        echo "$1, Memory Usage (Startup), $memory" >> results.csv

        #Make sure container runs normally
        checkContainer "$1"
    done

    #Stop container again
    disposeContainer "$1"
}

function disposeContainer() {
    docker-compose stop $1
    docker-compose rm -f $1
}

function startContainer() {
    docker-compose up -d $1
    cameUp=0
    for (( i=0; i<100; i++))
    do
        sleep 0.3
        curl -s http://localhost:8080/issue/550e8400-e29b-41d4-a716-446655440000/ | grep "This is a test" > /dev/null
        if [ $? -eq 0 ]
        then
            return;
        fi;
    done
    curl http://localhost:8080/issue/550e8400-e29b-41d4-a716-446655440000/ -v
    fail "Container could not start"
}

function checkContainer() {
    curl -s http://localhost:8080/issue/ | grep "This is a test" > /dev/null
    if [ $? -ne 0 ]
    then
        curl http://localhost:8080/issue/ -v
        fail "Failed GET ALL for $1"
    fi;

    #Create a new entry
    curl -X POST http://localhost:8080/issue/ \
        -d '{"id":"550e8400-e29b-41d4-a728-446655440000","name":"Test 123", "description":"Test 28"}' \
        -H "Content-Type: application/json" 
    curl -s http://localhost:8080/issue/ | grep "Test 28" > /dev/null
    if [ $? -ne 0 ]
    then
        curl http://localhost:8080/issue/ -v
        fail "Failed CREATE for $1"
    fi;

    #Patch new entry
    curl -X PATCH http://localhost:8080/issue/550e8400-e29b-41d4-a728-446655440000/ \
        -d '{"description":"Test NEW"}' \
       	-H "Content-Type: application/json" 
    curl -s http://localhost:8080/issue/ | grep "Test NEW" > /dev/null
    if [ $? -ne 0 ]
    then
        curl http://localhost:8080/issue/ -v
        fail "Failed PATCH for $1"
    fi;

    #Delete new entry
    curl -X DELETE http://localhost:8080/issue/550e8400-e29b-41d4-a728-446655440000/
    curl -s http://localhost:8080/issue/ | grep "Test NEW" > /dev/null
    if [ $? -eq 0 ]
    then
        curl http://localhost:8080/issue/ -v
        fail "Failed DELETE for $1"
    fi;
}

function fail() {
    echo "$1"  1>&2;
    exit -1
}

function load() {
    for (( load=0; load<LOAD_TIMES; load++))
    do
        prepareForLoad "$1" "$2"
        startNS=$(date +"%s%N")
        jmeter -n -t loadtest.jmx -j jmeter.out -l jmeter.log -L DEBUG
        endNS=$(date +"%s%N")
        memory=$(docker stats --format "{{.MemUsage}}" --no-stream "compare_$1_1" | awk 'match($0,/[0-9\.]+/) {print substr($0, RSTART, RLENGTH)}')
        loadtime=$(echo "scale=2;($endNS-$startNS)/1000000000" | bc)

        tail -1 jmeter.out | grep "Err: *0 ("
        if [ $? -ne 0 ]
        then
            echo "$1, Memory Usage (Load), FAIL" >> results.csv
            echo "$1, Load Time, FAIL" >> results.csv
        else
            echo "$1, Memory Usage (Load), $memory" >> results.csv
            echo "$1, Load Time, $loadtime" >> results.csv
        fi
    done
}

function prepareForLoad() {
    # We have to freshly set up the container (incl db) to avoid follow up effects
    prepareDocker "$1" "$2"
    sleep 10;
    startContainer "$1"
}

function cleanDocker() {
    docker stop compare_$1_1
    docker rm -f compare_$1_1
    docker rmi -f compare_$1
    docker image prune -f
    docker volume prune -f
}

#Check all needed software is installed
#sudo apt-get install \
#	jmeter docker docker-compose  \
#	golang-1.14-go  \
#	npm \
#	openjdk-11-jdk \
#	python3-venv


# Remove the old result file
rm -f results.csv
check "python-falcon"    "none"
check "python-django"    "none"    "python manage.py migrate"
#check "micronaut-graal"  "gradle"
#check "spring"           "mvn"
#check "node-js"          "npm"
#check "node-ts"          "npm"
#check "go"               "go"
cat results.csv;
