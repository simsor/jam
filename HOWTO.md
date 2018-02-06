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
