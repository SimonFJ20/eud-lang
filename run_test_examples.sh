
set -xe

clear
mypy parser.py

clear
go build

clear
./eud examples/functions-simple.eud --pyparse

clear
./eud examples/functions.eud --pyparse

clear
./eud examples/if-else.eud --pyparse

clear
./eud examples/if-statement.eud --pyparse

clear
./eud examples/initialisation.eud --pyparse

clear
./eud examples/math-simple.eud --pyparse

clear
./eud examples/sys-print-i32.eud --pyparse

clear
./eud examples/sys-put-i32.eud --pyparse

clear
./eud examples/variables.eud --pyparse

clear
./eud examples/while-print.eud --pyparse --nodebug

clear
./eud examples/while.eud --pyparse

set +xe

clear
echo "All example/* tests ran without error"
