
const list1 = {
    val: 1,
    next: {
      val: 2, 
      next: {
        val: 3,
        next: {
          val: 4
        }
      }
    }
  }
  const list2 = {
    val: 1,
    next: {
      val: 5, 
      next: {
        val: 5,
        next: {
          val: 3
        }
      }
    }
  }
  
  function listToArray(list) {
    const arr = []
    if(!list || list.next === null){
        return arr
    }
    let next = list.next
    let value = list.val || 0
    while(true){
        if(!next){
            break
        }
        arr.push(value)
        value = next.val
        next = next.next
    }
    arr.push(value)
    return arr;
  }
  
  function arrayToList(arr) {
    function ListNode(val, next) {
      this.val = (val === undefined ? 0 : val)
      this.next = (next === undefined ? null : next)
    }
  
    const head = arr.reverse().reduce((acc, curr) => {
      if (acc == null) {
        acc = new ListNode(curr);
    
      } else {
        acc = new ListNode(curr, acc);
      }
      return acc;
    }, null);
    return head
  }
  
  function mergeTwoLists(list1, list2) {
    const arr1 = listToArray(list1)
    const arr2 = listToArray(list2)
    const sortedArr = [...arr1, ...arr2].sort((a, b) => a - b)
    console.log(sortedArr)
    return arrayToList(sortedArr)
  };
  
  console.log(mergeTwoLists(list1, list2))