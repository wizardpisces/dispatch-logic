// ts-node generateParenthesis.ts
function fn(n:number){
    let result:string[] = []
    let str = ''

    function dfs(str:string,left:number,right:number){
        if(left === 0 && right === 0){
            return result.push(str)
        }
        if(left > right){
            return
        }
        if(left>0){ //优先添加左括号
            dfs(str+'(',left-1,right)
        }
        if(right>0){
            dfs(str+')',left,right-1)
        }
    }
    dfs(str,n,n)
    return result
}

type StackNode = { left: number, right: number, res: string }

function noneRecursive(n:number){
    let result: string[] = []
    let stackList: StackNode[] = [{left:n,right:n,res:''}]

    if(n===0) return []

    while (stackList.length>0){
        let stack = stackList.shift() as StackNode
        let left = stack.left,right = stack.right,str = stack.res
        if (left === 0 && right === 0) {
            result.push(str)
            continue
        }
        if (left > right) {
            continue
        }
        if (left > 0) { //优先添加左括号
            stackList.push({ res:str + '(', left: left - 1, right})
        }
        if (right > 0) {
            stackList.push({res:str + ')', left, right:right - 1})
        }
    }

    return result
}
// @ts-ignore
console.log(fn(3))
// @ts-ignore
console.log(noneRecursive(3))