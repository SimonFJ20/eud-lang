from enum import Enum, auto
import sys
from typing import List, Optional

# clear && mypy parser.py && python3 parser.py examples/math.eud

class File:
    def __init__(self, filename: str) -> None:
        self.filename = filename
    
    def text(self):
        with open(self.filename) as file:
            return file.read()

class Position:
    def __init__(self, row: int, col: int, filename: str) -> None:
        self.row = row
        self.col = col
        self.filename = filename
    
    def copy(self):
        return Position(self.row, self.col, self.filename)

    def next(self, linefeed: bool):
        if linefeed:
            self.row += 1
            self.col = 1
        else:
            self.col += 1

    def __repr__(self) -> str:
        return f'{self.filename}:{self.row}:{self.col}:'

    def to_json(self) -> str:
        return f'{{"type":"Position","row":{self.row},"col":{self.col},"filename":"{self.filename}"}}'

def fail(msg: str, pos: Optional[Position]):
    print(f'FAILED: {msg}')
    if pos: print(f'    at {pos}')
    exit(1)

class TT(Enum):
    EOF = auto()
    IDENTIFIER = auto()
    KEYWORD = auto()
    INT = auto()
    LPAREN = auto()
    RPAREN = auto()
    LBRACKET = auto()
    RBRACKET = auto()
    LBRACE = auto()
    RBRACE = auto()
    ADD_OP = auto()
    SUB_OP = auto()
    MUL_OP = auto()
    DIV_OP = auto()
    MOD_OP = auto()
    EXP_OP = auto()
    ASGN_OP = auto()
    CMP_LT_OP = auto()
    CMP_LTE_OP = auto()
    CMP_GT_OP = auto()
    CMP_GTE_OP = auto()
    CMP_EQ_OP = auto()
    CMP_NE_OP = auto()
    LNOT_OP = auto()
    COLON = auto()
    COMMA = auto()

def tokentype_to_string(t: TT) -> str:
    if t == TT.EOF:             return 'EOF'
    elif t == TT.IDENTIFIER:    return 'IDENTIFIER'
    elif t == TT.KEYWORD:       return 'KEYWORD'
    elif t == TT.INT:           return 'INT'
    elif t == TT.LPAREN:        return 'LPAREN'
    elif t == TT.RPAREN:        return 'RPAREN'
    elif t == TT.LBRACKET:      return 'LBRACKET'
    elif t == TT.RBRACKET:      return 'RBRACKET'
    elif t == TT.LBRACE:        return 'LBRACE'
    elif t == TT.RBRACE:        return 'RBRACE'
    elif t == TT.ADD_OP:        return 'ADD_OP'
    elif t == TT.SUB_OP:        return 'SUB_OP'
    elif t == TT.MUL_OP:        return 'MUL_OP'
    elif t == TT.DIV_OP:        return 'DIV_OP'
    elif t == TT.MOD_OP:        return 'MOD_OP'
    elif t == TT.EXP_OP:        return 'EXP_OP'
    elif t == TT.ASGN_OP:       return 'ASGN_OP'
    elif t == TT.CMP_LT_OP:     return 'CMP_LT_OP'
    elif t == TT.CMP_LTE_OP:    return 'CMP_LTE_OP'
    elif t == TT.CMP_GT_OP:     return 'CMP_GT_OP'
    elif t == TT.CMP_GTE_OP:    return 'CMP_GTE_OP'
    elif t == TT.CMP_EQ_OP:     return 'CMP_EQ_OP'
    elif t == TT.CMP_NE_OP:     return 'CMP_NE_OP'
    elif t == TT.LNOT_OP:       return 'LOG_NOT'
    elif t == TT.COLON:         return 'COLON'
    elif t == TT.COMMA:         return 'COMMA'
    else: raise Exception('unexhaustive')

class Token:
    def __init__(self, type: TT, value: str, fp: Position) -> None:
        self.type = type
        self.value = value
        self.fp = fp

    def __repr__(self) -> str:
        return f"[{str(self.type)[3:]}:'{self.value}']"

    def to_json(self) -> str:
        tstr = tokentype_to_string(self.type)
        fpstr = self.fp.to_json()
        return f'{{"type": "Token","tokenType":"{tstr}","value":"{self.value}","fp":{fpstr}}}'

