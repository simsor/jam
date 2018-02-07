# How to make music!

A music file is simply a succession of notes (one per line) with the corresponding duration.

```
A3 500
B2 150
PAUSE 100
A4 400
```

This will play an A on the third octave for 500ms, then B2 for 150, a pause for 100ms and finally A4 for 400ms before exiting.

The syntax supports comments: every line beginning with a ";" will be ignored

If you want to play a specific frequency, you can use the FREQ syntax:

``
FREQ 440 500
``

Will play a sound at 440Hz for 500ms

## Managing the tempo

You can use the ``TEMPO`` command to change the speed your music will be played at (in beats per minute, or bpm).

After issuing the ``TEMPO`` command, times will we treated as multiples of the tempo.

For example:

```
TEMPO 120
A3 1
B4 0.5
C4 0.5
PAUSE 2
```

Will set the tempo to 120bpm, then play A3 for 1 beat (500ms), B4 for half a beat, C4 for half a beat and pause for 2 beats
