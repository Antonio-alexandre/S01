10 INPUT "A = "; A
20 INPUT "B = "; B
30 INPUT "C = "; C

40 LET D = B * B - 4 * A * C

50 IF D < 0 THEN GOTO 120

60 LET SQRTD = SQR(D)

70 LET X1 = (-B - (-SQRTD)) / (2 * A)
  
80 LET X2 = (-B - SQRTD) / (2 * A)

90 PRINT "As raízes da fórmula são: "

100 PRINT X1
110 PRINT X2
  
120 END