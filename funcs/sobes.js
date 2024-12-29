function doSum(...nums) {
    let res = 0
    for (const num of nums) {
        res += num
    }
    
    return res
}

// task 1
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
class ClassA {
    classB = new ClassB()

    num1 = 1
    num2 = 2
    num3 = 3

    doThing() {
        // const doSumBound = ...
        // this.classB.doAnotherThing(callbackWithoutArgs)
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



// task 3
// sort order ->by num, if equals - then sort by word
const data = [
    {num: 3, word: 'zzzzzzz'},
    {num: 1, word: 'xxxxxxx'},
    {num: 1, word: 'aaaaaaa'},
    {num: 5, word: 'bbbbbb'},
    {num: 5, word: 'cccccc'},
    {num: 5, word: 'aaaaaa'},
]

// Answer:
data.sort((a,b) => {
    const sortByNum = a.num - b.num;
    const sortByWord = a.word > b.word ? 1 : -1
    return sortByNum || sortByWord
})
console.log('DATA ====> ', data)
