// https://developers.google.com/web/fundamentals/media/recording-audio/
function viewHandler() {
  console.log("Handle record view");

  var recorder = document.getElementById('recorder');
  var player = document.getElementById('player');

  recorder.addEventListener('change', function(e) {
    var file = e.target.files[0];
    // Do something with the audio file.
    player.src =  URL.createObjectURL(file);
  });
}

window.addEventListener('load', viewHandler);
