  package initial
  
  // Const is a constant equal to 5. It's the best constant I've ever written. 😹
  const Const = 5
//      ^^^^^ definition Const.
//      documentation Const is a constant equal to 5. It's the best constant I've ever written. 😹
  
  // Docs for the const block itself.
  const (
   // ConstBlock1 is a constant in a block.
   ConstBlock1 = 1
// ^^^^^^^^^^^ definition ConstBlock1.
// documentation Docs for the const block itself.
  
   // ConstBlock2 is a constant in a block.
   ConstBlock2 = 2
// ^^^^^^^^^^^ definition ConstBlock2.
// documentation Docs for the const block itself.
  )
  
  // Var is a variable interface.
  var Var Interface = &Struct{Field: "bar!"}
//    ^^^ definition Var.
//    documentation Var is a variable interface.
//        ^^^^^^^^^ reference sg/initial/Interface#
//                     ^^^^^^ reference sg/initial/Struct#
//                            ^^^^^ reference sg/initial/Struct#Field.
  
  // unexportedVar is an unexported variable interface.
  var unexportedVar Interface = &Struct{Field: "bar!"}
//    ^^^^^^^^^^^^^ definition unexportedVar.
//    documentation unexportedVar is an unexported variable interface.
//                  ^^^^^^^^^ reference sg/initial/Interface#
//                               ^^^^^^ reference sg/initial/Struct#
//                                      ^^^^^ reference sg/initial/Struct#Field.
  
  // x has a builtin error type
  var x error
//    ^ definition x.
//    documentation x has a builtin error type
  
  var BigVar Interface = &Struct{
//    ^^^^^^ definition BigVar.
//           ^^^^^^^^^ reference sg/initial/Interface#
//                        ^^^^^^ reference sg/initial/Struct#
   Field: "bar!",
// ^^^^^ reference sg/initial/Struct#Field.
   Anonymous: struct {
// ^^^^^^^^^ reference sg/initial/Struct#Anonymous.
    FieldA int
//  ^^^^^^ definition local 0
    FieldB int
//  ^^^^^^ definition local 1
    FieldC int
//  ^^^^^^ definition local 2
   }{FieldA: 1337},
//   ^^^^^^ reference local 0
  }
  
  // What are docs, really?
  // I can't say for sure, I don't write any.
  // But look, a CAT!
  //
  //       |\      _,,,---,,_
  // ZZZzz /,`.-'`'    -.  ;-;;,_
  //      |,4-  ) )-,_. ,\ (  `'-'
  //     '---''(_/--'  `-'\_)
  //
  // It's sleeping! Some people write that as `sleeping` but Markdown
  // isn't allowed in Go docstrings, right? right?!
  var (
   // This has some docs
   VarBlock1 = "if you're reading this"
// ^^^^^^^^^ definition VarBlock1.
// documentation What are docs, really?
  
   VarBlock2 = "hi"
// ^^^^^^^^^ definition VarBlock2.
// documentation What are docs, really?
  )
  
  // Embedded is a struct, to be embedded in another struct.
  type Embedded struct {
//     ^^^^^^^^ definition sg/initial/Embedded#
//     documentation Embedded is a struct, to be embedded in another struct.
   // EmbeddedField has some docs!
   EmbeddedField string
// ^^^^^^^^^^^^^ definition sg/initial/Embedded#EmbeddedField.
   Field         string // conflicts with parent "Field"
// ^^^^^ definition sg/initial/Embedded#Field.
  }
  
  type Struct struct {
//     ^^^^^^ definition sg/initial/Struct#
   *Embedded
//  ^^^^^^^^ definition sg/initial/Struct#Embedded.
//  ^^^^^^^^ reference sg/initial/Embedded#
   Field     string
// ^^^^^ definition sg/initial/Struct#Field.
   Anonymous struct {
// ^^^^^^^^^ definition sg/initial/Struct#Anonymous.
    FieldA int
//  ^^^^^^ definition sg/initial/Struct#Anonymous.FieldA.
    FieldB int
//  ^^^^^^ definition sg/initial/Struct#Anonymous.FieldB.
    FieldC int
//  ^^^^^^ definition sg/initial/Struct#Anonymous.FieldC.
   }
  }
  
  // StructMethod has some docs!
  func (s *Struct) StructMethod() {}
//      ^ definition local 3
//         ^^^^^^ reference sg/initial/Struct#
//                 ^^^^^^^^^^^^ definition sg/initial/Struct#StructMethod().
//                 documentation StructMethod has some docs!
  
  func (s *Struct) ImplementsInterface() string { return "hi!" }
