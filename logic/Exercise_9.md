1.

|  p≡q  | q≡r  |  ∴p≡r  |
|-------|------|--------|
|  TTT  | TTT  |   FFF  |      <- no contradiction, invalid argument

2. 

Guess one, p = true
no contradiction

|  PvQ  |  ∴p·q  |
|-------|--------|
|  TTF  |   TFF  |

Guess two, P=false
No contradiction, Invalid argument

|  PvQ  |  ∴p·q  |
|-------|--------|
|  FTT  |   FFT  |


3. 
Guess one, P=true
No contradiction

|  p⊃q  |  q≡r  |  ∴p⊃r  |
|-------|-------|--------|
|  TTT  |  TTT  |   TFT  |

Guess two, P=false
No contradiction, invalid argument

|  p⊃q  |  q≡r  |  ∴p⊃r  |
|-------|-------|--------|
|  FTT  |  TTT  |   FFT  |



4. 

Guess one, Q=true
no contradiction

|  (P⊃Q)v(R⊃S)  |  PvR  |  ∴QvS  |
|---------------|-------|--------|
|   TTT T FTF   |  TTF  |   TFF  |

Guess two, Q=false
No contradiction statement=invalid

|  (P⊃Q)v(R⊃S)  |  PvR  |  ∴QvS  |
|---------------|-------|--------|
|   FTF T TFF   |  FTT  |   FFF  |

5.

Guess one, R=false

|  PvQ  |  ~[Q·(R⊃P)]  |  ∴~(P≡Q)  |
|-------|--------------|-----------|
|  TTT  |  T TF FTT    |   F TTT   |

Contradiction

Guess two, R=true
Contradiction, statement=valid

|  PvQ  |  ~[Q·(R⊃P)]  |  ∴~(P≡Q)  |
|-------|--------------|-----------|
|  TTT  |  T TF TTT    |   F TTT   |


6.
Guess one, Q=true

|  P⊃(Q⊃R)  |  Q⊃(P⊃R)  |  ∴(PvQ)⊃R  |
|-----------|-----------|------------|
|  FT TFF   |  TT FTF   |    FTT FF  |

No contradiction
Guess two, Q=false

|  P⊃(Q⊃R)  |  Q⊃(P⊃R)  |  ∴(PvQ)⊃R  |
|-----------|-----------|------------|
|  TT FTF   |  FT TFF   |    TTF FF  |

No contradiction, statement=invalid

