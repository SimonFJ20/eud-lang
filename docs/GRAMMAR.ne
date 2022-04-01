
program             ->  toplevel_statements

toplevel_statements ->  (_ toplevel_statement (_nl_ toplevel_statement):*):? _

toplevel_statement  ->  exported
                    |   import
                    |   statement

exported            ->  "export" __ exportable

exportable          ->  type_def
                    |   trait_def
                    |   struct_def
                    |   func_def
                    |   inferred_init_stmt
                    |   typed_init_stmt
                    |   declaration_stmt

import              ->  "import" __ STRING

statements          ->  (_ statement (_nl_ statement):*):?

statement           ->  type_def
                    |   trait_def
                    |   struct_def
                    |   func_def
                    |   return
                    |   match_stmt
                    |   switch_stmt
                    |   for
                    |   while
                    |   break
                    |   continue
                    |   if
                    |   if_else
                    |   inferred_init_stmt
                    |   typed_init_stmt
                    |   declaration_stmt
                    |   expression

type_def            ->  "typedef" __ identifier _ "=" _ type

trait_def           ->  "trait" __ identifier __ trait_declarations __ "end"

trait_declarations  ->  (_ typed_declaration (_ "," _ typed_declaration):* ",":?):? _

struct_def          ->  "struct" __ identifier _ "(" struct_traits ")" __ struct_properties __ "end"

struct_traits       ->  (_ identifier (_ "," _ identifier):* ",":?):? _

struct_properties   ->  (_ struct_property (_nl_ struct_property):*):?

struct_property     ->  typed_declaration
                    |   typed_init
                    |   inferred_init
                    |   method
                    |   private_method

private_method      ->  "private" __ method

method              ->  "func" __ identifier _ "(" _ identifier _ (", " typed_declarations):? ")" _ "->" _ type __ statements __ "end"

func_def            ->  "func" __ identifier _ "(" typed_declarations ")" _ "->" _ type __ statements __ "end"

return              ->  "return" __ expression
                    |   "return"

match_stmt          ->  "match" __ match

switch_stmt         ->  "switch" __ match

match               ->  expression __ match_cases __ "end"

match_cases         ->  (_ match_case (_ "," _ match_case):* ",":?):? _

match_case          ->  expression _ "=>" _ statements __ "end"

for                 ->  for_c_like
                    |   for_each

for_c_like          ->  "for" _ "(" _ (for_declaration _):? ";" _ (expression _):? ";" _ (expression _):? ")" _ "do" __ statements __ "end"

for_declaration     ->  inferred_init_stmt
                    |   typed_init_stmt

for_each            ->  "for" __ declareable __ "in" __ expression __ "do" __ statements __ "end"

while               ->  "while" __ expression __ "do" __ statements __ "end"

break               ->  "break"

continue            ->  "continue"

if                  ->  "if" __ expression __ "then" __ statements __ "end"

if_else             ->  "if" __ expression __ "then" __ statements __ "else" __ statements "end"

inferred_init_stmt  ->  "let" __ inferred_init

inferred_init       ->  declareable _ ":=" _ expression

typed_init_stmt     ->  "let" __ typed_init

typed_init          ->  typed_declaration _ "=" _ expression

declaration_stmt    ->  "let" __ typed_declaration

typed_object        ->  type _ object_literal

object_literal      ->  "{" key_value_pairs "}"

key_value_pairs     ->  (_ key_value_pair (_ "," _ key_value_pair):* ",":?):? _

key_value_pair      ->  identifier _ ":" _ expression

typed_array_literal ->  type _ array_literal

array_literal       ->  "[" expressions "]"

lamda_literal       ->  single_expr_lam_lit
                    |   multi_expr_lambda

single_expr_lambda  ->  identifier _ "=>" _ expression

multi_expr_lambda   ->  "(" declarations ")" _ "=>" expression

unpacked_array      ->  "[" declarations "]"

renamed_identifier  ->  identifier __ "as" __ identifier

