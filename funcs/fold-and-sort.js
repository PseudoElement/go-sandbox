function foldAndSort(arr) {
    const sortedArr = [...arr].sort((a, b) => a - b)
    const matrix = []
    for(let i = 0; i < sortedArr.length; i += 2){
        const el = sortedArr[i]
        const nextEl = sortedArr[i+1]
        matrix.push(nextEl ? [el, nextEl] : [el])
    } 

    return matrix.map(nums => {
        if(nums.length === 1) {
            return nums[0].toString()
        }
        return `${nums[0]}-${nums[1]}`
    }).join(",")
}



console.log(foldAndSort([1,5,7,6,9,3,2]))
