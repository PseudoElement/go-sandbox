package main

/**
* CONVERT DECIMAL TO BINARY
* 	13 / 2 = 6 // 1
*	6 / 2 = 3 // 0
* 	3 / 2 = 1 // 1
* 	1 / 2 = 0 // 1
*  Собрать остаток от деления в обратном порядке - 13(decimal) = 1101(binary)
 */

/**
 * CONVERT BINARY To DECIMAL
 * 1011(binary) = 11(decimal)
 * (1 * 2**0) + (1 * 2**1) + (0 * 2**2) + (1 * 2**3) = 1 + 2 + 0 + 8 = 11
 */

/**
 * LEFT SHIFT
 * Нужно добавить в конец количество нулей равное сдвигу
 * 1101(binary) << 2 = 110100
 * 13(decimal) << 2 = 13 * 2**2 = 110100
 */

/**
 * RIGHT SHIFT
 * Cut last bits from the end of binary
 * 11011(bin) >> 2 = 27 / 2**2 = 110(bin) = 6(dec)
 * 27(dec) >> 2 = 110
 */

func main() {}
