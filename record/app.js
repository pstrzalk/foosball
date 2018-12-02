function viewHandler() {
  var recordRTC;

  function successCallback(stream) {
    var options = {
      mimeType: 'audio',
      recorderType: StereoAudioRecorder,
      numberOfAudioChannels: 1,
    };
    recordRTC = RecordRTC(stream, options);
    recordRTC.startRecording();
  }

  function errorCallback(error) {
    console.warn(error);
  }

  var mediaConstraints = { video: false, audio: true };

  navigator.mediaDevices
           .getUserMedia(mediaConstraints)
           .then(successCallback)
           .catch(errorCallback);

  var btnStopRecording = document.querySelector('.btnStopRecording');
  btnStopRecording.onclick = function () {
    recordRTC.stopRecording(function (audioVideoWebMURL) {
      var recordedBlob = recordRTC.getBlob();
      recordRTC.save('foosball.wav');
    });
  };
}

window.addEventListener('load', viewHandler);
