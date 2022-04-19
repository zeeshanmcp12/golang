function find_max(nums) {
    let max_num = Number.NEGATIVE_INFINITY; // smaller than all other numbers
    for (let num of nums) {
    if (num > max_num) {
    // (Fill in the missing line here)
     num = max_num
     console.log(num)
    }
   }
    return max_num;
   }