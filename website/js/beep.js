var notes = [];
var tempo = 1;
var isRunning = false;
var shouldStop = false;
var notesFreqs = {
		"B0":  31,
		"C1":  33,
		"CS1": 35,
		"D1":  37,
		"DS1": 39,
		"E1":  41,
		"F1":  44,
		"FS1": 46,
		"G1":  49,
		"GS1": 52,
		"A1":  55,
		"AS1": 58,
		"B1":  62,
		"C2":  65,
		"CS2": 69,
		"D2":  73,
		"DS2": 78,
		"E2":  82,
		"F2":  87,
		"FS2": 93,
		"G2":  98,
		"GS2": 104,
		"A2":  110,
		"AS2": 117,
		"B2":  123,
		"C3":  131,
		"CS3": 139,
		"D3":  147,
		"DS3": 156,
		"E3":  165,
		"F3":  175,
		"FS3": 185,
		"G3":  196,
		"GS3": 208,
		"A3":  220,
		"AS3": 233,
		"B3":  247,
		"C4":  262,
		"CS4": 277,
		"D4":  294,
		"DS4": 311,
		"E4":  330,
		"F4":  349,
		"FS4": 370,
		"G4":  392,
		"GS4": 415,
		"A4":  440,
		"AS4": 466,
		"B4":  494,
		"C5":  523,
		"CS5": 554,
		"D5":  587,
		"DS5": 622,
		"E5":  659,
		"F5":  698,
		"FS5": 740,
		"G5":  784,
		"GS5": 831,
		"A5":  880,
		"AS5": 932,
		"B5":  988,
		"C6":  1047,
		"CS6": 1109,
		"D6":  1175,
		"DS6": 1245,
		"E6":  1319,
		"F6":  1397,
		"FS6": 1480,
		"G6":  1568,
		"GS6": 1661,
		"A6":  1760,
		"AS6": 1865,
		"B6":  1976,
		"C7":  2093,
		"CS7": 2217,
		"D7":  2349,
		"DS7": 2489,
		"E7":  2637,
		"F7":  2794,
		"FS7": 2960,
		"G7":  3136,
		"GS7": 3322,
		"A7":  3520,
		"AS7": 3729,
		"B7":  3951,
		"C8":  4186,
		"CS8": 4435,
		"D8":  4699,
		"DS8": 4978,
	};


// create web audio api context
var audioCtx = new (window.AudioContext || window.webkitAudioContext)();

function stringToNotes(s) {
    let notes = [];
    let lines = s.split("\n");
    for (let i=0; i < lines.length; i++) {
        let line = lines[i].trim();
        let parts = line.split(" ")

        if (parts[0] == ";" || line == "") {
            continue;
        } else if (parts[0] == "FREQ") {
            let freq = parseFloat(parts[1]);
            let duration = parseFloat(parts[2]);
            
            if (isNaN(freq) || isNaN(duration)) { alert("Error line " + (i+1) + ": not a number"); return; }

            notes.push([freq, duration]);
        } else if (parts[0] == "PAUSE") {
            let duration = parseFloat(parts[1]);

            if (isNaN(duration)) { alert("Error line " + (i+1) + ": not a number"); return; }

            notes.push([0, duration]);
        } else if (parts[0] == "TEMPO") {
            let t = parseFloat(parts[1]);

            if (isNaN(t)) { alert("Error line "+(i+1)+": not a number"); return;}

            notes.push(["tempo", bpmToMs(t)]);
        } else {
            let note = parts[0];
            let duration = parseFloat(parts[1]);

            if (isNaN(duration)) { alert("Error line " + (i+1) + ": not a number"); return; }
            
            let notesKeys = Object.keys(notesFreqs);
            let found = false;
            for (let j=0; j < notesKeys.length; j++) {
                if (notesKeys[j] == note) {
                    notes.push([notesFreqs[note], duration]);
                    found = true;
                    break;
                }
            }

            if (!found) {
                alert("Error line " + (i+1) + ": unknown note " + note);
                return;
            }
        }
    }

    return notes;
}

function playMelody(){
	if (notes.length > 0){
		note = notes.pop();
                if (note[0] == "tempo") {
                    tempo = note[1];
                    playMelody();
                    return;
                }
		playNote(note[0], note[1]*tempo);
	} else {
            isRunning = false;
            shouldStop = false;
        }
}

function playNote(frequency, duration) {
	// create Oscillator node
	var oscillator = audioCtx.createOscillator();
	
	oscillator.type = 'square';
	oscillator.frequency.value = frequency; // value in hertz
	oscillator.connect(audioCtx.destination);
	oscillator.start();
		
	setTimeout(
		function(){
			oscillator.stop();
                        if (shouldStop) {
                            isRunning = false;
                            shouldStop = false;
                            notes = []
                            return;
                        }
                        playMelody();
		}, duration);
}

function bpmToMs(bpm) {
   return 1.0 / (bpm/60.0) * 1000; 
}
