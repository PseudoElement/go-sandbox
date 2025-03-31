class Parent {
    value;

    constructor(value) {
        this.value = value
        console.log('value ==> ', this.value)
    }
}

class Child extends Parent {
    constructor(value) {
        super(value)
        this.value = 'xxx'
    }
}

console.log("Child ==> ", new Child('zzz'))