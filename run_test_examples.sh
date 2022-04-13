
set -xe

clear
mypy parser.py

clear
go build

clear
./eud examples/functions-simple.eud

clear
./eud examples/functions.eud

clear
./eud examples/if-else.eud

clear
./eud examples/if-statement.eud

clear
./eud examples/initialisation.eud

clear
./eud examples/math-simple.eud

clear
./eud examples/sys-print-i32.eud

clear
./eud examples/sys-put-i32.eud

clear
./eud examples/variables.eud

clear
./eud examples/while-print.eud --nodebug

clear
./eud examples/while.eud

set +xe

clear
echo "All example/* tests ran without error"