KEYWORDS: List[str] = [
    'if',
    'else',
    'for',
    'while',
    'break',
    'func',
    'return',
    'let',
    'u8',
    'u16',
    'u32',
    'u64',
    'i8',
    'i16',
    'i32',
    'i64',
    'char',
    'usize',
    'uptr',
]

class Lexer:
    def tokenize(self, file: File) -> List[Token]:
        text = file.text()
        if len(text) <= 0:
            raise fail('no text to be tokenized', None)
        self.text = text
        self.pos = 0
        self.c = text[0]
        self.done = False
        self.fp = Position(1, 1, file.filename)
        tokens: List[Token] = []
        while not self.done:
            if self.c in '\n\t ':
                self.next()
            elif self.c in 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_':
                tokens.append(self.make_name())
            elif self.c in '123456789':
                tokens.append(self.make_int())
            elif self.c == '0':
                tokens.append(Token(TT.INT, self.c, self.fp.copy()))
                self.next()
            elif self.c == '(':
                tokens.append(Token(TT.LPAREN, self.c, self.fp.copy()))
                self.next()
            elif self.c == ')':
                tokens.append(Token(TT.RPAREN, self.c, self.fp.copy()))
                self.next()
            elif self.c == '[':
                tokens.append(Token(TT.LBRACKET, self.c, self.fp.copy()))
                self.next()
            elif self.c == ']':
                tokens.append(Token(TT.RBRACKET, self.c, self.fp.copy()))
                self.next()
            elif self.c == '{':
                tokens.append(Token(TT.LBRACE, self.c, self.fp.copy()))
                self.next()
            elif self.c == '}':
                tokens.append(Token(TT.RBRACE, self.c, self.fp.copy()))
                self.next()
            elif self.c == '+':
                tokens.append(Token(TT.ADD_OP, self.c, self.fp.copy()))
                self.next()
            elif self.c == '-':
                tokens.append(Token(TT.SUB_OP, self.c, self.fp.copy()))
                self.next()
            elif self.c == '*':
                tokens.append(self.make_mul_or_exp_op())
            elif self.c == '/':
                tokens.append(Token(TT.DIV_OP, self.c, self.fp.copy()))
                self.next()
            elif self.c == '%':
                tokens.append(Token(TT.MOD_OP, self.c, self.fp.copy()))
                self.next()
            elif self.c == '=':
                tokens.append(self.make_asgn_or_eq_op())
            elif self.c == ':':
                tokens.append(Token(TT.COLON, self.c, self.fp.copy()))
                self.next()
            elif self.c == ',':
                tokens.append(Token(TT.COMMA, self.c, self.fp.copy()))
                self.next()
            elif self.c == '<':
                tokens.append(self.make_lt_or_lte_op())
            elif self.c == '>':
                tokens.append(self.make_gt_or_gte_op())
            elif self.c == '!':
                tokens.append(self.make_log_not_or_ne_op())
            else:
                raise fail(f"unexpected character '{self.c}'", self.fp.copy())
        tokens.append(Token(TT.EOF, '\0', self.fp.copy()))
        return tokens

    def make_name(self) -> Token:
        value = self.c
        self.next()
        while not self.done and self.c in 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_':
            value += self.c
            self.next()
        if value in KEYWORDS:
            return Token(TT.KEYWORD, value, self.fp.copy())
        else:
            return Token(TT.IDENTIFIER, value, self.fp.copy())
    
    def make_int(self) -> Token:
        value = self.c
        self.next()
        while not self.done and self.c in '1234567890':
            value += self.c
            self.next()
        return Token(TT.INT, value, self.fp.copy())

    def make_mul_or_exp_op(self) -> Token:
        value = self.c
        self.next()
        if self.c == '*':
            value += self.c
            self.next()
            return Token(TT.EXP_OP, value, self.fp.copy())
        else:
            return Token(TT.MUL_OP, value, self.fp.copy())

    def make_asgn_or_eq_op(self) -> Token:
        value = self.c
        self.next()
        if self.c == '=':
            value += self.c
            self.next()
            return Token(TT.CMP_EQ_OP, value, self.fp.copy())
        else:
            return Token(TT.ASGN_OP, value, self.fp.copy())

    def make_lt_or_lte_op(self) -> Token:
        value = self.c
        self.next()
        if self.c == '=':
            value += self.c
            self.next()
            return Token(TT.CMP_LTE_OP, value, self.fp.copy())
        else:
            return Token(TT.CMP_LT_OP, value, self.fp.copy())

    def make_gt_or_gte_op(self) -> Token:
        value = self.c
        self.next()
        if self.c == '=':
            value += self.c
            self.next()
            return Token(TT.CMP_GTE_OP, value, self.fp.copy())
        else:
            return Token(TT.CMP_GT_OP, value, self.fp.copy())

    def make_log_not_or_ne_op(self) -> Token:
        value = self.c
        self.next()
        if self.c == '=':
            value += self.c
            self.next()
            return Token(TT.CMP_NE_OP, value, self.fp.copy())
        else:
            return Token(TT.LNOT_OP, value, self.fp.copy())

    def next(self):
        self.pos += 1
        if self.pos < len(self.text):
            self.c = self.text[self.pos]
        else:
            self.done = True
            self.c = '\0'
        self.fp.next(self.c == '\n')


