
# EUD Project

Dette projekt laves primært af mig selv, men Theis har også tænkt sig at være med. Mikkel har også spurgt om han kan hjælpe, så han kommer nok også med laver lidt.

Jeg har tænkt mig at lave en language interpreter. Selve programmeringssproget jeg vil implementere bruger simpel C-like syntax. Dette projekt er delt op i dele:
- Language specification, grammar specifikation og AST specifikation. Disse beskriver hvad sproget er og hvordan det bruges, hvordan sproget er bygget op rent tekst-mæssigt og hvordan teksten skal parse'es, så vi kan bruge det i interpreteren. Dette kan laves på ca 1 dag.
- Lexer og parser. Lexeren laver tekst-input om til 'Tokens', så fx `a = 5` bliver til `IDENTIFIER:"a" ASSG_OP INT:"5"` Parseren laver disse tokens om til AST (Abstract Syntax Tree). Det gør den ved at samle tokens'ne fra en liste til et 'træ', så fx bliver `LPAREN INT:"3" ADD_OP INT:"4" RPAREN MUL_OP INT:"6"` bliver til `Multiplication
 {left: Addition {left: IntLiteral("3"), right: IntLiteral("4")}, right: IntLiteral("5")}`. Dette kan laves på 1-3 dage og er noget jeg har gjort før.
- (Mulighed, mindre sansynlig) AST interpreter. Dette er et 'simpelt' program, som løber igennem AST'en og evaluere den. AST-evaluering er langsomt, og derfor vil jeg hellere lave bytecode compilation, jeg har bare dette med som en option, hvis der er mangel på tid. Dette kan laves på 1-2 dage, og er noget jeg har gjort før.
- (Mulighed, mere sansynlig) IR/bytecode specifikation. Dette er lidt komplekst, da bytecode'en både skal kunne køres og man parseren skal kunne generere det. Tager 0.5-2 dage, men laves sammen med kompilering og runtime VM. Har kun prøvet lidt før. 
- (Mulighed, mere sansynlig) IR/bytecode kompilering. Tager 2-3 dage.
- (Mulighed, mere sansynlig) Bytecode runtime VM. Relativt nemt. Tager 1-2 dage.
- (Mulighed, meget lidt sansynlig) X86-64 Kompilering. Kun hvis der er tid.
- CLI tool. Tager et par timer.

Jeg føler selv at dette er et realistisk projekt, og det er klart noget jeg har lyst til. Hvis man skulle kigge på relevansen i forhold til min uddannelse, vil dette nok gå ind under programmering generelt, da der ikke er noget specifikt web design, app development, database maintenance eller noget library/framework/technologi specifikt. Er dette projekt lidt ambitiøst? ja, men jeg prøver ikke at lave det nye C++. Jeg er som sådan ligeglad med performance, ease of use, maintainability, det eneste jeg vil, er at have en semi-fungurende interpreter. I forhold til implementeringssprog, tænker jeg Golang, da vi godt kan lide sproget og begge har brugt det før. Sproget er ikke et endeligt valg, vi kunne godt finde på at skifte efter forholdene. Det er et CLI projekt, fordi vi (Mikkel) laver en CLI-frontend. 

## Lidt om implementering

### Memory management

Lidt uvigtigt, da det ikke er et seriøst sprog. Tænker bare [Reference counting](https://en.wikipedia.org/wiki/Reference_counting), da det er markant nemmere end [Tracing garbage collection](https://en.wikipedia.org/wiki/Tracing_garbage_collection).

### Eksempel på kode

```
// explicit typing
let a: i32 = 3 + 4 * 5

let b: u32 = (3 + 4) * 5

// lineending ';' optional
let pi: f64 = 3.14;

if a > 4 then
    print("yes")
end

func myFunc() -> i32
    let h: string = "hello world"
    print(h)
    return 1
end
```

### Eksempel på AST

```
Program [
    Declaration {
        name: "a",
        type: "i32",
        value: Addition {
            left: IntLiteral("3"),
            right: Multiplication {
                left: IntLiteral("4"),
                right: IntLiteral("5"),
            },
        },
    },
    Declaration {
        name: "b",
        type: "u32",
        value: Multiplication {
            left: Addition {
                left: IntLiteral("3"),
                right: IntLiteral("4"),
            },
            right: IntLiteral("5"),
        },
    },
    Declaration {
        name: "pi",
        type: "f64",
        value: FloatLiteral("3.14"),
    },
    IfStatement {
        condition: CmpLessThan {
            left: VarAccess("a"),
            right: IntLiteral("5");
        },
        body: Statements [
            Declaration {
                name: "h",
                type: "string",
                value: StringLiteral("hello world"),
            },
            FuncCall {
                name: "print",
                args: Values [
                    VarAccess("h"),
                ],
            },
            Return {
                value: IntLiteral("1"),
            },
        ],
    }
]
```

### Eksempel på IR/Bytecode

Det er usikkert om det endelige produkt kommer til at bruge IR eller bytecode, det kommer lige an på hvor lang tid det andet tager.

Dette er heller ikke det endelige produkt, bare en approksimering af hvordan det kommer til at se ud.

Compileren laver det heller ikke til tekst, men bare til objekter af forskellige arter.

```
extern print

myString:
    define u64 3        ; length
    define char "yes"   ; data
    
myString2:
    define u64 11
    define char "hello world"

func void _start()
    ; i32 a = 3 + 4 * 5
    declare i32 0
    push i32 3      ; [] -> [3]
    push i32 4      ; [3] -> [3 4]
    push i32 5      ; [3 4] -> [3 4 5]
    multiply i32    ; [3 4 5] -> [3 20]
    add i32         ; [3 20] -> [23]
    store i32 0
    
    ; u32 b = (3 + 4) * 5
    declare u32 1
    push i32 3
    push i32 4
    push i32 5
    add i32
    multiply i32
    store u32 1
    
    ; if a > 4 then
    load 0          ; [] -> [3]
    push i32 4      ; [3] -> [3 4]
    cmp             ; [3 4] -> [], certain flags get set, enabling use of  jlte (jump if less than)
    jlt ._start_ifend_0
    
    ;     print("yes")
    push myString
    call print
._start_ifend_0:

    ; func myFunc() -> i32
func i32 myFunc()
    ; string h = "hello world"
    declare ptr 3
    push myString2
    store ptr 3
    
    ; print(h)
    load ptr 3
    call print
    
    ; return 1
    push i32 1
    return
```

### AST interpreter

Den simpleste måde at køre koden er AST walking, dette er slow af, men meget nemt at implementere.

```java
Value evaluateMultiplicationNode(MultiplicationNode node) {
    Value left = evaluateExpression(node.left);
    Value right = evaluateExpression(node.right);
    if (!(left instanceof Int32Value) || !(right.type instanceof Int32Value))
        throw new FuckYouCrashAndBurnException();
    return new Int32Value(left.value() * right.value());
}
```




