#include <stdio.h>
#include <math.h>

int isprime(int n)
{
    int i;

    if (n < 2) return 0;
    if (n == 2) return 1;
    if (n % 2 == 0) return 0;
    for(i = 3; i <= sqrt(n); i += 2) {
        if (n % i == 0) return 0;
    }
    return 1;
}

int main() {
    int b, c, g, h = 0;
    //    a = 1;
    b = 57 * 100 + 100000;
    c = b + 17000;
    do {
        if (!isprime(b)) {
            h++;
        }
        g = b - c;
        b += 17;
    } while (g != 0);
    printf("%d\n", h);
    return 0;
}
