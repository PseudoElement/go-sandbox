function doSum(...nums) {
    let res = 0
    for (const num of nums) {
        res += num
    }
    
    return res
}

// task 1
// Fix code, to see in terminal 'Sum is 6'
const fetchNumbers = () => new Promise(res => 
    setTimeout(() => res([1, 2, 3]), 1_000)
)

function functionCall() {
    const respNums = fetchNumbers()
    const sum = doSum(...respNums)
    console.log('Sum is ', sum)
}

// functionCall()

// task 2
// You need to call method doSum inside method doAnotherThing,
// doAnotherThing takes 1 argument callback doSum method
// Restriction: you need to call doAnotherThing method inside doThing method
class ClassA {
    classB = new ClassB()

    num1 = 1
    num2 = 2
    num3 = 3

    doThing() {
        // const doSumBound = 
        return this.classB.doAnotherThing(doSumBound)
    }

    doSum(...args) {
        return doSum(...args)
    }
}

class ClassB {
    doAnotherThing(callback) {
        callback()
    }
}

console.log('ANSWER ===> ', new ClassA().doThing())


// task 3
// sort order -> firstly by num, if equals - then sort by word
const data = [
    {num: 3, word: 'zzzzzzz'},
    {num: 1, word: 'xxxxxxx'},
    {num: 1, word: 'aaaaaaa'},
    {num: 5, word: 'bbbbbb'},
    {num: 5, word: 'cccccc'},
    {num: 10, word: 'ddddddd'},
    {num: 5, word: 'aaaaaa'},
]

function sorter() {
    // write implementation
}

// Answer:
data.sort((a,b) => {
    const sortByNum = a.num - b.num;
    const sortByWord = a.word > b.word ? 1 : -1
    return sortByNum || sortByWord
})
console.log('DATA ====> ', data)

// task 4
// returns rejected or resolved promise
function maybeThrowError() {
    return new Promise((res, rej) => {
        Math.random() > 0.5 
            ? res('Hello, Abshishek.') 
            : rej('More than 5.')
    })
}

async function asyncCall() {
    const value =  maybeThrowError()
    console.log('Next code execution... Value is ', value)
    // some logic...
}
// asyncCall()
