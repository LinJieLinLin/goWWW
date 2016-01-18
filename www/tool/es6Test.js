function foo() {
  let bar = true;

  if (bar) {
    var baz = 'hi!';
  }

  console.log(baz); // hi
}

foo();