(module
  (type $t0 (func))
  (type $t1 (func (param i64 i64 i64 i64)))
  (type $t2 (func (param i64 i64)))
  (import "env" "ask_external_data" (func $ask_external_data (type $t1)))
  (import "env" "set_return_data" (func $set_return_data (type $t2)))
  (func $prepare (export "prepare") (type $t0)
    (local $l0 i64)
    i64.const 1
    i64.const 1
    i64.const 1024
    tee_local $l0
    i64.const 4
    call $ask_external_data
    i64.const 2
    i64.const 2
    get_local $l0
    i64.const 4
    call $ask_external_data
    i64.const 3
    i64.const 3
    get_local $l0
    i64.const 4
    call $ask_external_data)
  (func $execute (export "execute") (type $t0)
    i64.const 1024
    i64.const 4
    call $set_return_data)
  (table $T0 1 1 anyfunc)
  (memory $memory (export "memory") 17)
  (data (i32.const 1024) "new beeb"))