//      ^ definition local 4
//         ^^^^^^ reference sg/initial/Struct#
//                 ^^^^^^^^^^^^^^^^^^^ definition sg/initial/Struct#ImplementsInterface().
  
  func (s *Struct) MachineLearning(
//      ^ definition local 5
//         ^^^^^^ reference sg/initial/Struct#
//                 ^^^^^^^^^^^^^^^ definition sg/initial/Struct#MachineLearning().
   param1 float32, // It's ML, I can't describe what this param is.
// ^^^^^^ definition local 6
  
   // We call the below hyperparameters because, uhh, well:
   //
   //    ,-.       _,---._ __  / \
   //   /  )    .-'       `./ /   \
   //   (  (   ,'            `/    /|
   //    \  `-"             \'\   / |
   //     `.              ,  \ \ /  |
   //   /`.          ,'-`----Y   |
   //     (            ;        |   '
   //     |  ,-.    ,-'         |  /
   //     |  | (   |        hjw | /
   //     )  |  \  `.___________|/
   //     `--'   `--'
   //
   hyperparam2 float32,
// ^^^^^^^^^^^ definition local 7
   hyperparam3 float32,
// ^^^^^^^^^^^ definition local 8
  ) float32 {
   // varShouldNotHaveDocs is in a function, should not have docs emitted.
   var varShouldNotHaveDocs int32
//     ^^^^^^^^^^^^^^^^^^^^ definition local 9
  
   // constShouldNotHaveDocs is in a function, should not have docs emitted.
   const constShouldNotHaveDocs = 5
//       ^^^^^^^^^^^^^^^^^^^^^^ definition local 10
  
   // typeShouldNotHaveDocs is in a function, should not have docs emitted.
   type typeShouldNotHaveDocs struct{ a string }
//      ^^^^^^^^^^^^^^^^^^^^^ definition local 11
//                                    ^ definition local 12
  
   // funcShouldNotHaveDocs is in a function, should not have docs emitted.
   funcShouldNotHaveDocs := func(a string) string { return "hello" }
// ^^^^^^^^^^^^^^^^^^^^^ definition local 13
//                               ^ definition local 14
  
   return param1 + (hyperparam2 * *hyperparam3) // lol is this all ML is? I'm gonna be rich
//        ^^^^^^ reference local 6
//                  ^^^^^^^^^^^ reference local 7
//                                 ^^^^^^^^^^^ reference local 8
  }
  
  // Interface has docs too
  type Interface interface {
//     ^^^^^^^^^ definition sg/initial/Interface#
//     documentation Interface has docs too
   ImplementsInterface() string
// ^^^^^^^^^^^^^^^^^^^ definition sg/initial/Interface#ImplementsInterface.
  }
  
  func NewInterface() Interface { return nil }
//     ^^^^^^^^^^^^ definition sg/initial/NewInterface().
//                    ^^^^^^^^^ reference sg/initial/Interface#
  
  var SortExportedFirst = 1
//    ^^^^^^^^^^^^^^^^^ definition SortExportedFirst.
  
  var sortUnexportedSecond = 2
//    ^^^^^^^^^^^^^^^^^^^^ definition sortUnexportedSecond.
  
  var _sortUnderscoreLast = 3
//    ^^^^^^^^^^^^^^^^^^^ definition _sortUnderscoreLast.
  
  // Yeah this is some Go magic incantation which is common.
  //
  //  ,_     _
  //  |\\_,-~/
  //  / _  _ |    ,--.
  // (  @  @ )   / ,-'
  //  \  _T_/-._( (
  // /         `. \
  // |         _  \ |
  // \ \ ,  /      |
  //  || |-_\__   /
  // ((_/`(____,-'
  var _ = Interface(&Struct{})
//        ^^^^^^^^^ reference sg/initial/Interface#
//                   ^^^^^^ reference sg/initial/Struct#
  
  type _ = struct{}
  
  // crypto/tls/common_string.go uses this pattern..
  func _() {
  }
  
  // Go can be fun
  type (
   // And confusing
   X struct {
// ^ definition sg/initial/X#
// documentation Go can be fun
    bar string
//  ^^^ definition sg/initial/X#bar.
   }
  
   Y struct {
// ^ definition sg/initial/Y#
// documentation Go can be fun
    baz float64
//  ^^^ definition sg/initial/Y#baz.
   }
  )
  
