

|  p  |  q  |  r  |
|-----|-----|-----|
|  T  |  T  |  T  |
|  T  |  T  |  F  |
|  T  |  F  |  T  |
|  T  |  F  |  F  |
|  F  |  T  |  T  |
|  F  |  T  |  F  |
|  F  |  F  |  T  |
|  F  |  F  |  F  |


|  p·q  |  PvQ  |  p⊃q  |  P≡Q  |
|-------|-------|-------|-------|
|  TTT  |  TTT  |  TTT  |  TTT  |
|  TFF  |  TTF  |  TFF  |  TFF  |
|  FFT  |  FTT  |  FTT  |  FFT  |
|  FFF  |  FFF  |  FTF  |  FTF  |

|  P⊃Q  |  ≡  |  P·Q  |
|-------|-----|-------|
|  TTT  |  T  |  TTT  |
|  TFF  |  T  |  TFF  |
|  FTT  |  F  |  FFT  |
|  FTF  |  F  |  FFF  |


##ARGUMENTS :(========)







|  P⊃Q  |  P  |  ∴Q  |
|-------|-----|------|
|  TTT  |  T  |   T  |
|  TFF  |  T  |   F  |
|  FTT  |  F  |   T  |
|  FTF  |  F  |   F  |



|  ⊃(⊃)  |  Q  |  ∴⊃  |
|--------|-----|------|
|  T T   |  T  |   T  |
|  F F   |  T  |   T  |
|  T T   |  F  |   T  |
|  T T   |  F  |   T  |
|  T T   |  T  |   F  |   <- INVALIDO 
|  T F   |  T  |   T  |
|  T T   |  F  |   F  |
|  T T   |  F  |   T  |



|  p=     |  ~[Pv(B·Y)]v[(PvB)·(PvY)]  |
|---------|----------------------------|
| P=true  |    TT TFF  T  TTT T TTF    |
| P=false |    TFFTFF  T  FTT F FFF    |

IF P=TRU
