var Record = function() {
  var recording = false;
  var recordingNumnber = 1;
  var recordRTC;
  var stop = false;

  function parseRecording() {
    if (!recording) {
      return;
    }

    recordRTC.stopRecording(function (audioVideoWebMURL) {
      var recordedBlob = recordRTC.getBlob();
      recordRTC.save(`foosball-${recordingNumnber}.wav`);
      recordingNumnber += 1;
      recording = false;
    });

    if (stop) {
      stop = false;
      return;
    } else {
      startRecording();
    }
  }

  function successCallback(stream) {
    var options = {
      mimeType: 'audio',
      recorderType: StereoAudioRecorder,
      numberOfAudioChannels: 1,
    };
    recordRTC = RecordRTC(stream, options);
    recordRTC.startRecording();
    recording = true;

    setTimeout(parseRecording, 5000);
  }

  function stopRecording() {
    stop = true;
    parseRecording();
  }

  function startRecording() {
    if (recording) {
      return;
    }
    recording = true;

    navigator.mediaDevices
             .getUserMedia({ video: false, audio: true })
             .then(successCallback)
             .catch(console.warn);
  }

  return {
    start: startRecording,
    stop: stopRecording
  }
}();
