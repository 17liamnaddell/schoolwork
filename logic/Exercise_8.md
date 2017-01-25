#Exercise 8

1. S⊃P W⊃F (S·W)⊃(PvF)

|  S⊃P  |  W⊃F  |  (S·W)⊃(PvF)  |
|-------|-------|---------------|
|  TTF  |  TTF  |   TTT F FFF   |   <-valid


2. 
B⊃R P R⊃~P ~B

|  B⊃R  |  P  |  R⊃~P  |  ~B  |
|-------|-----|--------|------|
|  TTT  |  T  |  TTFT  |  FT  |    <- valid 

3. 
M⊃P ~M ~P

|  M⊃P  |  ~M  |  ~P  |
|-------|------|------|
|  FTT  |  TF  |  TF  |      <- invalid


4. 

S⊃(PvA) H⊃L P⊃H ~L S A

|  S⊃(PvA)  |  H⊃L  |  P⊃H  |  ~L  |  S  |  A  |
|-----------|-------|-------|------|-----|-----|
|  TT FFF   |  FTF  |  FTF  |  TF  |  T  |  F  |    <-invalid


5. 

A⊃P U⊃J S⊃(AvU) S PvJ


|  A⊃P  |  U⊃J  |  S⊃(AvU)  |  S  |  PvJ  |
|-------|-------|-----------|-----|-------|
|  FTF  |  FTF  |  TT FTF   |  T  |  FFF  |     <- valid



6.
O⊃~C ~O C
FTTF TF F   <- Invalid

 
7. 

|  P  |  ∴~PvQ  |
|-----|---------|
|  T  |   FTTF  |    <- invalid argument

8.

|  P⊃Q  |  ∴~Q⊃~P  |
|-------|----------|
|  TTF  |   TFFFT  |  <- contradiction, valid argument


9.

|  P⊃Q  |  ~q  |  ∴p≡q  |
|-------|------|--------|
|  TTF  |  TF  |   TFF  | contradiction, valid argument

10. 

|  P⊃(Q⊃R)  |  Q  | ∴R⊃P  |
|-----------|-----|-------|
|  FT TTT   |  T  |  TFF  |    no contradiction, invalid argument

11. 
|  P⊃(~Q⊃R)  |  P  |  ∴~R⊃Q  |
|------------|-----|---------|
|  TT TFFF   |  T  |   TFFF  | contradiction, valid argument

12

|  (p⊃q)·[(p·q)⊃r]  |  p⊃(r⊃s)  |  ∴p⊃s  |
|-------------------|-----------|--------|
|   TTT T  TTT FF   |  TT FTF   |   TFF  |
        ^                                         Contradiction, valid argument

