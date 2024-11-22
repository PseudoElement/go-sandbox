function findElInArrayIterator(arr, cond) {
    let idx = 0;
    
    const iter = () => {
        return {
            next() {
                if (idx > arr.length) return {value: null, done: true}
                
                const currEl = arr[idx];
                const res = {
                    value: {el: currEl, idx: idx}, 
                    done: cond(currEl, idx, arr)
                }
                idx++;
                
                return res;
            },
            [Symbol.iterator](){
               return this;
            }
        } 
    }
    
    let res
    let next = iter().next()
    while(!next.done) {
        next = iter().next()
        res = next.value
    }
    
    return res
}

const res = findElInArrayIterator(
    [1,2,3,5,6,7,8,9,0,0,100, 23434.4, 23545, 457, 567678],
    (el, i, _) => el == 5676780
)
console.log(res)