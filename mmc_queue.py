import math

lamb = 1/4
mu = 1/3
c = 1
rho = 0
p0 = 0
P = 0
L_Q = 0

def calculate_p0():
    global p0, c, rho

    sum = 0
    for n in range(c):
        value = (c*rho)**n / math.factorial(n)
        sum = sum + value

    value1 = (c*rho)**c
    value2 = 1/math.factorial(c)
    value3 = 1/(1-rho)
   
    sum = sum + (value1*value2*value3)
    p0 = 1/sum

def calculate_P():
    global rho, c, mu, p0, P

    value1 = (c*rho)**c * p0
    value2 = math.factorial(c) * (1-rho)
    P = value1/value2


def main():
    global rho, c, mu, p0, P

    rho = lamb / (c * mu)
    print(rho)

    calculate_p0()
    print(p0)

    calculate_P()
    print(P)

    L_Q = (rho * P) / (1 - rho)
    print(L_Q)

    waste = c*(1-rho)*10 + L_Q*15
    print("WASTE: ", waste)


if __name__ == "__main__":
    main()