declarations        ->  (_ declareable (_ "," _ declareable):* ",":?):? _

typed_declarations  ->  (_ typed_declaration (_ "," _ typed_declaration):* ",":?):? _

typed_declaration   ->  declareable _ ":" _ type

declareable         ->  unpacked_array
                    |   unpacked_object
                    |   renamed_identifier
                    |   identifier

type                ->  array_type
                    |   object_type
                    |   lamda_type
                    |   KEYWORD
                    |   identifier

array_type          ->  type _ "[" _ "]"

object_type         ->  "{" typed_declarations "}"

function_type       ->  "(" typed_declarations ")" _ "->" _ type

expressions         ->  (_ expression (_ "," _ expression):* ",":?):? _

expression          ->  assignment

assignment          ->  ternary _ "=" _ assignment
                    |   ternary

ternary             ->  logical_or _ "?" _ ternary _ ":" _ ternary
                    |   logical_or

logical_or          ->  logical_and _ "||" _ logical_or
                    |   logical_and

logical_and         ->  bitwise_or _ "&&" _ logical_and
                    |   bitwise_or

bitwise_or          ->  bitwise_xor _ "|" _ bitwise_or
                    |   bitwise_xor

bitwise_xor         ->  bitwise_and _ "^" _ bitwise_xor
                    |   bitwise_and

bitwise_and         ->  equality _ "&" _ bitwise_and
                    |   equality

equality            ->  inequality _ "==" _ equality
                    |   inequality

inequality          ->  compare_lt _ "!=" _ inequality
                    |   compare_lt

compare_lt          ->  compare_lte _ "<" _ compare_lt
                    |   compare_lte

compare_lte         ->  compare_gt _ "<=" _ compare_lte
                    |   compare_gt

compare_gt          ->  compare_gte _ ">" _ compare_gt
                    |   compare_gte

compare_gte         ->  bit_shift_left _ ">=" _ compare_gte
                    |   bit_shift_left

bit_shift_left      ->  bit_shift_right _ "<<" _ bit_shift_left
                    |   bit_shift_right

bit_shift_right     ->  addition _ ">>" _ bit_shift_right
                    |   addition

addition            ->  subtraction _ "+" _ addition
                    |   subtraction

subtraction         ->  multiplication _ "-" _ subtraction
                    |   multiplication

multiplication      ->  division _ "*" _ multiplication
                    |   division

division            ->  remainder _ "/" _ division
                    |   remainder

remainder           ->  exponentation _ "%" _ remainder
                    |   exponentation

exponentation       ->  increment _ "**" _ exponentation
                    |   increment

logical_not         ->  "!" _ logical_not
                    |   bitwise_not

bitwise_not         ->  "~" _ bitwise_not
                    |   unary_plus

unary_plus          ->  "+" _ unary_plus
                    |   unary_negation

unary_negation      ->  "-" _ unary_negation
                    |   pre_increment

pre_increment       ->  "++" pre_increment
                    |   pre_decrement

pre_decrement       ->  "--" pre_decrement
                    |   post_increment

post_increment      ->  post_decrement _ "++"
                    |   post_decrement

post_decrement      ->  member_access _ "--"
                    |   member_access

member_access       ->  computed_member _ "." _ identifier
                    |   computed_member

computed_member     ->  func_call _ "[" _ expression _ "]"
                    |   func_call

func_call           ->  value _ "(" _ expressions _ ")"
                    |   value

value               ->  int_literal
                    |   float_literal
                    |   char_literal
                    |   string_literal
                    |   typed_array_literal
                    |   typed_object
                    |   var_access
                    |   "(" _ expression _ ")"
                    |   array_literal
                    |   object_literal

var_access          ->  identifier

int_literal         ->  INT
float_literal       ->  FLOAT
char_literal        ->  CHAR
string_literal      ->  STRING

identifier          ->  IDENTIFIER

_nl_                ->  (_ /[\n;]/ (_ /[\n;/]):*):? _
_                   ->  __:?
__                  ->  /[\s]+/

