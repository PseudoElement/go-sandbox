// 2(abc)3(ef)dd -> abcabcefefefdd
function multiplyStrings(str) {
    let res = ''
    let tempStr = ''
    let multiplier = 1
    for(let i = 0; i < str.length; i++) {
        const char = str[i]
        const nextChar = str[i+1]
        const prevChar = str[i-1]
        const num = getNum(char)

        if(char === '(') continue
        if(char === ')'){
            for(let j = 0; j < multiplier; j++) res += tempStr
            tempStr = ''
            continue
        }
        if(num && nextChar === '(') {
            multiplier = num
            continue
        }
        if(prevChar === '(' || tempStr.length) {
            tempStr += char
            continue
        }

        res += char
    }

    return res
}

function getNum(val) {
    const num = Number(val)
    return isNaN(num) ? null : num
}

console.log(multiplyStrings('2(abc)5(ef)dd55'))
