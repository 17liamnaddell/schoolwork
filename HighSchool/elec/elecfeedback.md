# Electronic feedback

This allows us to tame opamp gain
Electronic feedback is feedback without moving parts such as motors
To find the gain of a opamp with negative feedback in the negative terminal, we need to find vout/vin
Warning: math and graphical explanations alert: I cannot really explain what is going on here lol
After a lot of math we get that vout=-(Rf/Rin)Vin
OpAmp rules

1. No current flows into the opamp inputs
2. W/ negative feedback, v+=v-
3. B/c the negative feedback will deminish the current, the current will be multiplied by the opamp, and then it will be replenished, the cycle repeats.

this whole process allows us to tame the gain of the opamps(which is huge). Now we can amplify any current to almost any degree as we chose

