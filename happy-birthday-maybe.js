'use strict';

class Maybe {

  constructor (x) {
    this.$val = x;
  }
  
  static of (x) {
    return new Maybe(x);
  }
  
  get isNothing () {
    return this.$val !== 'c9s';
  }
  
  map(fn) {
    return this.isNothing ? this : new Maybe.of(fn(this.$val));
  }
  
  inspect() {
    return this.isNothing ? 'Maybe(Nothing)' : `Maybe(${inspect(this.$value)})`;
  }
}

let birthdayGreeting = function (who) {
  return 'Happy Birthday, ' + who + '🎉';
};

Maybe.of('c9s')
.map(birthdayGreeting)
.map(console.log);
