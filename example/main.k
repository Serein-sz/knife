let a = 4;
let b = 2;
func subtract(x, y) {
    return x - y;
}
func add(x, y, subtract) {
    return x + y + subtract(x, y)
}
print(add(a, b, subtract) == 8, print() == null)