class Node:
    def __init__(self, fp: Position) -> None:
        self.fp = fp
    
    def typestr(self) -> str:
        return f'{type(self).__name__}Node'

    def __repr__(self) -> str: return f'{type(self).__name__}'

    def to_json(self) -> str:
        raise NotImplemented

class Statement(Node):
    def __init__(self, fp: Position) -> None:
        super().__init__(fp)

class Expression(Statement):
    def __init__(self, fp: Position) -> None:
        super().__init__(fp)

class Type(Node):
    def __init__(self, token: Token) -> None:
        super().__init__(token.fp)
        self.token = token

    def __repr__(self) -> str: return super().__repr__() + f'({self.token.value})'

    def to_json(self) -> str:
        return f'{{"type":"{self.typestr()}","token":{self.token.to_json()},"fp":{self.fp.to_json()}}}'

class TypedDecl(Node):
    def __init__(self, target: Token, type: Type) -> None:
        super().__init__(target.fp)
        self.target = target
        self.type = type

    def __repr__(self) -> str: return super().__repr__() + f'({self.target.value}, {self.type})'

    def to_json(self) -> str:
        return f'{{"type":"{self.typestr()}","target":{self.target.to_json()},"valueType":{self.type.to_json()},"fp":{self.fp.to_json()}}}'

class FuncDef(Statement):
    def __init__(self, target: Token, type: Type, params: List[TypedDecl], body: List[Statement], fp: Position) -> None:
        super().__init__(fp)
        self.target = target
        self.params = params
        self.type = type
        self.body = body
    
    def __repr__(self) -> str:
        paramstr = ','.join(map(lambda x:x.__repr__(), self.params))
        bodystr = ','.join(map(lambda x:x.__repr__(), self.body))
        return super().__repr__() + f'({self.target.value}, {self.type}, [{paramstr}], [{bodystr}])'

    def to_json(self) -> str:
        tstr = self.target.to_json()
        vtstr = self.type.to_json()
        paramstr = ','.join(map(lambda x:x.to_json(), self.params))
        bodystr = ','.join(map(lambda x:x.to_json(), self.body))
        return f'{{"type":"{self.typestr()}","target":{tstr},"valueType":{vtstr},"params":[{paramstr}],"body":[{bodystr}],"fp":{self.fp.to_json()}}}'

class Return(Expression):
    def __init__(self, value: Expression, fp: Position) -> None:
        super().__init__(fp)
        self.value = value

    def __repr__(self) -> str: return super().__repr__() + f'({self.value})'

    def to_json(self):
        return f'{{"type":"{self.typestr()}","value":{self.value.to_json()},"fp":{self.fp.to_json()}}}'

class While(Statement):
    def __init__(self, condition: Expression, body: List[Statement], fp: Position) -> None:
        super().__init__(fp)
        self.condition = condition
        self.body = body
    
    def __repr__(self) -> str:
        bodystr = ','.join(map(lambda x:x.__repr__(), self.body))
        return super().__repr__() + f'({self.condition}, [{bodystr}])'

    def to_json(self) -> str:
        bodystr = ','.join(map(lambda x:x.to_json(), self.body))
        cstr = self.condition.to_json()
        return f'{{"type":"{self.typestr()}","condition":{cstr},"body":[{bodystr}],"fp":{self.fp.to_json()}}}'

