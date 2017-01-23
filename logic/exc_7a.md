#exercise 7


|  P  |  ⊃   |  q  |  ∴p  |  ⊃  |  (p  |  ·  |  q)  |
|-----|------|-----|------|-----|------|-----|------|
|  T  |  T   |  T  |   T  |  T  |  T   |  T  |  T   |      <- valid
|  T  |  F   |  F  |   T  |  F  |  T   |  F  |  F   |
|  F  |  T   |  T  |   F  |  T  |  F   |  F  |  T   |      <- valid
|  F  |  T   |  F  |   F  |  T  |  F   |  F  |  F   |      <- valid


Valid


#exercise 14

T⊃P D⊃P TvD P
         V                 V                 V                        

|  T  |  ⊃  |  P  |  D  |  ⊃  |  P  |  T  |  v  |  D  |  P  |
|-----|-----|-----|-----|-----|-----|-----|-----|-----|-----|
|  T  |  T  |  T  |  T  |  T  |  T  |  T  |  T  |  T  |  T  |    <-Valid
|  T  |  T  |  T  |  F  |  T  |  T  |  T  |  T  |  F  |  T  |    <-Valid
|  T  |  F  |  F  |  T  |  F  |  F  |  T  |  T  |  T  |  F  |    <-nothing
|  T  |  F  |  F  |  F  |  T  |  F  |  T  |  T  |  F  |  F  |    <-nothing
|  F  |  T  |  T  |  T  |  T  |  T  |  F  |  T  |  T  |  T  |    <-Valid
|  F  |  T  |  T  |  F  |  T  |  T  |  F  |  F  |  F  |  T  |    <-Valid
|  F  |  T  |  F  |  T  |  F  |  F  |  F  |  T  |  T  |  F  |    <-Nothing
|  F  |  T  |  F  |  F  |  T  |  F  |  F  |  F  |  F  |  F  |    <-Nothing

valid :0
