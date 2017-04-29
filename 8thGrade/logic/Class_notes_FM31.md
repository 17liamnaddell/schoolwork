1. Decompose non-branching members first
  1. p·q
  4. P        1·D
  5. q        1·D
2. Decompose members that result in the closing of one/more branches 
3. Stop when the truth tree answers the question being asked 
4. Decompose more complex Propositions first.
```
1.                            ~P           SM
2.                            Q            SM
3.                            ~(P·Q)       SM
                              /     \
4.                           ~P     ~Q 
                                    2x4
```
```
1.                             P·Q          SM
2.                             Pv~Q         SM
3.                             P            1·D
4.                             Q            1·D
                            /     \
5.                         P      ~Q
6.                                4x5
```
```
1.                         ~(Pv~Q)        SM
2.                         P≡Q            SM
3.                         ~P             1~vD
4.                         ~~Q            1~vD
5.                          Q             4~~D
6.                        /   \
7.                       P     ~P         2≡D
8.                       Q     ~Q         2≡D
```
```
1.                          P⊃Q           SM
2.                          ~(P⊃R)        SM
3.                          ~(Q≡R)        SM
4.                          P             2~⊃D
5.                          ~R            2~⊃D
6.                        /    \
7.                      ~P      Q         1⊃D
8.                      4x6   /   \
9.                           R    ~R
10.                          ~Q    Q
11.                          6x7    
```