class IfElse(Statement):
    def __init__(self, condition: Expression, truthy: List[Statement], falsy: List[Statement], fp: Position) -> None:
        super().__init__(fp)
        self.condition = condition
        self.truthy = truthy
        self.falsy = falsy
    
    def __repr__(self) -> str:
        truthystr = ','.join(map(lambda x:x.__repr__(), self.truthy))
        falsystr = ','.join(map(lambda x:x.__repr__(), self.falsy))
        return super().__repr__() + f'({self.condition}, [{truthystr}], [{falsystr}])'

    def to_json(self) -> str:
        truthystr = ','.join(map(lambda x:x.to_json(), self.truthy))
        falsystr = ','.join(map(lambda x:x.to_json(), self.falsy))
        cstr = self.condition.to_json()
        return f'{{"type":"{self.typestr()}","condition":{cstr},"truthy":[{truthystr}],"falsy":[{falsystr}],"fp":{self.fp.to_json()}}}'

class If(Statement):
    def __init__(self, condition: Expression, body: List[Statement], fp: Position) -> None:
        super().__init__(fp)
        self.condition = condition
        self.body = body
    
    def __repr__(self) -> str:
        bodystr = ','.join(map(lambda x:x.__repr__(), self.body))
        return super().__repr__() + f'({self.condition}, [{bodystr}])'

    def to_json(self) -> str:
        bodystr = ','.join(map(lambda x:x.to_json(), self.body))
        cstr = self.condition.to_json()
        return f'{{"type":"{self.typestr()}","condition":{cstr},"body":[{bodystr}],"fp":{self.fp.to_json()}}}'

class VarInit(Statement):
    def __init__(self, target: Token, type: Type, value: Expression, fp: Position) -> None:
        super().__init__(fp)
        self.target = target
        self.type = type
        self.value = value
    
    def __repr__(self) -> str: return super().__repr__() + f'({self.target.value}, {self.type}, {self.value})'

    def to_json(self) -> str:
        return f'{{"type":"{self.typestr()}","target":{self.target.to_json()},"valueType":{self.type.to_json()},"value":{self.value.to_json()},"fp":{self.fp.to_json()}}}'

class VarDecl(Statement):
    def __init__(self, target: Token, type: Type, fp: Position) -> None:
        super().__init__(fp)
        self.target = target
        self.type = type
    
    def __repr__(self) -> str: return super().__repr__() + f'({self.target.value}, {self.type})'

    def to_json(self) -> str:
        return f'{{"type":"{self.typestr()}","target":{self.target.to_json()},"valueType":{self.type.to_json()},"fp":{self.fp.to_json()}}}'

class Assign(Expression):
    def __init__(self, target: Token, value: Expression, fp: Position) -> None:
        super().__init__(fp)
        self.target = target
        self.value = value

    def __repr__(self) -> str: return super().__repr__() + f'({self.target.value}, {self.value})'

    def to_json(self):
        return f'{{"type":"{self.typestr()}","target":{self.target.to_json()},"value":{self.value.to_json()},"fp":{self.fp.to_json()}}}'

class BinaryOperation(Expression):
    def __init__(self, left: Expression, right: Expression, fp: Position) -> None:
        super().__init__(fp)
        self.left = left
        self.right = right

    def __repr__(self) -> str: return super().__repr__() + f'({self.left}, {self.right})'

    def to_json(self) -> str:
        return f'{{"type":"{self.typestr()}","left":{self.left.to_json()},"right":{self.right.to_json()},"fp":{self.fp.to_json()}}}'

class NotEqual(BinaryOperation):
    def __init__(self, left: Expression, right: Expression, fp: Position) -> None:
        super().__init__(left, right, fp)

class Equal(BinaryOperation):
    def __init__(self, left: Expression, right: Expression, fp: Position) -> None:
        super().__init__(left, right, fp)

class GreaterThanOrEqual(BinaryOperation):
    def __init__(self, left: Expression, right: Expression, fp: Position) -> None:
        super().__init__(left, right, fp)

class LessThanOrEqual(BinaryOperation):
    def __init__(self, left: Expression, right: Expression, fp: Position) -> None:
        super().__init__(left, right, fp)

class GreaterThan(BinaryOperation):
    def __init__(self, left: Expression, right: Expression, fp: Position) -> None:
        super().__init__(left, right, fp)

class LessThan(BinaryOperation):
    def __init__(self, left: Expression, right: Expression, fp: Position) -> None:
        super().__init__(left, right, fp)

class Add(BinaryOperation):
    def __init__(self, left: Expression, right: Expression, fp: Position) -> None:
        super().__init__(left, right, fp)

