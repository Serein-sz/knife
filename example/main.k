let a = 4;
let b = 2;
func subtract(x, y) {
    return x - y;
}
func add(x, y, subtract) {
    return x + y + subtract(x, y)
}
func findNumbers(nums) {
    
}
print(add(a, b, subtract))
print(add(a, b, subtract) == 8)
