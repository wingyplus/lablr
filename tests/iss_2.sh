LABLR_HOME="`pwd`"

function teardown() 
{
	rm -rf "$LABLR_HOME/iss_2_result.xml"
}

function _eq()
{
	local actual=$1
	local expect=$2

	local different=`diff $actual $expect`

	local err=$?

	if [ $err -ne 0 ]
	then
		echo $different
	else
		echo "ISS #2 PASS"
	fi
}

go run cmd/lablr.go "$LABLR_HOME/tests/iss_2_Model.xml" --share-search-config 

_eq "./iss_2_result.xml" "./tests/iss_2.xml"

teardown