class Sub(BinaryOperation):
    def __init__(self, left: Expression, right: Expression, fp: Position) -> None:
        super().__init__(left, right, fp)

class Mul(BinaryOperation):
    def __init__(self, left: Expression, right: Expression, fp: Position) -> None:
        super().__init__(left, right, fp)

class Div(BinaryOperation):
    def __init__(self, left: Expression, right: Expression, fp: Position) -> None:
        super().__init__(left, right, fp)

class Mod(BinaryOperation):
    def __init__(self, left: Expression, right: Expression, fp: Position) -> None:
        super().__init__(left, right, fp)

class Exp(BinaryOperation):
    def __init__(self, left: Expression, right: Expression, fp: Position) -> None:
        super().__init__(left, right, fp)

class FuncCall(Expression):
    def __init__(self, target: Expression, args: List[Expression]) -> None:
        super().__init__(target.fp)
        self.target = target
        self.args = args
    
    def __repr__(self) -> str:
        argstr = ','.join(map(lambda x:x.__repr__(), self.args))
        return super().__repr__() + f'({self.target}, [{argstr}])'

    def to_json(self):
        argstr = ','.join(map(lambda x:x.to_json(), self.args))
        return f'{{"type":"{self.typestr()}","target":{self.target.to_json()},"args":[{argstr}],"fp":{self.fp.to_json()}}}'

class Int(Expression):
    def __init__(self, token: Token) -> None:
        super().__init__(token.fp)
        self.token = token
    
    def __repr__(self) -> str: return f'{super().__repr__()}({self.token.value})'

    def to_json(self):
        return f'{{"type":"{self.typestr()}","token":{self.token.to_json()},"fp":{self.fp.to_json()}}}'

class Var(Expression):
    def __init__(self, token: Token) -> None:
        super().__init__(token.fp)
        self.token = token
    
    def __repr__(self) -> str: return f'{super().__repr__()}({self.token.value})'

    def to_json(self):
        return f'{{"type":"{self.typestr()}","token":{self.token.to_json()},"fp":{self.fp.to_json()}}}'

class Parser:
    def parse(self, tokens: List[Token]) -> List[Statement]:
        if len(tokens) <= 0:
            raise fail('no tokens to be parsed', None)
        self.tokens = tokens
        self.pos = 0
        self.t = tokens[0]
        self.done = False

        ast = self.make_program()
        return ast

    def make_program(self) -> List[Statement]:
        return self.make_statements()

    def make_statements(self) -> List[Statement]:
        statements: List[Statement] = []
        while not self.done and self.t.type not in [TT.RBRACE, TT.EOF]:
            statements.append(self.make_statement())
        return statements

    def make_statement(self) -> Statement:
        if self.t.value == 'func':
            return self.make_func_def()
        elif self.t.value == 'return':
            return self.make_return()
        elif self.t.value == 'let':
            return self.make_declaration_or_initialization()
        elif self.t.value == 'while':
            return self.make_while()
        elif self.t.value == 'if':
            return self.make_if_or_if_else()
        else:
            return self.make_expression()

    def make_func_def(self) -> FuncDef:
        fp = self.t.fp
        self.next()
        if self.t.type != TT.IDENTIFIER:
            fail(f"expected identifier, got {self.t}", self.t.fp)
        target = self.t
        self.next()
        if self.t.type != TT.LPAREN:
            fail(f"expected '(', got {self.t}", self.t.fp)
        self.next()
        params: List[TypedDecl] = []
        while not self.done and self.t.type != TT.RPAREN:
            params.append(self.make_typed_declaration())
            if self.t.type == TT.RPAREN:
                break
            elif self.t.type == TT.COMMA:
                self.next()
            else:
              fail(f"expected ',', got {self.t}", self.t.fp)
        if self.t.type != TT.RPAREN:
            fail(f"expected ')', got {self.t}", self.t.fp)
        self.next()
        if self.t.type != TT.COLON:
            fail(f'expected \':\', got {self.t}', self.t.fp)
        self.next()
        if self.t.value not in ['u8', 'u16', 'u32', 'u64', 'i8', 'i16', 'i32', 'i64', 'char', 'usize', 'uptr']:
            fail(f'expected keyword, got {self.t}', self.t.fp)
        type = self.make_type()
        if self.t.type != TT.LBRACE:
            fail("expected '{'" + f', got {self.t}', self.t.fp)
        self.next()
        body = self.make_statements()
        if self.t.type != TT.RBRACE:
            fail("expected '}'" + f', got {self.t}', self.t.fp)
        self.next()
        return FuncDef(target, type, params, body, fp)

    def make_return(self) -> Return:
        fp = self.t.fp
        self.next()
        value = self.make_expression()
        return Return(value, fp)

    def make_while(self) -> While:
        fp = self.t.fp
        self.next()
        if self.t.type != TT.LPAREN:
            fail(f"expected '(', got {self.t}", self.t.fp)
        self.next()
        condition = self.make_expression()
        if self.t.type != TT.RPAREN:
            fail(f"expected ')', got {self.t}", self.t.fp)
        self.next()
        if self.t.type != TT.LBRACE:
            fail("expected '{'" + f', got {self.t}', self.t.fp)
        self.next()
        body = self.make_statements()
        if self.t.type != TT.RBRACE:
            fail("expected '}'" + f', got {self.t}', self.t.fp)
        self.next()
        return While(condition, body, fp)

    def make_if_or_if_else(self) -> Statement:
        fp = self.t.fp
        self.next()
        if self.t.type != TT.LPAREN:
            fail(f"expected '(', got {self.t}", self.t.fp)
        self.next()
        condition = self.make_expression()
        if self.t.type != TT.RPAREN:
            fail(f"expected ')', got {self.t}", self.t.fp)
        self.next()
        if self.t.type != TT.LBRACE:
            fail("expected '{'" + f', got {self.t}', self.t.fp)
        self.next()
        body = self.make_statements()
        if self.t.type != TT.RBRACE:
            fail("expected '}'" + f', got {self.t}', self.t.fp)
        self.next()
        if self.t.value == 'else':
            self.next()
            if self.t.type != TT.LBRACE:
                fail("expected '{'" + f', got {self.t}', self.t.fp)
            self.next()
            falsy = self.make_statements()
            if self.t.type != TT.RBRACE:
                fail("expected '}'" + f', got {self.t}', self.t.fp)
            self.next()
            return IfElse(condition, body, falsy, fp)
        else:
            return If(condition, body, fp)

    def make_typed_declaration(self) -> TypedDecl:
        target = self.t
        self.next()
        if self.t.type != TT.COLON:
            fail(f'expected \':\', got {self.t}', self.t.fp)
        self.next()
        if self.t.value not in ['u8', 'u16', 'u32', 'u64', 'i8', 'i16', 'i32', 'i64', 'char', 'usize', 'uptr']:
            fail(f'expected keyword, got {self.t}', self.t.fp)
        type = self.make_type()
        return TypedDecl(target, type)

    def make_declaration_or_initialization(self) -> Statement:
        fp = self.t.fp
        self.next()
        if self.t.type != TT.IDENTIFIER:
            fail(f'expected identifier, got {self.t}', self.t.fp)
        target = self.t
        self.next()
        if self.t.type != TT.COLON:
            fail(f"expected ':', got {self.t}", self.t.fp)
        self.next()
        type = self.make_type()
        if self.t.type == TT.ASGN_OP:
            self.next()
            value = self.make_expression()
            return VarInit(target, type, value, fp)
        else:
            return VarDecl(target, type, fp)

    def make_type(self) -> Type:
        if self.t.value not in ['u8', 'u16', 'u32', 'u64', 'i8', 'i16', 'i32', 'i64', 'char', 'usize', 'uptr']:
            fail(f'expected keyword, got {self.t}', self.t.fp)
        token = self.t
        self.next()
        return Type(token)

    def make_expression(self) -> Expression:
        return self.make_assignment()

    def make_assignment(self) -> Expression:
        target = self.t
        self.next()
        if target.type == TT.IDENTIFIER and self.t.type == TT.ASGN_OP:
            self.next()
            value = self.make_expression()
            return Assign(target, value, target.fp)
        else:
            self.previus()
            return self.make_not_equal()

    def make_not_equal(self) -> Expression:
        left = self.make_equal()
        if self.t.type == TT.CMP_NE_OP:
            self.next()
            right = self.make_not_equal()
            return NotEqual(left, right, left.fp)
        else:
            return left

    def make_equal(self) -> Expression:
        left = self.make_greater_than_or_equal()
        if self.t.type == TT.CMP_EQ_OP:
            self.next()
            right = self.make_equal()
            return Equal(left, right, left.fp)
        else:
            return left

    def make_greater_than_or_equal(self) -> Expression:
        left = self.make_less_than_or_equal()
        if self.t.type == TT.CMP_GTE_OP:
            self.next()
            right = self.make_greater_than_or_equal()
            return GreaterThanOrEqual(left, right, left.fp)
        else:
            return left

    def make_less_than_or_equal(self) -> Expression:
        left = self.make_greater_than()
        if self.t.type == TT.CMP_LTE_OP:
            self.next()
            right = self.make_less_than_or_equal()
            return LessThanOrEqual(left, right, left.fp)
        else:
            return left

    def make_greater_than(self) -> Expression:
        left = self.make_less_than()
        if self.t.type == TT.CMP_GT_OP:
            self.next()
            right = self.make_greater_than()
            return GreaterThan(left, right, left.fp)
        else:
            return left

    def make_less_than(self) -> Expression:
        left = self.make_addition()
        if self.t.type == TT.CMP_LT_OP:
            self.next()
            right = self.make_less_than()
            return LessThan(left, right, left.fp)
        else:
            return left

    def make_addition(self) -> Expression:
        left = self.make_subtraction()
        if self.t.type == TT.ADD_OP:
            self.next()
            right = self.make_addition()
            return Add(left, right, left.fp)
        else:
            return left

    def make_subtraction(self) -> Expression:
        left = self.make_multiplication()
        if self.t.type == TT.SUB_OP:
            self.next()
            right = self.make_subtraction()
            return Sub(left, right, left.fp)
        else:
            return left

    def make_multiplication(self) -> Expression:
        left = self.make_division()
        if self.t.type == TT.MUL_OP:
            self.next()
            right = self.make_multiplication()
            return Mul(left, right, left.fp)
        else:
            return left

    def make_division(self) -> Expression:
        left = self.make_modulus()
        if self.t.type == TT.DIV_OP:
            self.next()
            right = self.make_division()
            return Div(left, right, left.fp)
        else:
            return left

    def make_modulus(self) -> Expression:
        left = self.make_exponentation()
        if self.t.type == TT.MOD_OP:
            self.next()
            right = self.make_modulus()
            return Mod(left, right, left.fp)
        else:
            return left

    def make_exponentation(self) -> Expression:
        left = self.make_func_call()
        if self.t.type == TT.EXP_OP:
            self.next()
            right = self.make_exponentation()
            return Exp(left, right, left.fp)
        else:
            return left

    def make_func_call(self) -> Expression:
        target = self.make_value()
        if self.t.type == TT.LPAREN:
            self.next()
            args: List[Expression] = []
            while not self.done and self.t.type != TT.RPAREN:
                args.append(self.make_expression())
                if self.t.type == TT.RPAREN:
                    break
                elif self.t.type == TT.COMMA:
                    self.next()
                else:
                    fail(f'expected \',\', got {self.t}', self.t.fp)
            if self.t.type != TT.RPAREN:
                fail(f"expected ')', got {self.t}", self.t.fp)
            self.next()
            return FuncCall(target, args)
        else:
            return target

    def make_value(self) -> Expression:
        token = self.t
        self.next()
        if token.type == TT.INT:
            return Int(token)
        elif token.type == TT.IDENTIFIER:
            return Var(token)
        elif token.type == TT.LPAREN:
            expression = self.make_expression()
            if self.t.type != TT.RPAREN:
                fail('parenthesis not closed', token.fp)
            self.next()
            return expression
        else:
            raise fail(f"unexpected token {token}", token.fp.copy())

    def next(self):
        self.pos += 1
        if self.pos < len(self.tokens):
            self.t = self.tokens[self.pos]
        else:
            self.done = True
            self.t = Token(TT.EOF, '\0', Position(-1, -1, ''))

    def previus(self):
        self.pos -= 1
        self.t = self.tokens[self.pos]

def main():
    if len(sys.argv) <= 1:
        print('python3 parser.py <file>')
    filename = sys.argv[1]
    tokens = Lexer().tokenize(File(filename))
    # for i in tokens: print(i)
    ast = Parser().parse(tokens)
    # for i in ast: print(i.to_json())
    # print('{"type":"AST","body":[' + ','.join([i.to_json() for i in ast]) + ']}')
    res = '[' + ','.join([i.to_json() for i in ast]) + ']'
    if '-ofile' in sys.argv:
        with open('ast.temp.json', 'w') as f:
            f.write(res)
    else:
        print(res)

if __name__ == '__main__':
    main()
